# Mocks

Generate and use [mockery](https://github.com/vektra/mockery) mocks for this project's interfaces,
following the conventions already in `mocks/`.

See also:
- [unit-tests.md](unit-tests.md) for the general test conventions (naming, AAA, parallelism)
  that apply to every test that uses a mock built with this skill.
- [contribution.md](contribution.md) for the full pull request checklist, including `make dod`
  which also runs `make mock-lint` to verify `mocks/` is up to date.

## Instructions

## 1. Generating mocks

- Mocks are generated with `mockery`, configured in `.mockery.yaml`:
  ```yaml
  all: True
  dir: "mocks/"
  outpkg: "mocks"
  filename: "{{.InterfaceName}}.go"
  mockname: "{{.InterfaceName}}"
  with-expecter: True
  packages:
    github.com/johnfercher/maroto/v2:
      config:
        recursive: True
  ```
- Never hand-write or hand-edit a file in `mocks/`. Run `make mocks` instead, which does:
  ```
  rm -R mocks || true
  go tool -modfile=tools/go.mod mockery
  make fmt
  ```
  `mockery` isn't installed globally — it's a versioned tool dependency declared in
  [`tools/go.mod`](../../tools/go.mod) (see [contribution.md](contribution.md) for why tool
  dependencies live in that separate module). `go tool -modfile=tools/go.mod mockery` builds and
  runs the exact version pinned there; there's no need to `go install` it yourself.
- `make mocks` regenerates **every** mock from scratch. Run it after adding, removing, or
  changing the signature of any exported interface — do not try to regenerate a single mock by hand.
- Every generated mock lives at `mocks/<InterfaceName>.go`, package `mocks`, and exposes:
  - A struct `<InterfaceName>` embedding `mock.Mock`.
  - A `<InterfaceName>_Expecter` returned by `(_m *<InterfaceName>) EXPECT()`.
  - One `<InterfaceName>_<Method>_Call` type per method, with `Run`, `Return`, and `RunAndReturn`.
  - A constructor:
    ```go
    func New<InterfaceName>(t interface {
        mock.TestingT
        Cleanup(func())
    },
    ) *<InterfaceName> {
        mock := &<InterfaceName>{}
        mock.Mock.Test(t)
        t.Cleanup(func() { mock.AssertExpectations(t) })
        return mock
    }
    ```
    This constructor already asserts expectations on cleanup — never call `mock.AssertExpectations(t)` yourself.

## 2. Using mocks in tests

- Create with the generated constructor: `dep := mocks.NewDependency(t)` — never `new(...)` or `&mocks.Dependency{}`.
- Set expectations with the type-safe EXPECT API, not `.On(...)`:
  ```go
  dep.EXPECT().MethodName(arg1, arg2).Return(val1, val2)
  ```
- Example from `internal/providers/gofpdf/font_test.go`:
  ```go
  fpdf := mocks.NewFpdf(t)
  fpdf.EXPECT().SetFont(family, string(style), size)

  font := gofpdf.NewFont(fpdf, size, family, style)
  ```
- Avoid `mock.Anything` — build the real expected argument whenever possible. It's only allowed
  when it's not possible to construct the argument (e.g. a reader/writer instance created
  internally). When you have to use it, warn the user and call it out explicitly in the commit
  and pull request:
  ```go
  // mock.Anything have to be avoided
  pdf.EXPECT().RegisterImageOptionsReader(mock.Anything, options, bytes.NewReader(img.Bytes)).Return(nil)
  ```
  (see `internal/providers/gofpdf/image_test.go`).
- Assert call count only when it matters, e.g. when the dependency is called inside a loop:
  ```go
  inner.AssertNumberOfCalls(t, "AddRows", 2)
  ```
  (see `metricsdecorator_test.go`).
- For methods expected **not** to be called, do not set an EXPECT — `testify/mock` fails
  automatically on an unexpected call. If a dependency ends up with zero EXPECT calls in a given
  subtest, don't pass the mock at all: pass `nil` to the `sut` constructor instead.

## References

- Mockery docs: https://vektra.github.io/mockery/latest/
- Testify mock docs: https://pkg.go.dev/github.com/stretchr/testify/mock
