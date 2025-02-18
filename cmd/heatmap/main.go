package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"log"
)

func main() {
	xMax := 620
	yMax := 200
	heat := buildHeat(xMax, yMax)

	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageSize(pagesize.A4).
		Build()

	m := maroto.New(cfg)

	m.AddRows(
		row.New(200).Add(
			chart.NewHeatMapCol(12, "Efficiency", heat, props.HeatMap{
				TransparentValues: []int{0},
				InvertScale:       false,
				HalfColor:         false,
			}),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/heatmap.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func buildHeat(x, y int) [][]int {
	var heat [][]int
	for i := 0; i < x; i++ {
		var line []int
		for j := 0; j < y; j++ {
			w := i + j
			wp := float64(w) / 100
			line = append(line, int(wp))
		}
		heat = append(heat, line)
	}
	return heat
}
