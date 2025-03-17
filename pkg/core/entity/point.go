package entity

import "github.com/johnfercher/maroto/v2/pkg/props"

type Point struct {
	X float64
	Y float64
}

func NewPoint(x, y float64) Point {
	return Point{
		X: x,
		Y: y,
	}
}

type TimeSeries struct {
	Values []Point
	Color  props.Color
}

func NewTimeSeries(color props.Color, values ...Point) TimeSeries {
	return TimeSeries{
		Values: values,
		Color:  color,
	}
}
