package trace

import (
	"context"
	"fmt"

	otlgrpc "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"

	// otlhttp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const traceEnabled = true

func NewTraceProvider(ctx context.Context, serviceName string) (trcer trace.TracerProvider, shutdown func(context.Context) error, err error) {
	if !traceEnabled {
		return noop.NewTracerProvider(), func(ctx context.Context) error { return nil }, nil
	}
	exp, err := otlgrpc.New(ctx, otlgrpc.WithInsecure())
	if err != nil {
		return nil, nil, fmt.Errorf("create http exporter: %w", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(serviceName),
			semconv.DeploymentEnvironmentKey.String("production"),
		)),
	)
	return tp, func(ctx context.Context) error {
		if err := exp.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown exporter: %w", err)
		}
		if err := tp.Shutdown(ctx); err != nil {
			return fmt.Errorf("shutdown trace provider: %w", err)
		}
		return nil
	}, nil
}
