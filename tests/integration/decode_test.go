package integration_test

import (
	"errors"
	"reflect"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func parseDecodeDocument(t *testing.T, text string, mode yini.Mode) *yini.Document {
	t.Helper()
	result, err := yini.ParseString(text, yini.Options{Mode: mode})
	if err != nil {
		t.Fatalf("ParseString: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	return result.Document
}

func TestDecodeDocumentAndSection(t *testing.T) {
	type label string
	type cache struct {
		Size int8 `yini:"size"`
	}
	type service struct {
		Port    uint16         `yini:"port"`
		Name    label          `yini:"name"`
		Ports   []uint16       `yini:"ports"`
		Details map[string]any `yini:"details"`
		Cache   *cache         `yini:"Cache"`
	}
	type config struct {
		Service service `yini:"Service"`
	}
	text := "^ Service\nport = 8080\nname = 'Echo'\nports = [80, 443]\ndetails = { owner: 'Mira', enabled: true }\n^^ Cache\nsize = 4\n/END"
	want := service{
		Port: 8080, Name: "Echo", Ports: []uint16{80, 443},
		Details: map[string]any{"owner": "Mira", "enabled": true},
		Cache:   &cache{Size: 4},
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			doc := parseDecodeDocument(t, text, mode)
			var whole config
			if err := doc.Decode(&whole); err != nil {
				t.Fatalf("Document.Decode: %v", err)
			}
			if !reflect.DeepEqual(whole.Service, want) {
				t.Fatalf("whole.Service = %#v, want %#v", whole.Service, want)
			}
			var part service
			if err := doc.Sections[0].Decode(&part); err != nil {
				t.Fatalf("Section.Decode: %v", err)
			}
			if !reflect.DeepEqual(part, want) {
				t.Fatalf("part = %#v, want %#v", part, want)
			}
			part.Details["owner"] = "changed"
			original := doc.Values()["Service"].(map[string]any)["details"].(map[string]any)["owner"]
			if original != "Mira" {
				t.Fatalf("decoding modified document value: %#v", original)
			}
		})
	}
}

func TestDecodeLenientRootAndNulls(t *testing.T) {
	type service struct {
		Port int64 `yini:"port"`
	}
	type config struct {
		Label   string         `yini:"label"`
		Service service        `yini:"Service"`
		Ptr     *int64         `yini:"ptr"`
		Items   []string       `yini:"items"`
		Meta    map[string]any `yini:"meta"`
		Any     any            `yini:"any"`
		Absent  string
	}
	text := "label = 'North'\nptr = null\nitems = null\nmeta = null\nany = null\n^ Service\nport = 9"
	doc := parseDecodeDocument(t, text, yini.ModeLenient)
	dst := config{Label: "old", Ptr: new(int64), Items: []string{"old"}, Meta: map[string]any{"old": true}, Any: "old", Absent: "old"}
	if err := doc.Decode(&dst); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	want := config{Label: "North", Service: service{Port: 9}}
	if !reflect.DeepEqual(dst, want) {
		t.Fatalf("decoded = %#v, want %#v", dst, want)
	}
}

func TestDecodeTypedMapsAndCopiedInterface(t *testing.T) {
	type ratio float64
	type config struct {
		Counts map[string]int64 `yini:"counts"`
		Data   any              `yini:"data"`
		Ratio  ratio            `yini:"ratio"`
		Extra  map[string]int64 `yini:"Extra"`
	}
	text := "^ Config\ncounts = { left: 1, right: 2 }\ndata = [{ tags: ['a', 'b'] }]\nratio = 0.25\n^^ Extra\ncount = 5\n/END"
	doc := parseDecodeDocument(t, text, yini.ModeStrict)
	var dst config
	if err := doc.Sections[0].Decode(&dst); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if !reflect.DeepEqual(dst.Counts, map[string]int64{"left": 1, "right": 2}) ||
		!reflect.DeepEqual(dst.Extra, map[string]int64{"count": 5}) || dst.Ratio != ratio(0.25) {
		t.Fatalf("decoded map/float values = %#v", dst)
	}
	items := dst.Data.([]any)
	items[0].(map[string]any)["tags"].([]any)[0] = "changed"
	original := doc.Values()["Config"].(map[string]any)["data"].([]any)[0].(map[string]any)["tags"].([]any)[0]
	if original != "a" {
		t.Fatalf("interface value shares document storage: %#v", original)
	}
}

func TestDecodeExactFieldNames(t *testing.T) {
	type config struct{ Name string }
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			valid := parseDecodeDocument(t, "^ Config\nName = 'Ada'\n/END", mode)
			var dst config
			if err := valid.Sections[0].Decode(&dst); err != nil || dst.Name != "Ada" {
				t.Fatalf("exact field name = %#v, %v", dst, err)
			}
			invalid := parseDecodeDocument(t, "^ Config\nname = 'Ada'\n/END", mode)
			var detail *yini.DecodeError
			if err := invalid.Sections[0].Decode(&dst); !errors.As(err, &detail) || detail.Path != "Config.name" || detail.Position != (yini.Position{Line: 2, Column: 1}) {
				t.Fatalf("case mismatch = %v; want Config.name at 2:1", err)
			}
			if dst.Name != "Ada" {
				t.Fatalf("failed decode changed destination: %#v", dst)
			}
		})
	}
}

