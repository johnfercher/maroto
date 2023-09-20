package main

import (
	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"testing"
)

func TestDocument_GetStructure(t *testing.T) {
	// Arrange
	m := pkg.NewMaroto()

	m.AddRow(10,
		code.NewBarCol(4, "barcode"),
		code.NewMatrixCol(4, "matrixcode"),
		code.NewQrCol(4, "qrcode"),
	)

	m.AddRow(10,
		image.NewFromFileCol(3, "barcode"),
		image.NewFromBase64Col(3, "base64string", extension.Png),
		signature.NewCol(3, "signature"),
		text.NewCol(3, "text"),
	)

	// nolint: lll
	test.New(t).Assert(m).JSON(`{"type":"pkg","nodes":[{"type":"page","nodes":[{"type":"row","nodes":[{"type":"col","nodes":[{"type":"barcode"}]},{"type":"col","nodes":[{"type":"matrixcode"}]},{"type":"col","nodes":[{"type":"qrcode"}]}]},{"type":"row","nodes":[{"type":"col","nodes":[{"type":"fileimage"}]},{"type":"col","nodes":[{"type":"base64image"}]},{"type":"col","nodes":[{"type":"signature"}]},{"type":"col","nodes":[{"type":"text"}]}]},{"type":"row","nodes":[{"type":"col"}]}]}]}`)
}
