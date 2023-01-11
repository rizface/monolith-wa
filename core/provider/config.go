package provider

import "go.opentelemetry.io/otel"

var tracer = otel.Tracer("service")