func TestDecodeValueErrors(t *testing.T) {
	tests := []struct {
		name    string
		body    string
		dst     func() any
		path    string
		line    int
		column  int
		message string
	}{
		{"string mismatch", "count = 'oops'", func() any {
			return &struct {
				Count int64 `yini:"count"`
			}{}
		}, "Config.count", 2, 1, "YINI string into Go int64"},
		{"signed overflow", "count = 128", func() any {
			return &struct {
				Count int8 `yini:"count"`
			}{}
		}, "Config.count", 2, 1, "YINI integer into Go int8"},
		{"negative unsigned", "count = -1", func() any {
			return &struct {
				Count uint8 `yini:"count"`
			}{}
		}, "Config.count", 2, 1, "YINI integer into Go uint8"},
		{"float to integer", "count = 1.5", func() any {
			return &struct {
				Count int64 `yini:"count"`
			}{}
		}, "Config.count", 2, 1, "YINI float into Go int64"},
		{"null to scalar", "count = null", func() any {
			return &struct {
				Count int64 `yini:"count"`
			}{}
		}, "Config.count", 2, 1, "YINI null into Go int64"},
		{"list item", "items = [1, 'bad']", func() any {
			return &struct {
				Items []int64 `yini:"items"`
			}{}
		}, "Config.items[1]", 2, 1, "YINI string into Go int64"},
		{"typed map value", "counts = { good: 1, wrong: 'bad' }", func() any {
			return &struct {
				Counts map[string]int64 `yini:"counts"`
			}{}
		}, "Config.counts.wrong", 2, 1, "YINI string into Go int64"},
		{"unknown member", "count = 1", func() any { return &struct{ Other int64 }{} }, "Config.count", 2, 1, "no matching field"},
		{"section mismatch", "^^ Child\nvalue = 1", func() any { return &struct{ Child int64 }{} }, "Config.Child", 2, 1, "YINI section into Go int64"},
		{"unknown section", "^^ Child\nvalue = 1", func() any { return &struct{ Other int64 }{} }, "Config.Child", 2, 1, "no matching field"},
		{"quoted name path", "`ui.theme` = true", func() any {
			return &struct {
				Theme string `yini:"ui.theme"`
			}{}
		}, "Config[\"ui.theme\"]", 2, 1, "YINI boolean into Go string"},
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, test := range tests {
			t.Run(string(mode)+"/"+test.name, func(t *testing.T) {
				doc := parseDecodeDocument(t, "^ Config\n"+test.body+"\n/END", mode)
				dst := test.dst()
				err := doc.Sections[0].Decode(dst)
				var detail *yini.DecodeError
				if !errors.As(err, &detail) {
					t.Fatalf("Decode error = %v; want *DecodeError", err)
				}
				if detail.Path != test.path || detail.Position != (yini.Position{Line: test.line, Column: test.column}) {
					t.Fatalf("error path/location = %q %#v; want %q %d:%d", detail.Path, detail.Position, test.path, test.line, test.column)
				}
				if !strings.Contains(detail.Error(), test.message) || !strings.Contains(detail.Error(), test.path) {
					t.Fatalf("error text = %q; want %q and path", detail.Error(), test.message)
				}
			})
		}
	}
}

func TestDecodeDeterministicOrderAndAtomicity(t *testing.T) {
	type child struct {
		Value int64 `yini:"value"`
	}
	type config struct {
		Child child `yini:"Child"`
		Late  int64 `yini:"late"`
	}
	doc := &yini.Document{
		Members: []*yini.Member{{Name: "late", Value: "bad", Position: yini.Position{Line: 4, Column: 1}}},
		Sections: []*yini.Section{{Name: "Child", Position: yini.Position{Line: 2, Column: 1}, Members: []*yini.Member{
			{Name: "value", Value: "bad", Position: yini.Position{Line: 3, Column: 1}},
		}}},
	}
	dst := config{Child: child{Value: 12}, Late: 13}
	before := dst
	err := doc.Decode(&dst)
	var detail *yini.DecodeError
	if !errors.As(err, &detail) || detail.Path != "Child.value" || detail.Position != (yini.Position{Line: 3, Column: 1}) {
		t.Fatalf("first error = %v; want Child.value at 3:1", err)
	}
	if dst != before {
		t.Fatalf("destination changed on error: %#v", dst)
	}

	parsed := parseDecodeDocument(t, "^ Config\nitem = { z: true, a: 'bad' }\n/END", yini.ModeStrict)
	var inline struct {
		Item struct {
			A int64 `yini:"a"`
			Z int64 `yini:"z"`
		} `yini:"item"`
	}
	err = parsed.Sections[0].Decode(&inline)
	if !errors.As(err, &detail) || detail.Path != "Config.item.a" || detail.Position != (yini.Position{Line: 2, Column: 1}) {
		t.Fatalf("inline object error = %v; want Config.item.a at 2:1", err)
	}
}

