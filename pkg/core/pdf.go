package core

import (
	"encoding/base64"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/merge"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/johnfercher/maroto/v2/pkg/time"
)

type pdf struct {
	bytes  []byte
	report *metrics.Report
}

func NewPDF(bytes []byte, report *metrics.Report) Document {
	return &pdf{
		bytes:  bytes,
		report: report,
	}
}

func (p *pdf) GetBytes() []byte {
	return p.bytes
}

func (p *pdf) GetBase64() string {
	return base64.StdEncoding.EncodeToString(p.bytes)
}

func (p *pdf) GetReport() *metrics.Report {
	return p.report
}

func (p *pdf) Save(file string) error {
	return os.WriteFile(file, p.bytes, os.ModePerm)
}

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
