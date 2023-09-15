package v2

import (
	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/pkg/v2/domain"
	"github.com/johnfercher/maroto/pkg/v2/metrics"
	"time"
)

type metricsDecorator struct {
	label         string
	addRowTime    []*metrics.Time
	addPageTime   []*metrics.Time
	generateTime  *metrics.Time
	structureTime *metrics.Time
	inner         domain.Maroto
}

func NewMetricsDecorator(inner domain.Maroto) domain.Maroto {
	return &metricsDecorator{
		inner: inner,
	}
}

func (d *metricsDecorator) Generate() (domain.Document, error) {
	var document domain.Document
	var err error

	timeSpent := d.getTimeSpent(func() {
		document, err = d.inner.Generate()
	})
	d.generateTime = timeSpent

	report := d.buildMetrics().Normalize()
	return domain.NewDocument(document.GetBytes(), report), err
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

func (d *metricsDecorator) getTimeSpent(closure func()) *metrics.Time {
	start := time.Now()
	closure()
	return &metrics.Time{
		Value: float64(time.Now().Sub(start).Nanoseconds()),
		Scale: metrics.Nano,
	}
}

func (d *metricsDecorator) buildMetrics() *metrics.Report {
	var report metrics.Report

	if d.structureTime != nil {
		report = append(report, metrics.Metric{
			Key:   "get_tree_structure",
			Times: []*metrics.Time{d.structureTime},
			Avg:   d.structureTime,
		})
	}

	if d.generateTime != nil {
		report = append(report, metrics.Metric{
			Key:   "generate",
			Times: []*metrics.Time{d.generateTime},
			Avg:   d.generateTime,
		})
	}

	if len(d.addPageTime) > 0 {
		report = append(report, metrics.Metric{
			Key:   "add_page",
			Times: d.addPageTime,
			Avg:   d.getAVG(d.addPageTime),
		})
	}

	if len(d.addRowTime) > 0 {
		report = append(report, metrics.Metric{
			Key:   "add_row",
			Times: d.addRowTime,
			Avg:   d.getAVG(d.addRowTime),
		})
	}

	return &report
}

func (d *metricsDecorator) getAVG(times []*metrics.Time) *metrics.Time {
	if len(times) == 0 {
		return nil
	}

	var sum float64
	for _, time := range times {
		sum += time.Value
	}

	return &metrics.Time{
		Value: sum / float64(len(times)),
		Scale: times[0].Scale,
	}
}
