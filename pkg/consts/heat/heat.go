package heat

type Style string

const (
	RectStyle   Style = "rect"
	CircleStyle Style = "circle"
)

func (s Style) IsValid() bool {
	return s == RectStyle || s == CircleStyle
}
