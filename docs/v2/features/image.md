# Image

The Image component renders a raster image (JPEG or PNG) inside a cell. Images can be loaded from a file path on disk or from a `[]byte` slice already in memory.

Both sources expose the same set of constructors — `New`, `NewCol`, `NewRow`, and `NewAutoRow` — so images integrate into the grid exactly like any other component.

## Props (`props.Rect`)

| Field | Type | Default | Description |
|-------|------|---------|-------------|
| `Percent` | `float64` | `100` | How much of the cell the image occupies (0–100) |
| `Center` | `bool` | `false` | Horizontally and vertically center the image |
| `Left` | `float64` | `0` | Left offset in mm — ignored when `Center` is true |
| `Top` | `float64` | `0` | Top offset in mm — ignored when `Center` is true |
| `JustReferenceWidth` | `bool` | `false` | Scale only by width, ignore available height — required for correct auto-row sizing |

## Usage notes

- Set `JustReferenceWidth: true` when placing an image inside `AddAutoRow` so that the row height is calculated from the image's proportional height relative to its width.
- `Percent` controls the largest dimension; the image is always scaled proportionally.
- When `Center` is `true`, `Left` and `Top` have no effect.

## GoDoc
* [constructor : NewFromBytes](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromBytes)
* [constructor : NewFromBytesCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromBytesCol)
* [constructor : NewFromBytesRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromBytesRow)
* [constructor : NewFromFile](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromFile)
* [constructor : NewFromFileCol](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromFileCol)
* [constructor : NewFromFileRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#NewFromFileRow)
* [props : Rect](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Rect)
* [component : BytesImage](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#BytesImage)
* [component : FileImage](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/image#FileImage)

## Code Example
[filename](../../assets/examples/imagegrid/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/imagegridv2.pdf
```

## Time Execution
[filename](../../assets/text/imagegridv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/imagegrid.json  ':include :type=code')
