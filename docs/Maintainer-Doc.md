# Maintainer Documentation

This document describes the maintainer workflow for building, checking, and publishing new releases of `yini-parser-go`.

The Go module and import path is `github.com/YINI-lang/yini-parser-go`. Go users obtain a release from a Git tag in the public repository. There is no separate package upload. A GitHub Release page is optional; it does not publish the Go module. See the [Go module publishing guide](https://go.dev/doc/modules/publishing).

---

## Release checklist

Before tagging a release, check that:
- `VERSION` contains the intended Go module version, such as `v1.0.0-alpha.2`.
- `CHANGELOG.md` has notes for that version, and the README badge and installation command name the same version.
- All intended release files are committed and the working tree is clean.
- `task check` passes.
  Note: If generated parser files or generation settings changed, run `task generate` first.

`VERSION` is the source of truth for the release version. It is not stored in `go.mod`; the `go` directive there sets the required Go version. The release task trims surrounding whitespace from `VERSION` and checks that its value matches the committed file.

---

## Create and publish a release

1. Update `VERSION`, the changelog, the README, and any other release documentation. Review the changes and run the project checks from the repository root:

   ```sh
   git status --short
   task check
   ```

2. Commit the intended release changes and push that commit to the release branch. Confirm that `git status --short` shows a clean working tree. The release task will refuse to tag uncommitted changes or a version that differs from the committed `VERSION` file.

3. Create the annotated tag locally:

   ```sh
   task release
   git tag --points-at HEAD
   ```

   `task release` runs `task check` again, then creates a local tag named by the trimmed `VERSION` value. It does not push the tag.

4. Push the exact tag printed by the task. For example, if `VERSION` contains `v1.0.0-alpha.2`:

   ```sh
   git push origin v1.0.0-alpha.2
   ```

   Pushing the tag makes that version available to Go tools. To publish release notes on GitHub, optionally create a GitHub Release from the existing tag and mark alpha, beta, or RC versions as pre-releases.

---

## Verify the published version

Request the version through the Go module proxy. Replace the example version with the tag just pushed:

```sh
go list -m github.com/YINI-lang/yini-parser-go@v1.0.0-alpha.2
task versions
```

The first command should return the exact module path and version. `task versions` lists versions visible to Go. A consumer can then run `go get github.com/YINI-lang/yini-parser-go@v1.0.0-alpha.2` and import the root `yini` package.

Never move a tag after users have fetched it. If the tagged code needs a fix, update `VERSION` and publish a new tag; Go verifies downloaded module content against the version it first observed. See the [Go module publishing guide](https://go.dev/doc/modules/publishing) for this rule.

---

^YINI ≡

