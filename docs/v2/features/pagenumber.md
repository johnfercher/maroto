# Page Number

`WithPageNumber` adds automatic page numbering to every page. The number is rendered in the margin area at a configurable position and can be styled like any other text.

## `props.PageNumber` fields

| Field | Type | Description |
|-------|------|-------------|
| `Pattern` | `string` | Text pattern; `{current}` and `{total}` are replaced at render time |
| `Place` | `props.Place` | Where on the page to render the number |
| `Family` | `string` | Font family |
| `Style` | `fontstyle.Type` | Font style (Normal, Bold, Italic, BoldItalic) |
| `Size` | `float64` | Font size in points |
| `Color` | `*props.Color` | Text color; `nil` uses black |

## Placement options (`props.Place`)

| Constant | Description |
|----------|-------------|
| `props.LeftTop` | Top-left corner |
| `props.Top` | Top centre |
| `props.RightTop` | Top-right corner |
| `props.LeftBottom` | Bottom-left corner |
| `props.Bottom` | Bottom centre |
| `props.RightBottom` | Bottom-right corner |

## Usage notes

- The `{current}` placeholder is replaced with the current page number; `{total}` is replaced with the total page count.
- Page numbers are injected during the final rendering pass, so `{total}` is always accurate even for multi-page documents.
- Page number rendering does not overlap body content — it is drawn in the margin.

## GoDoc
* [builder : WithPageNumber](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithPageNumber)
* [props : Place](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Place)

## Code Example
[filename](../../assets/examples/pagenumber/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/pagenumberv2.pdf
```

## Time Execution
[filename](../../assets/text/pagenumberv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/pagenumber.json  ':include :type=code')