# Add Page

`AddPages` lets you push complete, pre-built page objects into the document. A `page.Page` is a container for rows — it behaves just like a normal document section but is logically isolated. Rows added to a page that overflow its usable area are automatically split across additional physical pages.

This is useful when you want to control pagination explicitly: for example, forcing a chapter to always start on a new page, or composing a document from independently generated sections.

## GoDoc
* [constructor : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/page#New) 
* [interface : Page](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core#Page)
* [props : Page](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/props#Page)

## Code Example
[filename](../../assets/examples/addpage/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/addpagev2.pdf
```

## Time Execution
[filename](../../assets/text/addpagev2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/addpage.json  ':include :type=code')