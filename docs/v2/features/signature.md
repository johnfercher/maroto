# Signature

The Signature component renders a signature field: a centered label above a horizontal line. It is commonly used in contracts, invoices, and forms where a handwritten signature is required.

The row height is automatically calculated from the font height plus the `SafePadding` and the line thickness, so the component works correctly with both fixed-height and auto rows.

## Props (`props.Signature`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `FontFamily` | `string` | `fontfamily.Arial` | Font family for the label |
| `FontStyle` | `fontstyle.Type` | `fontstyle.Bold` | Font style for the label |
| `FontSize` | `float64` | `8` | Font size in points |
| `FontColor` | `*props.Color` | `nil` (black) | Label text color |
| `LineStyle` | `linestyle.Type` | `linestyle.Solid` | `Solid` or `Dashed` |
| `LineThickness` | `float64` | `0.2` | Signature line thickness in mm |
| `LineColor` | `*props.Color` | `nil` (black) | Signature line color |
| `SafePadding` | `float64` | `1.5` | Gap between the label and the line in mm |

## Usage notes

- The component always renders the label **above** the line, centered within the cell.
- Customize `LineColor` and `FontColor` independently for branded documents.
- Use `linestyle.Dashed` for a more informal look.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/signature#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/signature#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/signature#NewRow)
* [props : Signature](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Signature)
* [component : Signature](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/signature#Signature)

## Code Example
[filename](../../assets/examples/signaturegrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/signaturegridv2.pdf
```

## Time Execution
[filename](../../assets/text/signaturegridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/signaturegrid.json  ':include :type=code')