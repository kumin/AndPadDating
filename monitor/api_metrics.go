package monitor

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/metric"
)

var meter = otel.Meter("bity-service")
var (
	StatusCode = "status_code"
	ApiUrl     = "api_url"
)

func LatencyHistorgram() (metric.Float64Histogram, error) {
	return meter.Float64Histogram(
		"api_response_latency",
		metric.WithDescription("the latency of services"),
		metric.WithUnit("millisecond"),
	)
}
