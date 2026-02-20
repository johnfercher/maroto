# Rich Text

Rich Text allows rendering multiple text styles (bold, italic, colors, sizes) within a single
flowing paragraph. Unlike the standard `text` component which applies one style to the entire
string, `richtext` accepts a slice of `Chunk`s, each with its own style, and handles word
wrapping across style boundaries automatically.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#New)
* [constructor : NewCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewCol)
* [constructor : NewRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewRow)
* [constructor : NewAutoRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewAutoRow)
* [constructor : NewChunk](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#NewChunk)
* [type : Chunk](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/richtext#Chunk)
* [props : Text](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Text)

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
