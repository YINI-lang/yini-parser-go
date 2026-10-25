# Using the parser

Import `github.com/YINI-lang/yini-parser-go`. Public Go files are at the module root, so the import path has no `/src` suffix. The ANTLR-generated package is private under `internal/`.

## Entry points

| Function | Input | Result |
| --- | --- | --- |
| `ParseString(text, options)` | Inline YINI text | `Result, error` |
| `ParseFile(path, options)` | UTF-8 file | `Result, error` |
| `ParseReader(reader, options)` | `io.Reader` | `Result, error` |
| `ParseForTooling(text, options)` | Inline text | `Result, error`; syntax/validation issues stay in diagnostics |

Use `Options{}` for lenient mode or `Options{Mode: yini.ModeStrict}` for strict mode. An invalid mode returns an error. The filename and `@yini` declaration never change the selected mode. `ParseFile` warns when a `.strict.yini` file is parsed leniently. A mismatched `@yini strict` declaration is an error in lenient mode; `@yini lenient` is a warning in strict mode.

## Document and values

`Result.Document` contains `Members` at the root and `Sections`, each with its name, level, source position, members, and nested sections. The document may be partial if parsing reported errors. `Document.Values()` produces a deep copy of the parsed values. It mounts accepted lenient orphan members beside top-level sections; a name collision is an error.

The value mapping is:
| YINI value | Go value |
| --- | --- |
| Quoted string | `string` |
| Boolean | `bool` |
| Null or lenient empty member | `nil` |
| Integer in decimal, binary, octal, dozenal, or hexadecimal | `int64` |
| Fraction or exponent form, including `3e4` | `float64` |
| List | `[]any` |
| Inline object or section | `map[string]any` |

Integer conversion is exact within signed 64-bit range. Fractions and exponents use IEEE 754 binary64. Overflow and nonzero underflow to zero are errors, so callers never receive a silently rounded large integer. Floating values may have ordinary binary64 rounding.

## Modes and recovery

| Rule | Lenient | Strict |
| --- | --- | --- |
| Top-level structure | Sections and orphan members | Exactly one top-level section; no orphans |
| `/END` | Optional | Required |
| Duplicate key or sibling section | Warning; first wins | Error |
| Empty member value | Null | Error; write `null` |
| Trailing list/object comma | Accepted | Error |
| Inline object `=` | Accepted | Error; use `:` |
| Concatenation operands after the first string | Strings, numbers, booleans, null | Strings only |

A duplicate section's whole block is ignored. A misplaced `#!` line is treated as a comment and reported once: warning in lenient mode, error in strict mode. EOF may end the last assignment or section header; strict `/END` also works at EOF.

Check both the Go error and `Result.Diagnostics`. Each diagnostic has a severity, code, plain-language message, and one-based line/column position. Common syntax errors describe what is missing or misplaced. `ParseFile` prefixes parse errors with the file path; use `errors.As` to retrieve the wrapped `*yini.ParseError`. `ParseForTooling` keeps diagnostics and suppresses the Go error only for ordinary parse failures; invalid options still return an error. `ParseFile` and `ParseReader` return I/O errors directly.

^YINI ≡
