package gopdfwrapper

import (
	"bytes"
	"context"
	"fmt"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/model"
)

// addKeywords writes pdf to w with keywords in its information dictionary,
// gopdf has no way to write them itself.
func addKeywords(pdf []byte, keywords string, protection *passwords, w io.Writer) error {
	conf := model.NewDefaultConfiguration()
	conf.ValidationMode = model.ValidationRelaxed
	// Keeps the producer and dates written by gopdf.
	conf.PreserveInfoDict = true
	if protection != nil {
		conf.UserPW = protection.user
		conf.OwnerPW = protection.owner
	}

	err := api.AddKeywords(context.Background(), bytes.NewReader(pdf), w, []string{keywords}, conf)
	if err != nil {
		return fmt.Errorf("could not add keywords: %w", err)
	}

	return nil
}

type passwords struct {
	user  string
	owner string
}
