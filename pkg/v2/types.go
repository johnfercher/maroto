package v2

const (
	Document = "document"
	Row      = "row"
	Page     = "page"
	Col      = "col"
	Image    = "image"
	Text     = "text"
)

type DocumentType string

func (t DocumentType) String() string {
	return string(t)
}

func (t DocumentType) Accept(dt string) bool {
	_, ok := buildAcceptedMap()[dt]
	return ok
}

func buildAcceptedMap() map[string][]string {
	return map[string][]string{
		Document: {Row},
		Page:     {Row},
		Row:      {Col},
		Col:      {Row, Image, Text},
	}
}
