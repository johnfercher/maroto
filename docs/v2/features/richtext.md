# Rich Text

The Rich Text component renders multiple styled chunks inside a single flowing paragraph. It is useful when one sentence needs partial emphasis such as bold words, mixed colors, italic fragments, or different font sizes without splitting the content across multiple columns.

`richtext` keeps wrapping behavior across chunk boundaries, so the paragraph still behaves like one text block instead of several disconnected components.

## Props (`props.Text` per chunk)

Paragraph props are read from the first chunk only; setting them on later chunks has no effect. Font props apply to every chunk separately.

| Field | Type | Default | Scope | Description |
|-------|------|---------|-------|-------------|
| `Family` | `string` | global font | every chunk | Font family (e.g. `fontfamily.Arial`) |
| `Style` | `fontstyle.Type` | global font | every chunk | `Normal`, `Bold`, `Italic`, `BoldItalic`, `Strikethrough` |
| `Size` | `float64` | global font | every chunk | Font size in points |
| `Color` | `*props.Color` | global font | every chunk | Font color |
| `Hyperlink` | `*string` | `nil` | every chunk | URL — makes the chunk a clickable link (rendered in blue) |
| `Align` | `align.Type` | `align.Left` | paragraph | `Left`, `Center`, `Right`, `Justify` |
| `Top` | `float64` | `0` | paragraph | Top offset inside the cell (mm) |
| `Bottom` | `float64` | `0` | paragraph | Bottom offset — used by auto rows only (mm) |
| `Left` | `float64` | `0` | paragraph | Left margin inside the cell (mm) |
| `Right` | `float64` | `0` | paragraph | Right margin inside the cell (mm) |
| `VerticalPadding` | `float64` | `0` | paragraph | Extra spacing between lines (mm) |
| `PreserveLineBreaks` | `bool` | `false` | paragraph | Treat `\n`, `\r\n` and `\r` as hard line breaks instead of collapsing them into spaces |
| `BreakLineStrategy` | `breakline.Strategy` | — | — | Ignored: rich text always wraps at whitespace |

## Usage notes

- Lines wrap only at whitespace; chunks with no space between them stay on the same line, and a word wider than the column overflows.
- Each line is as tall as the largest chunk font on it. `Justify` stretches the spaces of every wrapped line except the last line of each paragraph.
- With `PreserveLineBreaks`, a trailing line break keeps an empty last line, and a `\r\n` split across two chunks counts as one break.
- A `RichText` without chunks renders nothing and has a height of `0`.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewRow)
* [constructor : NewAutoRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewAutoRow)
* [constructor : NewChunk](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewChunk)
* [type : Chunk](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#Chunk)

## Code Example
[filename](../../assets/examples/richtextgrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/richtextgridv2.pdf
```

## Time Execution
[filename](../../assets/text/richtextgridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/richtextgrid.json  ':include :type=code')
