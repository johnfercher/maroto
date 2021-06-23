/*
Package docs is the documentation of maroto. Maroto is a package which
provide a simple way to generate PDF documents. The project is inspired
in Bootstrap and uses gofpdf. Simple and Fast


Features and Components

-   Grid system with rows and columns

-   Automatic page breaks

-   Inclusion of JPEG, PNG, GIF, TIFF and basic path-only SVG images

-   Lines

-   Barcodes

-   Qrcodes

-   DataMatrix codes

-   Signatures

Maroto has only gofpdf dependency. All tests pass on Linux and Mac.


Installation

To install the package on your system, run

    go get github.com/johnfercher/maroto

Later, to receive updates, run

    go get -u -v github.com/johnfercher/maroto/...


Quick Start

The following Go Code generates a simple PDF file.


    m := pdf.NewMaroto(consts.Portrait, consts.Letter)

    m.Row(10, func() {
		m.Col(12, func() {
			m.Text(props.Text{
				Size: 18,
				Style: consts.Bold,
				Align: consts.Center,
				Top: 9,
			})
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

Maroto is released under the MIT License.


Acknowledgments

This package’s Code and documentation are based on gofpdf.


Roadmap


-   Improve test coverage as reported by the coverage tool.
*/
package docs
