package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/components/text"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"log"
	"math"
)

func main() {
	//heat := buildHeat(50, 50)
	//circleHeat := buildCircleHeat(50, 50)

	cfg := config.NewBuilder().
		//WithDebug(true).
		WithPageSize(pagesize.A4).
		Build()

	m := maroto.New(cfg)

	m.AddRows(text.NewRow(5, "HeatMap"))

	m.AddRows(
		row.New(70).Add(
			chart.NewHeatMapCol(12, "Efficiency", buildCircleHeat(100, 100), props.HeatMap{
				Chart: props.Chart{
					Scale: props.ChartScale{
						X: []float64{0, 25, 50, 75, 100},
						Y: []float64{0, 25, 50, 75, 100},
					},
					Title: props.ChartTitle{
						Text: "HeatMap",
					},
				},
			}),
		),
	)

	/*m.AddRows(
		row.New(70).Add(
			chart.NewHeatMapCol(5, "Efficiency", buildHeat(100, 100)),
			col.New(2),
			chart.NewHeatMapCol(5, "Efficiency", buildCircleHeat(100, 100), props.HeatMap{
				Chart: props.Chart{
					Scale: props.ChartScale{
						X: []float64{0, 25, 50, 75, 100},
						Y: []float64{0, 25, 50, 75, 100},
					},
					Title: props.ChartTitle{
						Text: "HeatMap",
					},
				},
			}),
		),
	)

	m.AddRows(row.New(10))

	m.AddRows(
		row.New(35).Add(
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Purple,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Blue,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Cyan,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Green,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Yellow,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Orange,
				},
			}),
		),
	)

	m.AddRows(row.New(5))

	m.AddRows(
		row.New(35).Add(
			chart.NewHeatMapCol(2, "Efficiency", heat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Pink,
					End:   hsv.Red,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", circleHeat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Blue,
					End:   hsv.Cyan,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", circleHeat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Blue,
					End:   hsv.Green,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", circleHeat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Blue,
					End:   hsv.Yellow,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", circleHeat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Blue,
					End:   hsv.Orange,
				},
			}),
			chart.NewHeatMapCol(2, "Efficiency", circleHeat, props.HeatMap{
				HSVScale: props.HSVScale{
					Begin: hsv.Blue,
					End:   hsv.Red,
				},
			}),
		),
	)*/

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
			wp := w
			line = append(line, wp)
		}
		heat = append(heat, line)
	}
	return heat
}

func buildCircleHeat(x, y int) [][]int {
	var heat [][]int

	for i := 0; i < x; i++ {
		var line []int
		for j := 0; j < y; j++ {
			iRad := float64(i) * math.Pi / 180.0
			jRad := float64(j) * math.Pi / 180.0
			x := math.Cos(iRad) * 360
			y := math.Sin(jRad) * 360
			line = append(line, int(y)+int(x))
		}
		heat = append(heat, line)
	}

	return heat
}
