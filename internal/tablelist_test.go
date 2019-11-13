package internal_test

import (
	"fmt"
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewTableList(t *testing.T) {
	// Act
	tableList := internal.NewTableList(nil, nil)

	// Assert
	assert.NotNil(t, tableList)
	assert.Equal(t, fmt.Sprintf("%T", tableList), "*internal.tableList")
}

func TestTableList_Create_WhenHeaderIsNil(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	_, contents := getContents()

	// Act
	sut.Create(nil, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenHeaderIsEmpty(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	_, contents := getContents()

	// Act
	sut.Create([]string{}, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenContentIsNil(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, nil)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_WhenContentIsEmpty(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, [][]string{})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
}

func TestTableList_Create_Happy(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	marotoGrid := &mocks.Maroto{}
	marotoGrid.On("Row", mock.Anything, mock.Anything).Return(nil)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(marotoGrid)

	headers, contents := getContents()

	// Act
	sut.Create(headers, contents)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	marotoGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	marotoGrid.AssertNumberOfCalls(t, "Row", 22)
}

func getContents() ([]string, [][]string) {
	header := []string{"j = 0", "j = 1", "j = 2", "j = 4"}

	contents := [][]string{}
	for i := 0; i < 20; i++ {
		content := []string{}
		for j := 0; j < 4; j++ {
			content = append(content, fmt.Sprintf("i = %d, j = %d", i, j))
		}
		contents = append(contents, content)
	}

	return header, contents
}
