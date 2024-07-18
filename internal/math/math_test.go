package math_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/stretchr/testify/assert"
)

func TestNewMath(t *testing.T) {
	sut := math.New()

	assert.NotNil(t, sut)
	assert.Equal(t, "*math.math", fmt.Sprintf("%T", sut))
}

func TestMath_GetCenterCorrection(t *testing.T) {
	t.Run("should get center correction correctly", func(t *testing.T) {
		// Arrange
		sut := math.New()
		outerSize := 100.0
		innerSize := 50.0

		// Act
		correction := sut.GetCenterCorrection(outerSize, innerSize)

		// Assertcenter
		assert.Equal(t, 25.0, correction)
	})
}

func TestMath_Resize(t *testing.T) {
	t.Run("When inner and outer have the same size and 100% is set, inner should be returned with 100 percent of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})

	t.Run("When inner and outer have the same size and 75% is set, inner should be returned with '75%' of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})

	t.Run("When inner is smaller and has the same proportion as outer and 100% is set, inner should be returned with '100%' of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})

	t.Run("When inner is smaller and has the same proportion as outer and 75% is set, inner should be returned with '75%' of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})

	t.Run("When inner is greater and has the same proportion as outer and 100% is set, inner should be returned with '100%' of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})

	t.Run("When inner is greater and has the same proportion as outer and 75% is set, inner should be returned with '75%' of outer", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 120, Height: 120}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})

	t.Run("When internal height is smaller and proportion is smaller with '100%' sent, it must return internal with '100%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})

	t.Run("When internal height is smaller and proportion is smaller with '75%' sent, it must return internal with '75%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 100, Height: 80}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})

	t.Run("When internal width is smaller and proportion is greater with '100%' sent, it must return internal with '100%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})

	t.Run("When internal width is smaller and proportion is greater with '75%' sent, it must return internal with '75%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 80, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})

	t.Run("When internal height is greater and proportion is greater with '100%' sent, it must return internal with '100%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 80.0, cell.Width)
		assert.Equal(t, 100.0, cell.Height)
	})

	t.Run("When internal height is greater and proportion is greater with '75%' sent, it must return internal with '75%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 100, Height: 125}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 60.0, cell.Width)
		assert.Equal(t, 75.0, cell.Height)
	})

	t.Run("When internal width is greater and proportion is smaller with '100%' sent, it must return internal with '100%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 100.0, cell.Width)
		assert.Equal(t, 80.0, cell.Height)
	})

	t.Run("When internal width is greater and proportion is smaller with '75%' sent, it must return internal with '75%' of the external maintaining its proportion", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		inner := &entity.Dimensions{Width: 125, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.Resize(inner, outer, percent, false)

		// Assert
		assert.Equal(t, 75.0, cell.Width)
		assert.Equal(t, 60.0, cell.Height)
	})

	// se justreferencewidith for true e altura proporcional da imagem extrapola celula, então a imagem é ajustada de acordo altura disponível
	// se jus for true e for possível ajustar imagem a altura disponível, percent se aplica apenas a largura

	t.Run("when justReferenceWidth is true and inner extrapolates external height, it should resize image based on available height", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 100.0
		justReferenceWidth := true
		inner := &entity.Dimensions{Width: 50, Height: 50}
		outer := &entity.Dimensions{Width: 100, Height: 50}

		// Act
		cell := sut.Resize(inner, outer, percent, justReferenceWidth)

		// Assert
		assert.Equal(t, 50.0, cell.Width)
		assert.Equal(t, 50.0, cell.Height)
	})
	t.Run("when justReferenceWidth is true and inner does not extrapolate external height, it should resize image only based on available width", func(t *testing.T) {
		// Arrange
		sut := math.New()

		percent := 75.0
		justReferenceWidth := true
		inner := &entity.Dimensions{Width: 100, Height: 55}
		outer := &entity.Dimensions{Width: 100, Height: 40.99999999999999}

		// Act
		cell := sut.Resize(inner, outer, percent, justReferenceWidth)

		// Assert
		assert.Equal(t, 74.54545454545453, cell.Width)
		assert.Equal(t, 40.99999999999999, cell.Height)
	})
}

func TestMath_GetInnerCenterCell(t *testing.T) {
	t.Run("there is not side-effect", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		_ = sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 100.0, inner.Width)
		assert.Equal(t, 100.0, inner.Height)
		assert.Equal(t, 100.0, outer.Width)
		assert.Equal(t, 100.0, outer.Height)
	})

	t.Run("when inner and outer have the same size, should return the center", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
	})

	t.Run("when inner is smaller than outer and has equal proportion, the center of the cell must be returned", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 80, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
	})

	t.Run("when the internal one has a smaller height and smaller proportion than the external one, the center of the cell must be returned", func(t *testing.T) {
		// Arrange
		sut := math.New()

		outer := &entity.Dimensions{Width: 100, Height: 100}
		inner := &entity.Dimensions{Width: 75.0, Height: 60.0}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 12.5, cell.X)
		assert.Equal(t, 20.0, cell.Y)
	})

	t.Run("when internal has a smaller width and greater proportion than external, the center of the cell must be returned", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 80.0, Height: 100}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 10.0, cell.X)
		assert.Equal(t, 0.0, cell.Y)
	})

	t.Run("when internal has greater height and proportion than external, the center of the cell must be returned", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 60.0, Height: 75.0}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 20.0, cell.X)
		assert.Equal(t, 12.5, cell.Y)
	})

	t.Run("quando interno tiver largura maior e proporção menor que externa, the center of the cell must be returned", func(t *testing.T) {
		// Arrange
		sut := math.New()

		inner := &entity.Dimensions{Width: 100, Height: 80}
		outer := &entity.Dimensions{Width: 100, Height: 100}

		// Act
		cell := sut.GetInnerCenterCell(inner, outer)

		// Assert
		assert.Equal(t, 0.0, cell.X)
		assert.Equal(t, 10.0, cell.Y)
	})
}
