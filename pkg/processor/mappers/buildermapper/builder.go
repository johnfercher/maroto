package buildermapper

import "github.com/johnfercher/maroto/v2/pkg/processor/mappers/propsmapper"

type Builder struct {
	Dimensions           propsmapper.Dimensions   `json:"dimensions"`
	Margins              propsmapper.Margins      `json:"margins"`
	ChunkWorkers         int                      `json:"chunk_workers"`
	Debug                bool                     `json:"debug"`
	MaxGridSize          int                      `json:"max_grid_size"`
	DefaultFont          propsmapper.Font         `json:"default_font"`
	CustomFonts          []propsmapper.CustomFont `json:"custom_fonts"`
	PageNumber           propsmapper.PageNumber   `json:"page_number"`
	Protection           propsmapper.Protection   `json:"protection"`
	Compression          bool                     `json:"compression"`
	PageSize             string                   `json:"page_size"`
	Orientation          string                   `json:"orientation"`
	Metadata             propsmapper.Metadata     `json:"metadata"`
	DisableAutoPageBreak bool                     `json:"disable_auto_page_break"`
	GenerationMode       string                   `json:"generation_mode"`
}
