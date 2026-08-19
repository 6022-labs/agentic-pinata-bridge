package use_cases_test

import (
	"context"

	traces_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces/interfaces"
	traces_mocks "github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_mocks/traces_mocks/interfaces_mocks"
	"go.uber.org/mock/gomock"
)

// newNoopPinTracer returns a tracer mock whose spans record nothing; traces are never real in a unit test.
func newNoopPinTracer(mockController *gomock.Controller) *traces_mocks.MockPinTracerInterface {
	span := traces_mocks.NewMockSpan(mockController)
	span.EXPECT().End().AnyTimes()
	span.EXPECT().Fail(gomock.Any()).AnyTimes()

	pinTracer := traces_mocks.NewMockPinTracerInterface(mockController)
	pinTracer.EXPECT().StartPin(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, _ string) (context.Context, traces_interfaces.Span) {
			return ctx, span
		}).AnyTimes()
	pinTracer.EXPECT().StartSweep(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, _ string) (context.Context, traces_interfaces.Span) {
			return ctx, span
		}).AnyTimes()

	return pinTracer
}
