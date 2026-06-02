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

// pdfcpuConfigMu serializes calls into pdfcpu's api.LoadConfiguration,
// which mutates package-level state inside pdfcpu and is therefore not
// safe for concurrent use. Without this guard, parallel callers (e.g.
// pkg/core tests running with t.Parallel()) trigger a data race in
// pdfcpu/pkg/pdfcpu/model. Holding the lock for the whole merge keeps
// the configuration stable for the duration of api.MergeRaw too.
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
	defer pdfcpuConfigMu.Unlock()

	conf := api.LoadConfiguration()
	conf.WriteXRefStream = false
	err := api.MergeRaw(readers, writer, dividerPage, conf)
	if err != nil {
		return fmt.Errorf("%w: %w", ErrCannotMergePDFs, err)
	}

	return nil
}
