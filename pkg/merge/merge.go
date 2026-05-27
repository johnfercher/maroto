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

// pdfcpuMu serializes calls into the pdfcpu library. pdfcpu lazily
// initializes (and reads) package-global configuration state on the
// first call to api.LoadConfiguration; concurrent first-time calls
// from parallel callers are not race-safe. Holding this mutex around
// every merge keeps that initialization deterministic without
// exposing pdfcpu's internals to our callers.
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
