/*
Package Maroto provide a simple way to generate PDF documents.
Maroto is inspired in Bootstrap and uses gofpdf. Simple and Fast

Features & Components

-   Grid system with rows and columns

-   Automatic page breaks

-   Inclusion of JPEG, PNG, GIF, TIFF and basic path-only SVG images

-   Lines

-   Barcodes

-   Qrcodes

-   Signature

Maroto has only gofpdf dependency. All tests pass on Linux and Mac.


Installation

To install the package on your system, run

    go get github.com/johnfercher/maroto

Later, to receive updates, run

    go get -u -v github.com/johnfercher/maroto/...


Quick Start

The following Go code generates a simple PDF file.


    m := maroto.NewMaroto(maroto.Portrait, maroto.A4)

    m.Row("MyRow", 10, func() {
		m.Col("MyCol", func() {
			m.Text("MyText", maroto.Arial, maroto.Bold, 15, 7.5, maroto.Center)
		})
	})

	m.OutputFileAndClose("maroto.pdf")

See the functions in the maroto_test.go file (shown as examples in this
documentation) for more advanced PDF examples.


Conversion Notes

This package is an high level API from gofpdf. The original API
names have been slightly adapted. And the package search to be
simpler to use.

The main contribution upside gofpdf is the grid system with
high level components.


License

Maroto is released under the GPL3 License.


Acknowledgments

This package’s code and documentation are based on gofpdf.


Roadmap


-   Improve test coverage as reported by the coverage tool.
*/
package maroto
