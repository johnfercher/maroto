# Basics

Maroto uses a **grid-based layout** to compose PDF documents. Every page is divided into a fixed number of columns (default: 12) and an unlimited number of rows. Components such as text, images, barcodes and lines are placed inside columns, and columns are grouped into rows.

The entry point is `maroto.New()`, which accepts an optional `*entity.Config` produced by `config.NewBuilder()`. Once the document is configured, content is added via `AddRow`, `AddRows`, `AddAutoRow`, or `AddPages`. Finally, `Generate()` returns a `Document` that can be saved to disk or exported as bytes.

## GoDoc
* [maroto : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2#New)
* [maroto : AddRow](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.AddRow)
* [maroto : AddRows](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.AddRows)
* [maroto : Generate](https://pkg.go.dev/github.com/johnfercher/maroto/v2#Maroto.Generate)
* [row : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/row#New)
* [row : Add](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/row#Row.Add)
* [col : New](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/col#New)
* [col : Add](https://pkg.go.dev/github.com/johnfercher/maroto/v2/pkg/components/col#Col.Add)

### Components
* [Barcode](v2/features/barcode?id=barcode)
* [DataMatrix](v2/features/datamatrix?id=data-matrix)
* [Image](v2/features/image?id=image)
* [Line](v2/features/line?id=line)
* [QrCode](v2/features/qrcode?id=qrcode)
* [Signature](v2/features/signature?id=signature)
* [Text](v2/features/text?id=text)
* [Checkbox](v2/features/checkbox?id=checkbox)

## Code Example
[filename](../../assets/examples/simplest/v2/main.go  ':include :type=code')

## PDF Generated
```pdf
	assets/pdf/simplestv2.pdf
```
