package instrumentation

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
)

func InitTracer(ctx context.Context,
	exporter trace.SpanExporter,
	resource *resource.Resource) func(ctx context.Context) error {
	traceProvider := trace.NewTracerProvider(
		trace.WithBatcher(
			exporter,
		),
		trace.WithResource(resource),
	)
	otel.SetTracerProvider(traceProvider)

	return traceProvider.Shutdown
}
