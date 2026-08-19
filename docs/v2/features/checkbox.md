# Checkbox

The Checkbox component renders a square checkbox with an optional label to its right. When `Checked` is true, an X mark is drawn inside the box. This component is useful for forms, questionnaires, and agreements.

The row height for auto-row usage is `Size + Top`.

## Props (`props.Checkbox`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Checked` | `bool` | `false` | Whether the checkbox is marked with an X |
| `Size` | `float64` | `5.0` | Side length of the checkbox square in mm |
| `Top` | `float64` | `0` | Space between the upper cell limit and the checkbox (mm). Accepts negative values to nudge the checkbox above the cell top. |
| `Left` | `float64` | `0` | Space between the left cell boundary and the checkbox (mm). Accepts negative values. |

## Usage notes

- The label is rendered to the right of the box using the document's default font; font styling is derived from the active `core.Font`.
- `Top` and `Left` accept negative values (rendered slightly outside the cell on that side); `Size` is still required to be > 0 and falls back to the default `5.0` otherwise.
- For forms with multiple options, place several `NewCol` checkboxes side-by-side in the same row.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/checkbox#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/checkbox#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/checkbox#NewRow)
* [constructor : NewAutoRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/checkbox#NewAutoRow)
* [props : Checkbox](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Checkbox)
* [component : Checkbox](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/checkbox#Checkbox)

## Code Example
[filename](../../assets/examples/checkbox/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/checkboxv2.pdf
```

## Time Execution
[filename](../../assets/text/checkboxv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/checkbox.json  ':include :type=code')
