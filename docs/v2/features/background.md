# Add Background

`WithBackgroundImage` stamps a single image on every page behind all other content. The image is stretched to fill the entire page area, making it ideal for letterheads, watermarks, and branded templates.

## Usage notes

- The background is rendered before any rows or columns, so it never obscures content.
- Combine with `WithOrientation` and custom margins to align the template with the layout.
- Any image format supported by the underlying PDF library (JPEG, PNG) can be used; the extension argument tells maroto how to decode the file.

## GoDoc
* [builder : WithBackgroundImage](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithBackgroundImage)
* [consts : Extension](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/consts/extension#Type)

## Code Example
[filename](../../assets/examples/background/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/backgroundv2.pdf
```

## Time Execution
[filename](../../assets/text/backgroundv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/background.json  ':include :type=code')