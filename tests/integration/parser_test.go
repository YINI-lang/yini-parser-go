package integration_test

import (
	"errors"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	yini "github.com/YINI-lang/yini-parser-go"
)

func TestParseFilesAndValues(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name        string
		file        string
		mode        yini.Mode
		wantError   bool
		want        map[string]any
		diagnostics []yini.Diagnostic
	}{
		{
			name: "lenient workshop",
			file: filepath.Join("lenient", "workshop.yini"),
			want: map[string]any{
				"label":      "north wing",
				"retryCount": int64(3),
				"Workshop": map[string]any{
					"active":     true,
					"tags":       []any{"wood", "steel"},
					"details":    map[string]any{"lead": "Mira", "floor": int64(2)},
					"Calibrator": map[string]any{"margin": 0.125},
				},
			},
		},
		{
			name: "lenient harbor",
			file: filepath.Join("lenient", "harbor.yini"),
			want: map[string]any{
				"region": "coast",
				"Harbor": map[string]any{
					"open":   true,
					"berths": []any{int64(3), int64(5)},
					"Dock":   map[string]any{"depth": 7.25},
				},
			},
		},
		{
			name:      "lenient dispatch",
			file:      filepath.Join("lenient", "dispatch.yini"),
			wantError: true,
			want: map[string]any{
				"Dispatch": map[string]any{
					"priority": "routine",
					"Courier": map[string]any{
						"name":  "Bramble",
						"zones": []any{"north", "west"},
					},
					"Backup": map[string]any{"enabled": true},
				},
			},
			diagnostics: []yini.Diagnostic{
				{Code: "duplicate-key", Severity: yini.SeverityWarning, Position: yini.Position{Line: 4, Column: 1}},
				{Code: "duplicate-section", Severity: yini.SeverityError, Position: yini.Position{Line: 10, Column: 1}},
			},
		},
		{
			name: "lenient field station",
			file: filepath.Join("lenient", "field-station.yini"),
			want: map[string]any{
				"station":     "ridge-17",
				"fallback":    nil,
				"pending":     []any{},
				"annotations": map[string]any{},
				"Weather": map[string]any{
					"enabled":     true,
					"samples":     []any{},
					"calibration": nil,
				},
				"Radio": map[string]any{
					"channels":    []any{int64(7), int64(11)},
					"credentials": nil,
				},
			},
		},
		{
			name: "strict observatory",
			file: filepath.Join("strict", "observatory.strict.yini"),
			mode: yini.ModeStrict,
			want: map[string]any{
				"Observatory": map[string]any{
					"title":       "Night survey",
					"count":       int64(17),
					"ready":       true,
					"coordinates": []any{12.5, -4.25},
					"Instrument":  map[string]any{"model": "Orion"},
				},
			},
		},
		{
			name: "strict archive",
			file: filepath.Join("strict", "archive.strict.yini"),
			mode: yini.ModeStrict,
			want: map[string]any{
				"Archive": map[string]any{
					"label":    "East vault",
					"sealed":   false,
					"labels":   []any{"paper", "film"},
					"metadata": map[string]any{"curator": "Ivo", "aisle": int64(4)},
					"Shelf":    map[string]any{"slots": int64(64)},
				},
			},
		},
		{
			name: "strict job runner",
			file: filepath.Join("strict", "job-runner.strict.yini"),
			mode: yini.ModeStrict,
			want: map[string]any{
				"JobRunner": map[string]any{
					"name":    "nightly-index",
					"workers": int64(6),
					"queues":  []any{"scan", "summarize"},
					"jobs": []any{
						map[string]any{"name": "index", "attempts": int64(3), "labels": []any{"local", "fast"}},
						map[string]any{"name": "summarize", "attempts": int64(1), "labels": []any{}},
					},
					"fallback": nil,
					"maxDelay": 12.5,
					"metadata": map[string]any{"owner": "operations", "region": "east"},
					"Pool": map[string]any{
						"enabled": true,
						"limits":  map[string]any{"pending": int64(64), "running": int64(6)},
						"Retry":   map[string]any{"delays": []any{int64(1), int64(3), int64(8)}, "jitter": 0.125},
					},
				},
			},
		},
		{
			name: "strict telemetry",
			file: filepath.Join("strict", "telemetry.strict.yini"),
			mode: yini.ModeStrict,
			want: map[string]any{
				"sensor.grid": map[string]any{
					"serial":        int64(9007199254740993),
					"minimum":       int64(math.MinInt64),
					"maximum":       int64(math.MaxInt64),
					"mask":          int64(195),
					"permissions":   int64(416),
					"batch":         int64(22),
					"signature":     int64(51966),
					"frequency":     float64(2500),
					"gain":          0.0625,
					"path":          `C:\Stations\Ridge`,
					"message":       "channel\tready\n\u03A9",
					"display.label": "Grid #4",
					"window A":      map[string]any{"samples": []any{0.0, -1.25, float64(200)}},
				},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			path := filepath.Join("testdata", test.file)
			options := yini.Options{Mode: test.mode}

			// 2. Act.
			result, err := yini.ParseFile(path, options)

			// 3. Assert.
			if (err != nil) != test.wantError {
				t.Fatalf("ParseFile error = %v; want error %v (diagnostics: %#v)", err, test.wantError, result.Diagnostics)
			}
			if len(result.Diagnostics) != len(test.diagnostics) {
				t.Fatalf("diagnostics = %#v, want %#v", result.Diagnostics, test.diagnostics)
			}
			for i, want := range test.diagnostics {
				got := result.Diagnostics[i]
				if got.Code != want.Code || got.Severity != want.Severity || got.Position != want.Position {
					t.Fatalf("diagnostic %d = %#v, want %#v", i, got, want)
				}
			}
			if got := result.Document.Values(); !reflect.DeepEqual(got, test.want) {
				t.Fatalf("Values() = %#v, want %#v", got, test.want)
			}
			t.Run("values copy", func(t *testing.T) {
				// 1. Assign.
				values := result.Document.Values()

				// 2. Act.
				values["changed"] = true

				// 3. Assert.
				if _, exists := result.Document.Values()["changed"]; exists {
					t.Fatal("Values() returned document-owned map")
				}
			})
			t.Run("nested values copy", func(t *testing.T) {
				// 1. Assign.
				values := result.Document.Values()

				// 2. Act.
				for _, section := range values {
					if nested, ok := section.(map[string]any); ok {
						nested["changed"] = true
					}
				}

				// 3. Assert.
				if !reflect.DeepEqual(result.Document.Values(), test.want) {
					t.Fatal("Values() returned document-owned nested map")
				}
			})
		})
	}
}

