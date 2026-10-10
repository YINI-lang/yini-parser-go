# AGENTS.md

> AI agent instructions for this repository.
>
> Read this before making any changes to the codebase.
>
> If any instruction in this file is unclear, ambiguous, or conflicts with the repository state, do not proceed with changes that depend on that instruction. Ask the human maintainer for clarification first.

See also: [Shared AI agent instructions for the YINI project family](../AGENTS.md)

## Project Overview

- **Name:** `yini-parser` (`yini-parser-go` repository).
- **Purpose:** Official GO-lang parser for YINI, a human-readable configuration format with explicit structure and predictable parsing.

## Repository Structure

```text
--TODO--
```

## Commands

--TODO--

## Required Checks

--TODO--

## Code Style

Follow the existing GO-lang style in nearby files.

- Keep changes small and focused.
- Prefer explicit, readable logic over compact cleverness.
- Preserve public APIs unless the task explicitly asks for an API change.
- Add or update tests when behavior changes.
- Do not reformat unrelated files.
- Do not rewrite unrelated parser, grammar, or fixture content.
- Keep comments useful and sparse; explain non-obvious parser behavior, not obvious assignments.

Naming conventions:
- Follow common GO-lang naming conventions.

## Strict vs Lenient Behavior

Default parsing is lenient.

- Lenient mode: `/END` is optional, duplicate keys/sections generally warn/recover according to options, and some loose syntax is accepted.
- Strict mode: `/END` is required by default, stricter structure rules apply, and strict-only syntax errors should be precise.
- `requireDocTerminator`, `onDuplicateKey`, and `treatEmptyValueAsNull` can override parts of the selected mode; this may make the effective mode `custom`.
- `failLevel` controls when parsing bails.
- `throwOnError` controls whether a bail-threshold issue throws.
- Library code should not write diagnostics unless `logDiagnostics` is enabled.
- `parseForTooling` should return structured diagnostics and avoid throwing for normal parse errors.

When fixing parser behavior, test both the intended valid path and the nearest invalid path. Prefer line/column assertions for diagnostic-location bugs.

## Grammar and Generated Files

Current grammar source is under `grammar/v1.0.0-rc.6/`.

ANTLR-generated files live in `src/grammar/generated/`. Do not hand-edit generated files unless the user explicitly asks for an emergency generated-output patch. Prefer editing grammar source and regenerating.

Do not edit files under `libs/` unless the task is specifically about vendored tooling.

## Testing Guidance

- Put focused regressions near the affected behavior.
- Use `tests/fixed-issues/` for named historical bug regressions.
- Use `tests/integration/` for parser features and public option behavior.
- Use `tests/smoke/` for broad end-to-end behavior.
- Use `tests/golden/` when parsing output for fixtures changes intentionally.
- Keep tests deterministic and avoid broad snapshot/golden churn unless the parser behavior intentionally changed.
- For file-location bugs, include fixtures or inline content that proves exact line/column behavior.

## Documentation Guidance

Update docs/examples/changelog when a change affects:
- public API,
- parser options,
- strict vs lenient behavior,
- diagnostics or error messages,
- package usage examples,
- generated grammar behavior visible to users.

Keep documentation concise and consistent with the README tone. Examples should be valid for the mode they claim to demonstrate.

## Dependency Policy

Do not add runtime dependencies unless clearly necessary.

Before adding any dependency, prefer:

1. Existing parser helpers or utilities.
2. GO-lang standard library functionality.
3. A small local helper.

## Safety and Scope Boundaries

### Always Do

- Run tests before submitting any change.
- Match the code patterns in the file you are editing.
- Keep changes focused — one concern per PR.
- When editing Markdown files, if a line introduces a bulleted list and ends with a colon (`:`), place the first bullet immediately on the next line. Do not insert a blank line between the introductory line and the first bullet.

### Ask First

- Before adding a new dependency.
- Before changing the public API or exported types.
- Before modifying CI/CD configuration.
- Before refactoring shared utilities used across multiple modules.

### Never Do

Do not modify:
- `.env` files, secrets, credentials, or keys,
- vendored dependencies under `libs/`,
- generated parser files unless grammar generation is part of the task,
- `dist/` by hand,
- lockfiles unless dependency changes require it,
- unrelated formatting, fixtures, or metadata.

Do not perform destructive operations such as:
- `git reset --hard`,
- force-pushing,
- publishing packages,
- creating releases,
- deleting large parts of the repository.

Do not create commits, tags, branches, package releases, or npm publishes unless explicitly requested.

## Project-Specific Priorities

YINI values clarity, readability, predictability, explicit structure, deterministic parsing, and precise diagnostics.

When changing code, tests, examples, or documentation:
- Prefer clarity over cleverness.
- Avoid implicit or magical behavior.
- Keep syntax examples human-readable.
- Preserve strict-mode and lenient-mode separation.
- Make parser behavior match the current YINI specification.
- Prefer structured diagnostics over vague errors.
- Keep source line/column reporting aligned with original user input, including BOM and shebang cases.
- Treat public types and options as user-facing contracts.
