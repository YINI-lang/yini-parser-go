# Development

This is a Go module with the importable `yini` package at its root. Go does not give `/src` special meaning; putting public Go files there would change the import path. Generated parser files are private in `internal/grammar/generated/`. The files under `grammar/` and `spec/` are read-only reference copies and must not be edited here.

Use [Task](https://taskfile.dev/) from the repository root:

| Command         | Action |
|-----------------|--------|
| `task generate` | Generate ANTLR Go files from `GRAMMAR_BASE` (currently `grammar/v1.0.0-rc.7/`) into `internal/grammar/generated/`. |
| `task build`    | Compile all Go packages and build the test adapter in the repository root. |
| `task clean`    | Remove local Go build artifacts and root-level test-adapter binaries. |
| `task fmt`      | Format Go packages. |
| `task vet`      | Analyze Go packages. |
| `task test`     | Run all tests. |
| `task test:verbose` | Run all tests without caching and show each test and subtest result. |
| `task test:summary` | Run all tests without caching and show passed, failed, and skipped counts by package. |
| `task check`    | Run format, build, vet, and test. |
| `task release`  | Run checks and create an annotated Git tag from `VERSION`. |

`task generate` uses the included `libs/antlr4/antlr-4.13.2-complete.jar` and requires Java. Generated files are checked in, so consumers need only Go. `go.mod` pins the separate ANTLR Go runtime. Run `task generate` only when generation inputs or settings change, and run `task check` before submitting work. `task vet` disables the unreachable-code analyzer because ANTLR emits unreachable branches.

`task build` compiles all packages and produces the standalone test adapter in the repository root: `yini-test-adapter.exe` on Windows or `yini-test-adapter` on Linux and macOS. The library is compiled by applications that import it.

`task clean` runs `go clean ./...` and removes `yini-test-adapter` and `yini-test-adapter.exe` from the repository root. Missing files are harmless. It preserves source files, generated parser source, grammar references, fixtures, vendored tooling, and Go's shared caches. It runs from the repository root even when called from a subdirectory.

`VERSION` is the parser release version. Commit it with the release changes before running `task release`. The task creates a local tag; push that tag to make the version available to Go users.

Focused unit tests use `package yini` and live beside the implementation files at the module root. They test private helpers for numeric and string conversion, identifiers, section headers, source positions, and deep copies. Integration tests and original YINI fixtures live in `tests/integration/` and cover public APIs and parse modes. Both are included in `task test` and `task check`.

`task test:summary` counts test functions and named subtests, including parent groups. It reports their totals separately below the table; these are reported test results, not a count of independent cases. Package durations come from Go, and the total duration is the time spent running `go test`, including compilation. Failed packages also print their detailed output, and any test or build failure makes the task fail. Go reports results by test name rather than by source file.

Add focused valid and invalid cases when changing semantic validation; assert line and column for source-location fixes. New fixtures must use original YINI content rather than copying content from sibling parser repositories. See [doc.go](../doc.go) for the package overview and source organization.

## Shared conformance suite

The Go adapter for `yini-test-suite` lives in `cmd/yini-test-adapter/`. Install the same pinned version as CI, build the adapter, and run the smoke suite in both modes:

```powershell
python -m pip install yini-test-suite==0.4.0rc1
task build
yini-test-suite smoke --all-modes --adapter ./yini-test-adapter.exe --input '{input}' --mode '{mode}'
```

The adapter writes parsed JSON to stdout and diagnostics to stderr. Replace `smoke` with `all` to run the full suite. All 542 cases in `0.4.0rc1` pass in both modes: 20 smoke cases and 522 golden cases. CI runs both the smoke and full conformance workflows.

---

^YINI ≡
