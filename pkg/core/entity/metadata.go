// Package contains all core entities.
package entity

import (
	"fmt"
	"time"
)

type Metadata struct {
	Author       *Utf8Text
	Creator      *Utf8Text
	Subject      *Utf8Text
	Title        *Utf8Text
	CreationDate *time.Time
}

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

	return mp
}

type Utf8Text struct {
	Text string
	UTF8 bool
}

func (u *Utf8Text) ToString() string {
	return fmt.Sprintf("Utf8Text(%s, %v)", u.Text, u.UTF8)
}