func TestDecodeDestinationValidationAndEmbedding(t *testing.T) {
	doc := &yini.Document{}
	var nilDocument *yini.Document
	var nilSection *yini.Section
	var nilStruct *struct{}
	for _, test := range []struct {
		name string
		run  func() error
	}{
		{"nil document", func() error { return nilDocument.Decode(&struct{}{}) }},
		{"nil section", func() error { return nilSection.Decode(&struct{}{}) }},
		{"nil destination", func() error { return doc.Decode(nilStruct) }},
		{"non-pointer destination", func() error { return doc.Decode(struct{}{}) }},
		{"map destination", func() error { return doc.Decode(&map[string]any{}) }},
		{"duplicate tags", func() error {
			return doc.Decode(&struct {
				A string `yini:"same"`
				B string `yini:"same"`
			}{})
		}},
		{"tag options", func() error {
			return doc.Decode(&struct {
				A string `yini:"a,inline"`
			}{})
		}},
		{"unsupported absent field", func() error { return doc.Decode(&struct{ A chan int }{}) }},
		{"float32", func() error { return doc.Decode(&struct{ A float32 }{}) }},
		{"uintptr", func() error { return doc.Decode(&struct{ A uintptr }{}) }},
		{"array", func() error { return doc.Decode(&struct{ A [2]int64 }{}) }},
		{"non-empty interface", func() error { return doc.Decode(&struct{ A error }{}) }},
		{"non-string map key", func() error { return doc.Decode(&struct{ A map[int]string }{}) }},
	} {
		t.Run(test.name, func(t *testing.T) {
			var detail *yini.DecodeError
			if err := test.run(); !errors.As(err, &detail) || detail.Position != (yini.Position{}) {
				t.Fatalf("error = %v; want unpositioned *DecodeError", err)
			}
		})
	}

	type hidden struct{ Secret string }
	type Public struct{ Value string }
	type target struct {
		hidden
		Public
		Name string
	}
	var value target
	var detail *yini.DecodeError
	if err := doc.Decode(&value); !errors.As(err, &detail) || !strings.Contains(err.Error(), "embedded") {
		t.Fatalf("exported embedded field error = %v", err)
	}
	type ignored struct {
		hidden
		Public `yini:"-"`
		Name   string
	}
	doc = parseDecodeDocument(t, "Name = 'Ada'", yini.ModeLenient)
	var accepted ignored
	if err := doc.Decode(&accepted); err != nil || accepted.Name != "Ada" {
		t.Fatalf("ignored embedded fields = %#v, %v", accepted, err)
	}
	doc = parseDecodeDocument(t, "Secret = 'visible'", yini.ModeLenient)
	if err := doc.Decode(&accepted); !errors.As(err, &detail) || detail.Path != "Secret" {
		t.Fatalf("unexported embedded field was promoted: %v", err)
	}
}

func TestDecodeRecursiveTypeAndNumericBoundaries(t *testing.T) {
	type node struct {
		Value string `yini:"value"`
		Next  *node  `yini:"next"`
	}
	type config struct {
		Node node   `yini:"node"`
		Max  uint64 `yini:"max"`
		Min  int64  `yini:"min"`
	}
	doc := parseDecodeDocument(t, "^ Config\nnode = { value: 'root', next: { value: 'leaf', next: null } }\nmax = 9223372036854775807\nmin = -9223372036854775808\n/END", yini.ModeStrict)
	var dst config
	if err := doc.Sections[0].Decode(&dst); err != nil {
		t.Fatalf("Decode: %v", err)
	}
	if dst.Max != uint64(9223372036854775807) || dst.Min != int64(-9223372036854775808) || dst.Node.Next == nil || dst.Node.Next.Value != "leaf" || dst.Node.Next.Next != nil {
		t.Fatalf("decoded boundaries/recursive type = %#v", dst)
	}
	type badNode struct {
		Next *badNode
		Bad  chan int
	}
	var bad struct{ Node badNode }
	var detail *yini.DecodeError
	if err := doc.Decode(&bad); !errors.As(err, &detail) || detail.Position != (yini.Position{}) {
		t.Fatalf("recursive schema with unsupported field = %v", err)
	}
}
