# Agent Instructions

This file guides any LLM, agent, or AI assistant (Claude, Copilot, Cursor, etc.)
working on this repository. If you are such a tool, read this file first.

## Skills

Standards and step-by-step conventions for common tasks live in
[`.github/skills`](.github/skills). Before performing a task covered by one of
these skills, read the matching file and follow it exactly instead of
improvising:

| Skill | Use when |
|-------|----------|
| [unit-tests.md](.github/skills/unit-tests.md) | Writing or regenerating `_test.go` files |
| [mocks.md](.github/skills/mocks.md) | Generating mocks with `mockery` or writing mock expectations in a test |
| [contribution.md](.github/skills/contribution.md) | Preparing any change (branch, code, docs, PR) to meet the pull request checklist |

As new skill files are added to `.github/skills`, treat them as mandatory for
the task they cover.

## Definition of Done

Before considering a change complete, follow [contribution.md](.github/skills/contribution.md),
which walks through the full [pull request checklist](pull_request_template.md) and ends with
running `make dod` (build + test + format + lint) with no `golangci-lint` issues.

If `make install-hooks` has been run in this clone, `make dod` also runs automatically as a
`pre-commit` git hook (see [`.githooks/pre-commit`](.githooks/pre-commit)) and blocks the commit
on failure — expect `git commit` itself to take as long as `make dod` does, and to fail the
commit if build/test/fmt/lint don't pass.

## Contributing commands

See the [Contributing](README.md#contributing) section of the README for the
`make` targets used to build, test, format, lint, and generate mocks.

## Tool dependencies

`goimports`, `gofumpt`, `golangci-lint`, `mockery`, and `godoc` are declared as tool dependencies
in [`tools/go.mod`](tools/go.mod), not the main `go.mod` — this keeps the ~200 extra transitive
dependencies pulled in mostly by `golangci-lint` out of the library's own `go.sum`. Run them with
`go tool -modfile=tools/go.mod <name>` (already wired into the Makefile); never `go install` them
globally or add them to the root `go.mod`.
