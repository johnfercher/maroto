# Line

The Line component draws a straight line inside a cell. Lines can be horizontal (default) or vertical, solid or dashed, and are useful as section dividers, table separators, or decorative rules.

For auto-row usage, the row height equals the line's `Thickness` value.

## Props (`props.Line`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Style` | `linestyle.Type` | `linestyle.Solid` | `linestyle.Solid` or `linestyle.Dashed` |
| `Thickness` | `float64` | `0.2` | Line width in mm |
| `Orientation` | `orientation.Type` | `orientation.Horizontal` | `orientation.Horizontal` or `orientation.Vertical` |
| `Color` | `*props.Color` | `nil` (black) | Line color |
| `OffsetPercent` | `float64` | `5` | Position along the perpendicular axis (5–95). `50` centers the line |
| `SizePercent` | `float64` | `90` | Length as a percentage of the cell's parallel dimension (0–100) |

## Usage notes

- `OffsetPercent` is clamped to the range [5, 95] by `MakeValid`.
- `SizePercent` is clamped to the range (0, 100] by `MakeValid`; values ≤ 0 default to 90.
- For a full-width horizontal divider with default styling, `line.NewCol(12)` with no props is sufficient.
- Vertical lines are drawn from top to bottom within the cell; `OffsetPercent` controls horizontal position.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/line#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/line#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/line#NewRow)
* [props : Line](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Line)
* [component : Line](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/line#Line)

## Code Example
[filename](../../assets/examples/line/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/linegridv2.pdf
```
## Time Execution
[filename](../../assets/text/linegridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/line.json  ':include :type=code')