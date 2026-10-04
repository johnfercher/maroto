# Code Quality

Standards for keeping code in this repository maintainable: what the static analyzer already
enforces, and the SOLID design principles this codebase follows in practice.

See also:
- [contribution.md](contribution.md) for the full PR checklist, including `make lint`.
- [unit-tests.md](unit-tests.md) and [mocks.md](mocks.md) for testing conventions that make the
  substitutability described below actually verifiable.

## Instructions

## 1. golangci-lint is the static reviewer

Run `make lint` (or `go tool -modfile=tools/go.mod golangci-lint run --config=.golangci.yml ./...`
directly) before treating a change as done — see [contribution.md](contribution.md#8-definition-of-done)
for how this fits into `make dod`. Don't hand-review for things the linter already checks
automatically; read `.golangci.yml` instead of guessing what's enabled.

Key facts about this repo's config:

- `linters.default: all` — every golangci-lint linter is on unless explicitly listed under
  `disable:`. Treat the disable list, not intuition, as the source of truth for what's allowed.
- Complexity/dead-code/duplication linters relevant to the principles below are **enabled**:
  `gocyclo`, `gocognit`, `nestif`, `maintidx`, `dupl`, `unused`, `unparam`, `prealloc`. A function
  that's hard to test in isolation, or a block copy-pasted across components, will get flagged
  here before it ever reaches review.
- `interfacebloat` is enabled with `settings.interfacebloat.max: 15` — an interface with more than
  15 methods fails lint. This is the automated backstop for Interface Segregation (§2.4 below).
- `cyclop` and `funlen` are **disabled** — this repo doesn't cap cyclomatic complexity or function
  length by line count; `gocyclo`/`gocognit`/`maintidx` cover complexity instead.
- `ireturn` is **disabled** — deliberately, because this codebase's component constructors return
  interfaces (`core.Component`, `core.Col`, `core.Row`) by design; see §2.3/§2.5. Don't "fix" a
  constructor to return a concrete type instead — that would fight the architecture, not improve it.
- `gochecknoglobals` is **disabled** — package-level vars (error sentinels, singletons like
  `pkg/test/test.go`'s `configSingleton`) are an accepted pattern here.
- `wrapcheck` ignores this module's own error returns (`ignore-package-globs:
  github.com/johnfercher/maroto/*`) — only unwrapped errors from *external* dependencies are flagged.
- Lint doesn't run on `docs/*`, `cmd/*`, or `_test.go` files (`exclusions.paths`).
- `formatters.enable: gofmt, gofumpt, goimports` — matches `make fmt` exactly.

If a linter flags something and the fix would go against a pattern documented below (e.g.
`interfacebloat` on `core.Provider`), that's a deliberate, already-reviewed exception — see the
`// nolint:interfacebloat` comment on `pkg/core/provider.go:2` before adding a new one elsewhere.

## 2. SOLID principles, with real examples from this codebase

### 2.1 Single Responsibility Principle

Every component package under `pkg/components/` does exactly one visual concern and delegates the
actual PDF drawing to the provider. `pkg/components/checkbox/checkbox.go` (79 lines, whole file)
is a clean example — it only holds a label/prop/config and renders by delegating:

```go
// Render renders a Checkbox into a PDF context.
func (c *Checkbox) Render(provider core.Provider, cell *entity.Cell) {
	provider.AddCheckbox(c.label, cell, &c.prop)
}
```

`text.go`, `line.go`, and every other file in `pkg/components/*` follow the same shape. When
adding a new component, keep this split: the component package owns *what* to draw and its
config; `core.Provider` owns *how* to draw it. Don't let a component reach into PDF internals
directly, and don't let unrelated concerns (e.g. validation logic for an unrelated prop) creep
into a component that isn't responsible for them.

### 2.2 Open/Closed Principle

`pkg/components/col/col.go`'s `Render` loop is agnostic to which components it holds:

```go
// Render renders a core.Col into a PDF context.
func (c *Col) Render(provider core.Provider, cell entity.Cell, createCell bool) {
	if createCell {
		provider.CreateCol(cell.Width, cell.Height, c.config, c.style)
	}

	for _, component := range c.components {
		component.Render(provider, &cell)
	}
}
```

Every existing component (`checkbox`, `text`, `image`, `line`, `code`, `signature`, ...) is a
sibling package that implements `core.Component`. Adding a brand-new component type requires zero
changes to `col.go`, `row.go`, or `page.go` — only a new package implementing the interface. When
building a new feature, prefer adding a new implementation of an existing interface over branching
on type inside shared rendering code (`col.go`, `row.go`, `page.go`).

### 2.3 Liskov Substitution Principle

The same loop above holds `components []core.Component` (`col.go:15`) — a slice of heterogeneous
concrete types (`*checkbox.Checkbox`, `*text.Text`, `*image.FileImage`, ...), all constructed via a
`func New(...) core.Component` in their own package (e.g. `pkg/components/text/text.go`). None of
them need special-casing in `Col.Render`, `Col.GetHeight`, or `Col.SetConfig` — every concrete type
behaves correctly wherever a `core.Component` is expected. If you add a component whose `Render` or
`GetHeight` needs the caller to know its concrete type to behave correctly, that breaks this
substitutability — keep behavior fully expressible through the interface's methods.

### 2.4 Interface Segregation Principle

`pkg/core/components.go` is full of small, single-purpose interfaces, e.g.:

```go
type Line interface {
	Add(cell *entity.Cell, prop *props.Line)
}
```

`Checkbox`, `Text`, `Code`, `Image`, `Math` in the same file are similarly narrow — each mockable
and testable independently (see the corresponding single-method mocks in `mocks/`).

Contrast this with `pkg/core/provider.go`'s `Provider` interface — 21 methods, and the file admits
the tension directly:

```go
// Package core contains all core interfaces and basic implementations.
// nolint:interfacebloat // there is no way to reduce this interface
package core
```

`.golangci.yml` caps interfaces at 15 methods (`interfacebloat.max: 15`). `Provider` isn't the
only exception — `internal/providers/gofpdf/gofpdfwrapper/fpdf.go:1` and `pkg/config/builder.go:2`
carry the same `// nolint:interfacebloat` comment, each for a different legitimate reason:
`gofpdfwrapper.Fpdf` wraps the external `gofpdf` package's own wide API one-to-one, and
`config.Builder` is a fluent builder whose method count mirrors the number of independent config
options, not accidental complexity. Treat all three as accepted, already-reviewed exceptions — not
precedent for adding more wide interfaces casually. When designing a new abstraction, default to
something the size of `Line`/`Checkbox`/`Text`, not the size of `Provider`.

### 2.5 Dependency Inversion Principle

`Maroto` depends on the `core.Provider` abstraction, not on the concrete `gofpdf` implementation:

```go
type Maroto struct {
	config   *entity.Config
	provider core.Provider
	cache    cache.Cache
	...
}
```

The one place that knows about the concrete `gofpdf` type is the composition root:

```go
func getProvider(cache cache.Cache, cfg *entity.Config) core.Provider {
	deps := gofpdf.NewBuilder().Build(cfg, cache)
	provider := gofpdf.New(deps)
	provider.SetMetadata(cfg.Metadata)
	provider.SetCompression(cfg.Compression)
	provider.SetProtection(cfg.Protection)
	return provider
}
```

Every component, and every component test, depends only on `core.Provider` — proven by how freely
`mocks.NewProvider(t)` substitutes for the real `gofpdf` provider throughout `pkg/components/*_test.go`
(see [mocks.md](mocks.md)). If a new feature needs a new external dependency, wire it in behind an
interface at this same boundary rather than importing the concrete package into component code.

## References

- golangci-lint config: [`.golangci.yml`](../../.golangci.yml)
- golangci-lint docs: https://golangci-lint.run/
