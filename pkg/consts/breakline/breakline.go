// Package breakline contains all break line strategies.
package breakline

// Strategy represents a break line strategy.
type Strategy string

const (
	// EmptyLineStrategy is a break line strategy that uses empty lines.
	EmptyLineStrategy Strategy = "empty_line_strategy"
	// DashStrategy is a break line strategy that uses dashes.
	DashStrategy Strategy = "dash_strategy"
)
