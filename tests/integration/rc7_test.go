package integration_test

import (
	"reflect"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestRC7IdentifiersAndSectionHeaders(t *testing.T) {
	// 1. Assign.
	content := "^ /* heading */ _Depot /* tail */\r\n# comment\r\n" +
		"TRUE = 1\r\nnull = 2\r\n`path\\\\part` = 3\r\n" +
		"^^`bay\\`name`\r\nvalue = {False: 4, NULL: 5, `a\\`b`: 6}\r\n/END"
	want := map[string]any{
		"_Depot": map[string]any{
			"TRUE": int64(1), "null": int64(2), `path\part`: int64(3),
			"bay`name": map[string]any{
				"value": map[string]any{"False": int64(4), "NULL": int64(5), "a`b": int64(6)},
			},
		},
	}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			t.Run("valid identifiers", func(t *testing.T) {
				// 1. Assign.
				options := yini.Options{Mode: mode}

				// 2. Act.
				result, err := yini.ParseString(content, options)

				// 3. Assert.
				if err != nil || !reflect.DeepEqual(result.Document.Values(), want) {
					t.Fatalf("ParseString = %#v, %v (%#v); want %#v", result.Document.Values(), err, result.Diagnostics, want)
				}
			})
			t.Run("invalid identifier escape", func(t *testing.T) {
				// 1. Assign.
				content := "^ Depot\n`bad\\q` = 1\n/END"

				// 2. Act.
				invalid, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || !invalid.HasErrors() {
					t.Fatalf("unsupported identifier escape accepted: %#v", invalid)
				}
			})
		})
	}
}

func TestRC7SignedNumbers(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			t.Run("valid numbers", func(t *testing.T) {
				// 1. Assign.
				content := "^ Scale\nminimum = -9223372036854775808\nvalues = [+0x2A, -0b10, +.25, -3e2]\n/END"
				want := map[string]any{"Scale": map[string]any{
					"minimum": int64(-9223372036854775808),
					"values":  []any{int64(42), int64(-2), 0.25, -300.0},
				}}

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err != nil || !reflect.DeepEqual(result.Document.Values(), want) {
					t.Fatalf("signed values = %#v, %v (%#v)", result.Document.Values(), err, result.Diagnostics)
				}
			})
			for _, literal := range []string{"- 7", "+\t7", "-/* gap */7", "-9223372036854775809"} {
				t.Run(literal, func(t *testing.T) {
					// 1. Assign.
					content := "^ Scale\nvalue = " + literal + "\n/END"

					// 2. Act.
					invalid, err := yini.ParseString(content, yini.Options{Mode: mode})

					// 3. Assert.
					if err == nil || !hasDiagnostic(invalid, "number") {
						t.Fatalf("invalid number %q accepted: %v (%#v)", literal, err, invalid.Diagnostics)
					}
					for _, diagnostic := range invalid.Diagnostics {
						if diagnostic.Code == "number" && diagnostic.Position != (yini.Position{Line: 2, Column: 9}) {
							t.Fatalf("number %q position = %#v", literal, diagnostic.Position)
						}
					}
				})
			}
		})
	}
}

func TestRC7NumericConcatenation(t *testing.T) {
	// 1. Assign.
	content := "^ Tags\nlabel = 'n=' + +1_0.50 + ',' + -0 + ',' + 3E+04 + ',' + -0xA\n/END"
	t.Run("lenient", func(t *testing.T) {
		// 1. Assign.
		options := yini.Options{}

		// 2. Act.
		result, err := yini.ParseString(content, options)

		// 3. Assert.
		if err != nil || result.Document.Values()["Tags"].(map[string]any)["label"] != "n=10.50,-0,3e04,-10" {
			t.Fatalf("numeric concatenation = %#v, %v (%#v)", result.Document.Values(), err, result.Diagnostics)
		}
	})
	t.Run("strict", func(t *testing.T) {
		// 1. Assign.
		options := yini.Options{Mode: yini.ModeStrict}

		// 2. Act.
		strict, err := yini.ParseString(content, options)

		// 3. Assert.
		if err == nil || !hasDiagnostic(strict, "concat-strict") {
			t.Fatalf("strict numeric concatenation accepted: %v (%#v)", err, strict.Diagnostics)
		}
	})
}

