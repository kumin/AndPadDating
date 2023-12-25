package exporter

import (
	"go.opentelemetry.io/otel/exporters/prometheus"
	"go.opentelemetry.io/otel/sdk/metric"
)

func NewPromExporter() (metric.Reader, error) {
	return prometheus.New()
}
