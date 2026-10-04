package merge_test

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/merge"
)

func TestBytes(t *testing.T) {
	t.Parallel()
	t.Run("when no PDFs are provided, should return merge error", func(t *testing.T) {
		t.Parallel()
		// Act
		result, err := merge.Bytes()

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
	t.Run("when a nil PDF slice is expanded, should return merge error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		var pdfs [][]byte

		// Act
		result, err := merge.Bytes(pdfs...)

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
	t.Run("when an empty PDF slice is expanded, should return merge error", func(t *testing.T) {
		t.Parallel()
		// Arrange
		pdfs := [][]byte{}

		// Act
		result, err := merge.Bytes(pdfs...)

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
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
	t.Run("when nil PDF bytes are provided, should return wrapped error", func(t *testing.T) {
		t.Parallel()
		// Act
		result, err := merge.Bytes(nil)

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
	t.Run("when empty PDF bytes are provided, should return wrapped error", func(t *testing.T) {
		t.Parallel()
		// Act
		result, err := merge.Bytes([]byte{})

		// Assert
		assert.Nil(t, result)
		assert.ErrorIs(t, err, merge.ErrCannotMergePDFs)
	})
}

func TestBytes_ConcurrentCalls(t *testing.T) {
	t.Parallel()
	t.Run("when valid PDFs are merged concurrently, should return every result", func(t *testing.T) {
		t.Parallel()
		// Arrange
		document := maroto.New()
		document.AddRows(text.NewRow(10, "concurrent"))
		generated, err := document.Generate()
		if !assert.NoError(t, err) {
			return
		}
		pdf := generated.GetBytes()
		const workers = 16
		start := make(chan struct{})
		results := make(chan []byte, workers)
		errCh := make(chan error, workers)
		var waitGroup sync.WaitGroup

		// Act
		waitGroup.Add(workers)
		for range workers {
			go func() {
				defer waitGroup.Done()
				<-start
				result, err := merge.Bytes(pdf, pdf, pdf)
				if err != nil {
					errCh <- err
					return
				}
				results <- result
			}()
		}
		close(start)
		waitGroup.Wait()
		close(results)
		close(errCh)

		// Assert
		for err := range errCh {
			assert.NoError(t, err)
		}
		assert.Len(t, results, workers)
		for result := range results {
			assert.NotEmpty(t, result)
		}
	})
}
