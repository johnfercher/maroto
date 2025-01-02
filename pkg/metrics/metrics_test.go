package metrics_test

import (
	"testing"

	"github.com/johnfercher/maroto/v2/pkg/metrics"
	"github.com/stretchr/testify/assert"
)

func TestTime_Normalize(t *testing.T) {
	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 3000,
			Scale: metrics.Nano,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 3.0, time.Value)
		assert.Equal(t, metrics.Micro, time.Scale)
	})
	t.Run("when scale is nano, should divide by 1000 and change to micro", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Micro,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 2.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})
	t.Run("when scale is milli, should return false", func(t *testing.T) {
		// Arrange
		time := metrics.Time{
			Value: 2000,
			Scale: metrics.Milli,
		}

		// Act
		ok := time.Normalize()

		// Assert
		assert.False(t, ok)
		assert.Equal(t, 2000.0, time.Value)
		assert.Equal(t, metrics.Milli, time.Scale)
	})
}

func TestTime_String(t *testing.T) {
	// Arrange
	time := metrics.Time{
		Value: 2000,
		Scale: metrics.Milli,
	}

	// Act
	s := time.String()

	// Assert
	assert.Equal(t, "2000.00ms", s)
}

func TestSize_Normalize(t *testing.T) {
	t.Run("when scale is byte, should divide by 1000 and change to kilo", func(t *testing.T) {
		// Arrange
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.Byte,
		}

		// Act
		ok := size.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.KiloByte, size.Scale)
	})
	t.Run("when scale is kilo, should divide by 1000 and change to mega", func(t *testing.T) {
		// Arrange
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.KiloByte,
		}

		// Act
		ok := size.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.MegaByte, size.Scale)
	})
	t.Run("when scale is mega, should divide by 1000 and change to giga", func(t *testing.T) {
		// Arrange
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.MegaByte,
		}

		// Act
		ok := size.Normalize()

		// Assert
		assert.True(t, ok)
		assert.Equal(t, 3.0, size.Value)
		assert.Equal(t, metrics.GigaByte, size.Scale)
	})
	t.Run("when scale is giga, should return false", func(t *testing.T) {
		// Arrange
		size := metrics.Size{
			Value: 3000,
			Scale: metrics.GigaByte,
		}

		// Act
		ok := size.Normalize()

		// Assert
		assert.False(t, ok)
		assert.Equal(t, 3000.0, size.Value)
		assert.Equal(t, metrics.GigaByte, size.Scale)
	})
}

func TestSize_String(t *testing.T) {
	// Arrange
	size := metrics.Size{
		Value: 2000,
		Scale: metrics.KiloByte,
	}

	// Act
	s := size.String()

	// Assert
	assert.Equal(t, "2000.00Kb", s)
}

func TestTimeMetric_Normalize(t *testing.T) {
	// Arrange
	arr := []*metrics.Time{
		{
			Value: 10000,
			Scale: metrics.Nano,
		},
		{
			Value: 20000,
			Scale: metrics.Nano,
		},
		{
			Value: 30000,
			Scale: metrics.Nano,
		},
	}

	sum := 0.0
	for _, t := range arr {
		sum += t.Value
	}
	avg := sum / float64(len(arr))

	timeMetric := &metrics.TimeMetric{
		Key:   "keyMetric",
		Times: arr,
		Avg: &metrics.Time{
			Value: avg,
			Scale: metrics.Nano,
		},
	}

	// Act
	timeMetric.Normalize()

	// Assert
	assert.Equal(t, 10.0, timeMetric.Times[0].Value)
	assert.Equal(t, metrics.Micro, timeMetric.Times[0].Scale)
	assert.Equal(t, 20.0, timeMetric.Times[1].Value)
	assert.Equal(t, metrics.Micro, timeMetric.Times[1].Scale)
	assert.Equal(t, 30.0, timeMetric.Times[2].Value)
	assert.Equal(t, metrics.Micro, timeMetric.Times[2].Scale)
}

func TestTimeMetric_String(t *testing.T) {
	// Arrange
	arr := []*metrics.Time{
		{
			Value: 10000,
			Scale: metrics.Nano,
		},
		{
			Value: 20000,
			Scale: metrics.Nano,
		},
		{
			Value: 30000,
			Scale: metrics.Nano,
		},
	}

	sum := 0.0
	for _, t := range arr {
		sum += t.Value
	}
	avg := sum / float64(len(arr))

	timeMetric := &metrics.TimeMetric{
		Key:   "keyMetric",
		Times: arr,
		Avg: &metrics.Time{
			Value: avg,
			Scale: metrics.Nano,
		},
	}

	// Act
	s := timeMetric.String()

	// Assert
	assert.Equal(t, "keyMetric -> avg: 20000.00ns, executions: [10000.00ns, 20000.00ns, 30000.00ns]", s)
}

