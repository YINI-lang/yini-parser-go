# Decode YINI documents into Go structs

Date:   2026-10-25
Status: Proposal for API review.

## Goal

Let applications decode parsed YINI values into typed Go configuration structs without manual `map[string]any` assertions. Keep parsing, diagnostics, and `Document.Values()` as they are. Decoding is an explicit second step after a successful parse.

## Proposed API

```go
func (d *Document) Decode(dst any) error
func (s *Section) Decode(dst any) error

type DecodeError struct {
    Path     string
    Position Position
    Message  string
}
func (e *DecodeError) Error() string
```

Both methods require a non-nil pointer to a struct. A nil document or section receiver returns an error. `Document.Decode` maps the document root: accepted orphan members and top-level sections are peers, as in `Document.Values()`. `Section.Decode` maps that section's members and child sections directly to struct fields. Neither method automatically unwraps a document's single top-level section.

For example, given `^ Service` with `port = 8080`, callers may decode the whole document into a struct with a `Service` field or decode the `Service` section into a struct with a `port` field. The selected parse mode does not change decoding rules.

```go
type ServiceConfig struct {
    Port int64 `yini:"port"`
}

result, err := yini.ParseFile("service.strict.yini", yini.Options{Mode: yini.ModeStrict})
if err != nil {
    return err
}
var config ServiceConfig
if err := result.Document.Sections[0].Decode(&config); err != nil {
    return err
}
```

The example assumes a valid strict document, which has exactly one top-level section. To decode the whole document, use a wrapper field for that section:

```go
type Config struct {
    Service ServiceConfig `yini:"Service"`
}
var whole Config
err = result.Document.Decode(&whole)
```

## Mapping rules

| YINI value | Supported Go destination |
| --- | --- |
| String | `string` |
| Boolean | `bool` |
| Integer (`int64`) | `int`, `int8`–`int64`, or `uint`, `uint8`–`uint64`, only if the value fits |
| Float (`float64`) | `float64` |
| List | Slice of a supported destination type |
| Inline object or section | Struct or `map[string]T` |
| Null | Pointer, slice, map, or empty interface; sets it to nil |

Pointers to supported types are supported recursively. Named scalar types with one of the supported underlying kinds follow the same checks. An empty interface receives an independent copy of the existing native value (`string`, `bool`, `int64`, `float64`, `[]any`, or `map[string]any`). `float32`, non-empty interfaces, fixed arrays, `uintptr`, and unsupported kinds return an error. No string-to-number, integer-to-float, float-to-integer, or other implicit conversion occurs. This keeps the parser's exact integer values intact.

Only exported struct fields participate. An untagged field matches a YINI name exactly, including case. A `yini:"name"` tag sets the exact source name; `yini:"-"` excludes a Go field. Tags have no options in this version. Duplicate field mappings and malformed destination types are errors. Anonymous embedded fields are rejected unless tagged `yini:"-"`; they are not flattened.

Unknown YINI members or child sections are errors when decoding into a struct; a `map[string]T` accepts any source name. Missing fields remain at their Go zero value; there is no required-field tag in this version. A present null for a non-nullable destination is an error. This schema check is independent of lenient and strict parsing: parsing a valid document can succeed while decoding it into a particular struct fails.

Decoding builds a fresh destination value and assigns it only after all fields decode successfully. On failure, the caller's destination remains unchanged. On success, fields absent from YINI have zero values, including when the destination previously held data. Decoding does not mutate the source document.

## Errors and locations

Decoding returns a Go error and does not append to parse diagnostics. A `*DecodeError` reports the YINI path and the nearest available source position. Use `errors.As` to inspect it. For example, a bad second list element can report `Service.ports[1]` with the `ports` member's position. A section mismatch uses the section header position. Destination validation errors without a source location use a zero `Position`.

Return the first error deterministically. Traverse document members and sections in source order. Inline objects are currently stored as Go maps, so examine their keys in sorted order. Error text should identify the source path, source type, and required Go type in plain language. Do not print from library code.

The current document tree does not retain separate locations for inline object entries or list elements. This proposal uses the containing member's position for those values; adding finer locations is a separate parser change.

## Implementation plan

1. Add the public methods and `DecodeError` at the module root. Share one private recursive decoder for document and section content. Use the Go standard library (`reflect` and `sort`); add no runtime dependency or JSON round trip.
2. Validate the destination schema before reading values: pointer shape, exported fields, tag mappings, duplicate names, and supported kinds. Guard recursive Go types during validation. Build a fresh value so failures cannot partially modify the caller's struct.
3. Decode sections, members, inline objects, and lists recursively with explicit type and range checks. Preserve exact `int64` and `float64` categories, nil behavior, and deterministic error order.
4. Add public API integration tests in both parse modes. Cover nested structs, slices, maps, pointers, tags, case sensitivity, unknown and missing fields, nulls, numeric boundaries, source paths and positions, duplicate mappings, unsupported and nil destinations, recursive Go types, and unchanged destination on error. Include a valid case beside each important invalid case.
5. Add a concise README example, detailed rules in `docs/usage.md`, and a changelog entry. Run `task check`; grammar and generated files need no changes.

## Acceptance criteria

- A valid parsed document or section decodes to the expected typed struct without callers asserting `map[string]any` values.
- Any type mismatch, overflow, unknown source name, or invalid destination returns an inspectable error and leaves the destination unchanged.
- Errors identify the affected YINI path and provide a source position when the document stores one.
- Existing parse APIs, mode behavior, diagnostics, and `Document.Values()` output remain unchanged.

## API review decisions

The proposed methods and exported error type are additions to the public API, so the maintainer must approve their shape before implementation. In particular, this proposal recommends exact-case field matching and rejecting unknown YINI fields. Those defaults should be confirmed as part of API review because changing them after release would affect callers.
