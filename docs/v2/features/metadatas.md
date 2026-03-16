# Metadatas

PDF metadata fields are stored in the document's information dictionary and are visible in PDF viewer properties dialogs, search indexes, and archiving systems. maroto exposes them through builder methods.

## Usage notes

- Metadata does not appear in the rendered PDF content — it is only stored in the document's information dictionary.
- All fields are optional; omit any method for which you do not have a value.
- `WithCreationDate` accepts a `time.Time` value; maroto formats it according to the PDF specification.

## GoDoc
* [builder : WithAuthor](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithAuthor)
* [builder : WithCreationDate](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithCreationDate)
* [builder : WithCreator](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithCreator)
* [builder : WithSubject](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithSubject)
* [builder : WithTitle](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithTitle)
* [builder : WithKeywords](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/config#CfgBuilder.WithKeywords)

## Code Example
[filename](../../assets/examples/metadatas/v2/main.go ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/metadatasv2.pdf
```

## Time Execution
[filename](../../assets/text/metadatasv2.txt  ':include :type=code')

## Test File
[filename](https://raw.githubusercontent.com/johnfercher/maroto/master/test/maroto/examples/metadatas.json  ':include :type=code')