# Custom Page

`WithPageSize` selects a standard paper size from the `pagesize` package. maroto defaults to `pagesize.A4`; use this option to switch to any other ISO or North-American paper size.

## Available page sizes

| Constant           | Dimensions (mm) |
|--------------------|----------------|
| `pagesize.A1`      | 594 × 841 |
| `pagesize.A2`      | 419 × 594 |
| `pagesize.A3`      | 297 × 420 |
| `pagesize.A4`      | 210 × 297 (default) |
| `pagesize.A5`      | 148 × 210 |
| `pagesize.A6`      | 105 × 148 |
| `pagesize.Letter`  | 215.9 × 279.4 |
| `pagesize.Legal`   | 215.9 × 355.6 |
| `pagesize.Tabloid` | 279.4 × 431.8 |

## Usage notes

- For non-standard dimensions use [`WithDimensions`](customdimensions.md) instead.
- `WithPageSize` and `WithDimensions` are mutually exclusive; the last one called wins.
- Page orientation is separate: combine this with `WithOrientation(orientation.Horizontal)` to get landscape A4, for example.

## GoDoc
* [builder : WithPageSize](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithPageSize)
* [pagesize : Type](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/consts/pagesize)

## Code Example
[filename](../../assets/examples/custompage/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/custompagev2.pdf
```
## Time Execution
[filename](../../assets/text/custompagev2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/custompage.json  ':include :type=code')