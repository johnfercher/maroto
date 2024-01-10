package fixture

import (
	"github.com/johnfercher/go-tree/node"

	"github.com/johnfercher/maroto/v2/pkg/consts/align"
	"github.com/johnfercher/maroto/v2/pkg/consts/border"
	"github.com/johnfercher/maroto/v2/pkg/consts/breakline"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/linestyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/orientation"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

// TextProp is responsible to give a valid props.Text.
func TextProp() props.Text {
	fontProp := FontProp()

	google := "https://www.google.com"

	prop := props.Text{
		Top:               12,
		Left:              3,
		Family:            fontProp.Family,
		Style:             fontProp.Style,
		Size:              fontProp.Size,
		Align:             align.Right,
		BreakLineStrategy: breakline.DashStrategy,
		VerticalPadding:   20,
		Color:             fontProp.Color,
		Hyperlink:         &google,
	}
	prop.MakeValid(&fontProp)
	return prop
}

// FontProp is responsible to give a valid props.Text.
func FontProp() props.Font {
	colorProp := ColorProp()
	prop := props.Font{
		Family: fontfamily.Helvetica,
		Style:  fontstyle.Bold,
		Size:   14,
		Color:  &colorProp,
	}
	prop.MakeValid(fontfamily.Arial)
	return prop
}

// BarcodeProp is responsible to give a valid props.Barcode.
func BarcodeProp() props.Barcode {
	prop := props.Barcode{
		Top:     10,
		Left:    10,
		Percent: 98,
		Proportion: props.Proportion{
			Width:  16,
			Height: 9,
		},
		Center: false,
	}
	prop.MakeValid()
	return prop
}

// RectProp is responsible to give a valid props.Rect.
func RectProp() props.Rect {
	prop := props.Rect{
		Top:     10,
		Left:    10,
		Percent: 98,
		Center:  false,
	}
	prop.MakeValid()
	return prop
}

// CellEntity is responsible to give a valid entity.Cell.
func CellEntity() entity.Cell {
	return entity.Cell{
		X:      10,
		Y:      15,
		Width:  100,
		Height: 150,
	}
}

// MarginsEntity is responsible to give a valid entity.Margins.
func MarginsEntity() entity.Margins {
	return entity.Margins{
		Left:   10,
		Top:    10,
		Right:  10,
		Bottom: 10,
	}
}

// ImageEntity is responsible to give a valid entity.Image.
func ImageEntity() entity.Image {
	return entity.Image{
		Bytes:     []byte{1, 2, 3},
		Extension: extension.Png,
	}
}

// CellProp is responsible to give a valid props.Cell.
func CellProp() props.Cell {
	prop := props.Cell{
		BackgroundColor: &props.Color{
			Red:   255,
			Green: 100,
			Blue:  50,
		},
		BorderColor: &props.Color{
			Red:   200,
			Green: 80,
			Blue:  60,
		},
		BorderType:      border.Left,
		BorderThickness: 0.6,
		LineStyle:       linestyle.Dashed,
	}
	return prop
}

// Node is responsible to give a valid node.Node.
func Node(rootType string) *node.Node[core.Structure] {
	marotoNode := node.New[core.Structure](core.Structure{
		Type: rootType,
	})
	pageNode := node.New[core.Structure](core.Structure{
		Type: "page",
	})

	marotoNode.AddNext(pageNode)
	return marotoNode
}

// ColorProp is responsible to give a valid props.Color.
func ColorProp() props.Color {
	return props.Color{
		Red:   100,
		Green: 50,
		Blue:  200,
	}
}

// LineProp is responsible to give a valid props.Line.
func LineProp() props.Line {
	colorProp := ColorProp()
	prop := props.Line{
		Color:         &colorProp,
		Style:         linestyle.Dashed,
		Thickness:     1.1,
		Orientation:   orientation.Vertical,
		OffsetPercent: 50,
		SizePercent:   20,
	}
	prop.MakeValid()
	return prop
}

// SignatureProp is responsible to give a valid props.Signature.
func SignatureProp() props.Signature {
	textProp := TextProp()
	lineProp := LineProp()
	prop := props.Signature{
		FontFamily:    textProp.Family,
		FontStyle:     textProp.Style,
		FontSize:      textProp.Size,
		FontColor:     textProp.Color,
		LineColor:     lineProp.Color,
		LineStyle:     lineProp.Style,
		LineThickness: lineProp.Thickness,
	}
	prop.MakeValid(textProp.Family)
	return prop
}

// PageProp is responsible to give a valid props.Page.
func PageProp() props.Page {
	fontProp := FontProp()
	prop := props.Page{
		Pattern: "{current} / {total}",
		Place:   props.LeftBottom,
		Family:  fontProp.Family,
		Style:   fontProp.Style,
		Size:    fontProp.Size,
		Color:   fontProp.Color,
	}
	return prop
}
