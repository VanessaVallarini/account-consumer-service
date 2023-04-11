package metrics

import (
	"sync"

	"github.com/prometheus/client_golang/prometheus"
)

var (
	onceConfig    sync.Once
	metricConfigs = &Metrics{}
)

type Metrics struct {
	KafkaStrategyConsumedMessagesCounter *prometheus.CounterVec
	KafkaStrategyDlqMessagesCounter      *prometheus.CounterVec
	KafkaStrategyLagGauge                *prometheus.GaugeVec
}

func NewMetrics() *Metrics {
	onceConfig.Do(func() {
		metricConfigs.KafkaStrategyConsumedMessagesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "kafka_consumed_messages",
			Help: "Kafka consumed messages counter",
		}, []string{"topic", "consumer_group"})
		prometheus.MustRegister(metricConfigs.KafkaStrategyConsumedMessagesCounter)

		metricConfigs.KafkaStrategyDlqMessagesCounter = prometheus.NewCounterVec(prometheus.CounterOpts{
			Name: "kafka_dlq_messages",
			Help: "Kafka dlq messages counter",
		}, []string{"topic", "consumer_group"})
		prometheus.MustRegister(metricConfigs.KafkaStrategyDlqMessagesCounter)

		metricConfigs.KafkaStrategyLagGauge = prometheus.NewGaugeVec(prometheus.GaugeOpts{
			Name: "kafka_lag_messages",
			Help: "Kafka lag messages gauge",
		}, []string{"topic", "consumer_group"})
		prometheus.MustRegister(metricConfigs.KafkaStrategyLagGauge)
	})
	return metricConfigs
}
