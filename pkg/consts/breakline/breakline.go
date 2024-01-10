// Package contains all breakline strategies.
package breakline

type Strategy string

const (
	EmptyLineStrategy Strategy = "empty_line_strategy"
	DashStrategy      Strategy = "dash_strategy"
)
