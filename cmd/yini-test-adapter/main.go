package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"os"

	yini "github.com/YINI-lang/yini-parser-go"
)

func main() {
	os.Exit(run(os.Args[1:], os.Stdout, os.Stderr))
}

func run(args []string, stdout, stderr io.Writer) int {
	flags := flag.NewFlagSet("yini-test-adapter", flag.ContinueOnError)
	flags.SetOutput(stderr)
	input := flags.String("input", "", "path to a YINI file")
	mode := flags.String("mode", string(yini.ModeLenient), "parser mode: lenient or strict")
	if err := flags.Parse(args); err != nil {
		return 2
	}
	if *input == "" || flags.NArg() != 0 {
		fmt.Fprintln(stderr, "usage: yini-test-adapter --input <file> [--mode <lenient|strict>]")
		return 2
	}
	if *mode != string(yini.ModeLenient) && *mode != string(yini.ModeStrict) {
		fmt.Fprintf(stderr, "invalid mode %q: expected lenient or strict\n", *mode)
		return 2
	}

	result, err := yini.ParseFile(*input, yini.Options{Mode: yini.Mode(*mode)})
	for _, diagnostic := range result.Diagnostics {
		fmt.Fprintf(stderr, "%s %s at line %d, column %d: %s\n",
			diagnostic.Severity, diagnostic.Code, diagnostic.Position.Line,
			diagnostic.Position.Column, diagnostic.Message)
	}
	if err != nil {
		if len(result.Diagnostics) == 0 {
			fmt.Fprintln(stderr, err)
		}
		return 1
	}

	output, err := json.Marshal(result.Document.Values())
	if err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	if _, err := fmt.Fprintln(stdout, string(output)); err != nil {
		fmt.Fprintln(stderr, err)
		return 1
	}
	return 0
}
