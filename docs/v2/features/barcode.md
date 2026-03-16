# Barcode

The Barcode component renders a 1-D barcode inside a cell. The default type is `Code128`, which can encode any ASCII string. Other types such as `EAN` are also supported.

## Props (`props.Barcode`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Type` | `barcode.Type` | `barcode.Code128` | Barcode symbology |
| `Percent` | `float64` | `100` | How much of the cell the barcode occupies (0–100) |
| `Center` | `bool` | `false` | Horizontally and vertically center the barcode |
| `Left` | `float64` | `0` | Left offset in mm — ignored when `Center` is true |
| `Top` | `float64` | `0` | Top offset in mm — ignored when `Center` is true |
| `Proportion` | `props.Proportion` | `{Width:1, Height:0.2}` | Width-to-height ratio |

## Usage notes

- The barcode height is constrained to be between 10% and 20% of its width to ensure scannability; values outside this range are clamped by `MakeValid`.
- For EAN barcodes the content must be exactly 12 or 13 digits.
- Use `JustReferenceWidth` (via `props.Rect`) on QR / Matrix codes; barcodes use `props.Barcode` which controls proportion directly.

## GoDoc
* [constructor : NewBar](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewBar)
* [constructor : NewBarCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewBarCol)
* [constructor : NewBarRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewBarRow)
* [props : Barcode](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Barcode)
* [component : Barcode](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#Barcode)

## Code Example
[filename](../../assets/examples/barcodegrid/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/barcodegridv2.pdf
```

## Time Execution
[filename](../../assets/text/barcodegridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/barcodegrid.json  ':include :type=code')
