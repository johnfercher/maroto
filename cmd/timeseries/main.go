package main

import (
	"github.com/johnfercher/maroto/v2"
	"github.com/johnfercher/maroto/v2/pkg/components/chart"
	"github.com/johnfercher/maroto/v2/pkg/components/row"
	"github.com/johnfercher/maroto/v2/pkg/config"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontfamily"
	"github.com/johnfercher/maroto/v2/pkg/consts/fontstyle"
	"github.com/johnfercher/maroto/v2/pkg/consts/pagesize"
	"github.com/johnfercher/maroto/v2/pkg/core/entity"
	"github.com/johnfercher/maroto/v2/pkg/props"
	"log"
	"math"
)

func main() {
	width := 30.0
	step := 0.1
	timeSeriesList := SinCos(width, step)
	timeSeriesList = append(timeSeriesList, Log(width, step)...)

	prop := props.Chart{
		Scale: props.ChartScale{
			X: []float64{0, 10, 20, 30},
			Y: []float64{0, 1, 2, 3},
			Font: props.Font{
				Family: fontfamily.Arial,
				Style:  fontstyle.Normal,
				Size:   7,
			},
		},
		Title: props.ChartTitle{
			Text: "Time Series",
			Font: props.Font{
				Family: fontfamily.Arial,
				Style:  fontstyle.Normal,
				Size:   9,
			},
		},
	}

	cfg := config.NewBuilder().
		WithDebug(true).
		WithPageSize(pagesize.A4).
		Build()

	m := maroto.New(cfg)

	m.AddRows(
		row.New(100).Add(
			chart.NewTimeSeriesCol(12, timeSeriesList, prop),
		),
	)

	document, err := m.Generate()
	if err != nil {
		log.Fatal(err.Error())
	}

	err = document.Save("docs/assets/pdf/timeseries.pdf")
	if err != nil {
		log.Fatal(err.Error())
	}
}

func SinCos(width float64, step float64) []entity.TimeSeries {
	timeSeries := []entity.TimeSeries{}

	var maxSin entity.Point
	sin := []entity.Point{}
	for i := 0.0; i < width; i += step {
		y := math.Sin(i) + 1 + (i / 25)
		point := entity.NewPoint(i, y)
		if y > maxSin.Y {
			maxSin = point
		}
		sin = append(sin, point)
	}
	timeSeries = append(timeSeries, entity.NewTimeSeries(props.RedColor, sin, entity.NewLabel("Max", maxSin)))

	var maxCos entity.Point
	cos := []entity.Point{}
	for i := 0.0; i < width; i += step {
		y := math.Cos(i) + 1 + (i / 30)
		point := entity.NewPoint(i, y)
		if y > maxCos.Y {
			maxCos = point
		}
		cos = append(cos, point)
	}
	timeSeries = append(timeSeries, entity.NewTimeSeries(props.BlueColor, cos, entity.NewLabel("Max", maxCos)))

	return timeSeries
}

func Log(width float64, step float64) []entity.TimeSeries {
	timeSeries := []entity.TimeSeries{}

	var logMax entity.Point
	sin := []entity.Point{}
	for i := 0.0; i < width; i += step {
		v := math.Log(i)
		point := entity.NewPoint(i, v)
		if v > logMax.Y {
			logMax = point
		}
		if v > 0 {
			sin = append(sin, point)
		}
	}
	timeSeries = append(timeSeries, entity.NewTimeSeries(props.Color{
		Red:  150,
		Blue: 150,
	}, sin, entity.NewLabel("Max", logMax)))

	return timeSeries
}
