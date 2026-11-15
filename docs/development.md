# Development

This is a Go module with the importable `yini` package at its root. Go does not give `/src` special meaning; putting public Go files there would change the import path. Generated parser files are private in `internal/grammar/generated/`. The files under `grammar/` and `spec/` are read-only reference copies and must not be edited here.

Use [Task](https://taskfile.dev/) from the repository root:

| Command         | Action |
|-----------------|--------|
| `task generate` | Generate ANTLR Go files from `GRAMMAR_BASE` (currently `grammar/v1.0.0-rc.7/`) into `internal/grammar/generated/`. |
| `task build`    | Compile the library and command packages with `go build ./...`. |
| `task fmt`      | Format Go packages. |
| `task vet`      | Analyze Go packages. |
| `task test`     | Run all tests. |
| `task check`    | Run format, build, vet, and test. |
| `task release`  | Run checks and create an annotated Git tag from `VERSION`. |

`task generate` uses the included `libs/antlr4/antlr-4.13.2-complete.jar` and requires Java. Generated files are checked in, so consumers need only Go. `go.mod` pins the separate ANTLR Go runtime. Run `task generate` only when generation inputs or settings change, and run `task check` before submitting work. `task vet` disables the unreachable-code analyzer because ANTLR emits unreachable branches.

`task build` verifies compilation. This repository publishes a Go library, so it does not need a distributable build directory; Go applications compile it when they build. Use an explicit `go build -o` command below to produce the standalone test adapter.

`VERSION` is the parser release version. Commit it with the release changes before running `task release`. The task creates a local tag; push that tag to make the version available to Go users.

Integration tests and original YINI fixtures live in `tests/integration/`. Add focused valid and invalid cases when changing semantic validation; assert line and column for source-location fixes. New fixtures must use original YINI content rather than copying content from sibling parser repositories.

## Shared conformance suite

The Go adapter for `yini-test-suite` lives in `cmd/yini-test-adapter/`. Install the same pinned version as CI, build the adapter, and run the smoke suite in both modes:

```powershell
python -m pip install yini-test-suite==0.4.0rc1
go build -o yini-test-adapter.exe ./cmd/yini-test-adapter
yini-test-suite smoke --all-modes --adapter ./yini-test-adapter.exe --input '{input}' --mode '{mode}'
```

The adapter writes parsed JSON to stdout and diagnostics to stderr. Replace `smoke` with `all` to run the full suite. All 542 cases in `0.4.0rc1` pass in both modes: 20 smoke cases and 522 golden cases. CI runs both the smoke and full conformance workflows.

---

^YINI ≡
