package metrics

import (
	"fmt"
	"os"
)

type TimeScale string
type SizeScale string

const (
	Nano     TimeScale = "ns"
	Micro    TimeScale = "μs"
	Milli    TimeScale = "ms"
	Byte     SizeScale = "b"
	KiloByte SizeScale = "Kb"
	MegaByte SizeScale = "Mb"
	GigaByte           = "Gb"
)

type Time struct {
	Value float64
	Scale TimeScale
}

func (t *Time) Normalize() bool {
	if t.Scale == Nano {
		t.Scale = Micro
		t.Value /= 1000.0
		return true
	}

	if t.Scale == Micro {
		t.Scale = Milli
		t.Value /= 1000.0
		return true
	}

	return false
}

func (t *Time) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Scale)
}

type Size struct {
	Value float64
	Scale SizeScale
}

func (t *Size) Normalize() bool {
	if t.Scale == Byte {
		t.Scale = KiloByte
		t.Value /= 1000.0
		return true
	}

	if t.Scale == KiloByte {
		t.Scale = MegaByte
		t.Value /= 1000.0
		return true
	}

	if t.Scale == MegaByte {
		t.Scale = GigaByte
		t.Value /= 1000.0
		return true
	}

	return false
}

func (t *Size) String() string {
	return fmt.Sprintf("%.2f%s", t.Value, t.Scale)
}

type TimeMetric struct {
	Key   string
	Times []*Time
	Avg   *Time
}

func (m *TimeMetric) Normalize() {
	greaterThan1000 := m.hasGreaterThan1000(m.Times)
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

func (m *TimeMetric) hasGreaterThan1000(times []*Time) bool {
	for _, time := range times {
		if time.Value > 1000.0 {
			return true
		}
	}

	return false
}

func (m *TimeMetric) String() string {
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

type SizeMetric struct {
	Key  string
	Size Size
}

func (m *SizeMetric) Normalize() {
	if m.Size.Value < 1000.0 {
		return
	}

	m.Size.Normalize()
	m.Normalize()
}

func (r *SizeMetric) String() string {
	return r.Key + " -> " + r.Size.String()
}

type Report struct {
	TimeMetrics []TimeMetric
	SizeMetric  SizeMetric
}

func (r *Report) Normalize() *Report {
	for _, metric := range r.TimeMetrics {
		metric.Normalize()
	}

	r.SizeMetric.Normalize()

	return r
}

func (r *Report) Save(file string) error {
	var content string
	for _, metric := range r.TimeMetrics {
		content += metric.String() + "\n"
	}
	content += r.SizeMetric.String() + "\n"

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
	for _, metric := range r.TimeMetrics {
		content += metric.String()
	}
	return content
}
