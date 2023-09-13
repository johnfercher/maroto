package v2

import (
	"github.com/johnfercher/maroto/internal/fpdf"
	"github.com/johnfercher/maroto/pkg/v2/context"
)

type Component interface {
	Render(fpdf fpdf.Fpdf, ctx context.Context)
	GetType() string
	Add(components ...Component) Component
}
