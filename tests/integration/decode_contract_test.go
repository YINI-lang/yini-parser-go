package integration_test

import (
	"errors"
	"reflect"
	"strconv"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

type decodeCycleFirst *decodeCycleSecond
type decodeCycleSecond *decodeCycleFirst

func TestDecodeIntegerRanges(t *testing.T) {
	// 1. Assign.
	type signed int8
	type unsigned uint8
	type rangeTest struct {
		name     string
		zero     any
		accepted []string
		rejected []string
	}
	tests := []rangeTest{
		{"int8", int8(0), []string{"-128", "0", "127"}, []string{"-129", "128"}},
		{"int16", int16(0), []string{"-32768", "32767"}, []string{"-32769", "32768"}},
		{"int32", int32(0), []string{"-2147483648", "2147483647"}, []string{"-2147483649", "2147483648"}},
		{"int64", int64(0), []string{"-9223372036854775808", "9007199254740993", "9223372036854775807"}, nil},
		{"uint8", uint8(0), []string{"0", "255"}, []string{"-1", "256"}},
		{"uint16", uint16(0), []string{"0", "65535"}, []string{"-1", "65536"}},
		{"uint32", uint32(0), []string{"0", "4294967295"}, []string{"-1", "4294967296"}},
		{"uint64", uint64(0), []string{"0", "9007199254740993", "9223372036854775807"}, []string{"-1"}},
		{"named signed", signed(0), []string{"-128", "127"}, []string{"-129", "128"}},
		{"named unsigned", unsigned(0), []string{"0", "255"}, []string{"-1", "256"}},
	}
	if strconv.IntSize == 32 {
		tests = append(tests,
			rangeTest{"int", int(0), []string{"-2147483648", "2147483647"}, []string{"-2147483649", "2147483648"}},
			rangeTest{"uint", uint(0), []string{"0", "4294967295"}, []string{"-1", "4294967296"}},
		)
	} else {
		tests = append(tests,
			rangeTest{"int", int(0), []string{"-9223372036854775808", "9223372036854775807"}, nil},
			rangeTest{"uint", uint(0), []string{"0", "9223372036854775807"}, []string{"-1"}},
		)
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, test := range tests {
			for _, valid := range []bool{true, false} {
				literals := test.accepted
				if !valid {
					literals = test.rejected
				}
				for _, literal := range literals {
					t.Run(string(mode)+"/"+test.name+"/"+literal, func(t *testing.T) {
						// 1. Assign.
						doc := parseDecodeDocument(t, "^ Limits\nvalue = "+literal+"\n/END", mode)
						typ := reflect.StructOf([]reflect.StructField{{Name: "Value", Type: reflect.TypeOf(test.zero), Tag: `yini:"value"`}})
						dst := reflect.New(typ)
						field := dst.Elem().Field(0)
						field.Set(reflect.ValueOf(test.zero))
						var want int64
						if valid {
							var err error
							want, err = strconv.ParseInt(literal, 10, 64)
							if err != nil {
								t.Fatal(err)
							}
						}

						// 2. Act.
						err := doc.Sections[0].Decode(dst.Interface())

						// 3. Assert.
						if !valid {
							var detail *yini.DecodeError
							if !errors.As(err, &detail) || detail.Path != "Limits.value" || detail.Position != (yini.Position{Line: 2, Column: 1}) {
								t.Fatalf("overflow error = %v; want Limits.value at 2:1", err)
							}
							if !reflect.DeepEqual(field.Interface(), test.zero) {
								t.Fatalf("overflow changed destination to %v", field.Interface())
							}
							return
						}
						if err != nil {
							t.Fatal(err)
						}
						if field.Kind() >= reflect.Int && field.Kind() <= reflect.Int64 {
							if field.Int() != want {
								t.Fatalf("decoded = %v, want %d", field.Interface(), want)
							}
						} else if field.Uint() != uint64(want) {
							t.Fatalf("decoded = %v, want %d", field.Interface(), want)
						}
					})
				}
			}
		}
	}
}

func TestDecodeDoesNotCoerceValues(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name    string
		zero    any
		valid   string
		invalid []string
	}{
		{"string", "", "'ready'", []string{"true", "1", "1.0", "[]", "{}", "null"}},
		{"bool", false, "true", []string{"'true'", "1", "1.0", "null"}},
		{"integer", int64(0), "1", []string{"'1'", "true", "1.0", "null"}},
		{"float", float64(0), "3e4", []string{"'3e4'", "1", "true", "null"}},
		{"slice", []int64(nil), "[1, 2]", []string{"1", "{}", "[1.0]"}},
		{"map", map[string]int64(nil), "{ first: 1 }", []string{"1", "[]", "{ first: 1.0 }"}},
		{"struct", struct{ Count int64 }{}, "{ Count: 1 }", []string{"1", "[]", "null"}},
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, test := range tests {
			for _, literal := range append([]string{test.valid}, test.invalid...) {
				t.Run(string(mode)+"/"+test.name+"/"+literal, func(t *testing.T) {
					// 1. Assign.
					doc := parseDecodeDocument(t, "^ Types\nvalue = "+literal+"\n/END", mode)
					typ := reflect.StructOf([]reflect.StructField{{Name: "Value", Type: reflect.TypeOf(test.zero), Tag: `yini:"value"`}})
					dst := reflect.New(typ).Interface()

					// 2. Act.
					err := doc.Sections[0].Decode(dst)

					// 3. Assert.
					if literal == test.valid {
						if err != nil {
							t.Fatalf("valid value: %v", err)
						}
						return
					}
					var detail *yini.DecodeError
					if !errors.As(err, &detail) || detail.Path != "Types.value" && !strings.HasPrefix(detail.Path, "Types.value[") && !strings.HasPrefix(detail.Path, "Types.value.") {
						t.Fatalf("type mismatch = %v; want a DecodeError at Types.value", err)
					}
				})
			}
		}
	}
}

