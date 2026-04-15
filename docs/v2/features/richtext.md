# Rich Text

The Rich Text component renders multiple styled chunks inside a single flowing paragraph. It is useful when one sentence needs partial emphasis such as bold words, mixed colors, italic fragments, or different font sizes without splitting the content across multiple columns.

`richtext` keeps wrapping behavior across chunk boundaries, so the paragraph still behaves like one text block instead of several disconnected components.

## Props (`props.Text` per chunk)

| Scope | Fields |
|-------|--------|
| Paragraph, read from the first chunk only | `Align`, `Top`, `Bottom`, `Left`, `Right`, `VerticalPadding`, `PreserveLineBreaks` |
| Every chunk | `Family`, `Style`, `Size`, `Color`, `Hyperlink` (zero values take the config defaults) |

## Usage notes

- Lines wrap only at whitespace; chunks with no space between them stay on the same line, and a word wider than the column overflows. `BreakLineStrategy` is ignored.
- Each line is as tall as the largest chunk font on it. `Justify` stretches the spaces of every line except the last.
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
