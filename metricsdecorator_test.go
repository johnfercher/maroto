package maroto

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/stretchr/testify/assert"
)

func TestNewMetricsDecorator(t *testing.T) {
	// Act
	sut := NewMetricsDecorator(nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*maroto.metricsDecorator", fmt.Sprintf("%T", sut))
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
}
