package domain

import "fmt"

type TimeScale string

const (
	Nano  TimeScale = "ns"
	Micro TimeScale = "μs"
	Milli TimeScale = "ms"
)

type Time struct {
	Value float64
	Scale TimeScale
}

func (t *Time) Normalize() {
	if t.Scale == Nano {
		t.Scale = Micro
		t.Value = t.Value / 1000.0
		return
	}

	if t.Scale == Micro {
		t.Scale = Milli
		t.Value = t.Value / 1000.0
		return
	}
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
			time.Normalize()
		}
	}

	if greaterThan1000 {
		m.Normalize()
	}
}

func hasGreaterThan1000(times []*Time) bool {
	for _, time := range times {
		if time.Value > 1000 {
			return true
		}
	}

	return false
}

func (m *Metric) String() string {
	var content string
	content += m.Key + "\n"
	content += "executions: ["
	for i, time := range m.Times {
		content += time.String()
		if i < len(m.Times)-1 {
			content += ", "
		}
	}
	content += "]\n"
	content += "avg: " + m.Avg.String()
	return content
}

type Report []Metric

func (r *Report) Normalize() *Report {
	for _, metric := range *r {
		metric.Normalize()
	}
	return r
}

func (r *Report) Print() {
	for _, metric := range *r {
		fmt.Println(metric.String())
	}
}
