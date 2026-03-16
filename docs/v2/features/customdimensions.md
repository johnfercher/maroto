# Custom Dimensions

`WithDimensions` sets an arbitrary page width and height in millimetres, overriding the standard page size. Use this when you need non-standard formats such as receipt rolls, labels, or custom card stock.

## Usage notes

- `WithDimensions` takes precedence over `WithPageSize`. Do not combine them.
- The grid system still divides the usable width (page width minus left and right margins) into the configured number of columns (default 12).
- Very narrow pages may cause text to overflow if column widths become smaller than a single word.

## GoDoc
* [builder : WithDimensions](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithDimensions)

## Code Example
[filename](../../assets/examples/customdimensions/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/customdimensionsv2.pdf
```
## Time Execution
[filename](../../assets/text/customdimensionsv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/customdimensions.json  ':include :type=code')