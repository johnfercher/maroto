package entity

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMetadata_AppendMap(t *testing.T) {
	// Arrange
	sut := fixtureMetadata()
	m := make(map[string]interface{})

	// Act
	sut.AppendMap(m)

	// Assert
	assert.Equal(t, "Utf8Text(author, true)", m["config_metadata_author"])
	assert.Equal(t, "Utf8Text(creator, false)", m["config_metadata_creator"])
	assert.Equal(t, "Utf8Text(subject, true)", m["config_metadata_subject"])
	assert.Equal(t, "Utf8Text(title, true)", m["config_metadata_title"])
	assert.Equal(t, true, m["config_metadata_creation_date"])
}

func fixtureMetadata() Metadata {
	now := time.Now()
	return Metadata{
		Author: &Utf8Text{
			Text: "author",
			UTF8: true,
		},
		Creator: &Utf8Text{
			Text: "creator",
			UTF8: false,
		},
		Subject: &Utf8Text{
			Text: "subject",
			UTF8: true,
		},
		Title: &Utf8Text{
			Text: "title",
			UTF8: true,
		},
		CreationDate: &now,
	}
}
