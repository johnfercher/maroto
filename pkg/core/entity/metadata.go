package entity

import (
	"fmt"
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

// AppendMap appends the metadata to a map.
func (m *Metadata) AppendMap(mp map[string]interface{}) map[string]interface{} {
	if m.Author != nil {
		mp["config_metadata_author"] = m.Author.ToString()
	}

	if m.Creator != nil {
		mp["config_metadata_creator"] = m.Creator.ToString()
	}

	if m.Subject != nil {
		mp["config_metadata_subject"] = m.Subject.ToString()
	}

	if m.Title != nil {
		mp["config_metadata_title"] = m.Title.ToString()
	}

	if m.CreationDate != nil {
		mp["config_metadata_creation_date"] = true
	}

	if m.KeywordsStr != nil {
		mp["config_metadata_keywords"] = m.KeywordsStr.ToString()
	}

	return mp
}

// Utf8Text is the representation of a text with a flag to indicate if it's UTF8.
type Utf8Text struct {
	Text string
	UTF8 bool
}

// ToString returns a string representation of the text.
func (u *Utf8Text) ToString() string {
	return fmt.Sprintf("Utf8Text(%s, %v)", u.Text, u.UTF8)
}
