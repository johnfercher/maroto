# PDF Config

Configure a maroto document globally with `config.NewBuilder()`: page geometry, fonts, page
numbers, metadata, protection, compression, background and generation mode.

See also:
- [pdf-generation.md](pdf-generation.md) for where the config is consumed (`maroto.New(cfg)`).
- [pdf-layout.md](pdf-layout.md) for what margins, grid size and orientation mean for rows/cols.
- [pdf-components.md](pdf-components.md) for the per-component props that override these defaults.

## Instructions

### 1. Shape

```go
func buildConfig() *entity.Config {
	return config.NewBuilder().
		WithPageSize(pagesize.A4).
		WithOrientation(orientation.Vertical).
		WithLeftMargin(15).
		WithRightMargin(15).
		WithTopMargin(15).
		WithBottomMargin(20).
		WithDefaultFont(&props.Font{Family: fontfamily.Helvetica, Size: 10}).
		WithPageNumber(props.PageNumber{Pattern: "Page {current} of {total}", Place: props.RightBottom}).
		WithTitle("Monthly Report", true).
		WithAuthor("Reporting Service", true).
		WithCompression(true).
		Build()
}
```

- Every `With*` returns the `Builder`, so chain and finish with `Build()`.
- `Build()` returns `*entity.Config`. Pass it to `maroto.New(cfg)`; read it back later with
  `m.GetCurrentConfig()` (useful for `cfg.Dimensions`, `cfg.Margins`, `cfg.MaxGridSize`).
- **Invalid values are silently ignored**, not rejected: a negative margin, a zero dimension, an
  empty page size, a nil font, `chunkWorkers < 1`. Don't rely on an error to tell you a value
  was dropped — check the resulting `*entity.Config` in a test if it matters.

### 2. Defaults (what you get from `maroto.New()` with no config)

| Setting          | Default                                     |
|------------------|---------------------------------------------|
| Page size        | `pagesize.A4` (210 × 297 mm), portrait      |
| Margins          | left 10, right 10, top 10, bottom 20.0025 mm |
| Grid             | 12 units                                    |
| Font             | `fontfamily.Arial`, `fontstyle.Normal`, 10 pt, black |
| Page number      | none                                        |
| Generation mode  | sequential, 1 worker                        |
| Compression      | off                                         |
| Debug            | off                                         |

Useful area on default A4: `190 × 266.9975` mm.

### 3. Page geometry

| Method                                   | Notes                                                                 |
|------------------------------------------|-----------------------------------------------------------------------|
| `WithPageSize(pagesize.Type)`            | `A1`–`A6`, `Letter`, `Legal`, `Tabloid`. Dimensions in `pagesize.GetDimensions`. |
| `WithDimensions(width, height)`          | Arbitrary mm. **Overrides** `WithPageSize` whichever is called first. Both must be `> 0`. |
| `WithOrientation(orientation.Horizontal)`| Swaps width/height **only for `WithPageSize` sizes**. With `WithDimensions` pass width/height already in the desired orientation. |
| `WithLeftMargin(x)` …`WithBottomMargin(x)` | mm, each `>= 0`. There is no `WithMargins`; call the four methods. |
| `WithMaxGridSize(n)`                     | Number of grid units per row (default 12). Pick `n` divisible by the column counts you need, e.g. 20 for 5 or 10 equal columns. |

Labels, receipts, cards: `WithDimensions(100, 150)` plus small margins. Pages are always the same
size within one document — generate separately and `merge.Bytes` for mixed sizes.

### 4. Page numbers

```go
cfg := config.NewBuilder().
	WithPageNumber(props.PageNumber{
		Pattern: "Page {current} of {total}",
		Place:   props.Bottom,
		Family:  fontfamily.Courier,
		Style:   fontstyle.Bold,
		Size:    9,
		Color:   &props.Color{Red: 100, Green: 100, Blue: 100},
	}).
	Build()
```

- `WithPageNumber()` with no argument gives `"{current} / {total}"` at `props.Bottom`.
- A `Pattern` missing either `{current}` or `{total}` is **replaced** by the default pattern.
- `Place`: `props.LeftTop`, `props.Top`, `props.RightTop`, `props.LeftBottom`, `props.Bottom`,
  `props.RightBottom` (invalid → `props.Bottom`). Drawn in the margin area, not the grid.
- Unset font fields fall back to the default font.

### 5. Fonts

**Default font** — merges only the non-zero fields you pass:

```go
cfg := config.NewBuilder().
	WithDefaultFont(&props.Font{
		Family: fontfamily.Helvetica,
		Size:   11,
		Color:  &props.Color{Red: 30, Green: 30, Blue: 30},
	}).
	Build()
```

