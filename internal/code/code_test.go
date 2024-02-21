package code_test

import (
	"fmt"
	"github.com/johnfercher/maroto/v2/pkg/consts/barcode"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/internal/code"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("constructor", func(t *testing.T) {
		// Act
		sut := code.New()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*code.code", fmt.Sprintf("%T", sut))
	})
	t.Run("singleton is applied", func(t *testing.T) {
		// Act
		sut1 := code.New()
		sut2 := code.New()

		// Assert
		assert.NotNil(t, sut1)
		assert.NotNil(t, sut2)
	})
}

func TestCode_GenDataMatrix(t *testing.T) {
	t.Run("When cannot generate data matrix, should return error", func(t *testing.T) {
		// Arrange
		sut := code.New()

		data := genStringWithLength(5000)

		// Act
		bytes, err := sut.GenDataMatrix(data)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, bytes)
	})
	t.Run("When can generate data matrix, should return bytes", func(t *testing.T) {
		// Arrange
		sut := code.New()

		data := genStringWithLength(50)

		// Act
		bytes, err := sut.GenDataMatrix(data)

		// Assert
		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
}

func TestCode_GenBar(t *testing.T) {
	t.Run("When cannot generate bar code, should return error", func(t *testing.T) {
		// Arrange
		sut := code.New()

		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}

		prop := &props.Barcode{}
		prop.MakeValid()

		data := genStringWithLength(5000)

		// Act
		bytes, err := sut.GenBar(data, cell, prop)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, bytes)
	})
	t.Run("When can generate bar code, should return bytes", func(t *testing.T) {
		// Arrange
		sut := code.New()

		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}

		prop := &props.Barcode{}
		prop.MakeValid()

		data := genStringWithLength(60)

		// Act
		bytes, err := sut.GenBar(data, cell, prop)

		// Assert
		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
	t.Run("When is ean and can generate bar code, should return bytes", func(t *testing.T) {
		// Arrange
		sut := code.New()

		cell := &entity.Cell{
			X:      10,
			Y:      10,
			Width:  100,
			Height: 100,
		}

		prop := &props.Barcode{
			Type: barcode.EAN,
		}
		prop.MakeValid()

		// Act
		bytes, err := sut.GenBar("123456789123", cell, prop)

		// Assert
		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
}

func TestCode_GenQr(t *testing.T) {
	t.Run("When cannot generate qr code, should return error", func(t *testing.T) {
		// Arrange
		sut := code.New()

		data := genStringWithLength(5000)

		// Act
		bytes, err := sut.GenQr(data)

		// Assert
		assert.NotNil(t, err)
		assert.Nil(t, bytes)
	})
	t.Run("When can generate qr code, should return bytes", func(t *testing.T) {
		// Arrange
		sut := code.New()

		data := genStringWithLength(50)

		// Act
		bytes, err := sut.GenQr(data)

		// Assert
		assert.NotNil(t, bytes)
		assert.Nil(t, err)
	})
}

func genStringWithLength(length int) string {
	var content string
	for i := 0; i < length; i++ {
		content += "a"
	}
	return content
}
