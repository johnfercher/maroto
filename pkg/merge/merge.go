// Package merge implements PDF merge.
package merge

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

// Bytes merges PDFs from byte slices.
func Bytes(pdfs ...[]byte) ([]byte, error) {
	readers := make([]io.ReadSeeker, len(pdfs))
	for i, pdf := range pdfs {
		readers[i] = bytes.NewReader(pdf)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	if err := mergePdfs(readers, writer, false); err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer, dividerPage bool) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, dividerPage, conf)
}
