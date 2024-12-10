package main

import (
	"fmt"
	"log"

	"github.com/johnfercher/maroto/v2"

	"github.com/johnfercher/maroto/v2/pkg/components/list"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/props"

	"github.com/johnfercher/maroto/v2/pkg/components/text"
)

var background = &props.Color{
	Red:   200,
	Green: 200,
	Blue:  200,
}

func main() {
	m := GetMaroto()
	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/listv2.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/listv2.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func GetMaroto() core.Maroto {
	mrt := maroto.New()
	m := maroto.NewMetricsDecorator(mrt)

	myList1 := list.New(GetHeader("header 1"), props.List{MinimumRowsBypage: 3}).Add(GetObjects(120)...)
	m.AddRows(myList1.GetRows()...)

	myList := list.New(GetHeader("header 2"), props.List{MinimumRowsBypage: 5}).Add(GetObjects(100)...)
	myList.BuildListWithFixedHeader(m)

	return m
}

func GetObjects(max int) []core.Row {
	var objects []core.Row
	for i := 0; i < max; i++ {
		objects = append(objects, GetContent(i))
	}
	return objects
}

func GetContent(i int) core.Row {
	r := row.New(4).Add(
		text.NewCol(4, fmt.Sprintf("key %d", i)),
		text.NewCol(8, fmt.Sprintf("Value %d", i)),
	)

	if i%2 == 0 {
		r.WithStyle(&props.Cell{
			BackgroundColor: background,
		})
	}

	return r
}

func GetHeader(name string) core.Row {
	return row.New(20).Add(
		text.NewCol(4, "Key "+name, props.Text{Style: fontstyle.Bold}),
		text.NewCol(8, "value "+name, props.Text{Style: fontstyle.Bold}),
	)
}
