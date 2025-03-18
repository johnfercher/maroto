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
	prop := props.Chart{
		XLabels: []float64{0, 10, 20, 30},
		YLabels: []float64{0, 1, 2},
		Font: &props.Font{
			Family: fontfamily.Arial,
			Style:  fontstyle.Normal,
			Size:   7,
		},
	}

	cfg := config.NewBuilder().
		//WithDebug(true).
		WithPageSize(pagesize.A4).
		Build()

	m := maroto.New(cfg)

	m.AddRows(
		row.New(100).Add(
			chart.NewTimeSeriesCol(12, timeSeriesList, prop),
		),
		/*row.New(100).Add(
			chart.NewTimeSeriesCol(12, pointsMatrix),
		),*/
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

	sin := []entity.Point{}
	for i := 0.0; i < width; i += step {
		sin = append(sin, entity.NewPoint(i, math.Sin(i)+1))
	}
	timeSeries = append(timeSeries, entity.NewTimeSeries(props.RedColor, sin...))

	cos := []entity.Point{}
	for i := 0.0; i < width; i += step {
		cos = append(cos, entity.NewPoint(i, math.Cos(i)+1))
	}
	timeSeries = append(timeSeries, entity.NewTimeSeries(props.BlueColor, cos...))

	return timeSeries
}
