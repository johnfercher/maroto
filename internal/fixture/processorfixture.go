package fixture

import (
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/listmapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/pagemapper"
	"github.com/johnfercher/maroto/v2/pkg/processor/mappers/rowmapper"
)

func MapperRow() *rowmapper.Row {
	return &rowmapper.Row{
		Height: 0,
		Cols:   make([]mappers.Componentmapper, 0),
	}
}

func MapperPage() *pagemapper.Page {
	return &pagemapper.Page{
		SourceKey: "template_page_1",
		Rows:      make([]mappers.Componentmapper, 0),
	}
}

func MapperList() *listmapper.List {
	return &listmapper.List{}
}
