package collectors

import (
	"database/sql"
	"github.com/prometheus/client_golang/prometheus"
)

type ConnectionController struct {
	*baseCollector
	maxConnectionsDesc   *prometheus.Desc
	threadsConnectedDesc *prometheus.Desc
}

func NewConnectionController(db *sql.DB) *ConnectionController {
	maxConnectionsDesc := prometheus.NewDesc("mysql_global_variables_max_connections", "mysql global variables max connections", nil, nil)
	threadsConnectedDesc := prometheus.NewDesc("mysql_global_status_threads_connected", "mysql global status threads connected", nil, nil)

	return &ConnectionController{
		newBaseController(db),
		maxConnectionsDesc,
		threadsConnectedDesc,
	}
}

func (c *ConnectionController) Describe(desc chan<- *prometheus.Desc) {
	desc <- c.threadsConnectedDesc
	desc <- c.maxConnectionsDesc
}

func (c *ConnectionController) Collect(metrics chan<- prometheus.Metric) {
	metrics <- prometheus.MustNewConstMetric(c.maxConnectionsDesc, prometheus.GaugeValue, c.variables("max_connections"))
	metrics <- prometheus.MustNewConstMetric(c.threadsConnectedDesc, prometheus.GaugeValue, c.variables("threads_connected"))
}
