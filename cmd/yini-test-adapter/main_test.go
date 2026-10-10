package main

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestRun(t *testing.T) {
	tests := []struct {
		name       string
		content    string
		mode       string
		wantCode   int
		wantJSON   string
		wantStderr string
	}{
		{
			name:     "lenient by default",
			content:  "count = 9223372036854775807\n^ Depot\nopen = yes\n",
			wantJSON: `{"Depot":{"open":true},"count":9223372036854775807}` + "\n",
		},
		{
			name:     "strict",
			content:  "^ Depot\nopen = yes\n/END\n",
			mode:     "strict",
			wantJSON: `{"Depot":{"open":true}}` + "\n",
		},
		{
			name:       "warning",
			content:    "^ Depot\nopen = yes\nopen = no\n",
			wantJSON:   `{"Depot":{"open":true}}` + "\n",
			wantStderr: "warning duplicate-key at line 3, column 1",
		},
		{
			name:       "parse failure",
			content:    "^ Depot\nopen = yes\n",
			mode:       "strict",
			wantCode:   1,
			wantStderr: "error missing-end",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			path := filepath.Join(t.TempDir(), "input.yini")
			if err := os.WriteFile(path, []byte(test.content), 0600); err != nil {
				t.Fatal(err)
			}
			args := []string{"--input", path}
			if test.mode != "" {
				args = append(args, "--mode", test.mode)
			}
			var stdout, stderr bytes.Buffer
			if got := run(args, &stdout, &stderr); got != test.wantCode {
				t.Errorf("exit code = %d, want %d; stderr = %q", got, test.wantCode, stderr.String())
			}
			if got := stdout.String(); got != test.wantJSON {
				t.Errorf("stdout = %q, want %q", got, test.wantJSON)
			}
			if !strings.Contains(stderr.String(), test.wantStderr) {
				t.Errorf("stderr = %q, want to contain %q", stderr.String(), test.wantStderr)
			}
		})
	}
}

func TestRunRejectsInvalidArguments(t *testing.T) {
	for _, args := range [][]string{
		{},
		{"--input", "file.yini", "extra"},
		{"--input", "file.yini", "--mode", "other"},
	} {
		var stdout, stderr bytes.Buffer
		if got := run(args, &stdout, &stderr); got != 2 || stdout.Len() != 0 || stderr.Len() == 0 {
			t.Errorf("run(%q) = %d, stdout %q, stderr %q", args, got, stdout.String(), stderr.String())
		}
	}
}
