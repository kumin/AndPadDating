package instrumentation

import (
	"context"
	"errors"

	"github.com/kumin/BityDating/monitor/instrumentation/exporter"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.21.0"
)

const (
	ServiceName    = "bity_dating"
	ServiceVersion = "1.0.0"
)

func SetupInstrument(ctx context.Context) (shutdown func(ctx context.Context) error, err error) {
	var shutdownFuncs []func(ctx context.Context) error
	shutdown = func(ctx context.Context) error {
		var err error
		for _, fn := range shutdownFuncs {
			err = errors.Join(err, fn(ctx))
		}
		shutdownFuncs = nil
		return err
	}
	handleErr := func(inErr error) {
		err = errors.Join(inErr, shutdown(ctx))
	}
	resource, err := newResource(ServiceName, ServiceVersion)
	if err != nil {
		handleErr(err)
		return
	}
	prop := newPropagator()
	otel.SetTextMapPropagator(prop)
	jaegerExporter, err := exporter.NewJaegerExporter(ctx)
	//consoleExporter, err := exporter.NewConsoleExporter(ctx)
	if err != nil {
		handleErr(err)
		return
	}
	promExporter, err := exporter.NewPromExporter()
	if err != nil {
		handleErr(err)
		return
	}
	shutdownFuncs = append(shutdownFuncs,
		InitTracer(ctx, jaegerExporter, resource),
		InitMeter(ctx, promExporter, resource))

	return
}

func newResource(serviceName, serviceVersion string) (*resource.Resource, error) {
	return resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName(serviceName),
			semconv.ServiceVersion(serviceVersion),
		),
	)
}

func newPropagator() propagation.TextMapPropagator {
	return propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	)
}
