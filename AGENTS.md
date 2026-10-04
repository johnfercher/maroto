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
| [code-quality.md](.github/skills/code-quality.md) | Writing or reviewing non-test Go code — what `golangci-lint` enforces and the SOLID patterns this codebase follows |

### Generating PDFs with maroto (library usage)

A second group of skills teaches how to **use** maroto from Go code — for agents writing an
application that produces PDFs, or for agents extending the examples under `docs/assets/examples`.
They live in the [`.github/skills/pdf`](.github/skills/pdf) folder. Start with
`pdf/generation.md`; it routes to the others. [contribution.md §6](.github/skills/contribution.md#6-githubskillspdf-library-usage-skills)
requires them to be updated in the same PR as any user-facing change.

| Skill | Use when |
|-------|----------|
| [pdf/generation.md](.github/skills/pdf/generation.md) | Writing any Go code that generates a PDF with maroto — install, program skeleton, output, errors, checklist |
| [pdf/config.md](.github/skills/pdf/config.md) | Choosing `config.NewBuilder()` options: page size, margins, fonts (incl. UTF-8), page numbers, metadata, protection, generation mode |
| [pdf/layout.md](.github/skills/pdf/layout.md) | Placing content: grid columns, fixed/auto rows, cell styles, headers, footers, page breaks |
| [pdf/components.md](.github/skills/pdf/components.md) | Using text, images, barcodes, QR codes, data matrices, lines, checkboxes, signatures and their props |
| [pdf/tables.md](.github/skills/pdf/tables.md) | Rendering tabular data (invoices, reports) with `list.Build` or manual loops, striping, borders, pagination |
| [pdf/testing.md](.github/skills/pdf/testing.md) | Unit-testing the component tree with `pkg/test` fixtures and debugging layout problems |

As new skill files are added to `.github/skills`, treat them as mandatory for
the task they cover.

## Definition of Done

Before considering a change complete, follow [contribution.md](.github/skills/contribution.md),
which walks through the full [pull request checklist](pull_request_template.md) and ends with
running `make dod` (build + test + format + lint + codecov) with no `golangci-lint` issues.
`codecov` never fails the build, but its output (functions with 0% test coverage) is highly
recommended to address before opening the PR — see
[contribution.md §9](.github/skills/contribution.md#9-definition-of-done).

If the change adds, changes or removes anything a user of the library can observe (a component,
constructor, prop, default, config option, error or layout rule), the matching skill under
[`.github/skills/pdf/`](.github/skills/pdf) must be updated in the same PR — see
[contribution.md §6](.github/skills/contribution.md#6-githubskillspdf-library-usage-skills).

If `make install-hooks` has been run in this clone, `make dod` also runs automatically as a
`pre-commit` git hook (see [`.githooks/pre-commit`](.githooks/pre-commit)) and blocks the commit
on failure — expect `git commit` itself to take as long as `make dod` does, and to fail the
commit if build/test/fmt/lint don't pass (`codecov` output won't block the commit, but don't
ignore it).

## Contributing commands

See the [Contributing](README.md#contributing) section of the README for the
`make` targets used to build, test, format, lint, and generate mocks.

## Tool dependencies

`goimports`, `gofumpt`, `golangci-lint`, `mockery`, and `godoc` are declared as tool dependencies
in [`tools/go.mod`](tools/go.mod), not the main `go.mod` — this keeps the ~200 extra transitive
dependencies pulled in mostly by `golangci-lint` out of the library's own `go.sum`. Run them with
`go tool -modfile=tools/go.mod <name>` (already wired into the Makefile); never `go install` them
globally or add them to the root `go.mod`.
