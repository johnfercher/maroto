package maroto_test

import (
	"fmt"
	"runtime"
	"testing"
	"time"

	"github.com/johnfercher/maroto/v2/pkg/props"

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
	t.Run("when config is sent, it should create Maroto object", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			Build()

		// Act
		sut := maroto.New(cfg)

		// Assert
		assert.NotNil(t, sut)
		assert.Equal(t, "*maroto.Maroto", fmt.Sprintf("%T", sut))
	})
	t.Run("when config with an concurrent mode is sent, should create Maroto object", func(t *testing.T) {
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
	t.Run("when config with an low memory mode is sent, should create Maroto object", func(t *testing.T) {
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
	t.Run("When row height and available sapacing are equals, should add row in current page", func(t *testing.T) {
		cfg := config.NewBuilder().
			WithDimensions(20, 20).
			WithBottomMargin(0).
			WithTopMargin(0).
			WithLeftMargin(0).
			WithRightMargin(0).
			Build()
		sut := maroto.New(cfg)
		// Act
		sut.AddRow(19)
		sut.AddRow(1)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row.json")
	})
	t.Run("when col is not sent, should empty col is set", func(t *testing.T) {
		// Arrange
		sut := maroto.New()
		// Act
		sut.AddRow(10)

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_4.json")
	})
	t.Run("when one row is sent, should create one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_1.json")
	})
	t.Run("when two rows are sent, should create two rows", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))
		sut.AddRow(10, col.New(12))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_row_2.json")
	})
	t.Run("when rows do not fit on the current page, should create a new page", func(t *testing.T) {
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
	t.Run("when one row is sent, should create one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_1.json")
	})
	t.Run("when two rows are sent, should create two rows", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRows(row.New(15).Add(col.New(12)))
		sut.AddRows(row.New(15).Add(col.New(12)))

		// Assert
		test.New(t).Assert(sut.GetStructure()).Equals("maroto_add_rows_2.json")
	})
	t.Run("when rows do not fit on the current page, should create a new page", func(t *testing.T) {
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
	t.Run("when a new page is created, should add a page", func(t *testing.T) {
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
	t.Run("when two pages are created, should add two pages", func(t *testing.T) {
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
	t.Run("when the sent page uses two pages, two pages are created", func(t *testing.T) {
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
	t.Run("when one row is sent, should generate one row", func(t *testing.T) {
		// Arrange
		sut := maroto.New()

		// Act
		sut.AddRow(10, col.New(12))

		// Assert
		doc, err := sut.Generate()
		assert.Nil(t, err)
		assert.NotNil(t, doc)
	})
	t.Run("when two row are sent, should generate two row", func(t *testing.T) {
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
	t.Run("when rows do not fit on the current page, should generate two pages", func(t *testing.T) {
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
	t.Run("when rows do not fit on the current page and concurrent mode is active, should executed in parallel", func(t *testing.T) {
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
	t.Run("when two pages are sent and low memory mode is active, should executed in low memory mode", func(t *testing.T) {
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
	t.Run("when two pages are sent and sequential generation is active, should executed in sequential generation mode", func(t *testing.T) {
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
	t.Run("when two pages are sent and sequential low memory is active, should executed in sequential low memory mode", func(t *testing.T) {
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
	t.Run("when two pages are sent and concurrent mode is active, should executed in parallel", func(t *testing.T) {
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
	t.Run("goroutines do not leak after multiple generate calls on concurrent mode", func(t *testing.T) {
		// Arrange
		cfg := config.NewBuilder().
			WithConcurrentMode(10).
			Build()

		sut := maroto.New(cfg)

		// Act
		for i := 0; i < 30; i++ {
			sut.AddRow(10, col.New(12))
		}
		initialGoroutines := runtime.NumGoroutine()
		_, err1 := sut.Generate()
		_, err2 := sut.Generate()
		_, err3 := sut.Generate()
		time.Sleep(100 * time.Millisecond)
		finalGoroutines := runtime.NumGoroutine()

		// Assert
		assert.Nil(t, err1)
		assert.Nil(t, err2)
		assert.Nil(t, err3)
		assert.Equal(t, initialGoroutines, finalGoroutines)
	})
	t.Run("when two pages are sent and page number is active, should add page number", func(t *testing.T) {
		pageNumber := props.PageNumber{
			MarginTop: 1,
		}
		// Arrange
		cfg := config.NewBuilder().
			WithPageNumber(pageNumber).
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

func TestMaroto_FitlnCurrentPage(t *testing.T) {
	t.Run("when component is smaller should available size, should return false", func(t *testing.T) {
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
	t.Run("when component is larger should the available size, should return true", func(t *testing.T) {
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
	t.Run("when it have content with an automatic height of 10 and the height sent fits the current page, it should return true",
		func(t *testing.T) {
			sut := maroto.New(config.NewBuilder().
				WithDimensions(210.0, 297.0).
				Build())

			var rows []core.Row
			for i := 0; i < 10; i++ {
				rows = append(rows, row.New().Add(text.NewCol(12, "teste")))
			}

			sut.AddPages(page.New().Add(rows...))
			assert.True(t, sut.FitlnCurrentPage(40))
		})
}

func TestMaroto_GetCurrentConfig(t *testing.T) {
	t.Run("When GetCurrentConfig is called, should return the current settings", func(t *testing.T) {
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
