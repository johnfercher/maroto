package v2

import "github.com/johnfercher/maroto/internal/fpdf"

type Component interface {
	Render(fpdf fpdf.Fpdf, ctx *Context)
	GetType() string
	Add(component ...Component)
}
