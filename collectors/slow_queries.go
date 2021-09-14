package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type SlowQueriesCollector struct {
	*baseCollector
	desc *prometheus.Desc
}

func NewSlowQueriesCollector(db *sql.DB) *SlowQueriesCollector {
	desc := prometheus.NewDesc("mysql_slow_queries_total", "Mysql slow queries total", nil, nil)
	return &SlowQueriesCollector{newBaseController(db), desc}
}

func (c *SlowQueriesCollector) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.desc
}

func (c *SlowQueriesCollector) Collect(metrics chan<- prometheus.Metric) {
	count := c.status("slow_queries")
	metrics <- prometheus.MustNewConstMetric(c.desc, prometheus.CounterValue, count)
}
