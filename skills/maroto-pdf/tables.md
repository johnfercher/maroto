# PDF Tables

Render tabular data — invoices, statements, reports, catalogs — as repeated rows with a header,
consistent column widths, striping, borders, totals and correct pagination.

See also:
- [layout.md](layout.md) for the grid arithmetic, cell styling and header/footer
  mechanics used here.
- [components.md](components.md) for `props.Text` fields used in cells.
- [testing.md](testing.md) for asserting the generated row structure.

## Instructions

### 1. Decide the column plan first

Write the widths down before any code; they must sum to the grid size (default 12):

| Column      | Units | Align   |
|-------------|-------|---------|
| Description | 6     | left    |
| Qty         | 2     | right   |
| Unit price  | 2     | right   |
| Total       | 2     | right   |

If the plan needs 5, 7 or 10 equal columns, set `WithMaxGridSize(20)` (or another multiple) in
the config instead of fudging widths. Right-align numeric columns; keep text columns left.

Keep the plan in code as constants or a small struct so header and content rows can't drift
apart:

```go
type column struct {
	title string
	size  int
	align align.Type
}

var invoiceColumns = []column{
	{title: "Description", size: 6, align: align.Left},
	{title: "Qty", size: 2, align: align.Right},
	{title: "Unit price", size: 2, align: align.Right},
	{title: "Total", size: 2, align: align.Right},
}
```

### 2. Shared cell styles

Define the look once; both approaches below reuse these:

```go
var (
	headerBackground = &props.Color{Red: 55, Green: 55, Blue: 55}
	stripeBackground = &props.Color{Red: 235, Green: 235, Blue: 235}
	gridBorder       = &props.Color{Red: 200, Green: 200, Blue: 200}

	headerText = props.Text{Style: fontstyle.Bold, Size: 9, Color: &props.WhiteColor, Top: 1.5, Left: 1.5, Right: 1.5}
	cellText   = props.Text{Size: 8, Top: 1.5, Left: 1.5, Right: 1.5}
	cellBorder = &props.Cell{BorderType: border.Full, BorderColor: gridBorder, BorderThickness: 0.2}
)

func withAlign(base props.Text, a align.Type) props.Text {
	base.Align = a
	return base
}
```

`Top/Left/Right: 1.5` keeps text off the cell edges; without it, text touches borders and
background edges. Default `props.Text{}` has no padding.

### 3. Approach A — manual loop (full control)

Best when rows differ (subtotals, grouped sections, variable heights):

```go
type lineItem struct {
	Description string
	Qty         int
	UnitPrice   float64
}

func (l lineItem) total() float64 { return float64(l.Qty) * l.UnitPrice }

func addInvoiceTable(m core.Maroto, items []lineItem) {
	header := row.New(7)
	for _, c := range invoiceColumns {
		header.Add(text.NewCol(c.size, c.title, withAlign(headerText, c.align)))
	}
	m.AddRows(header.WithStyle(&props.Cell{BackgroundColor: headerBackground}))

	var sum float64
	for i, item := range items {
		sum += item.total()

		r := m.AddRow(6,
			text.NewCol(6, item.Description, withAlign(cellText, align.Left)),
			text.NewCol(2, strconv.Itoa(item.Qty), withAlign(cellText, align.Right)),
			text.NewCol(2, fmt.Sprintf("%.2f", item.UnitPrice), withAlign(cellText, align.Right)),
			text.NewCol(2, fmt.Sprintf("%.2f", item.total()), withAlign(cellText, align.Right)),
		)
		if i%2 == 1 {
			r.WithStyle(&props.Cell{BackgroundColor: stripeBackground})
		}
	}

	m.AddRow(8,
		col.New(8),
		text.NewCol(2, "Total", props.Text{Style: fontstyle.Bold, Size: 9, Align: align.Right, Top: 2}),
		text.NewCol(2, fmt.Sprintf("%.2f", sum), props.Text{Style: fontstyle.Bold, Size: 9, Align: align.Right, Top: 2, Right: 1.5}),
	)
}
```

- Format numbers/dates into strings before building the row; components take strings only.
- `m.AddRow` returns the row, so striping is a `WithStyle` on the returned value.
- For grid lines, put `.WithStyle(cellBorder)` on each `text.NewCol(...)` instead of the row (a
  row border only boxes the whole row).

### 4. Approach B — `list.Build` (data type knows its layout)

Best when a domain type is rendered the same way in several documents. Implement
`list.Listable` on the type:

