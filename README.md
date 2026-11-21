# yini-parser-go

[![Go 1.22+](https://img.shields.io/badge/Go-1.22%2B-00ADD8?logo=go&logoColor=white)](go.mod)
[![v1.0.0-beta.1](https://img.shields.io/badge/release-v1.0.0--beta.1-orange)](CHANGELOG.md)
[![MIT License](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![YINI tests](https://github.com/YINI-lang/yini-parser-go/actions/workflows/yini-test-suite.yml/badge.svg?branch=main&event=push)](https://github.com/YINI-lang/yini-parser-go/actions/workflows/yini-test-suite.yml)

[YINI](https://yini-lang.org/refs/specification/?utm_source=github&utm_campaign=yini_parser_go&utm_content=readme&utm_medium=referral) is a human-readable configuration format for application settings. Named section markers make its structure explicit, and values can be strings, numbers, booleans, lists, or objects. Indentation is for readability and does not affect parsing.

`yini-parser-go` lets Go applications and tools parse YINI files or text and access a structured document, Go values, and diagnostics. It supports lenient and strict parsing. Version `v1.0.0-beta.1` targets YINI 1.0.0 RC.7.

The module supports Go 1.22 or newer. It uses ANTLR-generated Go code and the pinned ANTLR Go runtime; applications need only the Go module, not Java.

Start with the [package overview](doc.go) for the library entry points and source organization.

Useful links:

- [Try YINI in Playground](https://yini-lang.org/playground/?utm_source=github&utm_campaign=yini_parser_go&utm_content=readme_useful_links&utm_medium=referral)
- [This yini-parser-go on GitHub](https://github.com/YINI-lang/yini-parser-go)
- [Report an Issue or give Feedback](https://github.com/YINI-lang/yini-parser-go/issues)

---

## Install

```sh
go get github.com/YINI-lang/yini-parser-go@v1.0.0-beta.1
```

---

## Parse inline text

```go
package main

import (
    "fmt"

    yini "github.com/YINI-lang/yini-parser-go"
)

func main() {
    result, err := yini.ParseString("^ Service\nport = 8080\n", yini.Options{})
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(result.Document.Values()["Service"].(map[string]any)["port"]) // 8080
}
```

`Options{}` selects lenient mode. Choose strict mode explicitly:

```go
result, err := yini.ParseFile("settings.strict.yini", yini.Options{Mode: yini.ModeStrict})
```

Strict mode requires one explicit top-level section and `/END`, and rejects orphan members. The `.strict.yini` filename suffix and `@yini` declaration do not select the mode; `Options.Mode` does. A file can end immediately after its final assignment in lenient mode or `/END` in strict mode.

To decode a parsed section into a typed struct, call `Decode` after checking the parse error:

```go
type ServiceConfig struct {
    Port int64 `yini:"port"`
}

var service ServiceConfig
if err := result.Document.Sections[0].Decode(&service); err != nil {
    fmt.Println(err)
    return
}
```

`Document.Decode` maps the whole document, including top-level section names. See [usage](docs/usage.md) for the mapping and error rules.

---

## Results and diagnostics

`ParseString`, `ParseFile`, and `ParseReader` return a `Result` and a Go `error`. The result contains a structured `Document` with section/member source positions and all diagnostics. On malformed input, the document may be partial and the error includes a line and column. `ParseFile` also includes the path; use `errors.As` to inspect its wrapped `*yini.ParseError`. Warnings remain in `Result.Diagnostics` without causing an error. Library calls do not print diagnostics.

`Document.Values()` returns a deep-copy `map[string]any` for application code. Lenient orphan members and top-level sections appear side by side in this map. Values use `string`, `bool`, `nil`, `int64`, `float64`, `[]any`, and `map[string]any`. Integers outside signed 64-bit range and floating literals that overflow or lose a nonzero value to zero produce errors.

For editors and other tools, `ParseForTooling` returns diagnostics and recoverable document content without returning an error for normal YINI syntax failures.

See [usage](docs/usage.md) for mode and diagnostic details and [development](docs/development.md) for the Taskfile commands and repository layout. The active grammar is in `grammar/v1.0.0-rc.7/`; `spec/` retains an older RC.6 reference.

---

## Shared conformance tests

The Go adapter in [`cmd/yini-test-adapter/`](cmd/yini-test-adapter/) runs this parser against the external `yini-test-suite`. It accepts a YINI file and an explicit lenient or strict mode, then writes parsed JSON to stdout and diagnostics to stderr. See the [development guide](docs/development.md#shared-conformance-suite) for build and run commands. CI pins `yini-test-suite` to `0.4.0rc1` and runs smoke and full conformance tests in both modes. All 542 cases currently pass.

---

## Help and bug reports

For questions, bug reports, or feature requests, [open an issue](https://github.com/YINI-lang/yini-parser-go/issues). For parser problems, include a small YINI example, the selected mode, the expected result, and the actual result or diagnostics.

---

## Links

- [YINI Cheat Sheet](https://yini-lang.org/refs/yini-cheat-sheet/?utm_source=github&utm_campaign=yini_parser_go&utm_content=readme_links&utm_medium=referral)
- [YINI Specification](https://yini-lang.org/refs/specification/?utm_source=github&utm_campaign=yini_parser_go&utm_content=readme_links&utm_medium=referral)
- [YINI Demo Apps](https://github.com/YINI-lang/yini-demo-apps/tree/main)

---

## License

This project is licensed under the [MIT License](LICENSE).

---

**^YINI ≡**  
> YINI is a human-readable configuration format designed for clarity, readability, explicit structure, predictability, and deterministic parsing.
> 
> See the specification for syntax and format details.  

[yini-lang.org](https://yini-lang.org/?utm_source=github&utm_campaign=yini_parser_go&utm_medium=referral&utm_content=readme_footer) · [YINI Project on GitHub](https://github.com/YINI-lang)  
