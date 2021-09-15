package main

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"log"
	"mysql_exporter/auth"
	"mysql_exporter/collectors"
	"mysql_exporter/config"
	"net/http"
)

func initConfig() *config.ExporterConfig {

	// 配置文件解析
	return &config.ExporterConfig{
		Web: &config.WebConfig{
			Addr: ":9999",
			Auth: &config.AuthConfig{"lhq", "123"},
		},
	}
}

func main() {

	config := initConfig()

	addr := ":9999"
	mysqlAddr := "localhost:3306"
	dsn := "root:root@tcp(localhost:3306)/mysql?charset=utf8mb4&loc=PRC&parseTime=true"

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		logrus.Fatal(err)
	}

	mysqlInfo := prometheus.NewGauge(prometheus.GaugeOpts{
		Name:        "mysql_info",
		Help:        "mysql_info",
		ConstLabels: prometheus.Labels{"addr": mysqlAddr},
	})
	mysqlInfo.Set(1)

	// 定义指标
	// 注册指标
	// 时间触发  业务请求触发  metrics请求触发
	// 可选择的方案：1，3；2不可以，与业务无关
	prometheus.MustRegister(collectors.NewUpCollector(db))
	prometheus.MustRegister(collectors.NewSlowQueriesCollector(db))
	prometheus.MustRegister(collectors.NewTrafficCollector(db))
	prometheus.MustRegister(collectors.NewConnectionController(db))
	prometheus.MustRegister(collectors.NewCommandCollector(db))

	prometheus.MustRegister(mysqlInfo)

	/*prometheus.MustRegister(prometheus.NewGaugeFunc(prometheus.GaugeOpts{
		Name:        "mysql_info",
		Help:        "mysql_info",
		ConstLabels: prometheus.Labels{"addr": mysqlAddr},
	}, func() float64 {
		return  1
	}))*/

	// 注册控制器

	http.Handle("/metrics", auth.BasicAuth(config.Web.Auth, promhttp.Handler()))
	// 启动web服务
	http.ListenAndServe(addr, nil)
}
