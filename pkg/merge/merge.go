package merge

import (
	"bytes"
	"io"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

func Bytes(pdfs ...[]byte) ([]byte, error) {
	readers := make([]io.ReadSeeker, len(pdfs))
	for i, pdf := range pdfs {
		readers[i] = bytes.NewReader(pdf)
	}

	var buf bytes.Buffer
	writer := io.Writer(&buf)
	err := mergePdfs(readers, writer)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func mergePdfs(readers []io.ReadSeeker, writer io.Writer) error {
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	return api.MergeRaw(readers, writer, conf)
}
