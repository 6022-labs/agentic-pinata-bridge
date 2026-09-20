package configurations

import (
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics"
	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces"
	traces_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/traces/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/use_cases"
	"go.uber.org/dig"
)

func AddPinataBridgeConfiguration(container *dig.Container) {
	// Metrics
	if err := container.Provide(
		metrics.NewPinMetrics,
		dig.As(new(metrics_interfaces.PinMetricsInterface)),
	); err != nil {
		panic(err)
	}

	// Settings
	if err := container.Provide(settings.NewChainsSettings); err != nil {
		panic(err)
	}

	// Traces
	if err := container.Provide(
		traces.NewPinTracer,
		dig.As(new(traces_interfaces.PinTracerInterface)),
	); err != nil {
		panic(err)
	}

	// Services
	serviceProviders := []struct {
		constructor any
		port        any
	}{
		{services.NewCidPinner, new(interfaces.CidPinnerInterface)},
		{services.NewAgentImagesPinner, new(interfaces.AgentImagesPinnerInterface)},
		{services.NewMintProposalImagesPinner, new(interfaces.MintProposalImagesPinnerInterface)},
		{services.NewAgentImageProposalImagePinner, new(interfaces.AgentImageProposalImagePinnerInterface)},
		{services.NewMintedEventHandler, new(interfaces.MintedEventHandlerInterface)},
		{services.NewMintProposalCreatedEventHandler, new(interfaces.MintProposalCreatedEventHandlerInterface)},
		{
			services.NewAgentImageProposalCreatedEventHandler,
			new(interfaces.AgentImageProposalCreatedEventHandlerInterface),
		},
		{services.NewAgentImageUpdatedEventHandler, new(interfaces.AgentImageUpdatedEventHandlerInterface)},
	}
	for _, provider := range serviceProviders {
		if err := container.Provide(provider.constructor, dig.As(provider.port)); err != nil {
			panic(err)
		}
	}

	// Use cases
	useCaseProviders := []any{
		use_cases.NewGetHealth,
		use_cases.NewListCollectionAddresses,
		use_cases.NewPushMissingImagesOfAgent,
		use_cases.NewPushImagesOfMintProposal,
		use_cases.NewPushImageOfAgentImageProposal,
		use_cases.NewPushMissingImageCids,
	}
	for _, provider := range useCaseProviders {
		if err := container.Provide(provider); err != nil {
			panic(err)
		}
	}
}
