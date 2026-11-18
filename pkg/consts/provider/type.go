// Package provider contains all document generation providers.
package provider

// Type is a representation of a document generation provider.
type Type string

const (
	// Gopdf represents the gopdf (github.com/signintech/gopdf) provider.
	Gopdf Type = "gopdf"
	// Gofpdf represents the gofpdf provider, which is no longer used.
	//
	// Deprecated: maroto renders with gopdf, use Gopdf.
	Gofpdf Type = "gofpdf"
)
