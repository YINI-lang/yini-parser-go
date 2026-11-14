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
	tests := []struct {
		name        string
		file        string
		mode        yini.Mode
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
			name: "lenient dispatch",
			file: filepath.Join("lenient", "dispatch.yini"),
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
				{Code: "duplicate-section", Severity: yini.SeverityWarning, Position: yini.Position{Line: 10, Column: 1}},
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
			result, err := yini.ParseFile(filepath.Join("testdata", test.file), yini.Options{Mode: test.mode})
			if err != nil {
				t.Fatalf("ParseFile: %v (diagnostics: %#v)", err, result.Diagnostics)
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
			values := result.Document.Values()
			values["changed"] = true
			if _, exists := result.Document.Values()["changed"]; exists {
				t.Fatal("Values() returned document-owned map")
			}
			for _, section := range values {
				if nested, ok := section.(map[string]any); ok {
					nested["changed"] = true
				}
			}
			if !reflect.DeepEqual(result.Document.Values(), test.want) {
				t.Fatal("Values() returned document-owned nested map")
			}
		})
	}
}

func TestLenientFixturesRejectStrictMode(t *testing.T) {
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
			result, err := yini.ParseFile(filepath.Join("testdata", "lenient", test.file), yini.Options{Mode: yini.ModeStrict})
			var parseError *yini.ParseError
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
			result, err := yini.ParseString(test.text, yini.Options{Mode: test.mode})
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
	text := "\uFEFF#!/usr/bin/env yini\n^ Fleet\nname = 'First'\nname = 'Second'\n^^ Vessel\nid = 1\n^^ Vessel\nid = 2\n"
	result, err := yini.ParseString(text, yini.Options{})
	if err != nil {
		t.Fatalf("ParseString: %v (diagnostics: %#v)", err, result.Diagnostics)
	}
	want := map[string]any{"Fleet": map[string]any{"name": "First", "Vessel": map[string]any{"id": int64(1)}}}
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
	text := "^ Bench\nmin = -9223372036854775808\nmax = 9223372036854775807\nhex = hex:FF\nduo = 0zX\nbin = %101\nexp = 3e4\nnegativeZero = -0.0\nlabel = 'count=' + 0x10\n"
	result, err := yini.ParseString(text, yini.Options{})
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
	for _, literal := range []string{"9223372036854775808", "-9223372036854775809", "1e400", "1e-400"} {
		result, err := yini.ParseString("^ Bench\nvalue = "+literal+"\n", yini.Options{})
		if err == nil || !result.HasErrors() {
			t.Errorf("%s should fail: %#v", literal, result.Diagnostics)
		}
	}
}

func TestEOFAndTooling(t *testing.T) {
	result, err := yini.ParseString("^ Last\nvalue = C'line\\nend'", yini.Options{})
	if err != nil {
		t.Fatalf("final assignment at EOF: %v", err)
	}
	if got := result.Document.Values()["Last"].(map[string]any)["value"]; got != "line\nend" {
		t.Fatalf("decoded string = %#v", got)
	}
	_, err = yini.ParseString("^ Last\nvalue = 1\n/END", yini.Options{Mode: yini.ModeStrict})
	if err != nil {
		t.Fatalf("/END at EOF: %v", err)
	}
	result, err = yini.ParseForTooling("^ Last\nvalue = [1,,2]", yini.Options{})
	if err != nil || !result.HasErrors() {
		t.Fatalf("tooling parse = %#v, %v", result.Diagnostics, err)
	}
	_, err = yini.ParseString("^ Last\nvalue = [1,,2]", yini.Options{})
	var parseErr *yini.ParseError
	if !errors.As(err, &parseErr) || !strings.Contains(parseErr.Error(), "line") {
		t.Fatalf("expected positioned ParseError, got %v", err)
	}
}

func TestStrictFilenameDoesNotSelectMode(t *testing.T) {
	path := filepath.Join(t.TempDir(), "schedule.strict.yini")
	if err := os.WriteFile(path, []byte("orphan = 1"), 0600); err != nil {
		t.Fatal(err)
	}
	result, err := yini.ParseFile(path, yini.Options{})
	if err != nil || !hasDiagnostic(result, "filename-mode") {
		t.Fatalf("lenient file: %v (%#v)", err, result.Diagnostics)
	}
	if result.Document.Values()["orphan"] != int64(1) {
		t.Fatalf("filename changed parsing mode: %#v", result.Document.Values())
	}
	strict, err := yini.ParseFile(path, yini.Options{Mode: yini.ModeStrict})
	if err == nil || !hasDiagnostic(strict, "orphan-member") || hasDiagnostic(strict, "filename-mode") {
		t.Fatalf("strict file: %v (%#v)", err, strict.Diagnostics)
	}
}

func TestReaderAndInvalidUTF8(t *testing.T) {
	result, err := yini.ParseReader(strings.NewReader("^ Passage\nopen = yes"), yini.Options{})
	if err != nil || result.Document.Values()["Passage"].(map[string]any)["open"] != true {
		t.Fatalf("ParseReader: %v (%#v)", err, result.Diagnostics)
	}
	result, err = yini.ParseString("^ Passage\rentry = \xff", yini.Options{})
	if err == nil || len(result.Diagnostics) != 1 || result.Diagnostics[0].Code != "invalid-utf8" || result.Diagnostics[0].Position.Line != 2 {
		t.Fatalf("invalid UTF-8: %v (%#v)", err, result.Diagnostics)
	}
}
