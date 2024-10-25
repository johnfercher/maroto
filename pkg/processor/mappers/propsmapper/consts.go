package propsmapper

func NewAlign(align string) string {
	mapAligns := map[string]string{
		"Left": "L", "Right": "R", "Center": "C", "Top": "T", "Bottom": "B", "Middle": "M", "Justify": "J",
	}
	align, ok := mapAligns[align]
	if !ok {
		return ""
	}
	return align
}

func NewBreakLineStrategy(strategy string) string {
	switch strategy {
	case "EmptySpaceStrategy":
		return "empty_space_strategy"
	case "DashStrategy":
		return "dash_strategy"
	}
	return ""
}

func NewCodeType(typeProtection string) string {
	if typeProtection == "EAN" {
		return "ean"
	}
	return "code128"
}

func NewLineStyle(style string) string {
	if style != "dashed" && style != "solid" {
		return ""
	}
	return style
}

func NewOrientation(orientation string) string {
	switch orientation {
	case "vertical":
		return "vertical"
	case "horizontal":
		return "horizontal"
	}
	return ""
}

func NewTypeProtection(typeProtection string) byte {
	switch typeProtection {
	case "Print":
		return 4
	case "Modify":
		return 8
	case "Copy":
		return 16
	case "AnnotForms":
		return 32
	}

	return 0
}

func NewFontStyle(fontType string) string {
	switch fontType {
	case "Bold":
		return "B"
	case "Italic":
		return "I"
	case "BoldItalic":
		return "BI"
	}

	return ""
}
