package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestReadResults(t *testing.T) {
	// 1. Assign.
	input := `{"Action":"start","Package":"example/parser"}
{"Action":"run","Package":"example/parser","Test":"TestParse"}
{"Action":"output","Package":"example/parser","Test":"TestParse/valid","Output":"test log\n"}
{"Action":"pass","Package":"example/parser","Test":"TestParse/valid","Elapsed":0.2}
{"Action":"skip","Package":"example/parser","Test":"TestParse/unsupported"}
{"Action":"fail","Package":"example/parser","Test":"TestParse/invalid","Elapsed":0.1}
{"Action":"fail","Package":"example/parser","Test":"TestParse","Elapsed":0.4}
{"Action":"fail","Package":"example/parser","Elapsed":0.5}
{"Action":"output","Package":"example/generated","Output":"[no test files]\n"}
{"Action":"skip","Package":"example/generated"}
{"Action":"build-output","ImportPath":"example/broken","Output":"compile error\n"}
{"Action":"build-fail","ImportPath":"example/broken"}
`
	var stderr bytes.Buffer

	// 2. Act.
	results, err := readResults(strings.NewReader(input), &stderr)

	// 3. Assert.
	if err != nil {
		t.Fatal(err)
	}
	if len(results) != 2 {
		t.Fatalf("packages = %d, want 2", len(results))
	}
	result := results["example/parser"]
	if result.passed != 1 || result.failed != 2 || result.skipped != 1 || result.tests != 1 || result.subtests != 3 {
		t.Fatalf("counts = %+v, want 1 passed, 2 failed, 1 skipped, 1 top-level test, 3 subtests", result)
	}
	if result.elapsed != 0.5 || result.status != "fail" || result.output.String() != "test log\n" {
		t.Errorf("package result = %+v, want package duration 0.5s, failure, and retained output", result)
	}
	generated := results["example/generated"]
	if generated.skipped != 0 || generated.tests != 0 || generated.status != "skip" {
		t.Errorf("package without tests inflated test counts: %+v", generated)
	}
	if stderr.String() != "compile error\n" {
		t.Errorf("build output = %q, want compile error", stderr.String())
	}
}

func TestReadResultsLongOutput(t *testing.T) {
	// 1. Assign.
	want := strings.Repeat("long test output ", 10000)
	var input bytes.Buffer
	if err := json.NewEncoder(&input).Encode(testEvent{Action: "output", Package: "example/parser", Output: want}); err != nil {
		t.Fatal(err)
	}

	// 2. Act.
	results, err := readResults(&input, &bytes.Buffer{})

	// 3. Assert.
	if err != nil {
		t.Fatal(err)
	}
	if got := results["example/parser"].output.String(); got != want {
		t.Errorf("retained %d output bytes, want %d", len(got), len(want))
	}
}

func TestWriteSummary(t *testing.T) {
	// 1. Assign.
	results := map[string]*packageResult{
		"example/z": {passed: 2, skipped: 1, tests: 1, subtests: 2, status: "pass", elapsed: 0.5},
		"example/a": {status: "skip"},
	}
	var output bytes.Buffer

	// 2. Act.
	err := writeSummary(&output, results, 2*time.Second, false)

	// 3. Assert.
	if err != nil {
		t.Fatal(err)
	}
	lines := strings.Split(output.String(), "\n")
	if len(lines) < 7 {
		t.Fatalf("incomplete summary: %q", output.String())
	}
	wantRows := []string{
		"Package Passed Failed Skipped Total Duration Result",
		"example/a 0 0 0 0 0.000s NO TESTS",
		"example/z 2 0 1 3 0.500s PASS",
		"Total 2 0 1 3 2.000s",
	}
	for i, want := range wantRows {
		if got := strings.Join(strings.Fields(lines[i]), " "); got != want {
			t.Errorf("row %d = %q, want %q", i, got, want)
		}
	}
	if !strings.Contains(output.String(), "1 top-level tests and 2 named subtests, including parent groups.") ||
		!strings.Contains(output.String(), "Result: PASS") {
		t.Errorf("missing count explanation or final result: %q", output.String())
	}
}

