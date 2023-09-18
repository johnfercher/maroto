package domain

import (
	"encoding/base64"
	"github.com/johnfercher/v2/maroto/metrics"
	"os"
)

type document struct {
	bytes  []byte
	report *metrics.Report
}

func NewDocument(bytes []byte, report *metrics.Report) Document {
	return &document{
		bytes:  bytes,
		report: report,
	}
}

func (r *document) GetBytes() []byte {
	return r.bytes
}

func (r *document) GetBase64() string {
	return base64.StdEncoding.EncodeToString(r.bytes)
}

func (r *document) GetReport() *metrics.Report {
	return r.report
}

func (r *document) Save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(r.bytes)
	if err != nil {
		return err
	}

	return nil
}
