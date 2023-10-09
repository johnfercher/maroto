package core

import (
	"github.com/johnfercher/maroto/v2/pkg/cache"
	"github.com/johnfercher/maroto/v2/pkg/consts/extension"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type Provider interface {
	// Grid
	CreateRow(height float64)
	CreateCol(width, height float64, config *entity.Config, prop *props.Cell)

	// Features
	AddLine(cell *entity.Cell, prop *props.Line)
	AddText(text string, cell *entity.Cell, prop *props.Text)
	GetTextHeight(prop *props.Font) float64
	AddSignature(text string, cell *entity.Cell, prop *props.Text)
	AddMatrixCode(code string, cell *entity.Cell, prop *props.Rect)
	AddQrCode(code string, cell *entity.Cell, rect *props.Rect)
	AddBarCode(code string, cell *entity.Cell, prop *props.Barcode)
	AddImageFromFile(value string, cell *entity.Cell, prop *props.Rect)
	AddImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type)
	AddBackgroundImageFromBytes(bytes []byte, cell *entity.Cell, prop *props.Rect, extension extension.Type)

	// General
	GenerateBytes() ([]byte, error)
	SetCache(cache cache.Cache)

	SetProtection(protection *entity.Protection)
	SetCompression(compression bool)
	SetMetadata(metadata *entity.Metadata)
}
