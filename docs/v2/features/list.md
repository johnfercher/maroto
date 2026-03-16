# List

The List helper generates a sequence of rows from a Go slice, automatically prepending a header row. It is ideal for tabular data such as invoices, reports, and catalogs where each data item maps to a fixed row layout.

Items must implement the `Listable` interface:

## Functions

| Function | Description |
|----------|-------------|
| `list.Build[T Listable](arr []T)` | Build rows from a value slice |
| `list.BuildFromPointer[T Listable](arr []*T)` | Build rows from a pointer slice |

Both functions return `([]core.Row, error)`. Errors: `ErrEmptyArray` (empty slice) and `ErrNilElementInArray` (nil pointer element).

## Usage notes

- Pass the returned `[]core.Row` directly to `m.AddRows(rows...)`.
- The header row is produced by calling `GetHeader()` on the **first** element.
- Both value and pointer variants exist because Go generics cannot automatically dereference pointers; use `BuildFromPointer` when your slice is `[]*T`.

## GoDoc
* [interface : Listable](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/list#Listable)
* [list : Build](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/list#Build)
* [list : BuildFromPointer](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/list#BuildFromPointer)

## Code Example
[filename](../../assets/examples/list/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/listv2.pdf
```
## Time Execution
[filename](../../assets/text/listv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/list.json  ':include :type=code')