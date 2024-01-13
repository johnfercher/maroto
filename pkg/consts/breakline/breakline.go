// Package breakline contains all break line strategies.
package breakline

// Strategy represents a break line strategy.
type Strategy string

const (
	// EmptySpaceStrategy is a break line strategy that counts the length of words to create a new line.
	// This strategy only works in languages that use spaces to divide words
	EmptySpaceStrategy Strategy = "empty_space_strategy"
	// DashStrategy is a break line strategy that counts the length for
	// a set of characters with no relation with the meaning of words.
	// This strategy is useful for languages that don't use space between words.
	// To divide the lines, is applied a dash in the end of the line.
	DashStrategy Strategy = "dash_strategy"
)