Built-in families: `fontfamily.Arial`, `Helvetica`, `Courier`, `Symbol`, `ZapBats`. They are
PDF core fonts encoded as cp1252: accented Latin characters work, **anything else (CJK, Arabic,
Cyrillic, Greek, Hebrew, emoji) renders as garbage or blanks**.

**Custom UTF-8 fonts** — required for other scripts, and for any branded typography. Register
each style you intend to use; an unregistered style falls back to the `Normal` file:

```go
func configWithUTF8Font(fontFile string) (*entity.Config, error) {
	const family = "noto-sans"

	fonts, err := fontrepository.New().
		AddUTF8Font(family, fontstyle.Normal, fontFile).
		AddUTF8Font(family, fontstyle.Bold, fontFile).
		AddUTF8Font(family, fontstyle.Italic, fontFile).
		AddUTF8Font(family, fontstyle.BoldItalic, fontFile).
		Load()
	if err != nil {
		return nil, err
	}

	cfg := config.NewBuilder().
		WithCustomFonts(fonts).
		WithDefaultFont(&props.Font{Family: family}).
		Build()

	return cfg, nil
}
```

- `AddUTF8Font(family, style, path)` reads a `.ttf` from disk at `Load()` time;
  `AddUTF8FontFromBytes(family, style, bytes)` takes bytes (pair with `//go:embed`).
- Use the same `family` string in `props.Text{Family: ...}` to select it per component.
- Fonts are embedded in the PDF: self-contained output, larger file. Only `.ttf` is supported.
- `WithDefaultFont` and `WithCustomFonts` are independent: registering a font does not make it
  the default.

### 6. Metadata, protection, compression

| Method                                                    | Notes                                                      |
|-----------------------------------------------------------|------------------------------------------------------------|
| `WithTitle(s, isUTF8)`, `WithAuthor`, `WithSubject`, `WithCreator`, `WithKeywords` | Document-info dictionary; empty strings are ignored. Pass `isUTF8 = true` unless the text is pure ASCII. |
| `WithCreationDate(time.Time)`                             | Zero time is ignored. Pass a fixed time in tests for reproducible bytes. |
| `WithProtection(protection.Type, userPassword, ownerPassword)` | Flags combine with `\|`: `protection.Print \| protection.Copy` **restricts** printing and copying. `protection.None` with a user password = open-password only. Empty string disables that password. RC4 128-bit via gofpdf: a deterrent, not a security boundary. |
| `WithCompression(true)`                                   | zlib on content streams. Cheap win for text-heavy documents; little effect on image-heavy ones. |

### 7. Background image and page breaks

```go
import _ "embed"

//go:embed letterhead.png
var letterhead []byte

func configWithLetterhead() *entity.Config {
	return config.NewBuilder().
		WithBackgroundImage(letterhead, extension.Png).
		WithTopMargin(40). // keep content clear of the printed header area
		Build()
}
```

- `WithBackgroundImage(bytes, extension.Png|Jpg|Jpeg)` stretches one image over every page,
  behind all content. Design the image at the page's aspect ratio.
- `WithDisableAutoPageBreak(true)` keeps everything on one logical page; content taller than the
  page runs off the bottom and is **not clipped**. Only use it for single-page templates whose
  height you control (certificates, labels), typically combined with `AddPages`.

### 8. Generation mode (performance)

| Method                                   | Behaviour                                                                    |
|------------------------------------------|------------------------------------------------------------------------------|
| `WithSequentialMode()`                   | Default. Whole document rendered by one provider.                            |
| `WithConcurrentMode(workers)`            | Pages split into chunks rendered by `workers` goroutines, then merged with pdfcpu. Fastest for 50+ pages; memory scales with workers. |
| `WithSequentialLowMemoryMode(chunks)`    | Chunks rendered one at a time and merged. Lowest peak memory, slightly slower. |

- The last mode method called wins.
- Both non-default modes return `maroto.ErrCannotGenerateInParallelMode` /
  `ErrCannotGenerateInLowMemoryMode` from `Generate()` if any chunk fails, and the merge step
  changes the PDF's internal structure (not its appearance).
- Start sequential. Switch only after measuring with `maroto.NewMetricsDecorator(m)` and
  `document.GetReport()`.

### 9. Debug

`WithDebug(true)` draws a thin border around every row and column. Turn it on while iterating on
layout, off before shipping — it is part of the rendered output, not a log.

## References

- Builder GoDoc: https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#Builder
- Feature pages: https://maroto.tech/#/v2/features/custompage, `customdimensions`, `margins`,
  `maxgridsum`, `orientation`, `pagenumber`, `customfont`, `metadatas`, `protection`,
  `compression`, `background`, `disablepagebreak`, `parallelism`, `lowmemory`
- Builder source with validation rules: [`pkg/config/builder.go`](../../pkg/config/builder.go)
