package domain

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/internal"
	"os"
)

type Document struct {
	Bytes  []byte
	Report *Report
}

func (r *Document) Save(file string) error {
	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.Write(r.Bytes)
	if err != nil {
		return err
	}

	return nil
}

type Maroto interface {
	Generate() (*Document, error)
	ForceAddPage(pages ...Page)
	Add(rows ...Row)
	GetStructure() *tree.Node[Structure]
}

type Node interface {
	Render(fpdf Provider, ctx internal.Cell)
	GetStructure() *tree.Node[Structure]
}

type Page interface {
	Node
	Add(rows ...Row) Page
	GetNumber() int
	SetNumber(number int)
}

type Row interface {
	Node
	Add(cols ...Col) Row
	GetHeight() float64
}

type Col interface {
	Node
	Add(nodes ...Node) Col
	AddInner(rows ...Row) Col
	GetSize() int
}

type Structure struct {
	Type  string
	Value string
	Props map[string]string
}
