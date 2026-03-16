# Data Matrix

The Data Matrix component renders a 2-D matrix barcode inside a cell. Data Matrix codes are square, high-density barcodes used in manufacturing, logistics, and healthcare. They can encode more data in a smaller space than QR codes and are particularly suited for small labels.

Like QR codes, Data Matrix codes use `props.Rect` for layout control.

## Props (`props.Rect`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Percent` | `float64` | `100` | How much of the cell the code occupies (0–100) |
| `Center` | `bool` | `false` | Horizontally and vertically center the code |
| `Left` | `float64` | `0` | Left offset in mm — ignored when `Center` is true |
| `Top` | `float64` | `0` | Top offset in mm — ignored when `Center` is true |
| `JustReferenceWidth` | `bool` | `false` | Scale by width only — required for correct auto-row height |

## Usage notes

- Data Matrix codes are always square; `Percent` governs the side length relative to the shorter cell dimension.
- Set `JustReferenceWidth: true` when using `AddAutoRow` to allow the row to grow to the code's natural height.

## GoDoc
* [constructor : NewMatrix](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewMatrix)
* [constructor : NewMatrixCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewMatrixCol)
* [constructor : NewMatrixRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewMatrixRow)
* [props : Rect](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Rect)
* [component : MatrixCode](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#MatrixCode)

## Code Example
[filename](../../assets/examples/datamatrixgrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/datamatrixgridv2.pdf
```

## Time Execution
[filename](../../assets/text/datamatrixgridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/datamatrixgrid.json  ':include :type=code')