func TestLenientFixturesRejectStrictMode(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		file        string
		diagnostics []yini.Diagnostic
	}{
		{
			file: "dispatch.yini",
			diagnostics: []yini.Diagnostic{
				{Code: "duplicate-key", Severity: yini.SeverityError, Position: yini.Position{Line: 4, Column: 1}},
				{Code: "duplicate-section", Severity: yini.SeverityError, Position: yini.Position{Line: 10, Column: 1}},
			},
		},
		{
			file: "field-station.yini",
			diagnostics: []yini.Diagnostic{
				{Code: "orphan-member", Severity: yini.SeverityError, Position: yini.Position{Line: 1, Column: 1}},
				{Code: "orphan-member", Severity: yini.SeverityError, Position: yini.Position{Line: 2, Column: 1}},
				{Code: "orphan-member", Severity: yini.SeverityError, Position: yini.Position{Line: 3, Column: 1}},
				{Code: "orphan-member", Severity: yini.SeverityError, Position: yini.Position{Line: 4, Column: 1}},
				{Code: "top-level-section", Severity: yini.SeverityError, Position: yini.Position{Line: 11, Column: 1}},
			},
		},
	}
	for _, test := range tests {
		t.Run(test.file, func(t *testing.T) {
			// 1. Assign.
			path := filepath.Join("testdata", "lenient", test.file)
			options := yini.Options{Mode: yini.ModeStrict}
			var parseError *yini.ParseError

			// 2. Act.
			result, err := yini.ParseFile(path, options)

			// 3. Assert.
			if !errors.As(err, &parseError) || !result.HasErrors() {
				t.Fatalf("strict ParseFile = %v (%#v); want a ParseError", err, result.Diagnostics)
			}
			if len(result.Diagnostics) != len(test.diagnostics) {
				t.Fatalf("diagnostics = %#v, want %#v", result.Diagnostics, test.diagnostics)
			}
			for i, want := range test.diagnostics {
				got := result.Diagnostics[i]
				if got.Code != want.Code || got.Severity != want.Severity || got.Position != want.Position {
					t.Fatalf("diagnostic %d = %#v, want %#v", i, got, want)
				}
			}
		})
	}
}

