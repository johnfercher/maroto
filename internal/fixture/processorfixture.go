package fixture

import (
	"time"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/buildermapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/colmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/imagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/linemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/rowmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/signaturemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/textmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"
)

func MapperRow() *rowmapper.Row {
	return &rowmapper.Row{
		Height: 0,
		Cols:   make([]mappers.Componentmapper, 0),
	}
}

func MapperPage() *pagemapper.Page {
	return &pagemapper.Page{
		SourceKey: "template_page_1",
		Rows:      make([]mappers.OrderedComponents, 0),
	}
}

func MapperList() *listmapper.List {
	return &listmapper.List{}
}

func Barcode() *codemapper.Barcode {
	return &codemapper.Barcode{Order: 1}
}

func Matrixcode() *codemapper.Matrixcode {
	return &codemapper.Matrixcode{Order: 1}
}

func Qrcode() *codemapper.Qrcode {
	return &codemapper.Qrcode{Order: 1}
}

func Image() *imagemapper.Image {
	return &imagemapper.Image{Order: 1}
}

func Line() *linemapper.Line {
	return &linemapper.Line{Order: 1}
}

func Signature() *signaturemapper.Signature {
	return &signaturemapper.Signature{Order: 1}
}

func Text() *textmapper.Text {
	return &textmapper.Text{Order: 1}
}

func Metadata() *propsmapper.Metadata {
	creation := time.Now()
	return &propsmapper.Metadata{
		Author: &propsmapper.Utf8Text{Text: "Author", UTF8: true}, Creator: &propsmapper.Utf8Text{Text: "Creator", UTF8: true},
		Subject: &propsmapper.Utf8Text{Text: "Subject", UTF8: true}, Title: &propsmapper.Utf8Text{Text: "Title", UTF8: true},
		CreationDate: &creation, KeywordsStr: &propsmapper.Utf8Text{Text: "KeywordsStr", UTF8: true},
	}
}

func BuilderProps() *buildermapper.Builder {
	time, _ := time.Parse("2006-01-02 15:04:05", "2024-10-09 14:30:00")
	return &buildermapper.Builder{
		Dimensions: &propsmapper.Dimensions{
			Width:  10.0,
			Height: 10.0,
		},
		Margins: &propsmapper.Margins{
			Left:   10.0,
			Right:  10.0,
			Top:    10.0,
			Bottom: 10.0,
		},
		SequentialMode:          false,
		ConcurrentMode:          10,
		SequentialLowMemoryMode: -1,
		Debug:                   true,
		MaxGridSize:             10,
		DefaultFont: &propsmapper.Font{
			Family: "Arial",
			Style:  "bold",
			Size:   10,
			Color: &propsmapper.Color{
				Red:   10,
				Green: 100,
				Blue:  150,
			},
		},
		PageNumber: &propsmapper.PageNumber{
			Pattern: "pattern_test",
			Place:   "place_test",
			Family:  "family_test",
			Style:   "style_test",
			Size:    10.0,
			Color: &propsmapper.Color{
				Red:   10,
				Green: 100,
				Blue:  150,
			},
		},
		CustomFonts: []*propsmapper.CustomFont{
			{Family: "family_test", Style: "style_test", File: "file_test"},
			{Family: "family_test2", Style: "style_test2", File: "file_test2"},
		},
		Protection: &propsmapper.Protection{
			Type:          4,
			UserPassword:  "senha123",
			OwnerPassword: "senha123",
		},
		Compression: true,
		PageSize:    "T",
		Orientation: "vertical",
		Metadata: &propsmapper.Metadata{
			Author:       &propsmapper.Utf8Text{Text: "user_test", UTF8: true},
			Creator:      &propsmapper.Utf8Text{Text: "user_test", UTF8: true},
			Subject:      &propsmapper.Utf8Text{Text: "test", UTF8: true},
			Title:        &propsmapper.Utf8Text{Text: "report", UTF8: true},
			CreationDate: &time,
			KeywordsStr:  &propsmapper.Utf8Text{Text: "test", UTF8: true},
		},
		DisableAutoPageBreak: true,
		GenerationMode:       "concurrent",
	}
}

func Row(sourceKeyRow, sourceKeyText string) *rowmapper.Row {
	col := colmapper.Col{
		Size:       12,
		Components: []mappers.OrderedComponents{},
	}

	return &rowmapper.Row{
		Height:    10,
		SourceKey: sourceKeyRow,
		Cols:      []mappers.Componentmapper{&col},
	}
}

func Page(sourceKeyPage, sourceKeyRow, sourceKeyText string) *pagemapper.Page {
	col := colmapper.Col{
		Size:       12,
		Components: []mappers.OrderedComponents{},
	}

	return &pagemapper.Page{
		SourceKey: sourceKeyPage,
		Rows: []mappers.OrderedComponents{
			&rowmapper.Row{
				Height:    10,
				SourceKey: sourceKeyRow,
				Cols:      []mappers.Componentmapper{&col},
			},
		},
	}
}
