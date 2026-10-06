# PDF Layout

Place content on the page with maroto's grid: rows, columns, automatic heights, spacing, cell
styling, headers, footers and page breaks.

See also:
- [generation.md](generation.md) for the program skeleton these rows go into.
- [config.md](config.md) for page size, margins and grid size (`WithMaxGridSize`).
- [components.md](components.md) for what goes inside a column.
- [tables.md](tables.md) for repeating the same row layout over a data set.

## Instructions

### 1. Columns — the horizontal grid

- `col.New(n)` spans `n` of the grid's units (default grid: 12). `col.New()` with no size spans
  the **whole** grid.
- The sizes of the columns in one row should add up to the grid size. Fewer leaves blank space on
  the right; more pushes columns off the page (not an error, just wrong output).
- A column's width in mm is `usefulWidth * n / gridSize` (default A4: `190 * n / 12`, so one unit
  ≈ 15.8 mm). Text that doesn't fit the column width wraps.
- Use empty columns as horizontal spacers: `col.New(3)` with nothing added.

```go
func threeColumns(m core.Maroto) {
	m.AddRow(10,
		text.NewCol(3, "left"),
		col.New(6), // spacer
		text.NewCol(3, "right", props.Text{Align: align.Right}),
	)
}
```

When 12 doesn't divide evenly (5, 7, 10 equal columns), change the grid globally with
`config.NewBuilder().WithMaxGridSize(20)` rather than approximating.

### 2. Rows — fixed height in millimetres

```go
func fixedRows(m core.Maroto) {
	// Build then attach.
	r := row.New(8).Add(
		text.NewCol(8, "Description"),
		text.NewCol(4, "Amount", props.Text{Align: align.Right}),
	)
	m.AddRows(r)

	// Or build and attach in one call; AddRow returns the row for chaining WithStyle.
	m.AddRow(8,
		text.NewCol(8, "Shipping"),
		text.NewCol(4, "$ 10.00", props.Text{Align: align.Right}),
	).WithStyle(&props.Cell{BackgroundColor: &props.Color{Red: 240, Green: 240, Blue: 240}})

	// Vertical spacer.
	m.AddRows(row.New(5))
}
```

Height budget: one line of `S` pt text is `S / 2.835` mm (10 pt ≈ 3.5 mm, 16 pt ≈ 5.6 mm).
Add `props.Text.Top` and, for wrapped text, `lines * fontHeight + (lines-1) * VerticalPadding`.
**A fixed row never grows and never clips**: text taller than the row draws over whatever comes
below. If the content length is not known at build time, use an auto row (§3).

A row with no columns (`row.New(5)`) is given one full-width empty column automatically.

### 3. Auto rows

`m.AddAutoRow(cols...)`, `row.New()` with no height, or any component's `NewAutoRow(...)` sizes
the row to its **tallest column**:

```go
func autoRows(m core.Maroto) {
	m.AddAutoRow(
		text.NewCol(4, "Terms", props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, longLegalText, props.Text{Align: align.Justify, Bottom: 2}),
	)

	m.AddAutoRow(
		image.NewFromFileCol(3, "logo.png", props.Rect{Center: true, Percent: 80, JustReferenceWidth: true}),
		text.NewCol(9, "Caption", props.Text{Top: 2}),
	)
}
```

How each component reports its height:

| Component       | Auto height                                                            |
|-----------------|------------------------------------------------------------------------|
| text            | `lines * fontHeight + (lines-1) * VerticalPadding + Top + Bottom`      |
| image / QR / data matrix | `cellWidth * Percent/100 * (imageHeight/imageWidth) + Top` — set `JustReferenceWidth: true` or the drawn size won't match the computed height |
| barcode         | `cellWidth * Percent/100 * Proportion.Height/Proportion.Width + Top`   |
| line            | `Thickness`                                                            |
| checkbox        | `Size + Top`                                                           |
| signature       | `LineThickness + fontHeight * SafePadding`                             |
| empty column    | `0`                                                                    |

- `props.Text.Bottom` only has an effect in auto rows — it is the only way to add padding below
  text.
- An auto row made only of empty columns has height 0; use `row.New(h)` for spacers.
- Auto rows still break pages normally: a single auto row taller than the useful page height is
  placed alone on a new page and overflows it. Split very long text into paragraphs, one auto row
  each.

### 4. Several components in one column

`col.New(n).Add(a, b, c)` draws every component at the **same origin** (top-left of the cell).
They overlap unless you offset them with each component's `Top`/`Left` props:

```go
func stackedInOneColumn(m core.Maroto) {
	m.AddRow(20,
		col.New(6).Add(
			code.NewBar("5123.151231.512314", props.Barcode{Percent: 70, Center: true}),
			text.New("5123.151231.512314", props.Text{Top: 15, Size: 8, Align: align.Center}),
		),
		col.New(6).Add(
			text.New("Company Name", props.Text{Style: fontstyle.Bold}),
			text.New("Street 1, City", props.Text{Top: 5}),
			text.New("Tel: 555-0100", props.Text{Top: 10}),
		),
	)
}
```

Prefer one component per column plus multiple rows when the vertical stacking is the point; use
in-column stacking for captions under images/codes and for address blocks.

### 5. Cell style — backgrounds and borders

`WithStyle(&props.Cell{...})` works on both `core.Row` and `core.Col`:

| Field             | Type              | Default            | Notes                                              |
|-------------------|-------------------|--------------------|----------------------------------------------------|
| `BackgroundColor` | `*props.Color`    | `nil` (none)       | Fills the cell                                     |
| `BorderType`      | `border.Type`     | `border.None`      | `border.Full`, or combine: `border.Left \| border.Bottom` |
| `BorderColor`     | `*props.Color`    | `nil` (black)      |                                                    |
| `BorderThickness` | `float64`         | `0.2` mm           |                                                    |
| `LineStyle`       | `linestyle.Type`  | `linestyle.Solid`  | or `linestyle.Dashed`                              |

