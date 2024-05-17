package generation

type Mode string

const (
	Sequential          Mode = "sequential"
	Concurrent          Mode = "concurrent"
	SequentialLowMemory Mode = "sequential_low_memory"
)
