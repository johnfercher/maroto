# Max Grid Sum

`WithMaxGridSize` changes the total number of grid columns that the page width is divided into. maroto defaults to **12 columns**, following the Bootstrap-style grid convention. Increasing or decreasing this value changes the granularity of the layout.

## Usage notes

- All `col.New(n)` calls in the document must use values relative to the configured grid size; columns wider than the grid size are clamped.
- The total column widths in a row should sum to `maxGridSize`. Rows whose columns sum to less than `maxGridSize` leave empty space on the right.
- Change the grid size only at the document level; it cannot be changed per row.

## GoDoc
* [builder : WithMaxGridSize](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithMaxGridSize)

## Code Example
[filename](../../assets/examples/maxgridsum/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/maxgridsumv2.pdf
```

## Time Execution
[filename](../../assets/text/maxgridsumv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/maxgridsum.json  ':include :type=code')