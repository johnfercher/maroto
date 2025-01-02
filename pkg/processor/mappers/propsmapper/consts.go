package propsmapper

func NewAlign(align string) string {
	mapAligns := map[string]string{
		"left": "L", "right": "R", "center": "C", "top": "T", "bottom": "B", "middle": "M", "justify": "J",
	}
	align, ok := mapAligns[align]
	if !ok {
		return ""
	}
	return align
}

func NewBorder(border string) string {
	mapBorder := map[string]string{
		"full": "1", "top": "T", "bottom": "B", "left": "L", "right": "R",
	}
	border, ok := mapBorder[border]
	if !ok {
		return ""
	}
	return border
}

func NewBreakLineStrategy(strategy string) string {
	switch strategy {
	case "empty_space_strategy":
		return strategy
	case "dash_strategy":
		return strategy
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
	case "print":
		return 4
	case "modify":
		return 8
	case "copy":
		return 16
	case "annot_forms":
		return 32
	}

	return 0
}

func NewFontStyle(fontType string) string {
	switch fontType {
	case "bold":
		return "B"
	case "italic":
		return "I"
	case "bold_italic":
		return "BI"
	}

	return ""
}
