// Package protection contains all protection types.
package protection

// Type is a representation of a protection type.
type Type byte

const (
	// None represents no protection.
	None Type = 0
	// Print represents a print protection.
	Print Type = 4
	// Modify represents modification protection.
	Modify Type = 8
	// Copy represents copy protection.
	Copy Type = 16
	// AnnotForms represents annotation and form protection.
	AnnotForms Type = 32
)
