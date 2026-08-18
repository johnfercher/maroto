# Contribution

Prepare a change in this repository (branch, code, docs, PR) so it satisfies the
[pull request checklist](../../pull_request_template.md) before it's opened for review.

See also:
- [unit-tests.md](unit-tests.md) for how to write the unit tests this checklist requires.
- [mocks.md](mocks.md) for generating/using mocks referenced by the checklist.

## Instructions

Apply every item below that is relevant to the change being made — the PR checklist itself is
opt-in per item ("check ONLY IF APPLIED"), but if an item applies, follow it; don't skip it
because it's inconvenient.

### 1. Branch naming
Name the branch after the kind of change:
- `feature/<name>` for new functionality.
- `fix/<name>` for bug fixes.

### 2. Method receiver naming style
Every method on a struct uses the struct's first letter (lowercase) as the receiver name:
```go
func (s *Font) GetFamily() string {
```
(see `internal/providers/gofpdf/font.go`). Do not use `this`, `self`, or the full struct name.

### 3. Unit tests
- Write unit tests for every new or changed exported function/method.
- Follow [unit-tests.md](unit-tests.md) for naming (`when <condition>, should <outcome>`), AAA
  comments, parallelism, and the `sut` variable name.
- Follow [mocks.md](mocks.md) for constructing mocks (`m := mocks.NewConstructor(t)`) and setting
  expectations (`m.EXPECT().MethodName(...)`).

### 4. Comments on new public API
Every new exported struct, interface, or method gets a comment above it explaining its
responsibility, starting with the identifier name. This is required because these comments are
what `godoc`/`pkg.go.dev` renders as the package documentation (see `make godoc` and the
[GoDoc](https://pkg.go.dev/github.com/johnfercher/maroto/v2) badge in `README.md`) — undocumented
exported identifiers show up blank there:
```go
// NewFont create a Font.
func NewFont(...) *Font { ... }

// GetFamily return the currently Font family configured.
func (s *Font) GetFamily() string { ... }
```
(see `internal/providers/gofpdf/font.go`).

### 5. `docs/*`
`docs/` is a [docsify](https://docsify.js.org/) site. Run `make docs` to preview it locally.

- Feature pages live at `docs/v2/features/<name>.md`, one per component/config option, and are
  listed in `docs/v2/features/_sidebar.md` — **a new feature page must be added there too**, or it
  will exist but be unreachable from the site nav.
- Follow the structure of existing pages, e.g. `docs/v2/features/checkbox.md`:
  1. `# <Title>` and a short description of the component/feature.
  2. `## Props (<props.Type>)` — a table of every prop field, its type, default, and description
     (skip this section for features with no props).
  3. `## Usage notes` — constraints, clamping behavior, gotchas.
  4. `## GoDoc` — links to `https://pkg.go.dev/...` for the constructor(s), props, and component.
  5. `## Code Example` — a docsify include of the runnable example:
     `[filename](../../assets/examples/<name>/v2/main.go ':include :type=code')`
  6. `## PDF Generated` / `## Time Execution` — includes of the generated
     `assets/pdf/<name>v2.pdf` and `assets/text/<name>v2.txt`.
  7. `## Test File` — a docsify include of the matching fixture under `test/maroto/examples/`.
- The runnable example referenced above lives at `docs/assets/examples/<name>/v2/main.go` (plus a
  `main_test.go` if it's also used as a unit test fixture). It must call `document.Save(...)` and
  `document.GetReport().Save(...)` to produce `docs/assets/pdf/<name>v2.pdf` and
  `docs/assets/text/<name>v2.txt`, and must be added to the `examples` target in the `Makefile`
  (`go run docs/assets/examples/<name>/v2/main.go`) so `make examples` regenerates it.
- If the change affects behavior documented on an *existing* page instead of adding a new
  feature, update that page (and its example under `docs/assets/examples/`) in place rather than
  creating a new one.

### 6. `example_test.go`
If the change adds or changes a public entry point that's useful to demonstrate, add or update an
`Example<Type>_<Method>` function in `example_test.go`:
```go
// ExampleMaroto_AddPages demonstrates how to add a new page in maroto.
func ExampleMaroto_AddPages() {
    ...
}
```

### 7. `README.md`
Update `README.md` if the change affects installation, the `make` command table, or anything else
described there.

### 8. Definition of Done
Before opening the PR, run:
```
make dod
```
which runs, in order: `make build`, `make test`, `make fmt`, `make lint` (the last of which also
runs `make mock-lint` to verify `mocks/` is up to date with `mockery`). Fix every issue
`golangci-lint` points out — the PR checklist requires this to have zero issues.

If the `pre-commit` git hook is installed (`make install` or `make install-hooks`, see
[`.githooks/pre-commit`](../../.githooks/pre-commit)), `make dod` already runs automatically
before every commit and blocks it on failure — a commit that succeeded locally has already
satisfied this item.

`make fmt`, `make lint`, `make mocks`, and `make godoc` don't rely on globally installed
binaries. `goimports`, `gofumpt`, `golangci-lint`, `mockery`, and `godoc` are declared as tool
dependencies in a separate [`tools/go.mod`](../../tools/go.mod) (a nested module, isolated so its
~200 extra transitive dependencies — mostly from `golangci-lint`'s linter engine — never touch the
main module's `go.sum`, which every consumer of this library downloads). The Makefile invokes them
with `go tool -modfile=tools/go.mod <name>`, which builds and runs the exact version pinned in
`tools/go.sum` on first use — no `go install`/`sudo cp` needed. If you need to add or upgrade one
of these tools, run `go get -tool <module>@<version>` **from inside `tools/`**, never from the
repo root (that would add it to the main `go.mod` instead).

### 9. Description and related issue
When opening the PR, fill in:
- **Description**: how the PR is useful, and any tricky technical detail.
- **Related Issue**: a reference to the issue it closes/relates to, if any.

## References

Full checklist for reference: [`pull_request_template.md`](../../pull_request_template.md)
