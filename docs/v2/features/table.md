# Table

`AddTable` adds a table: its header rows followed by its body rows. When the body doesn't fit on the current page and continues onto the next one, the header rows are repeated at the top of every page the table spans.

## Usage notes

- The header repeats **only while the table's body is being added**. Rows added after the table, and the pages of other tables, never get it.
- The header is never left alone at the bottom of a page: if it doesn't fit together with the first body row, the table starts on a new page.
- The header can be several rows, for example a title row above the column names.
- It stacks under the document header from `RegisterHeader()`, which still appears on every page.
- If the header and the row that broke the page don't fit on one page together, that page gets no table header.

## AddTable vs RegisterHeader

| | `RegisterHeader()` | `AddTable()` |
|---|---|---|
| **Repeats on** | Every page of the document | Only the pages the table spans |
| **Position** | Top of the page | Top of the table on its first page, top of the page after that |
| **Use for** | Logo, company name, document title | Column names of a table |

## Using it with List

`list.Build` returns the header row first, so split it off:

```go
rows, err := list.Build(products)
if err != nil {
	return err
}
m.AddTable(rows[:1], rows[1:]...)
```

## GoDoc
* [maroto : AddTable](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.AddTable)

## Code Example
[filename](../../assets/examples/table/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/tablev2.pdf
```

## Time Execution
[filename](../../assets/text/tablev2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/table.json  ':include :type=code')

## Related Features

- [Header](v2/features/header?id=header) - Document header on every page
- [Footer](v2/features/footer?id=footer) - Document footer on every page
- [List](v2/features/list?id=list) - Builds table rows from a slice
- [Cell Style](v2/features/cellstyle?id=cell-style) - Styling the header and body
