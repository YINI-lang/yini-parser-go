# Development

This is a Go module with the importable `yini` package at its root. Go does not give `/src` special meaning; putting public Go files there would change the import path. Generated parser files are private in `internal/grammar/generated/`. The files under `grammar/` and `spec/` are read-only reference copies and must not be edited here.

Use [Task](https://taskfile.dev/) from the repository root:

| Command         | Action |
|-----------------|--------|
| `task generate` | Generate ANTLR Go files from `grammar/v1.0.0-rc.6/` into `internal/grammar/generated/`. |
| `task fmt`      | Format Go packages. |
| `task vet`      | Analyze Go packages. |
| `task test`     | Run all tests. |
| `task check`    | Run format, vet, and test. |

`task generate` uses the included `libs/antlr4/antlr-4.13.2-complete.jar` and requires Java. Generated files are checked in, so consumers need only Go. `go.mod` pins the separate ANTLR Go runtime. Run `task generate` only when generation inputs or settings change, and run `task check` before submitting work. `task vet` disables the unreachable-code analyzer because ANTLR emits unreachable branches.

Integration tests and original YINI fixtures live in `tests/integration/`. Add focused valid and invalid cases when changing semantic validation; assert line and column for source-location fixes. New fixtures must use original YINI content rather than copying content from sibling parser repositories.

## Shared conformance suite

The Go adapter for `yini-test-suite` lives in `cmd/yini-test-adapter/`. Build it and run both modes against the installed test suite:

```powershell
go build -o yini-test-adapter.exe ./cmd/yini-test-adapter
yini-test-suite all --all-modes --adapter ./yini-test-adapter.exe --input '{input}' --mode '{mode}'
```

The adapter writes parsed JSON to stdout and diagnostics to stderr. The suite reports pass and fail totals even when some cases fail.

---

^YINI ≡
