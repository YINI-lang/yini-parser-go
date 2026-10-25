# Changelog

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
