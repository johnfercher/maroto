package v2_test

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code"
	"github.com/johnfercher/maroto/pkg/v2/grid/col"
	"github.com/johnfercher/maroto/pkg/v2/grid/row"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/signature"
	"github.com/johnfercher/maroto/pkg/v2/test"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDocument(t *testing.T) {
	// Act
	sut := v2.NewMaroto()

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*v2.maroto", fmt.Sprintf("%T", sut))
}

func TestDocument_GetStructure(t *testing.T) {
	// Arrange
	maroto := v2.NewMaroto()

	r1 := row.New(10)
	r1c1 := col.New(4).Add(code.NewBar("barcode"))
	r1c2 := col.New(4).Add(code.NewMatrix("matrixcode"))
	r1c3 := col.New(4).Add(code.NewQr("qrcode"))
	r1.Add(r1c1, r1c2, r1c3)

	r2 := row.New(10)
	r2c1 := col.New(3).Add(image.NewFromFile("file.png"))
	r2c2 := col.New(3).Add(image.NewFromBase64("base64string", consts.Png))
	r2c3 := col.New(3).Add(signature.New("signature"))
	r2c4 := col.New(3).Add(text.New("text"))
	r2.Add(r2c1, r2c2, r2c3, r2c4)

	maroto.AddRows(r1, r2)

	// Assert
	test.New(t).Assert(maroto).JSON(`{"type":"maroto","nodes":[{"type":"page","nodes":[{"type":"row","nodes":[{"type":"col","nodes":[{"type":"barcode"}]},{"type":"col","nodes":[{"type":"matrixcode"}]},{"type":"col","nodes":[{"type":"qrcode"}]}]},{"type":"row","nodes":[{"type":"col","nodes":[{"type":"fileimage"}]},{"type":"col","nodes":[{"type":"base64image"}]},{"type":"col","nodes":[{"type":"signature"}]},{"type":"col","nodes":[{"type":"text"}]}]},{"type":"row","nodes":[{"type":"col"}]}]}]}`)
}
