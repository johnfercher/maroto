# Skills

Skills for LLMs and coding agents (Claude, Copilot, Cursor, …) that **use maroto** to generate PDFs
from Go code. They are written for people building applications with the library, not for people
developing maroto itself — contributor conventions live in [`.github/skills`](../.github/skills).

| Folder | Contents | Start at |
|--------|----------|----------|
| [`pdf/`](pdf) | Generating PDFs with maroto: setup, config, grid layout, components, tables, testing | [`pdf/generation.md`](pdf/generation.md) |

## How to use them in your project

1. Copy the folder you need (e.g. `skills/pdf/`) into the place your agent reads instructions
   from — a `skills/`, `.claude/skills/` or `.cursor/rules/` directory, or wherever your tool
   expects it — or reference the files by URL.
2. Tell the agent to read the entry file first (`pdf/generation.md`); each file links to the others
   for the decisions it doesn't cover.
3. Pin the maroto version the skills describe: they document the public API of the release they ship
   with. If you upgrade the library, pull the matching `skills/` folder again.

Every Go snippet in these files compiles against the module version in this repository, and every
default or clamp quoted comes from the library's source, so an agent can treat them as ground truth
rather than re-deriving behaviour from the code.

## Keeping them accurate

Contributors to maroto must update the affected skill in the same pull request as any user-facing
change — see [contribution.md §6](../.github/skills/contribution.md#6-skillspdf-library-usage-skills).
