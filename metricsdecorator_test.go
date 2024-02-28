package maroto

import (
	"fmt"
	"testing"

	"github.com/johnfercher/go-tree/node"
	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/stretchr/testify/assert"
)

func TestNewMetricsDecorator(t *testing.T) {
	// Act
	sut := NewMetricsDecorator(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*maroto.MetricsDecorator", fmt.Sprintf("%T", sut))
}

func TestMetricsDecorator_AddPages(t *testing.T) {
	// Arrange
	pg := page.New()

	docToReturn := &mocks.Document{}
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := &mocks.Maroto{}
	inner.EXPECT().AddPages(pg)
	inner.EXPECT().Generate().Return(docToReturn, nil)

	sut := NewMetricsDecorator(inner)

	// Act
	sut.AddPages(pg)
	sut.AddPages(pg)

	// Assert
	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)

	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_page", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddPages", 2)
}

func TestMetricsDecorator_AddRow(t *testing.T) {
	// Arrange
	col := col.New(12)

	docToReturn := &mocks.Document{}
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := &mocks.Maroto{}
	inner.EXPECT().AddRow(10.0, col).Return(nil)
	inner.EXPECT().Generate().Return(docToReturn, nil)

	sut := NewMetricsDecorator(inner)

	// Act
	sut.AddRow(10, col)
	sut.AddRow(10, col)

	// Assert
	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)

	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_row", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRow", 2)
}

func TestMetricsDecorator_AddRows(t *testing.T) {
	// Arrange
	row := row.New(10).Add(col.New(12))

	docToReturn := &mocks.Document{}
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := &mocks.Maroto{}
	inner.EXPECT().AddRows(row)
	inner.EXPECT().Generate().Return(docToReturn, nil)

	sut := NewMetricsDecorator(inner)

	// Act
	sut.AddRows(row)
	sut.AddRows(row)

	// Assert
	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)

	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 2, len(report.TimeMetrics))
	assert.Equal(t, "generate", report.TimeMetrics[0].Key)
	assert.Equal(t, "add_rows", report.TimeMetrics[1].Key)
	assert.Equal(t, 2, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRows", 2)
}

func TestMetricsDecorator_GetStructure(t *testing.T) {
	// Arrange
	row := row.New(10).Add(col.New(12))

	docToReturn := &mocks.Document{}
	docToReturn.EXPECT().GetBytes().Return([]byte{1, 2, 3})
	inner := &mocks.Maroto{}
	inner.EXPECT().AddRows(row)
	inner.EXPECT().GetStructure().Return(&node.Node[core.Structure]{})
	inner.EXPECT().Generate().Return(docToReturn, nil)

	sut := NewMetricsDecorator(inner)
	sut.AddRows(row)

	// Act
	_ = sut.GetStructure()

	// Assert
	doc, err := sut.Generate()
	assert.Nil(t, err)
	assert.NotNil(t, doc)

	report := doc.GetReport()
	assert.NotNil(t, report)
	assert.Equal(t, 3, len(report.TimeMetrics))
	assert.Equal(t, "get_tree_structure", report.TimeMetrics[0].Key)
	assert.Equal(t, "generate", report.TimeMetrics[1].Key)
	assert.Equal(t, "add_rows", report.TimeMetrics[2].Key)
	assert.Equal(t, 1, len(report.TimeMetrics[1].Times))
	inner.AssertNumberOfCalls(t, "AddRows", 1)
	inner.AssertNumberOfCalls(t, "GetStructure", 1)
}

func TestMetricsDecorator_GetDimensions(t *testing.T) {
	// Arrange
	inner := &mocks.Maroto{}
	inner.EXPECT().GetDimensions().Return(entity.Dimensions{Width: 100, Height: 150})

	sut := NewMetricsDecorator(inner)

	// Assert
	dimensions := sut.GetDimensions()
	assert.Equal(t, float64(100), dimensions.Width)
	assert.Equal(t, float64(150), dimensions.Height)
	inner.AssertNumberOfCalls(t, "GetDimensions", 1)
}

func TestMetricsDecorator_GetCurrentHeight(t *testing.T) {
	// Arrange
	inner := &mocks.Maroto{}
	inner.EXPECT().GetCurrentHeight().Return(120)

	sut := NewMetricsDecorator(inner)

	// Assert
	currentHeight := sut.GetCurrentHeight()
	assert.Equal(t, float64(120), currentHeight)
	inner.AssertNumberOfCalls(t, "GetCurrentHeight", 1)
}
