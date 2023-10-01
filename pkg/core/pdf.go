package core

import (
	"encoding/base64"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/metrics"
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
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(p.bytes)
	if err != nil {
		return err
	}

	return nil
}
