# Changelog

## v1.0.0-alpha.1

Initial alpha release of the Go parser for YINI 1.0.0 RC.6.

- Add `ParseString`, `ParseFile`, and `ParseReader` for inline text, files, and readers, plus `ParseForTooling` for recoverable diagnostics.
- Support lenient and strict modes with structured diagnostics and source locations. `Document.Values()` provides a deep copy of parsed values with exact signed 64-bit integers.
- Include the generated ANTLR parser, original YINI fixtures, integration tests, and Taskfile checks.
- Add a Go adapter for the external `yini-test-suite` and run its lenient and strict smoke cases in CI.

The full external conformance suite still has known failures. The smoke suite passes in both modes.

---

^YINI ≡
