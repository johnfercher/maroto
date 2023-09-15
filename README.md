# Maroto V2

[![GoDoc](https://godoc.org/github.com/johnfercher/maroto?status.svg)](https://godoc.org/github.com/johnfercher/maroto)
[![Go Report Card](https://goreportcard.com/badge/github.com/johnfercher/maroto)](https://goreportcard.com/report/github.com/johnfercher/maroto)
[![Mentioned in Awesome Go](https://awesome.re/mentioned-badge.svg)](https://github.com/avelino/awesome-go#template-engines)  
[![CI](https://github.com/johnfercher/maroto/actions/workflows/goci.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/goci.yml)
[![Lint](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml/badge.svg)](https://github.com/johnfercher/maroto/actions/workflows/golangci-lint.yml)
[![Codecov](https://img.shields.io/codecov/c/github/johnfercher/maroto)](https://codecov.io/gh/johnfercher/maroto)

A Maroto way to create PDFs. Maroto is inspired in Bootstrap and uses [Gofpdf](https://github.com/jung-kurt/gofpdf). Fast and simple.

> Maroto definition: Brazilian expression, means an astute/clever/intelligent person.

You can write your PDFs like you are creating a site using Bootstrap. A Row may have many Cols, and a Col may have many components. 
Besides that, pages will be added when content may extrapolate the useful area. You can define a header which will be added
always when a new page appear, in this case, a header may have many rows, lines or tablelist. 

* You can see the full documentation [here](https://maroto.io/).
* The RFC is being created [here](https://docs.google.com/document/u/1/d/1OMOdtR2c6FYfl4X5p6Y_KTId_GeONfGbZSebnK0Pxbk/edit).
* Discussions are being addressed in [this issue](https://github.com/johnfercher/maroto/issues/257).

## Installation

* With `go get`:

```bash
go get -u github.com/johnfercher/maroto/internal
```

## Contributing

| Command        | Description                                       | Dependencies                                                 |
|----------------|---------------------------------------------------|--------------------------------------------------------------|
| `make build`   | Build project                                     | `go`                                                         |
| `make test`    | Run unit tests                                    | `go`                                                         |
| `make fmt`     | Format files                                      | `gofmt`, `gofumpt` and `goimports`                           |
| `make lint`    | Check files                                       | `golangci-lint` and `goreportcard-cli`                       |
| `make dod`     | (Definition of Done) Format files and check files | Same as`make build`, `make test`, `make fmt` and `make lint` | 
| `make install` | Install all dependencies                          | `go`, `curl` and `git`                                       |
| `make v1`      | Run all v1 examples                               | `go`                                                         |
| `make v2`      | Run all v2 examples                               | `go`                                                         |

### Example
![result](docs/assets/images/result.png)

### Code

```go
package main

import (
	"encoding/base64"
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/config"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/provider"
	"github.com/johnfercher/maroto/pkg/v2/signature"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"log"
	"os"
)

func main() {
	pdf := buildMarotoPDF()
	html := buildMarotoHTML()

	gen(pdf)
	gen(html)
}

func buildMarotoPDF() domain.MarotoMetrified {
	m := v2.NewMaroto("v2.pdf")
	return v2.NewMarotoMetrified(m)
}

func buildMarotoHTML() domain.MarotoMetrified {
	builder := config.NewBuilder().
		WithPageSize(config.A4).
		WithProvider(provider.HTML)

	m := v2.NewMaroto("v2.html", builder)
	return v2.NewMarotoMetrified(m)
}

func gen(m domain.MarotoMetrified) {
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())
	m.Add(buildCodesRow(), buildImagesRow(), buildTextsRow())

	report, err := m.GenerateWithReport()
	if err != nil {
		log.Fatal(err.Error())
	}

	report.Print()
}

func buildCodesRow() domain.Row {
	r := row.New(70)

	col1 := col.New(4)
	col1.Add(code.NewBar("barcode"))

	col2 := col.New(4)
	col2.Add(code.NewQr("qrcode"))

	col3 := col.New(4)
	col3.Add(code.NewMatrix("matrixcode"))

	r.Add(col1, col2, col3)
	return r
}

func buildImagesRow() domain.Row {
	row := row.New(70)

	col1 := col.New(6)
	col1.Add(image.NewFromFile("internal/assets/images/biplane.jpg"))

	byteSlices, err := os.ReadFile("internal/assets/images/gopherbw.png")
	if err != nil {
		fmt.Println("Got error while opening file:", err)
		os.Exit(1)
	}
	stringBase64 := base64.StdEncoding.EncodeToString(byteSlices)
	col2 := col.New(6)
	col2.Add(image.NewFromBase64(stringBase64, consts.Png))

	row.Add(col1, col2)

	return row
}

func buildTextsRow() domain.Row {
	row := row.New(70)

	colText := "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Donec ac condimentum sem."
	col1 := col.New(6)
	col1.Add(text.New(colText, props.Text{
		Align: consts.Center,
	}))

	col2 := col.New(6)
	col2.Add(signature.New("Fulano de Tal", props.Font{
		Style:  consts.Italic,
		Size:   20,
		Family: consts.Courier,
	}))

	row.Add(col1, col2)

	return row
}
```

## Stargazers over time

[![Stargazers over time](https://starchart.cc/johnfercher/maroto.svg)](https://starchart.cc/johnfercher/maroto)
