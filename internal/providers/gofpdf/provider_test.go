package gofpdf_test

import (
	"errors"
	"fmt"
	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/internal/merror"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Act
	sut := gofpdf.New(&gofpdf.Dependencies{})

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*gofpdf.provider", fmt.Sprintf("%T", sut))
}

func TestProvider_AddText(t *testing.T) {
	// Arrange
	txtContent := "text"
	cell := &entity.Cell{}
	prop := fixture.TextProp()

	text := &mocks.Text{}
	text.EXPECT().Add(txtContent, cell, &prop)

	dep := &gofpdf.Dependencies{
		Text: text,
	}
	sut := gofpdf.New(dep)

	// Act
	sut.AddText(txtContent, cell, &prop)

	// Assert
	text.AssertNumberOfCalls(t, "Add", 1)
}

func TestProvider_GetTextHeight(t *testing.T) {
	// Arrange
	fontHeightToReturn := 10.0
	prop := fixture.FontProp()

	font := &mocks.Font{}
	font.EXPECT().GetHeight(prop.Family, prop.Style, prop.Size).Return(fontHeightToReturn)

	dep := &gofpdf.Dependencies{
		Font: font,
	}
	sut := gofpdf.New(dep)

	// Act
	fontHeight := sut.GetTextHeight(&prop)

	// Assert
	font.AssertNumberOfCalls(t, "GetHeight", 1)
	assert.Equal(t, fontHeightToReturn, fontHeight)
}

func TestProvider_AddLine(t *testing.T) {
	// Arrange
	cell := &entity.Cell{}
	prop := fixture.LineProp()

	line := &mocks.Line{}
	line.EXPECT().Add(cell, &prop)

	dep := &gofpdf.Dependencies{
		Line: line,
	}
	sut := gofpdf.New(dep)

	// Act
	sut.AddLine(cell, &prop)

	// Assert
	line.AssertNumberOfCalls(t, "Add", 1)
}

func TestProvider_AddMatrixCode(t *testing.T) {
	t.Run("when cannot find image on cache and cannot generate data matrix, should apply error message", func(t *testing.T) {
		// Arrange
		codeContent := "code"
		cell := &entity.Cell{}
		prop := fixture.RectProp()

		cache := &mocks.Cache{}
		cache.EXPECT().GetImage(codeContent, extension.Jpg).Return(nil, errors.New("anyError1"))

		code := &mocks.Code{}
		code.EXPECT().GenDataMatrix(codeContent).Return(nil, errors.New("anyError2"))

		text := &mocks.Text{}
		text.EXPECT().Add("could not generate matrixcode", cell, merror.DefaultErrorText)

		dep := &gofpdf.Dependencies{
			Cache: cache,
			Code:  code,
			Text:  text,
		}

		sut := gofpdf.New(dep)

		// Act
		sut.AddMatrixCode(codeContent, cell, &prop)

		// Assert

	})
}
