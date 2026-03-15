// Package merge implements PDF merge.
package merge

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

var ErrCannotMergePDFs = errors.New("cannot merge PDFs")

// Bytes merges PDFs from byte slices.
func Bytes(pdfs ...[]byte) ([]byte, error) {
	readers := make([]io.ReadSeeker, len(pdfs))
	for i, pdf := range pdfs {
		readers[i] = bytes.NewReader(pdf)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	err := mergePdfs(readers, writer, false)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer, dividerPage bool) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	err := api.MergeRaw(readers, writer, dividerPage, conf)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotMergePDFs, err)
	}

	return nil
}
