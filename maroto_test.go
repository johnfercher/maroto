package maroto_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/col"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/test"

	"github.com/johnfercher/maroto/v2"

	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	t.Run("new default", func(t *testing.T) {
		// Act
		sut := maroto.New()

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*maroto.Maroto", fmt.Sprintf("%T", sut))
	})
	t.Run("new with config", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			Build()

		// Act
		sut := maroto.New(cfg)

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*maroto.Maroto", fmt.Sprintf("%T", sut))
	})
	t.Run("new with config an concurrent mode on", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithConcurrentMode(7).
			Build()

		// Act
		sut := maroto.New(cfg)

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*maroto.Maroto", fmt.Sprintf("%T", sut))
	})
	t.Run("new with config an low memory mode on", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()

		// Act
		sut := maroto.New(cfg)

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*maroto.Maroto", fmt.Sprintf("%T", sut))
	})
}

func TestMaroto_AddRow(t *testing.T) {
	t.Run("when col is not sent, should empty col is set", func(t *testing.T) {
		// Arrange
		sut := maroto.New()
		// Act
		sut.AddRow(10)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_4.json")
	})
	t.Run("add one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_1.json")
	})
	t.Run("add one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_1.json")
	})
	t.Run("add two rows", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_2.json")
	})
	t.Run("add rows until add new page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_3.json")
	})
}

func TestMaroto_AddRows(t *testing.T) {
	t.Run("when col is not sent, should empty col is set", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_4.json")
	})
	t.Run("add one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_1.json")
	})
	t.Run("add two rows", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_2.json")
	})
	t.Run("add rows until add new page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 20; i++ {
			sut.AddRows(row.New(15).Add(col.New(12)))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_3.json")
	})
}

func TestMaroto_AddPages(t *testing.T) {
	t.Run("add one page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddPages(
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
		)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_pages_1.json")
	})
	t.Run("add two pages", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddPages(
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
			page.New().Add(
				row.New(20).Add(col.New(12)),
			),
		)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_pages_2.json")
	})
	t.Run("add page greater than one page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()
		var rows []core.Row
		for i := 0; i < 15; i++ {
			rows = append(rows, row.New(20).Add(col.New(12)))
		}

		// Act
		sut.AddPages(page.New().Add(rows...))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_pages_3.json")
	})
}

func TestMaroto_Generate(t *testing.T) {
	t.Run("add one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
	t.Run("add two rows", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
	t.Run("add rows until add new page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
	t.Run("add rows until add new page, execute in parallel", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithConcurrentMode(7).
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
	t.Run("add rows until add new page, execute in low memory mode", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithSequentialLowMemoryMode(10).
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
}

func TestMaroto_FitlnCurrentPage(t *testing.T) {
	t.Run("when component is smaller should available size, then false", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithDimensions(210.0, 297.0).
			Build())

		var rows []core.Row
		for i := 0; i < 26; i++ {
			rows = append(rows, row.New(10).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))
		assert.False(t, sut.FitlnCurrentPage(40))
	})
	t.Run("when component is larger should the available size, then true", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithDimensions(210.0, 297.0).
			Build())

		var rows []core.Row
		for i := 0; i < 10; i++ {
			rows = append(rows, row.New(10).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))
		assert.True(t, sut.FitlnCurrentPage(40))
	})
}
