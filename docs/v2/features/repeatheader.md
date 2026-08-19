# Repeat Header on Page Breaks

`WithRepeatOnPageBreak` marks individual rows to automatically re-appear at the top of new pages when content overflows. This is essential for multi-page tables, reports, and data exports where headers must remain visible across page boundaries.

## Overview

When a document's content exceeds a single page, maroto automatically creates new pages. Without repeat headers, table headers disappear after the first page, making it difficult to understand subsequent data. The `WithRepeatOnPageBreak()` method ensures marked rows (typically headers) are automatically re-injected at the beginning of each new page.

## Use Cases

- **Table Headers** - Column labels in multi-page tables
- **Report Sections** - Section titles that span multiple pages
- **Data Lists** - Headers for contact directories, inventories, or catalogs
- **Forms** - Field labels in long surveys or questionnaires
- **Shipping Labels** - Item headers in bill of lading documents

## Key Differences from RegisterHeader

| Aspect | `RegisterHeader()` | `WithRepeatOnPageBreak()` |
|--------|-------------------|-------------------------|
| **Scope** | Global (every page) | Local (on overflow only) |
| **Purpose** | Document-wide header/footer | Table/section header |
| **Setup** | Once at document start | Per-row marking |
| **Repeats** | Automatically on all pages | Only when content breaks |

## Usage Notes

- Marked rows are copied when a page break occurs, not moved
- Performance impact is minimal (typically <1% overhead for reasonable row counts)
- Works seamlessly with `RegisterHeader()` and `RegisterFooter()`
- Compatible with the `list` component for building table-like structures
- Default behavior is non-repeating; must be explicitly enabled

## GoDoc

* [row : WithRepeatOnPageBreak](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/row#Row.WithRepeatOnPageBreak)
* [row : IsRepeatOnPageBreak](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/row#Row.IsRepeatOnPageBreak)

## Code Example

### Basic Table with Repeating Header

```go
package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	m := maroto.New()

	// Create header row that repeats on page breaks
	headerRow := row.New(8).
		Add(text.NewCol(3, "Item ID")).
		Add(text.NewCol(6, "Description")).
		Add(text.NewCol(3, "Amount")).
		WithRepeatOnPageBreak()

	m.AddRows(headerRow)

	// Add many data rows (will span multiple pages)
	for i := 1; i <= 100; i++ {
		m.AddRow(6,
			text.NewCol(3, fmt.Sprintf("ID-%d", i)),
			text.NewCol(6, fmt.Sprintf("Item %d", i)),
			text.NewCol(3, fmt.Sprintf("$%.2f", 10.50*float64(i))),
		)
	}

	doc, err := m.Generate()
	if err != nil {
		panic(err)
	}

	doc.Save("invoice.pdf")
}
```

### With List Component

```go
type Product struct {
	ID    string
	Name  string
	Price float64
}

func (p *Product) GetHeader() core.Row {
	return row.New(8).
		Add(col.New(3).Add(text.NewCol("Item ID"))).
		Add(col.New(6).Add(text.NewCol("Product Name"))).
		Add(col.New(3).Add(text.NewCol("Price"))).
		WithRepeatOnPageBreak()  // Header repeats on page breaks
}

func (p *Product) GetContent(i int) core.Row {
	return row.New(6).
		Add(col.New(3).Add(text.NewCol(p.ID))).
		Add(col.New(6).Add(text.NewCol(p.Name))).
		Add(col.New(3).Add(text.NewCol(fmt.Sprintf("$%.2f", p.Price))))
}

func main() {
	m := maroto.New()

	products := []Product{
		// 100+ products...
	}

	rows, err := list.Build(products)
	if err != nil {
		panic(err)
	}

	m.AddRows(rows...)
	doc, err := m.Generate()
	if err != nil {
		panic(err)
	}

	doc.Save("catalog.pdf")
}
```


## Related Features

- [Register Header](v2/features/header?id=header) - Global document header
- [Register Footer](v2/features/footer?id=footer) - Global document footer
- [List Component](v2/features/list?id=list) - Table/list builder
- [Cell Style](v2/features/cellstyle?id=cell-style) - Styling headers and data
