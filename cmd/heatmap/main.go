package main

import (
	"github.com/johnfercher/maroto/v2/pkg/props"
	"log"
	"math/rand"

	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
)

func main() {
	cfg := config.NewBuilder().
		WithPageNumber().
		WithDebug(true).
		Build()

	mrt := maroto.New(cfg)
	m := maroto.NewMetricsDecorator(mrt)

	xMax := 620
	yMax := 200
	max := 10
	heat := buildHeat(xMax, yMax, max)

	m.AddRows(text.NewRow(10, "HeatMap"))
	m.AddRows(chart.NewHeatMapRow(100, "map", heat, props.HeatMap{
		TransparentValues: []int{0},
	}))

	/*m.AddRow(100,
		chart.NewHeatMapCol(6, "map", heat, props.HeatMap{
			TransparentValues: []int{0},
			Chart: &props.Chart{
				Axis: true,
			},
		}),
		chart.NewHeatMapCol(6, "map", heat, props.HeatMap{
			TransparentValues: []int{0},
			Chart: &props.Chart{
				Axis: true,
			},
		}),
	)*/

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/heatmap.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.GetReport().Save("docs/assets/text/heatmap.txt")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildHeat(x, y, max int) [][]int {
	var heat [][]int
	for i := 0; i < x; i++ {
		var line []int
		for j := 0; j < y; j++ {
			line = append(line, rand.Intn(max))
		}
		heat = append(heat, line)
	}
	return heat
}
