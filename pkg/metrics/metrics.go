package metrics

import (
	"fmt"
	"os"
)

type TimeScale string

const (
	Thousand float64   = 1000
	Nano     TimeScale = "ns"
	Micro    TimeScale = "μs"
	Milli    TimeScale = "ms"
)

type Time struct {
	Value float64
	Scale TimeScale
}

func (t *Time) Normalize() bool {
	if t.Scale == Nano {
		t.Scale = Micro
		t.Value /= Thousand
		return true
	}

	if t.Scale == Micro {
		t.Scale = Milli
		t.Value /= Thousand
		return true
	}

	return false
}

func (t *Time) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Scale)
}

type Metric struct {
	Key   string
	Times []*Time
	Avg   *Time
}

func (m *Metric) Normalize() {
	greaterThan1000 := hasGreaterThan1000(m.Times)
	if greaterThan1000 {
		for _, time := range m.Times {
			done := time.Normalize()
			if !done {
				return
			}
		}
	}

	if greaterThan1000 {
		m.Normalize()
	}
}

func hasGreaterThan1000(times []*Time) bool {
	for _, time := range times {
		if time.Value > Thousand {
			return true
		}
	}

	return false
}

func (m *Metric) String() string {
	var content string
	content += m.Key + " -> avg: " + m.Avg.String() + ", executions: ["
	for i, time := range m.Times {
		content += time.String()
		if i < len(m.Times)-1 {
			content += ", "
		}
	}
	content += "]"
	return content
}

type Report []Metric

func (r *Report) Normalize() *Report {
	for _, metric := range *r {
		metric.Normalize()
	}
	return r
}

func (r *Report) Save(file string) error {
	var content string
	for _, metric := range *r {
		content += metric.String() + "\n"
	}

	f, err := os.Create(file)
	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		return err
	}

	return nil
}

func (r *Report) String() string {
	var content string
	for _, metric := range *r {
		content += metric.String()
	}
	return content
}
