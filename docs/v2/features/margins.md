# Custom Margins

`WithMargins` sets the page margins — the blank space between the physical page edge and the area where rows and columns are rendered. All values are in **millimetres**. maroto's default margin is 10 mm on every side.

## Usage notes

- Margins reduce the usable content area. A wider left margin, for example, narrows every column.
- The footer is drawn inside the bottom margin area; ensure the bottom margin is large enough to accommodate it.
- Setting a margin to `0` renders content flush to the physical page edge — some printers clip this area.

## GoDoc
* [builder : WithBottomMargin](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithBottomMargin)
* [builder : WithLeftMargin](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithLeftMargin)
* [builder : WithRightMargin](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithRightMargin)
* [builder : WithTopMargin](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithTopMargin)

## Code Example
[filename](../../assets/examples/margins/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/marginsv2.pdf
```

## Time Execution
[filename](../../assets/text/marginsv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/margins.json  ':include :type=code')