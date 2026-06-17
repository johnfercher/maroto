# Rich Text

The Rich Text component renders multiple styled chunks inside a single flowing paragraph. It is useful when one sentence needs partial emphasis such as bold words, mixed colors, italic fragments, or different font sizes without splitting the content across multiple columns.

`richtext` keeps wrapping behavior across chunk boundaries, so the paragraph still behaves like one text block instead of several disconnected components.

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
