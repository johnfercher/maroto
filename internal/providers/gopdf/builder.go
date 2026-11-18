package gopdf

import (
	"github.com/johnfercher/maroto/v2/internal/cache"
	"github.com/johnfercher/maroto/v2/internal/code"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/internal/providers/gopdf/cellwriter"
	"github.com/johnfercher/maroto/v2/internal/providers/gopdf/gopdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// Dependencies is the dependencies provider for gopdf.
type Dependencies struct {
	Fpdf       gopdfwrapper.Gopdf
	Font       core.Font
	Text       core.Text
	Code       core.Code
	Image      core.Image
	Line       core.Line
	Checkbox   core.Checkbox
	Cache      cache.Cache
	CellWriter cellwriter.CellWriter
	Cfg        *entity.Config
}

// Builder is the dependencies builder for gopdf.
type Builder interface {
	Build(cfg *entity.Config, cache cache.Cache) *Dependencies
}

type builder struct{}

// NewBuilder create a new Builder
func NewBuilder() Builder {
	return &builder{}
}

// Build create a new Dependencies.
func (b *builder) Build(cfg *entity.Config, cache cache.Cache) *Dependencies {
	fpdf := gopdfwrapper.New(gopdfwrapper.Options{
		Width:                cfg.Dimensions.Width,
		Height:               cfg.Dimensions.Height,
		Margins:              *cfg.Margins,
		DisableAutoPageBreak: cfg.DisableAutoPageBreak,
		Protection:           cfg.Protection,
		Fonts:                gopdfwrapper.NewFonts(cfg.CustomFonts),
	})

	font := NewFont(fpdf, cfg.DefaultFont.Size, cfg.DefaultFont.Family, cfg.DefaultFont.Style)
	math := math.New()
	code := code.New()
	text := NewText(fpdf, math, font)
	image := NewImage(fpdf, math)
	line := NewLine(fpdf)
	checkbox := NewCheckbox(fpdf, font)
	cellWriter := cellwriter.NewBuilder().
		Build(fpdf)

	return &Dependencies{
		Fpdf:       fpdf,
		Font:       font,
		Text:       text,
		Code:       code,
		Image:      image,
		Line:       line,
		Checkbox:   checkbox,
		CellWriter: cellWriter,
		Cfg:        cfg,
		Cache:      cache,
	}
}
