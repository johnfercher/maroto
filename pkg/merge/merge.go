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

// pdfcpuMu serialises calls into github.com/pdfcpu/pdfcpu, whose
// configuration loader (api.LoadConfiguration / NewDefaultConfiguration)
// writes to package-level globals without synchronisation. Concurrent
// callers of merge.Bytes therefore race on those globals, so we
// serialise entry into pdfcpu here.
var pdfcpuMu sync.Mutex

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
	pdfcpuMu.Lock()
	defer pdfcpuMu.Unlock()

	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	err := api.MergeRaw(readers, writer, dividerPage, conf)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotMergePDFs, err)
	}

	return nil
}
