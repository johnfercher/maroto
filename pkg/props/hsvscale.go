package props

import "github.com/johnfercher/maroto/v2/pkg/consts/hsv"

type HSVScale struct {
	Begin hsv.Value
	End   hsv.Value
}

func (h *HSVScale) MakeValid() {
	if h.Begin == 0 {
		h.Begin = hsv.Green
	}
	if h.End == 0 {
		h.End = hsv.Red
	}

	if h.End > h.Begin {
		h.Begin, h.End = h.End, h.Begin
	}
}
