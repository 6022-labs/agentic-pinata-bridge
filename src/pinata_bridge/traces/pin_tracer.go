package traces

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces/interfaces"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

const tracerName = "agentic_pinata_bridge"

type PinTracer struct {
	tracer trace.Tracer
}

func NewPinTracer() *PinTracer {
	return &PinTracer{tracer: otel.Tracer(tracerName)}
}

func (t *PinTracer) StartPin(ctx context.Context, cid string) (context.Context, interfaces.Span) {
	ctx, span := t.tracer.Start(ctx, "pinata.pin")
	span.SetAttributes(attribute.String("ipfs.cid", cid))

	return ctx, &otelSpan{span: span}
}

func (t *PinTracer) StartSweep(ctx context.Context, kind string) (context.Context, interfaces.Span) {
	ctx, span := t.tracer.Start(ctx, "pinata.sweep "+kind)
	span.SetAttributes(attribute.String("pinata.sweep_kind", kind))

	return ctx, &otelSpan{span: span}
}