func TestDecodeFailurePreservesReferencesAndParseResult(t *testing.T) {
	// 1. Assign.
	type config struct {
		Pointer *int64           `yini:"pointer"`
		List    []int64          `yini:"list"`
		Map     map[string]int64 `yini:"map"`
		Last    int64            `yini:"last"`
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			// 1. Assign.
			result, err := yini.ParseString("^ State\npointer = 5\nlist = [6, 7]\nmap = { next: 8 }\nlast = 'bad'\n/END", yini.Options{Mode: mode})
			if err != nil {
				t.Fatal(err)
			}
			values := result.Document.Values()
			diagnostics := append([]yini.Diagnostic(nil), result.Diagnostics...)
			old := int64(42)
			dst := config{Pointer: &old, List: []int64{43}, Map: map[string]int64{"old": 44}, Last: 45}
			list := dst.List
			mapping := dst.Map
			var detail *yini.DecodeError

			// 2. Act.
			err = result.Document.Sections[0].Decode(&dst)

			// 3. Assert.
			if !errors.As(err, &detail) || detail.Path != "State.last" {
				t.Fatalf("Decode error = %v; want State.last", err)
			}
			if dst.Pointer != &old || old != 42 || &dst.List[0] != &list[0] || dst.List[0] != 43 ||
				!reflect.DeepEqual(dst.Map, map[string]int64{"old": 44}) || dst.Last != 45 {
				t.Fatalf("failed decode changed destination: %#v", dst)
			}
			t.Run("map reference preserved", func(t *testing.T) {
				// 1. Assign.
				want := int64(46)

				// 2. Act.
				mapping["shared"] = want

				// 3. Assert.
				if dst.Map["shared"] != want {
					t.Fatal("failed decode replaced the destination map")
				}
			})
			if !reflect.DeepEqual(result.Document.Values(), values) || len(result.Diagnostics) != len(diagnostics) {
				t.Fatal("failed decode changed the parse result")
			}
			for i := range diagnostics {
				if result.Diagnostics[i] != diagnostics[i] {
					t.Fatal("failed decode changed a parse diagnostic")
				}
			}
		})
	}
}

func TestDecodePointerCyclesAreInvalidSchemas(t *testing.T) {
	// 1. Assign.
	type loop *loop
	for _, dst := range []any{
		&struct{ Value loop }{},
		&struct{ Value decodeCycleFirst }{},
		&struct{ Value []loop }{},
		&struct{ Value map[string]loop }{},
	} {
		for _, doc := range []*yini.Document{
			{},
			{Members: []*yini.Member{{Name: "Value", Value: nil}}},
			{Members: []*yini.Member{{Name: "Value", Value: int64(1)}}},
			{Sections: []*yini.Section{{Name: "Value"}}},
		} {
			// 1. Assign.
			var detail *yini.DecodeError

			// 2. Act.
			err := doc.Decode(dst)

			// 3. Assert.
			if !errors.As(err, &detail) || detail.Position != (yini.Position{}) || !strings.Contains(detail.Message, "pointer cycle") {
				t.Fatalf("Decode(%T) = %v; want an unpositioned pointer cycle error", dst, err)
			}
		}
	}
}

