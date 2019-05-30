package maroto_test

import (
	"fmt"
	"github.com/johnfercher/maroto"
	"github.com/johnfercher/maroto/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewImage(t *testing.T) {
	image := maroto.NewImage(&mocks.Pdf{}, &mocks.Math{})

	assert.NotNil(t, image)
	assert.Equal(t, fmt.Sprintf("%T", image), "*maroto.image")
}

func TestImage_AddFromPath(t *testing.T) {
	// ARRANGE
	_pdf := &mocks.Pdf{}
	_pdf.On("GetMargins").Return(10.0, 10.0, 10.0, 10.0)
	_pdf.On("ImageOptions", mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything, mock.Anything)

	_math := &mocks.Math{}
	_math.On("GetWidthPerCol", mock.Anything).Return(100.0)

	image := maroto.NewImage(_pdf, _math)

	// Act
	image.AddFromPath("AnyPath", 10.0, 1.0, 4.0, 5.0)

	// Assert
	_math.AssertNumberOfCalls(t, "GetWidthPerCol", 1)
	_math.AssertCalled(t, "GetWidthPerCol", 4.0)

	_pdf.AssertNumberOfCalls(t, "GetMargins", 1)

	_pdf.AssertNumberOfCalls(t, "ImageOptions", 1)
	//_pdf.AssertCalled(t, "ImageOptions", "AnyPath", 110.0, 20.0, 100.0, 0, false, gofpdf.ImageOptions{}, 0, "")
}
