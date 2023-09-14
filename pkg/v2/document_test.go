package v2_test

import (
	"fmt"
	"github.com/johnfercher/maroto/pkg/consts"
	v2 "github.com/johnfercher/maroto/pkg/v2"
	"github.com/johnfercher/maroto/pkg/v2/code/barcode"
	"github.com/johnfercher/maroto/pkg/v2/code/matrixcode"
	"github.com/johnfercher/maroto/pkg/v2/code/qrcode"
	"github.com/johnfercher/maroto/pkg/v2/col"
	"github.com/johnfercher/maroto/pkg/v2/image"
	"github.com/johnfercher/maroto/pkg/v2/row"
	"github.com/johnfercher/maroto/pkg/v2/signature"
	"github.com/johnfercher/maroto/pkg/v2/text"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewDocument(t *testing.T) {
	// Act
	sut := v2.NewMaroto("file.pdf")

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*v2.document", fmt.Sprintf("%T", sut))
}

func TestDocument_GetStructure(t *testing.T) {
	// Arrange
	p := v2.NewMaroto("file.txt")

	r1 := row.New(10)
	r1c1 := col.New(4).Add(barcode.New("barcode"))
	r1c2 := col.New(4).Add(matrixcode.New("matrixcode"))
	r1c3 := col.New(4).Add(qrcode.New("qrcode"))
	r1.Add(r1c1, r1c2, r1c3)

	r2 := row.New(10)
	r2c1 := col.New(3).Add(image.NewFromFile("file.png"))
	r2c2 := col.New(3).Add(image.NewFromBase64("base64string", consts.Png))
	r2c3 := col.New(3).Add(signature.New("signature"))
	r2c4 := col.New(3).Add(text.New("text"))
	r2.Add(r2c1, r2c2, r2c3, r2c4)

	p.Add(r1, r2)

	// Act
	nodeDocument := p.GetStructure()

	// Assert Document
	assert.NotNil(t, nodeDocument)
	_, document := nodeDocument.Get()
	assert.Equal(t, "document", document.Type)
	assert.Equal(t, "file.txt", document.Value)

	nodeRows := nodeDocument.GetNexts()
	assert.Equal(t, 2, len(nodeRows))
}
