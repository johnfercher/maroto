package gofpdf

import (
	"github.com/phpdave11/gofpdf"

	"github.com/johnfercher/maroto/v2/internal/cache"
	"github.com/johnfercher/maroto/v2/internal/code"
	"github.com/johnfercher/maroto/v2/internal/math"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/cellwriter"
	"github.com/johnfercher/maroto/v2/internal/providers/gofpdf/gofpdfwrapper"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
)

// Dependencies is the dependencies provider for gofpdf.
type Dependencies struct {
	Fpdf       gofpdfwrapper.Fpdf
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

// Builder is the dependencies builder for gofpdf.
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
	fpdf := gofpdfwrapper.NewCustom(&gofpdf.InitType{
		OrientationStr: "P",
		UnitStr:        "mm",
		FontDirStr:     "",
		Size: gofpdf.SizeType{
			Wd: cfg.Dimensions.Width,
			Ht: cfg.Dimensions.Height,
		},
	})

	for _, font := range cfg.CustomFonts {
		// The underlying gofpdf library stores this byte slice by reference
		// and mutates it later (during Output/Close -> putfonts ->
		// utf8FontFile.generateChecksum/assembleTables). When the same
		// *entity.Config (and therefore the same CustomFont byte slices) is
		// shared across multiple providers — as happens in concurrent and
		// low-memory generation modes where each worker builds its own Fpdf
		// instance — those workers would otherwise race on the same
		// underlying font bytes. Give every provider its own private copy
		// so gofpdf's in-place mutations are safely isolated.
		// See: https://github.com/johnfercher/maroto/issues/550
		original := font.GetBytes()
		fontBytes := make([]byte, len(original))
		copy(fontBytes, original)
		fpdf.AddUTF8FontFromBytes(font.GetFamily(), string(font.GetStyle()), fontBytes)
	}

	if cfg.DisableAutoPageBreak {
		fpdf.SetAutoPageBreak(false, 0)
	} else {
		fpdf.SetAutoPageBreak(true, cfg.Margins.Bottom)
	}

	fpdf.SetMargins(cfg.Margins.Left, cfg.Margins.Top, cfg.Margins.Right)
	fpdf.AddPage()

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
