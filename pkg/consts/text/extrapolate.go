package text

type ExtrapolateStrategy string

const (
	ExtrapolateStrategyNone    ExtrapolateStrategy = "none"
	ExtrapolateStrategyWords   ExtrapolateStrategy = "words"
	ExtrapolateStrategySymbols ExtrapolateStrategy = "symbols"
)
