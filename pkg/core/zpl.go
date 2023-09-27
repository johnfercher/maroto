package core

import (
	"encoding/base64"
	"errors"
	"github.com/johnfercher/maroto/v2/pkg/consts/documenttype"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/johnfercher/maroto/v2/pkg/time"
	"github.com/karmdip-mi/go-fitz"
	"log"
	"os"
	"simonwaldherr.de/go/zplgfa"
)

type zpl struct {
	bytes        []byte
	report       *metrics.Report
	documentType documenttype.DocumentType
}

func NewZPLFromPDF(document Document) Document {
	docBytes := document.GetBytes()

	var zplBytes []byte
	timeSpent := time.GetTimeSpent(func() {
		docFitz, err := fitz.NewFromMemory(docBytes)
		if err != nil {
			log.Fatal(err.Error())
		}

		page := 0
		img, err := docFitz.Image(page)
		if err != nil {
			panic(err)
		}

		flat := zplgfa.FlattenImage(img)
		gfimg := zplgfa.ConvertToZPL(flat, zplgfa.CompressedASCII)
		zplBytes = []byte(gfimg)
	})

	timeMetric := metrics.TimeMetric{
		Key: "pdf2pzl",
		Times: []*metrics.Time{
			timeSpent,
		},
		Avg: timeSpent,
	}

	timeMetric.Normalize()

	report := document.GetReport()
	report.TimeMetrics = append(report.TimeMetrics, timeMetric)

	return &zpl{
		bytes:        zplBytes,
		report:       report,
		documentType: documenttype.ZPL,
	}
}

func (z *zpl) GetBytes() []byte {
	return z.bytes
}

func (z *zpl) GetType() documenttype.DocumentType {
	return z.documentType
}

func (z *zpl) GetBase64() string {
	return base64.StdEncoding.EncodeToString(z.bytes)
}

func (z *zpl) GetReport() *metrics.Report {
	return z.report
}

func (z *zpl) Save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(z.bytes)
	if err != nil {
		return err
	}

	return nil
}

func (z *zpl) To(newType documenttype.DocumentType) (Document, error) {
	if newType == documenttype.ZPL {
		return z, nil
	}

	return nil, errors.New("invalid type conversion")
}
