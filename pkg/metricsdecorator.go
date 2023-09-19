package pkg

import (
	"time"

	"github.com/johnfercher/go-tree/tree"
	"github.com/johnfercher/maroto/v2/pkg/core"
	"github.com/johnfercher/maroto/v2/pkg/metrics"
)

type metricsDecorator struct {
	addRowTime    []*metrics.Time
	addColTime    []*metrics.Time
	addPageTime   []*metrics.Time
	headerTime    *metrics.Time
	footerTime    *metrics.Time
	generateTime  *metrics.Time
	structureTime *metrics.Time
	inner         core.Maroto
}

func NewMetricsDecorator(inner core.Maroto) core.Maroto {
	return &metricsDecorator{
		inner: inner,
	}
}

func (m *metricsDecorator) Generate() (core.Document, error) {
	var document core.Document
	var err error

	timeSpent := m.getTimeSpent(func() {
		document, err = m.inner.Generate()
	})
	m.generateTime = timeSpent

	report := m.buildMetrics().Normalize()
	if err != nil {
		return nil, err
	}

	return core.NewDocument(document.GetBytes(), report), nil
}

func (m *metricsDecorator) ForceAddPage(pages ...core.Page) {
	timeSpent := m.getTimeSpent(func() {
		m.inner.ForceAddPage(pages...)
	})

	m.addPageTime = append(m.addPageTime, timeSpent)
}

func (m *metricsDecorator) AddRows(rows ...core.Row) {
	timeSpent := m.getTimeSpent(func() {
		m.inner.AddRows(rows...)
	})

	m.addRowTime = append(m.addRowTime, timeSpent)
}

func (m *metricsDecorator) AddRow(rowHeight float64, cols ...core.Col) core.Row {
	var r core.Row
	timeSpent := m.getTimeSpent(func() {
		r = m.inner.AddRow(rowHeight, cols...)
	})

	m.addColTime = append(m.addColTime, timeSpent)
	return r
}

func (m *metricsDecorator) RegisterHeader(rows ...core.Row) error {
	var err error
	timeSpent := m.getTimeSpent(func() {
		err = m.inner.RegisterHeader(rows...)
	})
	m.headerTime = timeSpent
	return err
}

func (m *metricsDecorator) RegisterFooter(rows ...core.Row) error {
	var err error
	timeSpent := m.getTimeSpent(func() {
		err = m.inner.RegisterFooter(rows...)
	})
	m.footerTime = timeSpent
	return err
}

func (m *metricsDecorator) GetStructure() *tree.Node[core.Structure] {
	var tree *tree.Node[core.Structure]

	timeSpent := m.getTimeSpent(func() {
		tree = m.inner.GetStructure()
	})
	m.structureTime = timeSpent

	return tree
}

func (m *metricsDecorator) getTimeSpent(closure func()) *metrics.Time {
	start := time.Now()
	closure()
	return &metrics.Time{
		Value: float64(time.Since(start).Nanoseconds()),
		Scale: metrics.Nano,
	}
}

func (m *metricsDecorator) buildMetrics() *metrics.Report {
	var report metrics.Report

	if m.structureTime != nil {
		report = append(report, metrics.Metric{
			Key:   "get_tree_structure",
			Times: []*metrics.Time{m.structureTime},
			Avg:   m.structureTime,
		})
	}

	if m.generateTime != nil {
		report = append(report, metrics.Metric{
			Key:   "generate",
			Times: []*metrics.Time{m.generateTime},
			Avg:   m.generateTime,
		})
	}

	if m.headerTime != nil {
		report = append(report, metrics.Metric{
			Key:   "header",
			Times: []*metrics.Time{m.headerTime},
			Avg:   m.headerTime,
		})
	}

	if m.footerTime != nil {
		report = append(report, metrics.Metric{
			Key:   "footer",
			Times: []*metrics.Time{m.footerTime},
			Avg:   m.footerTime,
		})
	}

	if len(m.addPageTime) > 0 {
		report = append(report, metrics.Metric{
			Key:   "add_page",
			Times: m.addPageTime,
			Avg:   m.getAVG(m.addPageTime),
		})
	}

	if len(m.addRowTime) > 0 {
		report = append(report, metrics.Metric{
			Key:   "add_row",
			Times: m.addRowTime,
			Avg:   m.getAVG(m.addRowTime),
		})
	}

	if len(m.addColTime) > 0 {
		report = append(report, metrics.Metric{
			Key:   "add_cols",
			Times: m.addColTime,
			Avg:   m.getAVG(m.addColTime),
		})
	}

	return &report
}

func (m *metricsDecorator) getAVG(times []*metrics.Time) *metrics.Time {
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