func TestModesAndDiagnostics(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name     string
		text     string
		mode     yini.Mode
		code     string
		severity yini.Severity
	}{
		{"strict terminator", "^ Ship\nname = 'Dawn'", yini.ModeStrict, "missing-end", yini.SeverityError},
		{"strict orphan", "name = 'Dawn'\n^ Ship\n/END", yini.ModeStrict, "orphan-member", yini.SeverityError},
		{"strict extra root", "^ Ship\n^ Dock\n/END", yini.ModeStrict, "top-level-section", yini.SeverityError},
		{"declared strict", "@yini strict\n^ Ship\n", yini.ModeLenient, "mode-mismatch", yini.SeverityError},
		{"declared lenient", "@yini lenient\n^ Ship\n/END", yini.ModeStrict, "mode-mismatch", yini.SeverityWarning},
		{"strict trailing comma", "^ Ship\nports = [1, 2,]\n/END", yini.ModeStrict, "trailing-comma", yini.SeverityError},
		{"strict object separator", "^ Ship\nport = { number = 7 }\n/END", yini.ModeStrict, "object-separator", yini.SeverityError},
		{"strict empty value", "^ Ship\nname =\n/END", yini.ModeStrict, "empty-value", yini.SeverityError},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			options := yini.Options{Mode: test.mode}

			// 2. Act.
			result, err := yini.ParseString(test.text, options)

			// 3. Assert.
			if test.severity == yini.SeverityError && err == nil {
				t.Fatalf("expected parse error, got diagnostics %#v", result.Diagnostics)
			}
			if test.severity == yini.SeverityWarning && err != nil {
				t.Fatalf("unexpected parse error: %v", err)
			}
			for _, diagnostic := range result.Diagnostics {
				if diagnostic.Code == test.code && diagnostic.Severity == test.severity {
					return
				}
			}
			t.Fatalf("missing %s diagnostic %q: %#v", test.severity, test.code, result.Diagnostics)
		})
	}
}

