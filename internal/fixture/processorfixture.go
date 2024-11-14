package fixture

import (
	"time"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/codemapper"
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
		Rows:      make([]mappers.Componentmapper, 0),
	}
}

func MapperList() *listmapper.List {
	return &listmapper.List{}
}

func Barcode() *codemapper.Barcode {
	return &codemapper.Barcode{}
}

func Matrixcode() *codemapper.Matrixcode {
	return &codemapper.Matrixcode{}
}

func Qrcode() *codemapper.Qrcode {
	return &codemapper.Qrcode{}
}

func Image() *imagemapper.Image {
	return &imagemapper.Image{}
}

func Line() *linemapper.Line {
	return &linemapper.Line{}
}

func Signature() *signaturemapper.Signature {
	return &signaturemapper.Signature{}
}

func Text() *textmapper.Text {
	return &textmapper.Text{}
}

func Metadata() *propsmapper.Metadata {
	creation := time.Now()
	return &propsmapper.Metadata{
		Author: &propsmapper.Utf8Text{Text: "Author", UTF8: true}, Creator: &propsmapper.Utf8Text{Text: "Creator", UTF8: true},
		Subject: &propsmapper.Utf8Text{Text: "Subject", UTF8: true}, Title: &propsmapper.Utf8Text{Text: "Title", UTF8: true},
		CreationDate: &creation, KeywordsStr: &propsmapper.Utf8Text{Text: "KeywordsStr", UTF8: true},
	}
}
