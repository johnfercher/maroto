package types

import "slices"

const (
	Document   = "document"
	Row        = "row"
	Page       = "page"
	Col        = "col"
	Image      = "image"
	Text       = "text"
	Signature  = "signature"
	Barcode    = "barcode"
	QrCode     = "qrcode"
	MatrixCode = "matrixcode"
)

type DocumentType string

func (t DocumentType) String() string {
	return string(t)
}

func (t DocumentType) Accept(dt string) bool {
	if val, ok := buildAcceptedMap()[t.String()]; ok {
		return slices.Contains(val, dt)
	}
	return false
}

func buildAcceptedMap() map[string][]string {
	return map[string][]string{
		Document: {Row, Page},
		Page:     {Row},
		Row:      {Col},
		Col:      {Row, Image, Text, Signature, Barcode, QrCode, MatrixCode},
	}
}
