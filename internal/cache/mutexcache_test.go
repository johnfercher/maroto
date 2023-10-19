package cache_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/internal/cache"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/stretchr/testify/assert"
)

func TestNewMutexDecorator(t *testing.T) {
	// Act
	sut := cache.NewMutexDecorator(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*cache.mutexCache", fmt.Sprintf("%T", sut))
}

func TestMutexCache_AddImage(t *testing.T) {
	// Arrange
	value := "value1"
	img := &entity.Image{}

	innerMock := &mocks.Cache{}
	innerMock.EXPECT().AddImage(value, img)

	sut := cache.NewMutexDecorator(innerMock)

	// Act
	sut.AddImage(value, img)

	// Assert
	innerMock.AssertNumberOfCalls(t, "AddImage", 1)
}

func TestMutexCache_GetImage(t *testing.T) {
	// Arrange
	value := "value2"
	ext := extension.Jpg
	imgToReturn := &entity.Image{}
	errToReturn := errors.New("any error")

	innerMock := &mocks.Cache{}
	innerMock.EXPECT().GetImage(value, ext).Return(imgToReturn, errToReturn)

	sut := cache.NewMutexDecorator(innerMock)

	// Act
	img, err := sut.GetImage(value, ext)

	// Assert
	assert.Equal(t, imgToReturn, img)
	assert.Equal(t, errToReturn, err)
	innerMock.AssertNumberOfCalls(t, "GetImage", 1)
}

func TestMutexCache_LoadImage(t *testing.T) {
	// Arrange
	value := "value3"
	ext := extension.Jpg
	errToReturn := errors.New("any error")

	innerMock := &mocks.Cache{}
	innerMock.EXPECT().LoadImage(value, ext).Return(errToReturn)

	sut := cache.NewMutexDecorator(innerMock)

	// Act
	err := sut.LoadImage(value, ext)

	// Assert
	assert.Equal(t, errToReturn, err)
	innerMock.AssertNumberOfCalls(t, "LoadImage", 1)
}
