package col

import (
	"image"

	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/barcode"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/components/text"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/props/col"
)

type Col struct {
	Props   col.ColProps      `json:"props"`
	Text    []text.Text       `json:"text"`
	BarCode []barcode.BarCode `json:"bar_code"`
	Image   []image.Image     `json:"image"`
}
