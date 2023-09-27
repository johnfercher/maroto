package core

import (
	"encoding/base64"
	"errors"
	"os"

	"github.com/johnfercher/maroto/v2/pkg/consts/documenttype"

	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

type pdf struct {
	bytes        []byte
	report       *metrics.Report
	documentType documenttype.DocumentType
}

func NewPDF(bytes []byte, report *metrics.Report) Document {
	return &pdf{
		bytes:        bytes,
		report:       report,
		documentType: documenttype.PDF,
	}
}

func (p *pdf) GetBytes() []byte {
	return p.bytes
}

func (p *pdf) GetType() documenttype.DocumentType {
	return p.documentType
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

func (p *pdf) To(newType documenttype.DocumentType) (Document, error) {
	if newType == documenttype.PDF {
		return p, nil
	}

	if newType == documenttype.ZPL {
		return NewZPLFromPDF(p), nil
	}

	return nil, errors.New("invalid type conversion")
}
