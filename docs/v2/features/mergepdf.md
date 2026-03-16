# Merge PDF

Maroto provides two complementary ways to combine PDF documents:

1. **`merge.Bytes`** — merges two or more raw PDF byte slices into one. Useful when you have independently generated PDFs (e.g., from different sources or libraries).
2. **`Document.Merge`** — appends the pages of another `core.Document` to the current one at runtime, before calling `Generate`.

## Usage notes

- `merge.Bytes` accepts two or more byte slices and returns `([]byte, error)`. It wraps `merge.ErrCannotMergePDFs` on failure.
- `Document.Merge` mutates the receiver document in place.
- Both approaches produce a flat concatenation of pages — bookmarks and internal links from the source documents are preserved where the underlying library supports it.
- For documents generated entirely within maroto, using `AddPages` to compose sections before `Generate` is simpler and avoids a second merge step.

## GoDoc
* [merge : Bytes](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/merge#Bytes)
* [interface : Document](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core#Document)
* [pdf : Merge](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/core#Pdf.Merge)

## Code Example
[filename](../../assets/examples/mergepdf/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/mergepdfv2.pdf
```

## Time Execution
[filename](../../assets/text/mergepdfv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/mergepdf.json  ':include :type=code')