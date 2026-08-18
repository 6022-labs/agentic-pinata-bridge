package traces

import (
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
)

// otelSpan is the unexported OTel-backed Span; callers only see the interface.
type otelSpan struct {
	span trace.Span
}

func (s *otelSpan) Fail(err error) {
	s.span.RecordError(err)
	s.span.SetStatus(codes.Error, err.Error())
}

func (s *otelSpan) End() {
	s.span.End()
}
