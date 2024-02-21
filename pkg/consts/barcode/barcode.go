package barcode

type Type string

const (
	// Code128 represents the default barcode type of maroto.
	Code128 Type = "code128"
	// EAN represents the ean barcode type.
	EAN Type = "ean"
	// Codabar   Type = "codabar"
	// Code39    Type = "code39"
	// Code93    Type = "code93"
	// PDF417    Type = "pdf417"
	// TwoOfFive Type = "twooffive"
)
