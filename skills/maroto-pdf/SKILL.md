---
name: maroto-pdf
description: Generate PDFs from Go code with the maroto v2 library — install, program skeleton, page config, grid layout, components (text, image, barcode, QR code, signature), tables and testing. Use when writing or extending Go code that builds a PDF (invoices, reports, certificates) with github.com/johnfercher/maroto.
---

# Maroto PDF

Skill for LLMs and coding agents (Claude, Copilot, Cursor, …) that **use maroto** to generate PDFs
from Go code. It is written for people building applications with the library, not for people
developing maroto itself — contributor conventions live in [`.github/skills`](https://github.com/johnfercher/maroto/blob/master/.github/skills).

Start at [generation.md](generation.md); it routes to the others.

| File | Use when |
|------|----------|
| [generation.md](generation.md) | Writing any Go code that generates a PDF — install, program skeleton, output, errors, checklist. **Entry point.** |
| [config.md](config.md) | Choosing `config.NewBuilder()` options: page size, margins, fonts, page numbers, metadata, protection, generation mode |
| [layout.md](layout.md) | Placing content: grid columns, fixed/auto rows, cell styles, headers, footers, page breaks |
| [components.md](components.md) | Using text, images, barcodes, QR codes, data matrices, lines, checkboxes, signatures and their props |
| [tables.md](tables.md) | Rendering tabular data with `list.Build` or manual loops, striping, borders, pagination |
| [testing.md](testing.md) | Unit-testing the component tree with `pkg/test` fixtures and debugging layout problems |

This skill is also rendered on the documentation site:
[maroto.tech → Agent Skills](https://maroto.tech/#/v2/skills?id=agent-skills).

## How to use it in your project

1. Copy the whole `skills/maroto-pdf/` folder into your project, in the directory your agent
   loads skills from — the folder name must stay `maroto-pdf`, matching the `name` in the frontmatter
   above. The agent then discovers `SKILL.md` and loads it whenever you ask for PDF generation:

   ```bash
   # Claude Code
   cp -r skills/maroto-pdf <your-project>/.claude/skills/
   # Codex, Cursor and other agents that read .agents/skills
   cp -r skills/maroto-pdf <your-project>/.agents/skills/
   ```

2. Each file links to the others for the decisions it doesn't cover.
3. Pin the maroto version the skill describes: it documents the public API of the release it ships
   with. If you upgrade the library, pull the matching folder again.

Every Go snippet in these files compiles against the module version in this repository, and every
default or clamp quoted comes from the library's source, so an agent can treat them as ground truth
rather than re-deriving behaviour from the code.

## Keeping it accurate

Contributors to maroto must update the affected file in the same pull request as any user-facing
change — see [contribution.md §6](https://github.com/johnfercher/maroto/blob/master/.github/skills/contribution.md#6-skillspdf-library-usage-skills).
