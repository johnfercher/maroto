// Package core contains all core interfaces and basic implementations.
package core

import (
	"encoding/base64"
	"os"

	"github.com/johnfercher/maroto/v2/internal/time"
	"github.com/johnfercher/maroto/v2/pkg/merge"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

type pdf struct {
	bytes  []byte
	report *metrics.Report
}

// NewPDF is responsible to create a new instance of PDF.
func NewPDF(bytes []byte, report *metrics.Report) Document {
	return &pdf{
		bytes:  bytes,
		report: report,
	}
}

// GetBytes returns the PDF bytes.
func (p *pdf) GetBytes() []byte {
	return p.bytes
}

// GetBase64 returns the PDF bytes in base64.
func (p *pdf) GetBase64() string {
	return base64.StdEncoding.EncodeToString(p.bytes)
}

// GetReport returns the metrics.Report.
func (p *pdf) GetReport() *metrics.Report {
	return p.report
}

// Save saves the PDF in a file.
func (p *pdf) Save(file string) error {
	return os.WriteFile(file, p.bytes, os.ModePerm)
}

// Merge merges the PDF with another PDF.
func (p *pdf) Merge(bytes []byte) error {
	var mergedBytes []byte
	var err error

	timeSpent := time.GetTimeSpent(func() {
		mergedBytes, err = merge.Bytes(p.bytes, bytes)
	})
	if err != nil {
		return err
	}
	p.bytes = mergedBytes
	if p.report != nil {
		p.appendMetric(timeSpent)
	}

	return nil
}

func (p *pdf) appendMetric(timeSpent *metrics.Time) {
	timeMetric := metrics.TimeMetric{
		Key:   "merge_pdf",
		Times: []*metrics.Time{timeSpent},
		Avg:   timeSpent,
	}
	timeMetric.Normalize()
	p.report.TimeMetrics = append(p.report.TimeMetrics, timeMetric)

	p.report.SizeMetric = metrics.SizeMetric{
		Key: "file_size",
		Size: metrics.Size{
			Value: float64(len(p.bytes)),
			Scale: metrics.Byte,
		},
	}
	p.report.Normalize()
}
