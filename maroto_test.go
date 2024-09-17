package maroto_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/text"

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
	t.Run("when one row is sent,it should set one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_1.json")
	})
	t.Run("when two row is sent,it should set two row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_2.json")
	})
	t.Run("when rows exceed the page, should add a new page", func(t *testing.T) {
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

func TestMaroto_FillPageToAddNew(t *testing.T) {
	t.Run("when FillPageToAddNew is called, it should add components to the new page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15))
		sut.FillPageToAddNew()
		sut.AddRows(row.New(15))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_fill_page_1.json")
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
	t.Run("when one row is sent,it should set one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_1.json")
	})
	t.Run("when two row is sent,it should set two row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_2.json")
	})
	t.Run("when rows exceed the page, should add a new page", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 20; i++ {
			sut.AddRows(row.New(15).Add(col.New(12)))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_3.json")
	})

	t.Run("when autoRow is sent, should set autoRow", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 20; i++ {
			sut.AddRows(row.New().Add(text.NewCol(12, "teste")))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_5.json")
	})
}

func TestMaroto_AddAutoRow(t *testing.T) {
	t.Run("When 100 automatic rows are sent, it should create 2 pages", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		for i := 0; i < 150; i++ {
			sut.AddAutoRow(text.NewCol(12, "teste"))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_auto_row_1.json")
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
	t.Run("sequential generation", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithSequentialMode().
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_sequential.json")
	})
	t.Run("sequential low memory generation", func(t *testing.T) {
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
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_sequential_low_memory.json")
	})
	t.Run("sequential low memory generation", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithConcurrentMode(10).
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_concurrent.json")
	})
	t.Run("page number", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithPageNumber().
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_page_number.json")
	})
}

func TestMaroto_FitsOnCurrentPage(t *testing.T) {
	t.Run("when no row fits on the current page, should return 0", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithDimensions(210.0, 297.0).
			Build())

		var rows []core.Row
		for i := 0; i < 26; i++ {
			rows = append(rows, row.New(10).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))
		assert.Equal(t, 0, sut.FitsOnCurrentPage(rows...))
	})
	t.Run("when all rows fit on the current page, should return the total number of rows", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithDimensions(210.0, 297.0).
			Build())

		var rows []core.Row
		for i := 0; i < 10; i++ {
			rows = append(rows, row.New(10).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))
		assert.Equal(t, 10, sut.FitsOnCurrentPage(rows...))
	})
	t.Run("when half of the rows fit on the current page, should return the amount that fits", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithDimensions(210.0, 297.0).
			Build())

		var rows []core.Row
		for i := 0; i < 20; i++ {
			rows = append(rows, row.New(10).Add(col.New(12)))
		}

		sut.AddPages(page.New().Add(rows...))
		assert.Equal(t, 6, sut.FitsOnCurrentPage(rows...))
	})
	t.Run("when it have content with an automatic height of 10 and the height sent fits the current page, it should return 1",
		func(t *testing.T) {
			sut := maroto.New(config.NewBuilder().
				WithDimensions(210.0, 297.0).
				Build())

			var rows []core.Row
			for i := 0; i < 10; i++ {
				rows = append(rows, row.New().Add(text.NewCol(12, "test")))
			}

			sut.AddPages(page.New().Add(rows...))
			assert.Equal(t, 1, sut.FitsOnCurrentPage(text.NewRow(40, "test")))
		})
}

func TestMaroto_GetCurrentConfig(t *testing.T) {
	t.Run("When GetCurrentConfig is called then current settings are returned", func(t *testing.T) {
		sut := maroto.New(config.NewBuilder().
			WithMaxGridSize(20).
			Build())

		assert.Equal(t, sut.GetCurrentConfig().MaxGridSize, 20)
	})
}

// nolint:dupl // dupl is good here
func TestMaroto_RegisterHeader(t *testing.T) {
	t.Run("when header size is greater than useful area, should return error", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterHeader(row.New(1000))

		assert.NotNil(t, err)
		assert.Equal(t, "header height is greater than page useful area", err.Error())
	})
	t.Run("when header size is correct, should not return error and apply header", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterHeader(code.NewBarRow(10, "header"))

		var rows []core.Row
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("header.json")
	})
	t.Run("when autoRow is sent, should set autoRow", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterHeader(text.NewAutoRow("header"))

		var rows []core.Row
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("header_auto_row.json")
	})
}

// nolint:dupl // dupl is good here
func TestMaroto_RegisterFooter(t *testing.T) {
	t.Run("when footer size is greater than useful area, should return error", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterFooter(row.New(1000))

		assert.NotNil(t, err)
		assert.Equal(t, "footer height is greater than page useful area", err.Error())
	})
	t.Run("when header size is correct, should not return error and apply header", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterFooter(code.NewBarRow(10, "footer"))

		var rows []core.Row
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("footer.json")
	})
	t.Run("when autoRow is sent, should set autoRow", func(t *testing.T) {
		sut := maroto.New()

		err := sut.RegisterFooter(text.NewAutoRow("header"))

		var rows []core.Row
		for i := 0; i < 5; i++ {
			rows = append(rows, row.New(100).Add(col.New(12)))
		}

		sut.AddRows(rows...)

		// Assert
		assert.Nil(t, err)
		test.New(t).Assert(sut.GetStructure()).Equals("footer_auto_row.json")
	})
}
