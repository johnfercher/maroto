package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"time"
)

type metricsDecorator struct {
	label         string
	addRowTime    []*domain.Time
	addPageTime   []*domain.Time
	generateTime  *domain.Time
	structureTime *domain.Time
	inner         domain.Maroto
}

func NewMetricsDecorator(inner domain.Maroto) domain.Maroto {
	return &metricsDecorator{
		inner: inner,
	}
}

func (d *metricsDecorator) Generate() (*domain.Document, error) {
	var result *domain.Document
	var err error

	timeSpent := d.getTimeSpent(func() {
		result, err = d.inner.Generate()
	})
	d.generateTime = timeSpent

	result.Report = d.buildMetrics().Normalize()

	return result, err
}

func (d *metricsDecorator) ForceAddPage(pages ...domain.Page) {
	timeSpent := d.getTimeSpent(func() {
		d.inner.ForceAddPage(pages...)
	})

	d.addPageTime = append(d.addPageTime, timeSpent)
}

func (d *metricsDecorator) Add(rows ...domain.Row) {
	timeSpent := d.getTimeSpent(func() {
		d.inner.Add(rows...)
	})

	d.addRowTime = append(d.addRowTime, timeSpent)
}

func (d *metricsDecorator) GetStructure() *tree.Node[domain.Structure] {
	var tree *tree.Node[domain.Structure]

	timeSpent := d.getTimeSpent(func() {
		tree = d.inner.GetStructure()
	})
	d.structureTime = timeSpent

	return tree
}

func (d *metricsDecorator) getTimeSpent(closure func()) *domain.Time {
	start := time.Now()
	closure()
	return &domain.Time{
		Value: float64(time.Now().Sub(start).Nanoseconds()),
		Scale: domain.Nano,
	}
}

func (d *metricsDecorator) buildMetrics() *domain.Report {
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

func (d *metricsDecorator) getAVG(times []*domain.Time) *domain.Time {
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
