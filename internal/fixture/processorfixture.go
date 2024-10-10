package fixture

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/rowmapper"
)

func MapperRow() *rowmapper.Row {
	return &rowmapper.Row{
		Test: "1",
	}
}

func MapperPage() *pagemapper.Page {
	return &pagemapper.Page{
		Teste: "1",
	}
}

func MapperList() *listmapper.List {
	return &listmapper.List{}
}
