# Changelog

## Unreleased

- **Added:** Add `task build` and include compilation in `task check` and Go CI.
- **Updated:** Pin external CI to `yini-test-suite` 0.4.0rc1 and enable the full conformance workflow; all 542 smoke and golden cases pass in both modes.
- **Fixed:** Adapt the host parser to RC.7 grammar interfaces, including signed numbers, Boolean/null identifier names, compact inline-object keys named `hex`, escaped identifiers, and section-header comments.
- **Changed:** Reject duplicate sibling sections in both modes, retaining the first section only for recovery.
- **Fixed:** Support `E`/`e` escaped strings and command strings, normalize multiline source line endings to LF, preserve decimal spelling in lenient concatenation, and warn for deprecated `%` binary literals. Mixed source line endings warn in lenient mode and fail in strict mode.
- **Fixed:** Reject pointer-only recursive destination types during schema validation to prevent unbounded recursion while decoding.
- **Changed:** Warn when a lenient root-level or section-level member has an empty value, including when only whitespace or comments follow `=`. The value remains `null`; explicit `null` does not produce this warning.
- **Added:** Add `Document.Decode` and `Section.Decode` to map parsed values into Go structs with explicit `yini` field tags and exact numeric conversions.
- **Added:** Return `*DecodeError` with YINI paths and source positions; leave the destination unchanged if decoding fails.
- **Added:** Add struct-decoding regression tests in both parse modes for numeric boundaries, type mismatches, recursive types, deep-copy isolation, source locations, and unchanged destinations on failure.
- **Updated:** Expand integration test data from four to eight original YINI fixtures, covering duplicate recovery, root members and nulls, typed struct decoding, and numeric and string forms.

## v1.0.0-alpha.3

- Include the file path in `ParseFile` parse errors while keeping `*yini.ParseError` available through `errors.As`.
- Give more specific syntax diagnostics for common malformed lists, inline objects, concatenation, comments, and content after `/END`.

## v1.0.0-alpha.2
- Add `VERSION` as the source of truth for the parser release version.
- Add `task release` to run checks and create an annotated local Git tag from the trimmed `VERSION` value. It requires the version to be committed and the working tree to be clean before tagging.
- Update the README and development guide for the `v1.0.0-alpha.2` release workflow.

## v1.0.0-alpha.1

Initial alpha release of the Go parser for YINI 1.0.0 RC.6.

- Add `ParseString`, `ParseFile`, and `ParseReader` for inline text, files, and readers, plus `ParseForTooling` for recoverable diagnostics.
- Support lenient and strict modes with structured diagnostics and source locations. `Document.Values()` provides a deep copy of parsed values with exact signed 64-bit integers.
- Include the generated ANTLR parser, original YINI fixtures, integration tests, and Taskfile checks.
- Add a Go adapter for the external `yini-test-suite` and run its lenient and strict smoke cases in CI.

The full external conformance suite still has known failures. The smoke suite passes in both modes.

---

^YINI ≡
