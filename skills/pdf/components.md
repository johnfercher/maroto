# PDF Components

Put content into a column: text, images, barcodes, QR codes, data matrices, lines, checkboxes and
signatures — constructors, props, defaults and the gotchas of each.

See also:
- [layout.md](layout.md) for rows, columns, auto heights and cell styling around these
  components.
- [config.md](config.md) for document-wide defaults (default font, custom UTF-8 fonts).
- [tables.md](tables.md) for repeating text columns over a data set.

## Instructions

### 1. Constructor pattern (all components)

| Shape                         | Example (text)                               | Returns          |
|-------------------------------|----------------------------------------------|------------------|
| component                     | `text.New(value, props...)`                  | `core.Component` |
| column                        | `text.NewCol(size, value, props...)`         | `core.Col`       |
| fixed-height row              | `text.NewRow(height, value, props...)`       | `core.Row`       |
| auto-height row               | `text.NewAutoRow(value, props...)`           | `core.Row`       |

Props are variadic: pass zero or one struct value (not a pointer). Zero-valued fields take the
defaults listed below. Nothing here returns an error; bad input renders a red error message in the
cell (see [generation.md §6](generation.md#6-errors-to-handle)).

### 2. Text — `pkg/components/text`, `props.Text`

```go
func textExamples(m core.Maroto) {
	url := "https://maroto.tech"

	m.AddRows(
		text.NewRow(10, "Title", props.Text{Size: 14, Style: fontstyle.Bold}),
		text.NewRow(6, "Subtitle", props.Text{Size: 9, Style: fontstyle.Italic, Color: &props.Color{Red: 100, Green: 100, Blue: 100}}),
		text.NewRow(6, "Read the docs", props.Text{Hyperlink: &url}),
	)

	m.AddAutoRow(
		text.NewCol(12, paragraph, props.Text{Align: align.Justify, VerticalPadding: 1, Bottom: 3}),
	)
}
```

| Field               | Type                 | Default                 | Notes                                                                 |
|---------------------|----------------------|-------------------------|-----------------------------------------------------------------------|
| `Family`            | `string`             | config default font     | `fontfamily.Arial/Helvetica/Courier/Symbol/ZapBats` or a registered custom family |
| `Style`             | `fontstyle.Type`     | config default          | `Normal`, `Bold`, `Italic`, `BoldItalic`, `Underline`, `Strikethrough`; combine by concatenation: `fontstyle.Bold + fontstyle.Underline` |
| `Size`              | `float64` (pt)       | config default (10)     | Line height in mm is `Size / 2.835`                                   |
| `Color`             | `*props.Color`       | config default (black)  | Overridden to blue when `Hyperlink` is set                            |
| `Align`             | `align.Type`         | `align.Left`            | `Left`, `Center`, `Right`, `Justify`                                  |
| `Top`               | `float64` (mm)       | `0`                     | Offset from the top of the cell; clamped to the cell height           |
| `Left`, `Right`     | `float64` (mm)       | `0`                     | Inner horizontal padding; shrink the wrap width                       |
| `Bottom`            | `float64` (mm)       | `0`                     | **Auto rows only** — extra height below the text                      |
| `VerticalPadding`   | `float64` (mm)       | `0`                     | Extra space between wrapped lines                                     |
| `BreakLineStrategy` | `breakline.Strategy` | `EmptySpaceStrategy`    | `EmptySpaceStrategy` wraps on spaces; `DashStrategy` wraps per character and appends `-` (for languages without spaces) |
| `Hyperlink`         | `*string`            | `nil`                   | Clickable link; text turns blue                                       |
| `Rotation`          | `float64` (degrees)  | `0`                     | Positive rotates counter-clockwise, negative clockwise                |
| `RotationPivot`     | `rotationpivot.Pivot` | `{Center, Middle}`     | Anchor of the rotation: `Horizontal` `Start`/`Center`/`End`, `Vertical` `Top`/`Middle`/`Bottom` (of the whole block for multi-line text) |

Gotchas:
- Wrapping happens automatically when the string is wider than `cellWidth - Left - Right`, but the
  **row does not grow**; overflow draws over the next row. Size the row (§height budget in
  [layout.md §2](layout.md#2-rows--fixed-height-in-millimetres)) or use an auto row.
- A single word wider than the column is not broken under `EmptySpaceStrategy`; it overflows to
  the right. Use `DashStrategy` for identifiers/URLs in narrow columns.
- Built-in fonts render cp1252 only. Any other script needs a custom font
  ([config.md §5](config.md#5-fonts)) — otherwise characters silently come out wrong.
- `Top` is the only vertical positioning: there is no vertical-align. To bottom-align, compute
  `Top = rowHeight - lineHeight`.
- Rotated text needs an **auto row**: only an auto row grows to the rotated bounding box
  (`w·|sinθ| + h·|cosθ|`). A fixed-height row keeps its height and the text draws over the next
  row. Only the height is reserved, so a rotated box wider than its column spills sideways.

```go
func rotatedLabel(m core.Maroto) {
	m.AddAutoRow(
		text.NewCol(4, "DRAFT", props.Text{Rotation: 30, Size: 14}),
		text.NewCol(8, "Body", props.Text{
			Rotation:      -15,
			RotationPivot: rotationpivot.Pivot{Horizontal: rotationpivot.Start, Vertical: rotationpivot.Top},
		}),
	)
}
```

- Newlines (`\n`) in the value are not line breaks. Add one text component (or auto row) per line
  or paragraph.

### 3. Images — `pkg/components/image`, `props.Rect`

Two sources, same four shapes each:

| From file (path, extension inferred from the name) | From memory (`[]byte` + `extension.Png/Jpg/Jpeg`) |
|-----------------------------------------------------|----------------------------------------------------|
| `image.NewFromFile(path, props...)`                 | `image.NewFromBytes(b, ext, props...)`            |
| `image.NewFromFileCol(size, path, props...)`        | `image.NewFromBytesCol(size, b, ext, props...)`   |
| `image.NewFromFileRow(height, path, props...)`      | `image.NewFromBytesRow(height, b, ext, props...)` |
| `image.NewAutoFromFileRow(path, props...)`          | `image.NewAutoFromBytesRow(b, ext, props...)`     |

```go
import _ "embed"

//go:embed assets/logo.png
var logo []byte

func imageExamples(m core.Maroto) {
	m.AddRow(25,
		image.NewFromBytesCol(3, logo, extension.Png, props.Rect{Center: true, Percent: 90}),
		image.NewFromFileCol(3, "./photo.jpg", props.Rect{Left: 2, Top: 2, Percent: 80}),
	)

	m.AddAutoRow(
		image.NewFromBytesCol(12, logo, extension.Png, props.Rect{Center: true, Percent: 50, JustReferenceWidth: true}),
	)
}
```

| Field                | Default | Notes                                                                                      |
|----------------------|---------|--------------------------------------------------------------------------------------------|
| `Percent`            | `100`   | Size of the image relative to the cell (0–100]; the image is scaled proportionally to fit  |
| `Center`             | `false` | Center in both axes; when `true`, `Left`/`Top` are ignored                                 |
| `Left`, `Top`        | `0`     | Offset in mm from the cell's top-left                                                      |
| `JustReferenceWidth` | `false` | Scale from the cell **width** only. Set it in auto rows so the drawn height matches the computed row height |

Gotchas:
- Only PNG and JPEG. Convert SVG/GIF/WebP before building. The file extension must match the
  real format (`.jpg`/`.jpeg`/`.png`, case-insensitive).
- Prefer `NewFromBytes*` with `//go:embed` for assets shipped with the binary; file paths are
  resolved against the process working directory at `Generate()` time.
- A file that can't be read renders `could not load image` in red in the cell and `Generate()`
  still succeeds — check `os.Stat` first if a missing asset must fail the job.
- Images loaded by **path** (and generated barcodes/QR/matrix codes) are cached by path/value, so
  repeating the same logo on every page is cheap. `NewFromBytes*` images are decoded on every
  render — fine for a logo, wasteful for hundreds of identical thumbnails (use a file path then).

### 4. Barcodes — `pkg/components/code`, `props.Barcode`

`code.NewBar(value, props...)`, `NewBarCol(size, value, ...)`, `NewBarRow(height, value, ...)`,
`NewAutoBarRow(value, ...)`.

```go
func barcodeExamples(m core.Maroto) {
	m.AddRow(20,
		code.NewBarCol(6, "ORDER-2026-000123", props.Barcode{Center: true, Percent: 90}),
		code.NewBarCol(6, "123456789012", props.Barcode{Type: barcode.EAN, Center: true}),
	)
}
```

| Field        | Default            | Notes                                                                           |
|--------------|--------------------|---------------------------------------------------------------------------------|
| `Type`       | `barcode.Code128`  | `Code128` (any ASCII) or `barcode.EAN` (exactly 12 or 13 digits)                |
| `Proportion` | `{Width: 1, Height: 0.2}` | Width:height ratio. Height is clamped to 10–20 % of width for scannability |
| `Percent`, `Center`, `Left`, `Top` | as for images | Position/size inside the cell                                   |

Place the human-readable value as a `text.New(..., props.Text{Top: ...})` in the same column
([layout.md §4](layout.md#4-several-components-in-one-column)).

### 5. QR codes and data matrices — `pkg/components/code`, `props.Rect`

| QR                                         | Data Matrix                                   |
|--------------------------------------------|-----------------------------------------------|
| `code.NewQr(value, props...)`              | `code.NewMatrix(value, props...)`             |
| `code.NewQrCol(size, value, props...)`     | `code.NewMatrixCol(size, value, props...)`    |
| `code.NewQrRow(height, value, props...)`   | `code.NewMatrixRow(height, value, props...)`  |
| `code.NewAutoQrRow(value, props...)`       | `code.NewAutoMatrixRow(value, props...)`      |

```go
func qrExamples(m core.Maroto) {
	m.AddRow(30,
		code.NewQrCol(3, "https://example.com/invoice/123", props.Rect{Center: true}),
		col.New(6),
		code.NewMatrixCol(3, "LOT-7781-A", props.Rect{Center: true, Percent: 70}),
	)
}
```

Both are square; `Percent` applies to the shorter cell side. Use `Center: true` for a dedicated
code column and `JustReferenceWidth: true` inside auto rows. Any string content works (URLs,
JSON, vCards); keep QR payloads under a few hundred characters for reliable scanning.

### 6. Lines — `pkg/components/line`, `props.Line`

`line.New(props...)`, `NewCol(size, props...)`, `NewRow(height, props...)`, `NewAutoRow(props...)`.

```go
func lineExamples(m core.Maroto) {
	// Full-width divider centred in a 4 mm row.
	m.AddRows(line.NewRow(4, props.Line{OffsetPercent: 50, SizePercent: 100}))

	// Vertical separator between two columns.
	m.AddRow(20,
		text.NewCol(5, "left"),
		line.NewCol(2, props.Line{Orientation: orientation.Vertical, OffsetPercent: 50, Thickness: 0.4}),
		text.NewCol(5, "right"),
	)
}
```

| Field           | Default                  | Notes                                                                     |
|-----------------|--------------------------|---------------------------------------------------------------------------|
| `Orientation`   | `orientation.Horizontal` | or `orientation.Vertical`                                                 |
| `Style`         | `linestyle.Solid`        | or `linestyle.Dashed`                                                     |
| `Thickness`     | `0.2` mm                 |                                                                           |
| `Color`         | black                    |                                                                           |
| `OffsetPercent` | `5`                      | Position across the cell (0 = top/left edge … 100 = bottom/right); clamped to 5–95. **`50` centres it** |
| `SizePercent`   | `90`                     | Length relative to the cell; clamped to (0, 100]                          |

Default `OffsetPercent` puts a horizontal line near the **top** of its row — pass `50` for a
divider between rows. In an auto row the row height equals `Thickness`, so a divider with
breathing room needs a fixed-height row.

### 7. Checkboxes — `pkg/components/checkbox`, `props.Checkbox`

`checkbox.New(label, props...)`, `NewCol(size, label, ...)`, `NewRow(height, label, ...)`,
`NewAutoRow(label, ...)`.

```go
func checkboxExamples(m core.Maroto) {
	m.AddAutoRow(
		checkbox.NewCol(4, "I agree to the terms", props.Checkbox{Checked: true}),
		checkbox.NewCol(4, "Send me updates"),
		checkbox.NewCol(4, "Large box", props.Checkbox{Size: 8, Top: 1}),
	)
}
```

| Field     | Default | Notes                                              |
|-----------|---------|----------------------------------------------------|
| `Checked` | `false` | Draws an X inside the box                          |
| `Size`    | `5` mm  | Side of the square                                 |
| `Top`, `Left` | `0` | Offset in mm; negatives clamp to 0                 |

The label is drawn to the right of the box in the document's default font; it is not wrapped.
The component is a static drawing, not an interactive PDF form field.

### 8. Signatures — `pkg/components/signature`, `props.Signature`

`signature.New(label, props...)`, `NewCol(size, label, ...)`, `NewRow(height, label, ...)`,
`NewAutoRow(label, ...)`.

```go
func signatureExamples(m core.Maroto) {
	m.AddRow(20,
		signature.NewCol(5, "Customer", props.Signature{FontSize: 8}),
		col.New(2),
		signature.NewCol(5, "Supplier", props.Signature{FontSize: 8, LineStyle: linestyle.Dashed}),
	)
}
```

| Field            | Default           | Notes                                          |
|------------------|-------------------|------------------------------------------------|
| `FontFamily`     | config default    |                                                |
| `FontStyle`      | `fontstyle.Bold`  |                                                |
| `FontSize`       | `8` pt            |                                                |
| `FontColor`      | black             |                                                |
| `LineStyle`      | `linestyle.Solid` | or `Dashed`                                    |
| `LineThickness`  | `0.2` mm          |                                                |
| `LineColor`      | black             |                                                |
| `SafePadding`    | `1.5`             | Multiplier of the font height between line and label |

Renders a horizontal line at the bottom of the cell with the label centred **below** the line.
Leave the row tall enough (≥ 15 mm) for an actual handwritten signature above the line.

### 9. Colors

`props.Color{Red, Green, Blue int}` (0–255). Predefined values: `props.WhiteColor`,
`BlackColor`, `RedColor`, `GreenColor`, `BlueColor` — they are values, so pass `&props.RedColor`.
Define the document palette once as package-level `*props.Color` variables and reuse them; a
`nil` color always means "the default".

## References

- Feature pages: https://maroto.tech/#/v2/features/text, `image`, `barcode`, `qrcode`,
  `datamatrix`, `line`, `checkbox`, `signature`
- Props GoDoc: https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props
- Runnable grids showing every prop combination:
  [`docs/assets/examples/textgrid`](../../docs/assets/examples/textgrid/v2/main.go),
  [`imagegrid`](../../docs/assets/examples/imagegrid/v2/main.go),
  [`barcodegrid`](../../docs/assets/examples/barcodegrid/v2/main.go),
  [`qrgrid`](../../docs/assets/examples/qrgrid/v2/main.go),
  [`line`](../../docs/assets/examples/line/v2/main.go),
  [`checkbox`](../../docs/assets/examples/checkbox/v2/main.go),
  [`signaturegrid`](../../docs/assets/examples/signaturegrid/v2/main.go)
