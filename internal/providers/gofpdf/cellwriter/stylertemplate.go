package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type StylerTemplate struct {
	next CellWriter
	fpdf gofpdfwrapper.Fpdf
}

func (s *StylerTemplate) SetNext(next CellWriter) {
	s.next = next
}

func (s *StylerTemplate) GoToNext(width, height float64, config *entity.Config, prop *props.Cell) {
	if s.next == nil {
		return
	}

	s.next.Apply(width, height, config, prop)
}