func TestDuplicateRecoveryAndLocations(t *testing.T) {
	// 1. Assign.
	text := "\uFEFF#!/usr/bin/env yini\n^ Fleet\nname = 'First'\nname = 'Second'\n^^ Vessel\nid = 1\n^^ Vessel\nid = 2\n"
	want := map[string]any{"Fleet": map[string]any{"name": "First", "Vessel": map[string]any{"id": int64(1)}}}

	// 2. Act.
	result, err := yini.ParseString(text, yini.Options{})

	// 3. Assert.
	if err == nil || !result.HasErrors() {
		t.Fatalf("duplicate section must fail: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	if got := result.Document.Values(); !reflect.DeepEqual(got, want) {
		t.Fatalf("Values() = %#v, want %#v", got, want)
	}
	for _, check := range []struct {
		code string
		line int
	}{
		{"duplicate-key", 4}, {"duplicate-section", 7},
	} {
		found := false
		for _, diagnostic := range result.Diagnostics {
			if diagnostic.Code == check.code && diagnostic.Position.Line == check.line && diagnostic.Position.Column == 1 {
				found = true
			}
		}
		if !found {
			t.Fatalf("missing %s at %d:1: %#v", check.code, check.line, result.Diagnostics)
		}
	}
}

func TestNumbersAndStringConcatenation(t *testing.T) {
	t.Run("valid values", func(t *testing.T) {
		// 1. Assign.
		text := "^ Bench\nmin = -9223372036854775808\nmax = 9223372036854775807\nhex = hex:FF\nduo = 0zX\nbin = %101\nexp = 3e4\nnegativeZero = -0.0\nlabel = 'count=' + 0x10\n"

		// 2. Act.
		result, err := yini.ParseString(text, yini.Options{})

		// 3. Assert.
		if err != nil {
			t.Fatalf("ParseString: %v (diagnostics: %#v)", err, result.Diagnostics)
		}
		values := result.Document.Values()["Bench"].(map[string]any)
		for key, want := range map[string]any{
			"min": int64(math.MinInt64), "max": int64(math.MaxInt64),
			"hex": int64(255), "duo": int64(10), "bin": int64(5),
			"exp": float64(30000), "label": "count=16",
		} {
			if values[key] != want {
				t.Errorf("%s = %#v, want %#v", key, values[key], want)
			}
		}
		if !math.Signbit(values["negativeZero"].(float64)) {
			t.Error("negative floating-point zero lost its sign")
		}
	})
	for _, literal := range []string{"9223372036854775808", "-9223372036854775809", "1e400", "1e-400"} {
		t.Run(literal, func(t *testing.T) {
			// 1. Assign.
			text := "^ Bench\nvalue = " + literal + "\n"

			// 2. Act.
			result, err := yini.ParseString(text, yini.Options{})

			// 3. Assert.
			if err == nil || !result.HasErrors() {
				t.Errorf("%s should fail: %#v", literal, result.Diagnostics)
			}
		})
	}
}

func TestEOFAndTooling(t *testing.T) {
	t.Run("final assignment at EOF", func(t *testing.T) {
		// 1. Assign.
		text := "^ Last\nvalue = C'line\\nend'"

		// 2. Act.
		result, err := yini.ParseString(text, yini.Options{})

		// 3. Assert.
		if err != nil {
			t.Fatalf("final assignment at EOF: %v", err)
		}
		if got := result.Document.Values()["Last"].(map[string]any)["value"]; got != "line\nend" {
			t.Fatalf("decoded string = %#v", got)
		}
	})
	t.Run("terminator at EOF", func(t *testing.T) {
		// 1. Assign.
		text := "^ Last\nvalue = 1\n/END"
		options := yini.Options{Mode: yini.ModeStrict}

		// 2. Act.
		_, err := yini.ParseString(text, options)

		// 3. Assert.
		if err != nil {
			t.Fatalf("/END at EOF: %v", err)
		}
	})
	t.Run("tooling diagnostics", func(t *testing.T) {
		// 1. Assign.
		text := "^ Last\nvalue = [1,,2]"

		// 2. Act.
		result, err := yini.ParseForTooling(text, yini.Options{})

		// 3. Assert.
		if err != nil || !result.HasErrors() {
			t.Fatalf("tooling parse = %#v, %v", result.Diagnostics, err)
		}
	})
	t.Run("positioned parse error", func(t *testing.T) {
		// 1. Assign.
		text := "^ Last\nvalue = [1,,2]"
		var parseErr *yini.ParseError

		// 2. Act.
		_, err := yini.ParseString(text, yini.Options{})

		// 3. Assert.
		if !errors.As(err, &parseErr) || !strings.Contains(parseErr.Error(), "line") {
			t.Fatalf("expected positioned ParseError, got %v", err)
		}
	})
}

func TestStrictFilenameDoesNotSelectMode(t *testing.T) {
	// 1. Assign.
	path := filepath.Join(t.TempDir(), "schedule.strict.yini")
	if err := os.WriteFile(path, []byte("orphan = 1"), 0600); err != nil {
		t.Fatal(err)
	}
	t.Run("lenient", func(t *testing.T) {
		// 1. Assign.
		options := yini.Options{}

		// 2. Act.
		result, err := yini.ParseFile(path, options)

		// 3. Assert.
		if err != nil || !hasDiagnostic(result, "filename-mode") {
			t.Fatalf("lenient file: %v (%#v)", err, result.Diagnostics)
		}
		if result.Document.Values()["orphan"] != int64(1) {
			t.Fatalf("filename changed parsing mode: %#v", result.Document.Values())
		}
	})
	t.Run("strict", func(t *testing.T) {
		// 1. Assign.
		options := yini.Options{Mode: yini.ModeStrict}

		// 2. Act.
		strict, err := yini.ParseFile(path, options)

		// 3. Assert.
		if err == nil || !hasDiagnostic(strict, "orphan-member") || hasDiagnostic(strict, "filename-mode") {
			t.Fatalf("strict file: %v (%#v)", err, strict.Diagnostics)
		}
	})
}

func TestReaderAndInvalidUTF8(t *testing.T) {
	t.Run("reader", func(t *testing.T) {
		// 1. Assign.
		reader := strings.NewReader("^ Passage\nopen = yes")

		// 2. Act.
		result, err := yini.ParseReader(reader, yini.Options{})

		// 3. Assert.
		if err != nil || result.Document.Values()["Passage"].(map[string]any)["open"] != true {
			t.Fatalf("ParseReader: %v (%#v)", err, result.Diagnostics)
		}
	})
	t.Run("invalid UTF-8", func(t *testing.T) {
		// 1. Assign.
		text := "^ Passage\rentry = \xff"

		// 2. Act.
		result, err := yini.ParseString(text, yini.Options{})

		// 3. Assert.
		if err == nil || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "invalid-utf8" || result.Diagnostics[0].Position.Line != 2 {
			t.Fatalf("invalid UTF-8: %v (%#v)", err, result.Diagnostics)
		}
	})
}