```go
type Product struct {
	SKU   string
	Name  string
	Price float64
}

// GetHeader is called once, on the first element, to produce the header row.
func (p Product) GetHeader() core.Row {
	return row.New(7).Add(
		text.NewCol(3, "SKU", headerText),
		text.NewCol(6, "Name", headerText),
		text.NewCol(3, "Price", withAlign(headerText, align.Right)),
	).WithStyle(&props.Cell{BackgroundColor: headerBackground})
}

// GetContent is called for every element with its index, which allows striping.
func (p Product) GetContent(i int) core.Row {
	r := row.New(6).Add(
		text.NewCol(3, p.SKU, cellText),
		text.NewCol(6, p.Name, cellText),
		text.NewCol(3, fmt.Sprintf("%.2f", p.Price), withAlign(cellText, align.Right)),
	)
	if i%2 == 1 {
		r.WithStyle(&props.Cell{BackgroundColor: stripeBackground})
	}
	return r
}

func addProductTable(m core.Maroto, products []Product) error {
	rows, err := list.Build(products)
	if err != nil {
		return err // list.ErrEmptyArray — decide whether an empty table is an error for you
	}
	m.AddRows(rows...)
	return nil
}
```

- `list.Build[T Listable](arr []T)` for value slices; `list.BuildFromPointer[T Listable](arr []*T)`
  for pointer slices (returns `list.ErrNilElementInArray` on a nil element). Implement the
  interface on the **value** receiver for both to work.
- `Build` returns `list.ErrEmptyArray` for an empty slice. If an empty table should render just
  the header, handle that branch explicitly (`m.AddRows(Product{}.GetHeader())`).
- `GetHeader` is taken from element `0`, so it must not depend on the element's data.

### 5. Pagination: repeating the header on every page

`list.Build` and the manual loop print the header **once**. For long tables where the header must
repeat after each automatic page break, register it as the document header instead:

```go
func addLongTable(m core.Maroto, products []Product) error {
	if len(products) == 0 {
		return list.ErrEmptyArray
	}

	if err := m.RegisterHeader(products[0].GetHeader()); err != nil {
		return err
	}

	for i, p := range products {
		m.AddRows(p.GetContent(i))
	}

	return nil
}
```

Trade-off: the document header is global and `RegisterHeader` must precede all content
([layout.md §6](layout.md#6-header-and-footer)). Anything that should appear above the
table on page 1 only (title block, addresses) therefore can't be a normal row before the table.
Either include it in the header rows (it then repeats), or generate the title page and the table
as two documents and merge them ([generation.md §5](generation.md#5-output-options)).

Rows are never split across pages; a row that doesn't fit moves whole to the next page.
Totals rows are just more rows, so they may land alone on a new page — acceptable in most
reports; use `m.FitlnCurrentPage(h)` to force the last `n` rows together when it's not.

### 6. Variable-height cells

When a text column may wrap (descriptions, notes), use auto rows so the row grows to the longest
cell and striping/borders cover the full height:

```go
func addNotesTable(m core.Maroto, notes [][2]string) {
	for i, n := range notes {
		r := m.AddAutoRow(
			text.NewCol(3, n[0], withAlign(cellText, align.Left)).WithStyle(cellBorder),
			text.NewCol(9, n[1], props.Text{Size: 8, Top: 1.5, Bottom: 1.5, Left: 1.5, Right: 1.5}).WithStyle(cellBorder),
		)
		if i%2 == 1 {
			r.WithStyle(&props.Cell{BackgroundColor: stripeBackground})
		}
	}
}
```

Set `Bottom` on the wrapping column so the text doesn't sit on the bottom border — `Bottom` is
honoured only in auto rows. Fixed-height rows with wrapping text overflow into the next row.

### 7. Many columns

- More than ~8 columns on A4 portrait: switch to landscape
  (`WithOrientation(orientation.Horizontal)`) before shrinking fonts below 7 pt.
- Build the column slices in a loop when the count is dynamic, exactly as in
  [`docs/assets/examples/maxgridsum`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/maxgridsum/v2/main.go):

```go
func dynamicColumns(m core.Maroto, headers []string, gridSize int) {
	cols := make([]core.Col, 0, len(headers))
	for _, h := range headers {
		cols = append(cols, text.NewCol(gridSize/len(headers), h, headerText))
	}
	m.AddRow(7, cols...).WithStyle(&props.Cell{BackgroundColor: headerBackground})
}
```

Make sure `gridSize % len(headers) == 0`, or distribute the remainder to the first columns.

## References

- Feature page: https://maroto.tech/#/v2/features/list
- Runnable examples: [`docs/assets/examples/list`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/list/v2/main.go),
  [`billing`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/billing/v2/main.go) (full invoice with header, footer,
  striped table, totals and barcode),
  [`maxgridsum`](https://github.com/johnfercher/maroto/blob/master/docs/assets/examples/maxgridsum/v2/main.go)
- `list` GoDoc: https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/list
