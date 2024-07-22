package metrics

import "github.com/prometheus/client_golang/prometheus"

var MetricCollectors = []prometheus.Collector{deliveriesGauge}

var (
	deliveriesGauge = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "ozon_deliveries_total",
			Help: "Total number of deliveries",
		},
	)
)

func IncDeliveriesTotal() {
	deliveriesGauge.Inc()
}

func DecDeliveriesTotal() {
	deliveriesGauge.Dec()
}
