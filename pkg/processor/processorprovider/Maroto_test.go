package processorprovider_test

import (
	"fmt"
	"testing"

	"github.com/johnfercher/maroto/v2/mocks"
	"github.com/johnfercher/maroto/v2/pkg/components/code"
	"github.com/johnfercher/maroto/v2/pkg/components/page"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/processorprovider"
	"github.com/johnfercher/maroto/v2/pkg/processor/repository"
	"github.com/johnfercher/maroto/v2/pkg/test"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestAddPages(t *testing.T) {
	t.Run("when page is not sent, should not add header", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddPages()

		assert.Nil(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("when invalid component is sent, should return an error", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddPages(text.New("123"))

		assert.NotNil(t, err)
		assert.Nil(t, provider)
	})

	t.Run("when page is sent, should add page", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddPages(page.New().Add(row.New(10)))

		assert.NotNil(t, provider)
		assert.Nil(t, err)
		test.New(t).Assert(provider.GetStructure()).Equals("processor/provider/document_with_page.json")
	})
}

// nolint: dupl
func TestAddFooter(t *testing.T) {
	t.Run("when footer is not sent, should not add footer", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddFooter()

		assert.Nil(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("when invalid component is sent, should return an error", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddFooter(page.New())

		assert.NotNil(t, err)
		assert.Nil(t, provider)
	})

	t.Run("when footer is sent, should add footer", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddFooter(text.NewAutoRow("footer"))

		assert.NotNil(t, provider)
		assert.Nil(t, err)
		test.New(t).Assert(provider.GetStructure()).Equals("processor/provider/document_with_footer.json")
	})
}

// nolint: dupl
func TestAddHeader(t *testing.T) {
	t.Run("when header is not sent, should not add header", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddHeader()

		assert.Nil(t, err)
		assert.NotNil(t, provider)
	})

	t.Run("when invalid component is sent, should return an error", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddHeader(page.New())

		assert.NotNil(t, err)
		assert.Nil(t, provider)
	})

	t.Run("when header is sent, should add header", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)
		provider, _ := processorprovider.NewMaroto(repository)

		provider, err := provider.AddHeader(text.NewAutoRow("header"))

		assert.NotNil(t, provider)
		assert.Nil(t, err)
		test.New(t).Assert(provider.GetStructure()).Equals("processor/provider/document_with_header.json")
	})
}

func TestNewMaroto(t *testing.T) {
	t.Run("when builder is not sent, should not create Maroto builder", func(t *testing.T) {
		repository := mocks.NewProcessorRepository(t)

		provider, err := processorprovider.NewMaroto(repository)

		assert.NotNil(t, provider)
		assert.Nil(t, err)
	})
	t.Run("when is not possible create Maroto builder, should return an error", func(t *testing.T) {
		mapperFix := buildermapper.Builder{BackgroundImage: "file"}
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument(mock.Anything).Return("", nil, fmt.Errorf("any"))

		provider, err := processorprovider.NewMaroto(repository, mapperFix)

		assert.NotNil(t, err)
		assert.Nil(t, provider)
	})
	t.Run("when builder is sent, should create Maroto builder", func(t *testing.T) {
		mapperFix := buildermapper.Builder{BackgroundImage: "file"}
		repository := mocks.NewProcessorRepository(t)
		repository.EXPECT().GetDocument(mock.Anything).Return("", []byte{123}, nil)

		provider, err := processorprovider.NewMaroto(repository, mapperFix)

		assert.NotNil(t, provider)
		assert.Nil(t, err)
	})
}

func TestCreateBarCode(t *testing.T) {
	t.Run("when CreateBarCode is called, should generate a barcode", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateBarCode("code",
			&propsmapper.Barcode{
				Left: 10.0, Top: 10.0, Percent: 100.0,
				Proportion: propsmapper.Proportion{Width: 10, Height: 2},
				Center:     false, Type: propsmapper.NewCodeType("code128"),
			},
		)

		assert.IsType(t, barcode, &code.Barcode{})
		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/barcode.json")
	})
}

