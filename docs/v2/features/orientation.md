# Orientation

`WithOrientation` sets the page orientation. maroto defaults to **portrait** (taller than wide). Switching to **landscape** rotates the page 90°, making it wider than tall — useful for wide tables, timelines, or presentations.

## Usage notes

- When switching to landscape, maroto swaps the page width and height automatically, so the grid still divides the (now wider) page correctly.
- Combine with `WithPageSize` or `WithDimensions` to get the exact physical size and orientation you need.
- All pages in a single document share the same orientation; mixed-orientation documents are not directly supported. Use `merge.Bytes` to combine separately generated portrait and landscape documents.

## GoDoc
* [builder : WithOrientation](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithOrientation)
* [orientation : Type](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/consts/orientation)

## Code Example
[filename](../../assets/examples/orientation/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/orientationv2.pdf
```

## Time Execution
[filename](../../assets/text/orientationv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/orientation.json  ':include :type=code')