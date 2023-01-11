package app

import (
	"fmt"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.12.0"
)

func newJaegerExporter() *jaeger.Exporter {
	exp, err := jaeger.New(jaeger.WithCollectorEndpoint())
	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return exp
}

func newResource() *resource.Resource {
	r, err := resource.Merge(
		resource.Default(),
		resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String("monolith-wa"),
			semconv.ServiceVersionKey.String("0.0.1"),
		),
	)

	if err != nil {
		fmt.Println(err.Error())
		return nil
	}

	return r
}

func NewOtelProvider() *trace.TracerProvider {
	tp := trace.NewTracerProvider(
		trace.WithResource(newResource()),
		trace.WithBatcher(newJaegerExporter()),
		trace.WithSampler(trace.AlwaysSample()),
	)
	otel.SetTracerProvider(tp)
	return tp
}

func InitOtel() {

}
