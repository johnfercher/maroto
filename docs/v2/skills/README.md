# Agent Skills

Maroto ships a set of **skills**: Markdown instruction files that teach an LLM or coding agent
(Claude, Copilot, Cursor, Codex, …) how to write Go code that generates PDFs with maroto. They
live in the [`skills/`](https://github.com/johnfercher/maroto/tree/master/skills) folder of the
repository, apart from the contributor conventions in `.github/skills`, so you can copy them into
your own project.

Each skill is a short, prescriptive reference: the constructors to call, every prop with its real
default, the layout arithmetic, and the pitfalls an agent would otherwise rediscover by trial and
error (a fixed row never grows, the header must be registered before content, built-in fonts only
cover Latin-1, …). Every Go snippet compiles against the current module and every default quoted
comes from the library's source.

## The set

| Skill | Read it when |
|-------|--------------|
| [PDF Generation](v2/skills/generation.md?id=pdf-generation) | Starting any PDF code — install, mental model, program skeleton, output options, errors, a hand-over checklist. **Entry point**; it routes to the others. |
| [PDF Config](v2/skills/config.md?id=pdf-config) | Choosing `config.NewBuilder()` options: page size, margins, grid size, fonts (incl. UTF-8), page numbers, metadata, protection, compression, generation mode. |
| [PDF Layout](v2/skills/layout.md?id=pdf-layout) | Placing content: grid columns, fixed vs. auto rows and their height math, cell styles, headers, footers, page breaks. |
| [PDF Components](v2/skills/components.md?id=pdf-components) | Using text, images, barcodes, QR codes, data matrices, lines, checkboxes and signatures, with a props table for each. |
| [PDF Tables](v2/skills/tables.md?id=pdf-tables) | Rendering tabular data (invoices, reports) with `list.Build` or manual loops: column plans, striping, borders, repeating headers, variable-height cells. |
| [PDF Testing](v2/skills/testing.md?id=pdf-testing) | Unit-testing the component tree with `pkg/test` fixtures and debugging layout problems. |

## How to use them

1. **Copy or reference.** Copy `skills/pdf/` into the directory your agent reads instructions from
   (`skills/`, `.claude/skills/`, `.cursor/rules/`, a repository `AGENTS.md`, …), or give the agent
   the GitHub URLs of the files.
2. **Point the agent at the entry file.** Ask it to read
   [`skills/pdf/generation.md`](https://github.com/johnfercher/maroto/blob/master/skills/pdf/generation.md)
   before writing code; it links to the other five for the decisions it doesn't cover.
3. **Pin the version.** The skills describe the public API of the maroto release they ship with.
   When you upgrade `github.com/johnfercher/maroto/v2`, pull the matching `skills/` folder again.

A minimal prompt that works with most tools:

```text
Read skills/pdf/generation.md and the skills it links to, then write a Go program that
generates <describe the document> with maroto, following those skills exactly.
```

## Keeping them accurate

The skills are part of maroto's definition of done: any pull request that adds, changes or removes
a user-facing feature must update the affected skill in the same change (see
[`.github/skills/contribution.md` §6](https://github.com/johnfercher/maroto/blob/master/.github/skills/contribution.md#6-skillspdf-library-usage-skills)).
If you find a sentence that no longer matches the library, open an issue or a pull request against
the file under `skills/pdf/`.

The pages in this section render the skill files straight from the repository's `master` branch.
