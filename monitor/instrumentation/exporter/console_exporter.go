package exporter

import (
	"context"

	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/sdk/trace"
)

func NewConsoleExporter(ctx context.Context) (trace.SpanExporter, error) {
	return stdouttrace.New()
}
