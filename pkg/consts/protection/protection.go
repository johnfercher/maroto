// Package contains all protection types.
package protection

type Type byte

const (
	None       Type = 0
	Print      Type = 4
	Modify     Type = 8
	Copy       Type = 16
	AnnotForms Type = 32
)
