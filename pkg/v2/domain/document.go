package domain

import "os"

type Document interface {
	GetBytes() []byte
	GetBase64() string
	Save(file string) error
	GetReport() *Report
}

type document struct {
	bytes  []byte
	report *Report
}

func NewDocument(bytes []byte) Document {
	return &document{}
}

func (r *document) GetBytes() []byte {
	//TODO implement me
	panic("implement me")
}

func (r *document) GetBase64() string {
	//TODO implement me
	panic("implement me")
}

func (r *document) GetReport() *Report {
	//TODO implement me
	panic("implement me")
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
