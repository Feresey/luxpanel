package trace

import (
	"context"
	"fmt"

	otlhttp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
	"go.opentelemetry.io/otel/trace/noop"
)

const traceEnabled = false

func NewTraceProvider(ctx context.Context) (trace.TracerProvider, error) {
	if !traceEnabled {
		return noop.NewTracerProvider(), nil
	}
	exp, err := otlhttp.New(ctx, otlhttp.WithEndpoint(":8080"))
	if err != nil {
		return nil, fmt.Errorf("create http exporter: %w", err)
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithBatcher(exp),
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("todo-service"),
			semconv.DeploymentEnvironmentKey.String("production"),
		)),
	)
	return tp, nil
}
