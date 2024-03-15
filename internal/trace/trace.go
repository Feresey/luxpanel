package trace

import (
	"context"
	"fmt"

	otlhttp "go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func NewTraceProvider(ctx context.Context) (*sdktrace.TracerProvider, error) {
	exp, err := otlhttp.New(ctx, otlhttp.WithEndpoint("http://localhost:8080/trace"))
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