func TestRC7BinaryPrefixWarning(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			for _, literal := range []string{"%101", "-%101", "0b101"} {
				t.Run(literal, func(t *testing.T) {
					// 1. Assign.
					content := "^ Flags\nvalue = " + literal + "\n/END"
					wantWarning := strings.Contains(literal, "%")

					// 2. Act.
					result, err := yini.ParseString(content, yini.Options{Mode: mode})

					// 3. Assert.
					if err != nil {
						t.Fatal(err)
					}
					if hasDiagnostic(result, "deprecated-binary-prefix") != wantWarning {
						t.Fatalf("%q diagnostics = %#v", literal, result.Diagnostics)
					}
					if wantWarning {
						column := 9 + strings.Index(literal, "%")
						diagnostic := result.Diagnostics[0]
						if diagnostic.Severity != yini.SeverityWarning || diagnostic.Position != (yini.Position{Line: 2, Column: column}) {
							t.Fatalf("%q warning = %#v", literal, diagnostic)
						}
					}
				})
			}
		})
	}
	t.Run("lenient concatenation", func(t *testing.T) {
		// 1. Assign.
		content := "^ Flags\nlabel = 'bits=' + %101\n"

		// 2. Act.
		result, err := yini.ParseString(content, yini.Options{})

		// 3. Assert.
		if err != nil || !hasDiagnostic(result, "deprecated-binary-prefix") || result.Document.Values()["Flags"].(map[string]any)["label"] != "bits=5" {
			t.Fatalf("binary concatenation = %#v, %v (%#v)", result.Document.Values(), err, result.Diagnostics)
		}
	})
}

func TestRC7DuplicateSections(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			t.Run("duplicate siblings", func(t *testing.T) {
				// 1. Assign.
				content := "^ Depot\n^^ Bay\nfirst = 1\n^^ Bay\nsecond = 2\n^^^ Ignored\nvalue = 3\n^^ Store\nvalue = 4\n/END"
				want := map[string]any{"Depot": map[string]any{
					"Bay": map[string]any{"first": int64(1)}, "Store": map[string]any{"value": int64(4)},
				}}

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || len(result.Diagnostics) != 1 {
					t.Fatalf("duplicate section = %v (%#v)", err, result.Diagnostics)
				}
				diagnostic := result.Diagnostics[0]
				if diagnostic.Code != "duplicate-section" || diagnostic.Severity != yini.SeverityError || diagnostic.Position != (yini.Position{Line: 4, Column: 1}) {
					t.Fatalf("duplicate diagnostic = %#v", diagnostic)
				}
				if !reflect.DeepEqual(result.Document.Values(), want) {
					t.Fatalf("recovered values = %#v; want %#v", result.Document.Values(), want)
				}
			})
			t.Run("different parents", func(t *testing.T) {
				// 1. Assign.
				valid := "^ Depot\n^^ East\n^^^ Bay\nvalue = 1\n^^ West\n^^^ Bay\nvalue = 2\n/END"

				// 2. Act.
				result, err := yini.ParseString(valid, yini.Options{Mode: mode})

				// 3. Assert.
				if err != nil {
					t.Fatalf("different parent sections = %v (%#v)", err, result.Diagnostics)
				}
			})
		})
	}
}

func TestRC7HexObjectKeys(t *testing.T) {
	// 1. Assign.
	content := "^ Palette\nvalue = {hex:10, HEX:10.5, Hex:0x10, nested: [{hex:1e2}, {hex:-2}], bases: [hex:10]}\nnumber = hex:10\n/END"
	want := map[string]any{"Palette": map[string]any{
		"value": map[string]any{"hex": int64(10), "HEX": 10.5, "Hex": int64(16),
			"nested": []any{map[string]any{"hex": 100.0}, map[string]any{"hex": int64(-2)}},
			"bases":  []any{int64(16)},
		},
		"number": int64(16),
	}}
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode), func(t *testing.T) {
			t.Run("valid keys", func(t *testing.T) {
				// 1. Assign.
				options := yini.Options{Mode: mode}

				// 2. Act.
				result, err := yini.ParseString(content, options)

				// 3. Assert.
				if err != nil || !reflect.DeepEqual(result.Document.Values(), want) {
					t.Fatalf("hex keys = %#v, %v (%#v); want %#v", result.Document.Values(), err, result.Diagnostics, want)
				}
			})
			t.Run("invalid compact value", func(t *testing.T) {
				// 1. Assign.
				content := "^ Palette\nvalue = {hex:1_}\n/END"

				// 2. Act.
				invalid, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || !invalid.HasErrors() {
					t.Fatalf("invalid compact value accepted: %#v", invalid)
				}
			})
			t.Run("incomplete compact object", func(t *testing.T) {
				// 1. Assign.
				content := "^ Palette\nvalue = {hex:10"

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || !result.HasErrors() {
					t.Fatalf("incomplete compact object accepted: %#v", result)
				}
			})
			t.Run("duplicate compact key", func(t *testing.T) {
				// 1. Assign.
				content := "\uFEFF^ Palette\nvalue = {hex:10, hex:20}\n/END"

				// 2. Act.
				invalid, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if (err != nil) != (mode == yini.ModeStrict) || len(invalid.Diagnostics) != 1 || invalid.Diagnostics[0].Position != (yini.Position{Line: 2, Column: 18}) {
					t.Fatalf("compact key duplicate = %v (%#v)", err, invalid.Diagnostics)
				}
			})
			for _, ending := range []string{"\r\n", "\r"} {
				t.Run("Unicode/"+strings.ReplaceAll(ending, "\r", "CR"), func(t *testing.T) {
					// 1. Assign.
					content := "\uFEFF^ Palette" + ending + "value = {label:'é', hex:10, hex:20}" + ending + "/END"

					// 2. Act.
					result, err := yini.ParseString(content, yini.Options{Mode: mode})

					// 3. Assert.
					if (err != nil) != (mode == yini.ModeStrict) || len(result.Diagnostics) != 1 || result.Diagnostics[0].Position != (yini.Position{Line: 2, Column: 29}) {
						t.Fatalf("Unicode compact key duplicate = %v (%#v)", err, result.Diagnostics)
					}
				})
			}
		})
	}
}

