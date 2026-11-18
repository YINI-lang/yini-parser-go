// Package yini parses YINI configuration documents.
//
// Start in parse.go. Use [ParseString] to read text, [ParseFile] to read a file,
// or [ParseReader] to read from an io.Reader. Parsing is lenient by default.
// Set Options.Mode to [ModeStrict] to apply strict rules.
//
// # Results and diagnostics
//
// Check the returned error before using the parsed document. [Result] contains
// the document and any warnings or errors, with their line and column numbers.
// Invalid input may produce an incomplete document. [ParseForTooling] lets
// editors inspect that document and its diagnostics without returning a Go
// error for ordinary input problems. Callers decide how to display diagnostics.
//
// # Values and decoding
//
// Use [Document.Values] to get a map you can change without changing the
// document. To fill a Go struct, use [Document.Decode] or [Section.Decode].
// Fields match by name or by yini tags. If decoding fails, the struct is left
// unchanged.
//
// # Source organization
//
// types.go defines documents and diagnostics. decode.go handles Go structs.
// The other root files read grammar tokens, build values, and track source
// locations. Generated grammar code is under internal/grammar/generated.
// Unit tests live beside the code they test. Integration tests and their YINI
// input files live under tests/integration.
package yini