func TestSizeMetric_Normalize(t *testing.T) {
	// Arrange
	sizeMetric := &metrics.SizeMetric{
		Key: "keyMetric",
		Size: metrics.Size{
			Value: 2000,
			Scale: metrics.Byte,
		},
	}

	// Act
	sizeMetric.Normalize()

	// Assert
	assert.Equal(t, "keyMetric", sizeMetric.Key)
	assert.Equal(t, 2.0, sizeMetric.Size.Value)
	assert.Equal(t, metrics.KiloByte, sizeMetric.Size.Scale)
}

func TestSizeMetric_String(t *testing.T) {
	// Arrange
	sizeMetric := &metrics.SizeMetric{
		Key: "keyMetric",
		Size: metrics.Size{
			Value: 2000,
			Scale: metrics.Byte,
		},
	}

	// Act
	s := sizeMetric.String()

	// Assert
	assert.Equal(t, "keyMetric -> 2000.00b", s)
}

func TestReport_String(t *testing.T) {
	t.Run("when formating, then do it correctly", func(t *testing.T) {
		// Arrange
		report := &metrics.Report{
			TimeMetrics: []metrics.TimeMetric{
				{
					Key: "keyMetric",
					Times: []*metrics.Time{
						{
							Value: 60,
							Scale: metrics.Nano,
						},
						{
							Value: 100,
							Scale: metrics.Milli,
						},
					},
					Avg: &metrics.Time{
						Value: 100,
						Scale: metrics.Milli,
					},
				},
			},
		}

		// Act
		s := report.String()

		// Assert
		assert.Equal(t, "keyMetric -> avg: 100.00ms, executions: [60.00ns, 100.00ms]", s)
	})
}

func TestReport_Normalize(t *testing.T) {
	t.Run("when normalize, do it correctly", func(t *testing.T) {
		// Arrange
		report := &metrics.Report{
			TimeMetrics: []metrics.TimeMetric{
				{
					Key: "keyMetric",
					Times: []*metrics.Time{
						{
							Value: 6000,
							Scale: metrics.Nano,
						},
						{
							Value: 2000,
							Scale: metrics.Micro,
						},
					},
					Avg: &metrics.Time{
						Value: 100,
						Scale: metrics.Milli,
					},
				},
			},
			SizeMetric: metrics.SizeMetric{
				Key: "keyMetric",
				Size: metrics.Size{
					Value: 2000,
					Scale: metrics.Byte,
				},
			},
		}

		// Act
		normalized := report.Normalize()

		// Assert
		assert.Equal(t, 2.0, normalized.SizeMetric.Size.Value)
		assert.Equal(t, metrics.KiloByte, normalized.SizeMetric.Size.Scale)
		assert.Equal(t, 6.0, normalized.TimeMetrics[0].Times[0].Value)
		assert.Equal(t, metrics.Micro, normalized.TimeMetrics[0].Times[0].Scale)
		assert.Equal(t, 2.0, normalized.TimeMetrics[0].Times[1].Value)
		assert.Equal(t, metrics.Milli, normalized.TimeMetrics[0].Times[1].Scale)
	})
}

func TestReport_Save(t *testing.T) {
	t.Run("when saving, then do it correctly", func(t *testing.T) {
		// Arrange
		report := &metrics.Report{
			TimeMetrics: []metrics.TimeMetric{
				{
					Key: "keyMetric",
					Times: []*metrics.Time{
						{
							Value: 6000,
							Scale: metrics.Nano,
						},
						{
							Value: 2000,
							Scale: metrics.Micro,
						},
					},
					Avg: &metrics.Time{
						Value: 100,
						Scale: metrics.Milli,
					},
				},
			},
			SizeMetric: metrics.SizeMetric{
				Key: "keyMetric",
				Size: metrics.Size{
					Value: 2000,
					Scale: metrics.Byte,
				},
			},
		}

		// Act
		err := report.Save("report-test-demo.txt")

		// Assert
		assert.NoError(t, err)
	})
}
