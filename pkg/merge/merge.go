// Package merge implements PDF merge.
package merge

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"sync"

	"github.com/pdfcpu/pdfcpu/pkg/api"
)

var ErrCannotMergePDFs = errors.New("cannot merge PDFs")

// pdfcpuConfigMu serializes access to api.LoadConfiguration, which lazily
// initializes (and reads/writes) pdfcpu's package-level default
// configuration. Without it, parallel callers race on those globals — see
// data races reported by `go test -race` in pkg/core.
var pdfcpuConfigMu sync.Mutex

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
	pdfcpuConfigMu.Lock()
	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	pdfcpuConfigMu.Unlock()

	err := api.MergeRaw(readers, writer, dividerPage, conf)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotMergePDFs, err)
	}

	return nil
}
