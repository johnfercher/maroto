# Text

The Text component renders a string inside a cell. It supports custom fonts, alignment, colors, hyperlinks, and automatic line breaking when the text is wider than the available column.

Text can be created as a standalone `Component`, wrapped directly into a `Col`, or wrapped into a `Row` (fixed or automatic height). For long content that should expand the row height automatically, use `NewAutoRow` or `NewCol` inside `AddAutoRow`.

## Props

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Family` | `string` | global font | Font family (e.g. `fontfamily.Arial`) |
| `Style` | `fontstyle.Type` | global font | `Normal`, `Bold`, `Italic`, `BoldItalic`, `Strikethrough` |
| `Size` | `float64` | global font | Font size in points |
| `Color` | `*props.Color` | global font | Font color |
| `Align` | `align.Type` | `align.Left` | `Left`, `Center`, `Right`, `Justify` |
| `Top` | `float64` | `0` | Top offset inside the cell (mm) |
| `Bottom` | `float64` | `0` | Bottom offset — used by auto rows only (mm) |
| `Left` | `float64` | `0` | Left margin inside the cell (mm) |
| `Right` | `float64` | `0` | Right margin inside the cell (mm) |
| `BreakLineStrategy` | `breakline.Strategy` | `EmptySpaceStrategy` | `EmptySpaceStrategy` breaks on spaces; `DashStrategy` breaks mid-word with a hyphen; `CharacterStrategy` breaks at character boundaries without adding symbols |
| `VerticalPadding` | `float64` | `0` | Extra spacing between lines (mm) |
| `Hyperlink` | `*string` | `nil` | URL — makes the text a clickable link (rendered in blue) |

## Usage notes

- When `Hyperlink` is set, the text color is overridden with blue regardless of `Color`.
- `Top` and `Left`/`Right` are clamped to the cell dimensions if they exceed it.
- `BreakLineStrategy` only applies when the text does not fit on a single line.
- Use `CharacterStrategy` when a text should wrap without spaces and without inserting trailing hyphens.
- For justified text on the last line, spacing may revert to default space width to avoid stretching a few characters across the full width.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewRow)
* [props : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Text)
* [component : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#Text)


## Code Example
[filename](../../assets/examples/textgrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/textgridv2.pdf
```

## Time Execution
[filename](../../assets/text/textgridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/textgrid.json  ':include :type=code')
