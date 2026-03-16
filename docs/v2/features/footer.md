# Footer

`RegisterFooter` registers a row (or a set of rows) that is automatically printed at the bottom of **every** page. The footer is drawn after page content, just above the bottom margin. It is ideal for company names, legal disclaimers, page references, or any repeated bottom-of-page content.

## Usage notes

- `RegisterFooter` accepts one or more `core.Row` values; they are stacked in the order given.
- The footer rows consume vertical space — maroto deducts their total height from the usable page area so that body content never overlaps the footer.
- Call `RegisterFooter` once before generating any content; calling it again replaces the previous footer.
- Returns an error if the footer height exceeds the page's usable height.

## GoDoc
* [maroto : RegisterFooter](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.RegisterFooter)

## Code Example
[filename](../../assets/examples/footer/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/footerv2.pdf
```

## Time Execution
[filename](../../assets/text/footerv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/footer.json  ':include :type=code')