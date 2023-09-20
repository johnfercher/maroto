package pkg_test

import (
	"fmt"
	"testing"

	code2 "github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	image2 "github.com/johnfercher/maroto/v2/pkg/components/image"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

	"github.com/johnfercher/maroto/v2/pkg/consts/extension"

	"github.com/johnfercher/maroto/v2/pkg"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
)

func TestNewDocument(t *testing.T) {
	// Act
	sut := pkg.NewMaroto()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*pkg.maroto", fmt.Sprintf("%T", sut))
}

func TestDocument_GetStructure(t *testing.T) {
	// Arrange
	maroto := pkg.NewMaroto()

	r1 := row.New(10)
	r1c1 := col.New(4).Add(code2.NewBar("barcode"))
	r1c2 := col.New(4).Add(code2.NewMatrix("matrixcode"))
	r1c3 := col.New(4).Add(code2.NewQr("qrcode"))
	r1.Add(r1c1, r1c2, r1c3)

	r2 := row.New(10)
	r2c1 := col.New(3).Add(image2.NewFromFile("file.png"))
	r2c2 := col.New(3).Add(image2.NewFromBase64("base64string", extension.Png))
	r2c3 := col.New(3).Add(signature.New("signature"))
	r2c4 := col.New(3).Add(text.New("text"))
	r2.Add(r2c1, r2c2, r2c3, r2c4)

	maroto.AddRows(r1, r2)

	// nolint: lll
	test.New(t).Assert(maroto).JSON(`{"type":"pkg","nodes":[{"type":"page","nodes":[{"type":"row","nodes":[{"type":"col","nodes":[{"type":"barcode"}]},{"type":"col","nodes":[{"type":"matrixcode"}]},{"type":"col","nodes":[{"type":"qrcode"}]}]},{"type":"row","nodes":[{"type":"col","nodes":[{"type":"fileimage"}]},{"type":"col","nodes":[{"type":"base64image"}]},{"type":"col","nodes":[{"type":"signature"}]},{"type":"col","nodes":[{"type":"text"}]}]},{"type":"row","nodes":[{"type":"col"}]}]}]}`)
}