func TestCreateMatrixCode(t *testing.T) {
	t.Run("when CreateMatrixCode is called, should generate a matrixcode", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateMatrixCode("code",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/matrixcode.json")
	})
}

func TestCreateQRCode(t *testing.T) {
	t.Run("when CreateQrCode is called, should generate a qrcode", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateQrCode("code",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/qrcode.json")
	})
}

func TestCreateImage(t *testing.T) {
	t.Run("when CreateImage is called, should generate a image", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		image := m.CreateImage(make([]byte, 0), "png",
			&propsmapper.Rect{Left: 10.0, Top: 10.0, Percent: 100.0, JustReferenceWidth: false, Center: false},
		)

		test.New(t).Assert(image.GetStructure()).Equals("processor/provider/image.json")
	})
}

func TestCreateLine(t *testing.T) {
	t.Run("when CreateLine is called, should generate a line", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateLine(
			&propsmapper.Line{
				Color: &propsmapper.Color{Red: 10, Green: 10, Blue: 10}, Style: "solid", Thickness: 10.0,
				Orientation: "vertical", OffsetPercent: 50, SizePercent: 50,
			},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/line.json")
	})
}

func TestCreateSignature(t *testing.T) {
	t.Run("when CreateSignature is called, should generate a signature", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateSignature("signature",
			&propsmapper.Signature{
				FontFamily: "Arial", FontStyle: "bold", FontSize: 10.0, FontColor: &propsmapper.Color{Red: 10, Green: 10, Blue: 10},
				LineColor: &propsmapper.Color{Red: 10, Green: 10, Blue: 10}, LineStyle: "solid", LineThickness: 10.0, SafePadding: 10.0,
			},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/signature.json")
	})
}

func TestCreateText(t *testing.T) {
	t.Run("when CreateText is called, should generate a text", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		barcode := m.CreateText("text",
			&propsmapper.Text{
				Top: 10.0, Left: 10.0, Right: 10.0, Family: "Arial", Style: "bold", Size: 10.0, Align: "center", BreakLineStrategy: "dash_strategy",
				VerticalPadding: 10.0, Color: &propsmapper.Color{Red: 10, Green: 10, Blue: 10}, Hyperlink: "test",
			},
		)

		test.New(t).Assert(barcode.GetStructure()).Equals("processor/provider/text.json")
	})
}

func TestCreateCol(t *testing.T) {
	t.Run("when CreateCol is called, should generate a col", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		text := text.New("test")

		col, err := m.CreateCol(10, text)

		assert.Nil(t, err)
		assert.NotNil(t, col)
	})

	t.Run("when invalid components are sent, should return an error", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		page, err := m.CreatePage(text.NewCol(10, "10"))

		assert.Nil(t, page)
		assert.NotNil(t, err)
	})
}

func TestCreateRow(t *testing.T) {
	t.Run("when CreateRow is called, should generate a row", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		text := text.NewCol(12, "test")

		col, err := m.CreateRow(10, text)

		assert.Nil(t, err)
		assert.NotNil(t, col)
	})

	t.Run("when invalid components are sent, should return an error", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		page, err := m.CreatePage(text.New("10"))

		assert.Nil(t, page)
		assert.NotNil(t, err)
	})
}

func TestCreatePage(t *testing.T) {
	t.Run("when CreatePage is called, should generate a page", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		page, err := m.CreatePage(row.New(10))

		assert.Nil(t, err)
		assert.NotNil(t, page)
	})
	t.Run("when invalid components are sent, should return an error", func(t *testing.T) {
		m, _ := processorprovider.NewMaroto(repository.NewMemoryStorage())
		page, err := m.CreatePage(text.New("10"))

		assert.Nil(t, page)
		assert.NotNil(t, err)
	})
}
