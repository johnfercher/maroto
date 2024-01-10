// Package contains all border types.
package border

type Type string

const (
	None   Type = ""
	Full   Type = "1"
	Left   Type = "L"
	Top    Type = "T"
	Right  Type = "R"
	Bottom Type = "B"
)

func (t Type) IsValid() bool {
	return t == Full || t == Left || t == Top || t == Right || t == Bottom
}