func TestRunCommand(t *testing.T) {
	// 1. Assign.
	tests := []struct {
		name       string
		output     string
		stderr     string
		exit       int
		wantCode   int
		wantResult string
		wantError  string
	}{
		{
			name: "pass",
			output: `{"Action":"output","Package":"example/parser","Output":"passing log\n"}
{"Action":"pass","Package":"example/parser","Test":"TestParse"}
{"Action":"pass","Package":"example/parser","Elapsed":0.1}`,
			wantResult: "PASS",
		},
		{
			name: "failed test",
			output: `{"Action":"output","Package":"example/parser","Output":"parser_test.go:12: wrong value\n"}
{"Action":"fail","Package":"example/parser","Test":"TestParse"}
{"Action":"fail","Package":"example/parser"}`,
			exit: 1, wantCode: 1, wantResult: "FAIL", wantError: "parser_test.go:12: wrong value",
		},
		{
			name: "failed package without failed tests",
			output: `{"Action":"pass","Package":"example/parser","Test":"TestParse"}
{"Action":"fail","Package":"example/parser"}`,
			wantCode: 1, wantResult: "FAIL",
		},
		{
			name:   "build failure on stderr",
			output: `{"Action":"fail","Package":"example/parser"}`,
			stderr: "undefined: brokenSymbol\n", exit: 1, wantCode: 1, wantResult: "FAIL", wantError: "undefined: brokenSymbol",
		},
		{
			name: "process failure after passing events",
			output: `{"Action":"pass","Package":"example/parser","Test":"TestParse"}
{"Action":"pass","Package":"example/parser"}`,
			exit: 2, wantCode: 1, wantResult: "FAIL", wantError: "Run tests:",
		},
		{
			name: "truncated JSON", output: `{"Action":`,
			wantCode: 1, wantResult: "FAIL", wantError: "Read test results:",
		},
		{
			name: "incomplete package", output: `{"Action":"output","Package":"example/parser","Output":"panic: stopped before package result\n"}`,
			wantCode: 1, wantResult: "FAIL", wantError: "panic: stopped before package result",
		},
		{name: "no package results", wantCode: 1, wantResult: "FAIL"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// 1. Assign.
			command := exec.Command(os.Args[0], "-test.run=^TestSummaryProcessHelper$")
			command.Env = append(os.Environ(),
				"YINI_SUMMARY_PROCESS_HELPER=1",
				"YINI_SUMMARY_OUTPUT="+test.output,
				"YINI_SUMMARY_STDERR="+test.stderr,
				"YINI_SUMMARY_EXIT="+strconv.Itoa(test.exit))
			var stdout, stderr bytes.Buffer

			// 2. Act.
			code := runCommand(command, &stdout, &stderr)

			// 3. Assert.
			if code != test.wantCode {
				t.Errorf("exit code = %d, want %d (stderr: %s)", code, test.wantCode, stderr.String())
			}
			if !strings.Contains(stdout.String(), "Result: "+test.wantResult) {
				t.Errorf("summary = %q, want result %s", stdout.String(), test.wantResult)
			}
			if !strings.Contains(stderr.String(), test.wantError) {
				t.Errorf("stderr = %q, want %q", stderr.String(), test.wantError)
			}
			if test.wantCode == 0 && stderr.Len() != 0 {
				t.Errorf("passing run printed diagnostics: %q", stderr.String())
			}
		})
	}
}

func TestRunCommandStartFailure(t *testing.T) {
	// 1. Assign.
	command := exec.Command(filepath.Join(t.TempDir(), "missing-command"))
	var stdout, stderr bytes.Buffer

	// 2. Act.
	code := runCommand(command, &stdout, &stderr)

	// 3. Assert.
	if code != 1 || !strings.Contains(stderr.String(), "Start tests:") {
		t.Errorf("exit code = %d, stderr = %q; want start failure", code, stderr.String())
	}
}

type failingWriter struct{}

func (failingWriter) Write([]byte) (int, error) {
	return 0, errors.New("output unavailable")
}

func TestWriteSummaryFailure(t *testing.T) {
	// 1. Assign.
	output := failingWriter{}

	// 2. Act.
	err := writeSummary(output, nil, 0, false)

	// 3. Assert.
	if err == nil || !strings.Contains(err.Error(), "output unavailable") {
		t.Errorf("error = %v, want output failure", err)
	}
}

func TestSummaryProcessHelper(t *testing.T) {
	if os.Getenv("YINI_SUMMARY_PROCESS_HELPER") != "1" {
		return
	}
	fmt.Fprint(os.Stdout, os.Getenv("YINI_SUMMARY_OUTPUT"))
	fmt.Fprint(os.Stderr, os.Getenv("YINI_SUMMARY_STDERR"))
	code, err := strconv.Atoi(os.Getenv("YINI_SUMMARY_EXIT"))
	if err != nil {
		os.Exit(2)
	}
	os.Exit(code)
}
