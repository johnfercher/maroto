package propsmapper

import (
	"time"
)

// Metadata is the representation of a PDF metadata.
type Metadata struct {
	Author       *Utf8Text
	Creator      *Utf8Text
	Subject      *Utf8Text
	Title        *Utf8Text
	CreationDate *time.Time
	KeywordsStr  *Utf8Text
}

// Utf8Text is the representation of a text with a flag to indicate if it's UTF8.
type Utf8Text struct {
	Text string
	UTF8 bool
}

// NewUtf8Text is responsible for creating the Utf8Text, if the Utf8Text fields cannot be
// converted, an invalid value is set.
func NewUtf8Text(utf8Text interface{}) *Utf8Text {
	utf8TextMap, ok := utf8Text.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Utf8Text{
		Text: *convertFields(utf8TextMap["text"], ""),
		UTF8: *convertFields(utf8TextMap["utf8"], true),
	}
}

// NewMetadata is responsible for creating the metadata, if the metadata fields cannot be
// converted, an invalid value is set.
func NewMetadata(metadata interface{}) *Metadata {
	metadataMap, ok := metadata.(map[string]interface{})
	if !ok {
		return nil
	}

	return &Metadata{
		Author:       NewUtf8Text(metadataMap["author"]),
		Creator:      NewUtf8Text(metadataMap["creator"]),
		Subject:      NewUtf8Text(metadataMap["subject"]),
		Title:        NewUtf8Text(metadataMap["title"]),
		CreationDate: factoryTime(*convertFields(metadataMap["creation_date"], ""), "2006-01-02 15:04:05", time.Now()),
		KeywordsStr:  NewUtf8Text(metadataMap["keywords_str"]),
	}
}
