package signature_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/internal/fixture"
	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/signature"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNew(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := signature.New("signature")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := signature.New("signature", fixture.SignatureProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_custom_prop.json")
	})
}

func TestNewCol(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := signature.NewCol(12, "signature")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_col_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := signature.NewCol(12, "signature", fixture.SignatureProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_col_custom_prop.json")
	})
}

func TestNewRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := signature.NewRow(10, "signature")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := signature.NewRow(10, "signature", fixture.SignatureProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_row_custom_prop.json")
	})
}

func TestNewAutoRow(t *testing.T) {
	t.Run("when prop is not sent, should use default", func(t *testing.T) {
		// Act
		sut := signature.NewAutoRow("signature")

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_auto_row_default_prop.json")
	})
	t.Run("when prop is sent, should use the provided", func(t *testing.T) {
		// Act
		sut := signature.NewAutoRow("signature", fixture.SignatureProp())

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("components/signatures/new_signature_auto_row_custom_prop.json")
	})
}

func TestSignature_Render(t *testing.T) {
	t.Run("should call provider correctly", func(t *testing.T) {
		// Arrange
		label := "signature"
		cell := fixture.CellEntity()
		prop := fixture.SignatureProp()
		sut := signature.New(label, prop)

		provider := mocks.NewProvider(t)
		provider.On("AddText", mock.Anything, mock.Anything, mock.Anything).Return(10.0)
		provider.On("GetTextHeight", mock.Anything).Return(10.0)
		provider.On("AddLine", mock.Anything, mock.Anything)

		// Act
		sut.Render(provider, &cell)

		// Assert
		provider.AssertNumberOfCalls(t, "AddText", 1)
		provider.AssertNumberOfCalls(t, "GetTextHeight", 1)
		provider.AssertNumberOfCalls(t, "AddLine", 1)
	})
}

func TestSignature_SetConfig(t *testing.T) {
	t.Run("should call correctly", func(t *testing.T) {
		// Arrange
		prop := fixture.SignatureProp()
		sut := signature.New("signature", prop)

		// Act
		sut.SetConfig(nil)
	})
}

func TestSignature_GetHeight(t *testing.T) {

	t.Run("When signature has a height of 10, should return 10", func(t *testing.T) {
		cell := fixture.CellEntity()
		font := fixture.FontProp()

		sut := signature.New("signature", props.Signature{SafePadding: 1, FontFamily: font.Family, FontStyle: font.Style, FontSize: font.Size, FontColor: font.Color, LineThickness: 2})

		provider := mocks.NewProvider(t)
		provider.EXPECT().GetTextHeight(&font).Return(5.0)

		// Act
		height := sut.GetHeight(provider, &cell)
		assert.Equal(t, 7.0, height)
	})
}
