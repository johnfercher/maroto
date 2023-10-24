package cellwriter

import (
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
)

type stylerTemplate struct {
	next CellWriter
	fpdf gofpdfwrapper.Fpdf
	name string
}

func (s *stylerTemplate) SetNext(next CellWriter) {
	s.next = next
}

func (s *stylerTemplate) GetName() string {
	return s.name
}

func (s *stylerTemplate) GetNext() CellWriter {
	return s.next
}

func (s *stylerTemplate) GoToNext(width, height float64, config *entity.Config, prop *props.Cell) {
	if s.next == nil {
		return
	}

	s.next.Apply(width, height, config, prop)
}
