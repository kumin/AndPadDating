package instrumentation

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
)

func InitMeter(
	ctx context.Context,
	exporter metric.Reader,
	resource *resource.Resource,
) func(ctx context.Context) error {
	meterProvider := metric.NewMeterProvider(
		metric.WithReader(exporter),
		metric.WithResource(resource),
	)
	otel.SetMeterProvider(meterProvider)

	return meterProvider.Shutdown
}
