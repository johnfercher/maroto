# Header

`RegisterHeader` registers a row (or a set of rows) that is automatically printed at the top of **every** page. The header appears just below the top margin and is drawn before any body content on each page. Use it for logos, report titles, column labels, or any content that should repeat on every page.

## Usage notes

- `RegisterHeader` accepts one or more `core.Row` values; they are stacked top-to-bottom in the order given.
- Header rows consume vertical space — maroto deducts their total height from the usable area on every page so body content starts below the header.
- Call `RegisterHeader` once before generating any content; calling it again replaces the previous header.
- Returns an error if the header rows are taller than the page's usable height.

## GoDoc
* [maroto : RegisterHeader](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.RegisterHeader)

## Code Example
[filename](../../assets/examples/header/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/headerv2.pdf
```

## Time Execution
[filename](../../assets/text/headerv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/header.json  ':include :type=code')