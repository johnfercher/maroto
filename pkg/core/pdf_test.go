package core_test

import (
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

func TestNewPDF(t *testing.T) {
	// Act
	sut := core.NewPDF(nil, nil)

	// Assert
	assert.NotNil(t, sut)
	assert.Equal(t, "*core.Pdf", fmt.Sprintf("%T", sut))
}

func TestPdf_GetBase64(t *testing.T) {
	// Arrange
	sut := core.NewPDF([]byte{1, 2, 3}, nil)

	// Act
	b64 := sut.GetBase64()

	// Assert
	assert.Equal(t, "AQID", b64)
}

func TestPdf_GetBytes(t *testing.T) {
	// Arrange
	sut := core.NewPDF([]byte{1, 2, 3}, nil)

	// Act
	bytes := sut.GetBytes()

	// Assert
	assert.Equal(t, []byte{1, 2, 3}, bytes)
}

func TestPdf_GetReport(t *testing.T) {
	// Arrange
	sut := core.NewPDF(nil, &metrics.Report{SizeMetric: metrics.SizeMetric{
		Key: "key",
		Size: metrics.Size{
			Value: 10.0,
			Scale: metrics.Byte,
		},
	}})

	// Act
	report := sut.GetReport()

	// Assert
	assert.Equal(t, "key", report.SizeMetric.Key)
}

func TestPdf_Save(t *testing.T) {
	t.Run("when cannot save, should return error", func(t *testing.T) {
		// Arrange
		sut := core.NewPDF(nil, nil)

		// Act
		err := sut.Save("")

		// Assert
		assert.NotNil(t, err)
	})
	t.Run("when can save, should not return error", func(t *testing.T) {
		// Arrange
		bytes := []byte{1, 2, 3}
		file := buildPath("test.txt")
		sut := core.NewPDF(bytes, nil)

		// Act
		err := sut.Save(file)

		// Assert
		assert.Nil(t, err)
		savedBytes, _ := os.ReadFile(file)
		assert.Equal(t, bytes, savedBytes)
		_ = os.Remove(file)
	})
}

func buildPath(file string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}

	dir = strings.ReplaceAll(dir, "pkg/core/entity", "")
	return path.Join(dir, file)
}
