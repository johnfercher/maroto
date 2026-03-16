# Auto Row

Auto Row removes the need to hard-code row heights. When you call `m.AddAutoRow(cols...)` or use a convenience constructor such as `text.NewAutoRow(...)`, maroto measures each column's content and sets the row height to the tallest column automatically.

This is particularly useful for text blocks of unknown length, dynamic lists, or any content that varies between documents.

## Usage notes

- For **images and 2-D codes** inside auto rows, set `JustReferenceWidth: true` in `props.Rect`. Without it, the height calculation ignores the image's intrinsic aspect ratio.
- For **lines**, the auto-row height equals the line `Thickness`.
- For **signatures**, height is derived from the font size plus `SafePadding`.
- Multiple components in different columns are measured independently; the row expands to the maximum.
- Auto rows can be mixed freely with fixed-height rows in the same document.

## GoDoc
* [maroto : AddAutoRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.AddAutoRow)
* [row : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/row#New)
* [text : NewAutoRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/text#NewAutoRow)
* [image : NewAutoFromFileRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewAutoFromFileRow)

## Code Example
[filename](../../assets/examples/autorow/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/autorow.pdf
```

## Time Execution
[filename](../../assets/text/autorow.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/autorow.json  ':include :type=code')
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/page#New) 
* [interface : Page](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core#Page)
* [props : Page](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Page)

## Code Example
[filename](../../assets/examples/autorow/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/autorow.pdf
```

## Time Execution
[filename](../../assets/text/autorow.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/autorow.json  ':include :type=code')