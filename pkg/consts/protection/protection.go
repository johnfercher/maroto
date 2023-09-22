package protection

type Type byte

const (
	Print      Type = 4
	Modify     Type = 8
	Copy       Type = 16
	AnnotForms Type = 32
)
