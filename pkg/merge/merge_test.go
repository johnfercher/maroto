package merge_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/merge"
)

func TestBytes(t *testing.T) {
	t.Parallel()
	t.Run("when valid PDFs are provided, should merge and return bytes", func(t *testing.T) {
		t.Parallel()
		// Arrange
		m1 := maroto.New()
		m1.AddRows(text.NewRow(10, "text1"))
		doc1, _ := m1.Generate()
		doc1Bytes := doc1.GetBytes()

		m2 := maroto.New()
		m2.AddRows(text.NewRow(10, "text2"))
		doc2, _ := m2.Generate()
		doc2Bytes := doc2.GetBytes()

		// Act
		result, err := merge.Bytes(doc1Bytes, doc2Bytes)

		// Assert
		assert.Nil(t, err)
		assert.InDelta(t, len(doc1Bytes)+len(doc2Bytes), len(result), 500)
	})
	t.Run("when invalid PDF bytes are provided, should return wrapped error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		invalidPDF := []byte("not a valid pdf")

		// Act
		result, err := merge.Bytes(invalidPDF, invalidPDF)

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
}