```go
func styledCells(m core.Maroto) {
	header := &props.Cell{
		BackgroundColor: &props.Color{Red: 55, Green: 55, Blue: 55},
	}
	boxed := &props.Cell{
		BorderType:      border.Full,
		BorderColor:     &props.Color{Red: 180, Green: 180, Blue: 180},
		BorderThickness: 0.3,
	}

	m.AddRow(8, text.NewCol(12, "Section", props.Text{Color: &props.WhiteColor, Style: fontstyle.Bold, Left: 2, Top: 1.5})).
		WithStyle(header)

	m.AddRow(8,
		text.NewCol(6, "Boxed cell", props.Text{Left: 2, Top: 1.5}).WithStyle(boxed),
		text.NewCol(6, "Boxed cell", props.Text{Left: 2, Top: 1.5}).WithStyle(boxed),
	)
}
```

- A row style paints the full row width; a column style paints only that column. Applying both
  works (column on top of row).
- Text sits flush against the cell edge by default — give styled cells `props.Text{Left: 1.5,
  Top: 1.5}` (or similar) so it doesn't touch the border.
- There is no padding/margin prop on `props.Cell`; spacing is done with `Top/Left/Right` on the
  component or with spacer rows/columns.

### 6. Header and footer

```go
func registerHeaderAndFooter(m core.Maroto) error {
	if err := m.RegisterHeader(
		row.New(15).Add(
			image.NewFromFileCol(3, "logo.png", props.Rect{Center: true, Percent: 80}),
			col.New(6),
			text.NewCol(3, "ACME Corp.", props.Text{Align: align.Right, Style: fontstyle.Bold, Top: 5}),
		),
		line.NewRow(2, props.Line{OffsetPercent: 50}),
	); err != nil {
		return err
	}

	return m.RegisterFooter(
		text.NewRow(8, "Confidential — generated automatically", props.Text{Size: 7, Align: align.Center, Top: 2}),
	)
}
```

- Call both **before the first `AddRow*`/`AddRows`/`AddPages`**. `RegisterHeader` places the
  header rows on the current page immediately; `RegisterFooter` reserves its height on every page
  from that point on. Registering late puts the header mid-page and lets earlier rows collide with
  the footer.
- Each accepts one or more rows, stacked in order. Their total height is subtracted from the
  useful area of every page.
- Each returns an error when the rows are taller than the page; handle it.
- Registering again replaces the previous header/footer for subsequent pages — don't.
- Headers and footers repeat on auto-break pages and on `AddPages` pages. With
  `WithDisableAutoPageBreak(true)` they appear on the first page only.
- Page numbers are not part of the footer; they come from `WithPageNumber` in the config and are
  drawn in the margin.

### 7. Pages and page breaks

- **Automatic**: a row that doesn't fit the remaining height goes to a new page (header first).
  Rows are never split; a tall row simply starts the next page, leaving blank space.
- **Check before adding**: `m.FitlnCurrentPage(height)` (note the library's spelling, `Fitln`)
  reports whether `height` mm still fits. Use it to keep a title with its first content row:

```go
func keepTogether(m core.Maroto, titleHeight, firstRowHeight float64, title core.Row, first core.Row) {
	if !m.FitlnCurrentPage(titleHeight + firstRowHeight) {
		m.AddPages(page.New()) // force a break; an empty page.New() starts a fresh page
	}
	m.AddRows(title, first)
}
```

- **Explicit pages**: `m.AddPages(page.New().Add(rows...), ...)` closes the current page and
  starts each given page on a new sheet. Pages whose rows exceed one sheet are split
  automatically. Use it for chapters/sections that must start on a new page, or when composing a
  document from independently built blocks.
- `AddPages` on an empty document still starts page 1 normally (no leading blank page).
- Orientation and size are per document, not per page; merge separately generated documents for
  mixed layouts ([generation.md §5](generation.md#5-output-options)).

### 8. Worked layout: title block

```go
func titleBlock(m core.Maroto) {
	m.AddRows(
		text.NewRow(12, "Quarterly Statement", props.Text{Size: 18, Style: fontstyle.Bold}),
		text.NewRow(6, "Q3 2026 · Account 00123", props.Text{Size: 9, Color: &props.Color{Red: 120, Green: 120, Blue: 120}}),
		line.NewRow(4, props.Line{OffsetPercent: 50, SizePercent: 100}),
		row.New(4),
	)

	m.AddRow(14,
		col.New(6).Add(
			text.New("Billed to", props.Text{Style: fontstyle.Bold, Size: 9}),
			text.New("Jane Doe", props.Text{Top: 4.5}),
			text.New("42 Example Road", props.Text{Top: 9}),
		),
		col.New(6).Add(
			text.New("Payment due", props.Text{Style: fontstyle.Bold, Size: 9, Align: align.Right}),
			text.New("2026-10-31", props.Text{Top: 4.5, Align: align.Right}),
		),
	)
}
```

## References

- Grid explanation: https://maroto.tech/#/README?id=maroto-columns-and-rows
- Feature pages: https://maroto.tech/#/v2/features/autorow, `cellstyle`, `header`, `footer`,
  `addpage`, `disablepagebreak`
- Runnable examples: [`docs/assets/examples/autorow`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/autorow/v2/main.go),
  [`cellstyle`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/cellstyle/v2/main.go),
  [`header`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/header/v2/main.go),
  [`addpage`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/addpage/v2/main.go)
