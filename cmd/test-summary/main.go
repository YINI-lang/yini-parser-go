package main

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"os/exec"
	"sort"
	"strings"
	"text/tabwriter"
	"time"
)

func main() {
	os.Exit(runCommand(exec.Command("go", "test", "-json", "-count=1", "./..."), os.Stdout, os.Stderr))
}

type testEvent struct {
	Action  string
	Package string
	Test    string
	Elapsed float64
	Output  string
}

type packageResult struct {
	passed   int
	failed   int
	skipped  int
	tests    int
	subtests int
	elapsed  float64
	status   string
	output   strings.Builder
}

func runCommand(command *exec.Cmd, stdout, stderr io.Writer) int {
	started := time.Now()
	command.Stderr = stderr
	pipe, err := command.StdoutPipe()
	if err != nil {
		fmt.Fprintf(stderr, "Read test output: %v\n", err)
		return 1
	}
	if err := command.Start(); err != nil {
		fmt.Fprintf(stderr, "Start tests: %v\n", err)
		return 1
	}
	results, readErr := readResults(pipe, stderr)
	if readErr != nil {
		// Let the child finish even if its remaining output cannot be decoded.
		_, _ = io.Copy(io.Discard, pipe)
	}
	waitErr := command.Wait()
	failed := readErr != nil || waitErr != nil || len(results) == 0
	for _, result := range results {
		if result.status == "" || result.status == "fail" || result.failed > 0 {
			failed = true
		}
	}
	writeErr := writeSummary(stdout, results, time.Since(started), failed)
	for _, name := range sortedPackages(results) {
		result := results[name]
		if result.status == "" || result.status == "fail" || result.failed > 0 {
			fmt.Fprint(stderr, result.output.String())
		}
	}
	if readErr != nil {
		fmt.Fprintf(stderr, "Read test results: %v\n", readErr)
	}
	if waitErr != nil {
		fmt.Fprintf(stderr, "Run tests: %v\n", waitErr)
	}
	if writeErr != nil {
		fmt.Fprintf(stderr, "Write test summary: %v\n", writeErr)
	}
	if failed || writeErr != nil {
		return 1
	}
	return 0
}

func readResults(input io.Reader, stderr io.Writer) (map[string]*packageResult, error) {
	results := make(map[string]*packageResult)
	decoder := json.NewDecoder(input)
	for {
		var event testEvent
		if err := decoder.Decode(&event); err != nil {
			if err == io.EOF {
				return results, nil
			}
			return results, err
		}
		if strings.HasPrefix(event.Action, "build-") {
			fmt.Fprint(stderr, event.Output)
			continue
		}
		if event.Package == "" {
			continue
		}
		result := results[event.Package]
		if result == nil {
			result = &packageResult{}
			results[event.Package] = result
		}
		result.output.WriteString(event.Output)
		if event.Test == "" {
			if event.Action == "pass" || event.Action == "fail" || event.Action == "skip" {
				result.status = event.Action
				result.elapsed = event.Elapsed
			}
			continue
		}
		switch event.Action {
		case "pass":
			result.passed++
		case "fail":
			result.failed++
		case "skip":
			result.skipped++
		default:
			continue
		}
		if strings.Contains(event.Test, "/") {
			result.subtests++
		} else {
			result.tests++
		}
	}
}

func sortedPackages(results map[string]*packageResult) []string {
	names := make([]string, 0, len(results))
	for name := range results {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func writeSummary(output io.Writer, results map[string]*packageResult, elapsed time.Duration, failed bool) error {
	table := tabwriter.NewWriter(output, 0, 4, 2, ' ', 0)
	fmt.Fprintln(table, "Package\tPassed\tFailed\tSkipped\tTotal\tDuration\tResult")
	var passed, failures, skipped, tests, subtests int
	for _, name := range sortedPackages(results) {
		result := results[name]
		status := strings.ToUpper(result.status)
		if result.status == "skip" && result.tests+result.subtests == 0 {
			status = "NO TESTS"
		} else if status == "" {
			status = "INCOMPLETE"
		}
		fmt.Fprintf(table, "%s\t%d\t%d\t%d\t%d\t%.3fs\t%s\n", name,
			result.passed, result.failed, result.skipped,
			result.passed+result.failed+result.skipped, result.elapsed, status)
		passed += result.passed
		failures += result.failed
		skipped += result.skipped
		tests += result.tests
		subtests += result.subtests
	}
	fmt.Fprintf(table, "Total\t%d\t%d\t%d\t%d\t%.3fs\n", passed, failures, skipped,
		passed+failures+skipped, elapsed.Seconds())
	if err := table.Flush(); err != nil {
		return err
	}
	status := "PASS"
	if failed {
		status = "FAIL"
	}
	_, err := fmt.Fprintf(output, "\nCounts include %d top-level tests and %d named subtests, including parent groups.\nResult: %s\n", tests, subtests, status)
	return err
}
