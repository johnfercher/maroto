package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"time"
)

type documentMetrics struct {
	label         string
	addRowTime    []*domain.Time
	addPageTime   []*domain.Time
	generateTime  *domain.Time
	structureTime *domain.Time
	inner         domain.Maroto
}

func NewMarotoMetrified(inner domain.Maroto) domain.MarotoMetrified {
	return &documentMetrics{
		inner: inner,
	}
}

func (d *documentMetrics) Generate() error {
	var err error

	timeSpent := d.getTimeSpent(func() {
		err = d.inner.Generate()
	})
	d.generateTime = timeSpent

	return err
}

func (d *documentMetrics) ForceAddPage(pages ...domain.Page) {
	timeSpent := d.getTimeSpent(func() {
		d.inner.ForceAddPage(pages...)
	})

	d.addPageTime = append(d.addPageTime, timeSpent)
}

func (d *documentMetrics) Add(rows ...domain.Row) {
	timeSpent := d.getTimeSpent(func() {
		d.inner.Add(rows...)
	})

	d.addRowTime = append(d.addRowTime, timeSpent)
}

func (d *documentMetrics) GetStructure() *tree.Node[domain.Structure] {
	var tree *tree.Node[domain.Structure]

	timeSpent := d.getTimeSpent(func() {
		tree = d.inner.GetStructure()
	})
	d.structureTime = timeSpent

	return tree
}

func (d *documentMetrics) GenerateWithReport() (*domain.Report, error) {
	var err error

	timeSpent := d.getTimeSpent(func() {
		err = d.inner.Generate()
	})
	d.generateTime = timeSpent

	return d.buildMetrics().Normalize(), err
}

func (d *documentMetrics) getTimeSpent(closure func()) *domain.Time {
	start := time.Now()
	closure()
	return &domain.Time{
		Value: float64(time.Now().Sub(start).Nanoseconds()),
		Scale: domain.Nano,
	}
}

func (d *documentMetrics) buildMetrics() *domain.Report {
	var report domain.Report

	if d.structureTime != nil {
		report = append(report, domain.Metric{
			Key:   "get_tree_structure",
			Times: []*domain.Time{d.structureTime},
			Avg:   d.structureTime,
		})
	}

	if d.generateTime != nil {
		report = append(report, domain.Metric{
			Key:   "generate",
			Times: []*domain.Time{d.generateTime},
			Avg:   d.generateTime,
		})
	}

	if len(d.addPageTime) > 0 {
		report = append(report, domain.Metric{
			Key:   "add_page",
			Times: d.addPageTime,
			Avg:   d.getAVG(d.addPageTime),
		})
	}

	if len(d.addRowTime) > 0 {
		report = append(report, domain.Metric{
			Key:   "add_row",
			Times: d.addRowTime,
			Avg:   d.getAVG(d.addRowTime),
		})
	}

	return &report
}

func (d *documentMetrics) getAVG(times []*domain.Time) *domain.Time {
	if len(times) == 0 {
		return nil
	}

	var sum float64
	for _, time := range times {
		sum += time.Value
	}

	return &domain.Time{
		Value: sum / float64(len(times)),
		Scale: times[0].Scale,
	}
}
