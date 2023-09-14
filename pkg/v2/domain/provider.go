package domain

import (
	"github.com/johnfercher/maroto/internal"
	"github.com/johnfercher/maroto/pkg/props"
)

type Provider interface {
	CreateRow(height float64)
	CreateCol(width, height float64)
	GetDimensions() (width float64, height float64)
	GetMargins() (left float64, top float64, right float64, bottom float64)
	AddText(text string, cell internal.Cell, prop props.Text)
	Generate(file string) error
}