func TestDecodeRecursiveCollectionsAndPointers(t *testing.T) {
	// 1. Assign.
	type tree map[string]tree
	type list []list
	type number int64
	type pointer *number
	type enabled bool
	type config struct {
		Tree    tree      `yini:"tree"`
		List    list      `yini:"list"`
		Number  **pointer `yini:"number"`
		Enabled enabled   `yini:"enabled"`
		Items   []*struct {
			Name string `yini:"name"`
		} `yini:"items"`
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			// 1. Assign.
			doc := parseDecodeDocument(t, "^ Recursive\ntree = { branch: { leaf: null } }\nlist = [[[], null]]\nnumber = 17\nenabled = true\nitems = [{ name: 'birch' }, null]\n/END", mode)
			var dst config

			// 2. Act.
			err := doc.Sections[0].Decode(&dst)

			// 3. Assert.
			if err != nil {
				t.Fatal(err)
			}
			if !reflect.DeepEqual(dst.Tree, tree{"branch": {"leaf": nil}}) || !reflect.DeepEqual(dst.List, list{list{list{}, nil}}) ||
				dst.Number == nil || *dst.Number == nil || **dst.Number == nil || ***dst.Number != 17 || !bool(dst.Enabled) ||
				len(dst.Items) != 2 || dst.Items[0] == nil || dst.Items[0].Name != "birch" || dst.Items[1] != nil {
				t.Fatalf("decoded recursive values = %#v", dst)
			}
		})
	}
}

func TestDecodeSectionInterfaceCopyAndRootShape(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			// 1. Assign.
			doc := parseDecodeDocument(t, "^ Service\nport = 8080\n^^ Extra\nitems = [{ labels: ['original'] }]\n/END", mode)
			t.Run("root shape", func(t *testing.T) {
				// 1. Assign.
				var flat struct {
					Port int64 `yini:"port"`
				}
				var detail *yini.DecodeError

				// 2. Act.
				err := doc.Decode(&flat)

				// 3. Assert.
				if !errors.As(err, &detail) || detail.Path != "Service" || detail.Position != (yini.Position{Line: 1, Column: 1}) {
					t.Fatalf("root was unwrapped: %v", err)
				}
			})
			t.Run("section interface", func(t *testing.T) {
				// 1. Assign.
				var dst struct{ Service *any }
				want := doc.Values()["Service"]

				// 2. Act.
				err := doc.Decode(&dst)

				// 3. Assert.
				if err != nil {
					t.Fatal(err)
				}
				if dst.Service == nil || !reflect.DeepEqual(*dst.Service, want) {
					t.Fatalf("section interface = %#v, want %#v", dst.Service, want)
				}
				t.Run("independent copy", func(t *testing.T) {
					// 1. Assign.
					copied := (*dst.Service).(map[string]any)

					// 2. Act.
					copied["Extra"].(map[string]any)["items"].([]any)[0].(map[string]any)["labels"].([]any)[0] = "changed"

					// 3. Assert.
					if !reflect.DeepEqual(doc.Values()["Service"], want) {
						t.Fatal("section interface shares storage with source")
					}
				})
			})
		})
	}
}

func TestDecodeOriginalSourceLocations(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, ending := range []string{"\n", "\r\n", "\r"} {
			t.Run(string(mode)+"/"+strconv.Quote(ending), func(t *testing.T) {
				// 1. Assign.
				text := "\ufeff#!/usr/bin/env yini" + ending + "^ `ui.theme`" + ending + "  ports = [80, 'bad']" + ending + "/END"
				doc := parseDecodeDocument(t, text, mode)
				var dst struct {
					Ports []int64 `yini:"ports"`
				}
				var detail *yini.DecodeError

				// 2. Act.
				err := doc.Sections[0].Decode(&dst)

				// 3. Assert.
				if !errors.As(err, &detail) || detail.Path != `["ui.theme"].ports[1]` || detail.Position != (yini.Position{Line: 3, Column: 3}) {
					t.Fatalf("original source error = %v; want [\"ui.theme\"].ports[1] at 3:3", err)
				}
			})
		}
	}
}
