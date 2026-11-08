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

## Decode into Go structs

After a successful parse, use `Document.Decode(&dst)` to map the whole document or `Section.Decode(&dst)` to map one section's contents. `dst` must be a non-nil pointer to a struct. `Document.Decode` keeps orphan members and top-level sections at the same level; it does not unwrap a single top-level section.

```go
type Service struct {
    Port int64 `yini:"port"`
}
type Config struct {
    Service Service `yini:"Service"`
}

func decodeConfig(result yini.Result) (Config, error) {
    var config Config
    if err := result.Document.Decode(&config); err != nil {
        return Config{}, err
    }
    return config, nil
}
```

Only exported, named fields participate. Untagged fields match YINI names exactly, including case. Use `yini:"name"` to choose a name and `yini:"-"` to ignore a field. Unexported fields, including embedded ones, are ignored. Exported anonymous embedded fields must be ignored with `yini:"-"` or replaced with named fields; decoding does not flatten them. Tag options and duplicate field mappings are errors. Unknown YINI names are errors when decoding to a struct; missing Go fields remain zero. `map[string]T` accepts any source name.

Strings decode to strings; booleans to bools; integers to signed or unsigned integer types if they fit; floats to `float64`; lists to slices; and objects or sections to structs or `map[string]T`. Pointers to supported types and named scalar types are supported. Null sets a pointer, slice, map, or empty interface to nil and is an error for other types. An empty interface receives a deep copy of the native value. There are no implicit numeric or string conversions. Unsupported types, including `float32`, arrays, and non-empty interfaces, are errors even if the corresponding field is absent from the document.

Decoding replaces the destination only after every value succeeds. On failure, the destination remains unchanged; on success, absent fields are reset to zero. `*DecodeError` provides a YINI path and the nearest available source position through `errors.As`. List elements and inline object entries use their containing member's position. Decode errors do not change parse diagnostics or `Document.Values()`.

## Modes and recovery

| Rule | Lenient | Strict |
| --- | --- | --- |
| Top-level structure | Sections and orphan members | Exactly one top-level section; no orphans |
| `/END` | Optional | Required |
| Duplicate key or sibling section | Warning; first wins | Error |
| Empty member value | Warning; treated as null | Error; write `null` |
| Trailing list/object comma | Accepted | Error |
| Inline object `=` | Accepted | Error; use `:` |
| Concatenation operands after the first string | Strings, numbers, booleans, null | Strings only |

A duplicate section's whole block is ignored. A misplaced `#!` line is treated as a comment and reported once: warning in lenient mode, error in strict mode. EOF may end the last assignment or section header; strict `/END` also works at EOF.

Check both the Go error and `Result.Diagnostics`. Each diagnostic has a severity, code, plain-language message, and one-based line/column position. Common syntax errors describe what is missing or misplaced. `ParseFile` prefixes parse errors with the file path; use `errors.As` to retrieve the wrapped `*yini.ParseError`. `ParseForTooling` keeps diagnostics and suppresses the Go error only for ordinary parse failures; invalid options still return an error. `ParseFile` and `ParseReader` return I/O errors directly.

^YINI ≡
