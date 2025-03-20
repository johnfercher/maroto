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

type Label struct {
	Value string
	Point Point
}

func NewLabel(value string, point Point) Label {
	return Label{
		Value: value,
		Point: point,
	}
}

type TimeSeries struct {
	Points []Point
	Labels []Label
	Color  props.Color
}

func NewTimeSeries(color props.Color, points []Point, labels ...Label) TimeSeries {
	return TimeSeries{
		Points: points,
		Color:  color,
		Labels: labels,
	}
}