func TestRC7CommandStringDiagnosticLocations(t *testing.T) {
	// 1. Assign.
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, test := range []struct {
			name     string
			lines    []string
			position yini.Position
		}{
			{"insufficient indentation", []string{"    first", "  short"}, yini.Position{Line: 4, Column: 3}},
			{"leading blank indentation", []string{"     ", "    first"}, yini.Position{Line: 3, Column: 5}},
		} {
			t.Run(string(mode)+"/"+test.name, func(t *testing.T) {
				// 1. Assign.
				content := "\uFEFF^ Notes\r\nvalue = command\"\"\"\r\n" + strings.Join(test.lines, "\r\n") + "\r\n\"\"\"\r\n/END"

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "string" || result.Diagnostics[0].Position != test.position {
					t.Fatalf("command string error = %v (%#v); want %#v", err, result.Diagnostics, test.position)
				}
			})
		}
	}
}

func TestRC7StringsAndLineEndings(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		for _, ending := range []string{"\n", "\r\n", "\r"} {
			t.Run(string(mode)+"/"+strings.ReplaceAll(ending, "\r", "CR"), func(t *testing.T) {
				// 1. Assign.
				content := "^ Notes\nraw = R\"\"\"one\ntwo\"\"\"\nescaped = E\"\"\"one\n\\r\\n\\`two\"\"\"\nscript = command\"\"\"\n  first\n    second\n  \n\"\"\"\n/END"
				content = strings.ReplaceAll(content, "\n", ending)
				want := map[string]any{"Notes": map[string]any{
					"raw": "one\ntwo", "escaped": "one\n\r\n`two", "script": "first\n  second\n",
				}}

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err != nil || !reflect.DeepEqual(result.Document.Values(), want) {
					t.Fatalf("strings = %#v, %v (%#v)", result.Document.Values(), err, result.Diagnostics)
				}
			})
		}
		for _, literal := range []string{
			"command\"\"\"\n  first\n second\n\"\"\"",
			"command\"\"\"\n   \n  first\n\"\"\"",
			"command\"\"\"\n  first\n\"\"\" trailing",
			"E'bad\\z'",
		} {
			t.Run(string(mode)+"/invalid/"+literal, func(t *testing.T) {
				// 1. Assign.
				content := "^ Notes\nvalue = " + literal + "\n/END"

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || !result.HasErrors() {
					t.Fatalf("invalid string %q accepted: %#v", literal, result)
				}
			})
		}
		t.Run(string(mode)+"/mixed endings", func(t *testing.T) {
			// 1. Assign.
			content := "^ Notes\r\nvalue = 1\n/END"

			// 2. Act.
			mixed, err := yini.ParseString(content, yini.Options{Mode: mode})

			// 3. Assert.
			if (err != nil) != (mode == yini.ModeStrict) || len(mixed.Diagnostics) != 1 || mixed.Diagnostics[0].Code != "mixed-line-endings" || mixed.Diagnostics[0].Position != (yini.Position{Line: 2, Column: 10}) {
				t.Fatalf("mixed endings = %v (%#v)", err, mixed.Diagnostics)
			}
		})
	}
}

func TestRC7MultilineTriviaPlacement(t *testing.T) {
	for _, mode := range []yini.Mode{yini.ModeLenient, yini.ModeStrict} {
		t.Run(string(mode)+"/valid comments", func(t *testing.T) {
			// 1. Assign.
			content := "^ Catalog\nitems = [\n; list note\n-- skipped\n1,\n2\n]\n/END"

			// 2. Act.
			result, err := yini.ParseString(content, yini.Options{Mode: mode})

			// 3. Assert.
			if err != nil {
				t.Fatalf("valid list comments = %v (%#v)", err, result.Diagnostics)
			}
		})
		for _, invalid := range []string{
			"items = [1 ; inline\n]",
			"items = [1 -- inline\n]",
			"items = [1\n,2]",
			"items = /* line\nbreak */ [1]",
		} {
			t.Run(string(mode)+"/invalid/"+invalid, func(t *testing.T) {
				// 1. Assign.
				content := "^ Catalog\n" + invalid + "\n/END"

				// 2. Act.
				result, err := yini.ParseString(content, yini.Options{Mode: mode})

				// 3. Assert.
				if err == nil || !result.HasErrors() {
					t.Fatalf("invalid placement %q accepted: %#v", invalid, result)
				}
			})
		}
	}
}
