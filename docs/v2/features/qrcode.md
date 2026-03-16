# QrCode

The QrCode component renders a square QR code inside a cell. QR codes can encode URLs, plain text, contact cards, or any binary-safe string and are widely supported by modern smartphone cameras.

QR codes share the same `props.Rect` struct as images, giving them identical positioning and sizing semantics.

## Props (`props.Rect`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Percent` | `float64` | `100` | How much of the cell the QR code occupies (0–100) |
| `Center` | `bool` | `false` | Horizontally and vertically center the code |
| `Left` | `float64` | `0` | Left offset in mm — ignored when `Center` is true |
| `Top` | `float64` | `0` | Top offset in mm — ignored when `Center` is true |
| `JustReferenceWidth` | `bool` | `false` | Scale by width only — required for correct auto-row height |

## Usage notes

- Set `JustReferenceWidth: true` inside `AddAutoRow` so the row expands proportionally to the code size.
- `Center: true` is the most common setting for a dedicated QR code column.

## GoDoc
* [constructor : NewQr](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewQr)
* [constructor : NewQrCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewQrCol)
* [constructor : NewQrRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#NewQrRow)
* [props : Rect](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Rect)
* [component : QrCode](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/code#QrCode)

## Code Example
[filename](../../assets/examples/qrgrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/qrgridv2.pdf
```

## Time Execution
[filename](../../assets/text/qrgridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/qrgrid.json  ':include :type=code')