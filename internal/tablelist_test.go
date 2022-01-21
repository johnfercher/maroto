package internal_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/internal/mocks"
	"github.com/johnfercher/maroto/pkg/color"
	"github.com/johnfercher/maroto/pkg/consts"
	"github.com/johnfercher/maroto/pkg/props"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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
	sut.Create(nil, contents, consts.Arial)

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
	sut.Create([]string{}, contents, consts.Arial)

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
	sut.Create(headers, nil, consts.Arial)

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
	sut.Create(headers, [][]string{}, consts.Arial)

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
	marotoGrid.On("Line", mock.Anything).Return(nil)
	marotoGrid.On("SetBackgroundColor", mock.Anything).Return(nil)
	marotoGrid.On("GetPageMargins").Return(10.0, 10.0, 10.0, 10.0)
	marotoGrid.On("GetPageSize").Return(200.0, 600.0)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(marotoGrid)

	headers, contents := getContents()

	// Act
	sut.Create(headers, contents, consts.Arial, props.TableList{
		Line: true,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	marotoGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	marotoGrid.AssertNumberOfCalls(t, "Row", 22)
	marotoGrid.AssertNumberOfCalls(t, "Line", 20)
	marotoGrid.AssertNotCalled(t, "SetBackgroundColor")
}

func TestTableList_Create_HappyWithBackgroundColor(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	marotoGrid := &mocks.Maroto{}
	marotoGrid.On("Row", mock.Anything, mock.Anything).Return(nil)
	marotoGrid.On("Line", mock.Anything).Return(nil)
	marotoGrid.On("SetBackgroundColor", mock.Anything).Return(nil)
	marotoGrid.On("GetPageMargins").Return(10.0, 10.0, 10.0, 10.0)
	marotoGrid.On("GetPageSize").Return(200.0, 600.0)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(marotoGrid)

	headers, contents := getContents()
	color := color.Color{
		Red:   200,
		Green: 200,
		Blue:  200,
	}

	// Act
	sut.Create(headers, contents, consts.Arial, props.TableList{
		AlternatedBackground: &color,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	marotoGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	marotoGrid.AssertNumberOfCalls(t, "Row", 22)

	marotoGrid.AssertNotCalled(t, "Line")

	marotoGrid.AssertCalled(t, "SetBackgroundColor", color)
	marotoGrid.AssertNumberOfCalls(t, "SetBackgroundColor", 20)
}

func TestTableList_Create_Happy_Without_Line(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)

	font := &mocks.Font{}
	font.On("GetFont").Return(consts.Arial, consts.Bold, 1.0)
	font.On("GetScaleFactor").Return(1.5)

	marotoGrid := &mocks.Maroto{}
	marotoGrid.On("Row", mock.Anything, mock.Anything).Return(nil)
	marotoGrid.On("Line", mock.Anything).Return(nil)
	marotoGrid.On("GetPageMargins").Return(10.0, 10.0, 10.0, 10.0)
	marotoGrid.On("GetPageSize").Return(200.0, 600.0)

	sut := internal.NewTableList(text, font)
	sut.BindGrid(marotoGrid)

	headers, contents := getContents()

	// Act
	sut.Create(headers, contents, consts.Arial)

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	text.AssertNumberOfCalls(t, "GetLinesQuantity", 84)

	font.AssertCalled(t, "GetFont")
	font.AssertNumberOfCalls(t, "GetFont", 21)

	marotoGrid.AssertCalled(t, "Row", mock.Anything, mock.Anything)
	marotoGrid.AssertNumberOfCalls(t, "Row", 22)
	marotoGrid.AssertNumberOfCalls(t, "Line", 0)
}

func TestTableList_Create_WhenContentIsEmptyWithLine(t *testing.T) {
	// Arrange
	text := &mocks.Text{}
	text.On("GetLinesQuantity", mock.Anything, mock.Anything, mock.Anything).Return(1)
	sut := internal.NewTableList(text, nil)

	marotoGrid := mocks.Maroto{}
	marotoGrid.On("Line", mock.Anything).Return(nil)

	headers, _ := getContents()

	// Act
	sut.Create(headers, [][]string{}, consts.Arial, props.TableList{
		Line: true,
	})

	// Assert
	text.AssertNotCalled(t, "GetLinesQuantity")
	marotoGrid.AssertNotCalled(t, "Line")
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
