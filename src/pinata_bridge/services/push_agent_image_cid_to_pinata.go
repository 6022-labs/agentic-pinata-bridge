package services

import (
	"context"
	"math/big"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/metrics/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

const hostAddressesMaxRetries = 3

type PushAgentImageCidToPinata struct {
	logger                           *zap.Logger
	chainsSettings                   *settings.ChainsSettings
	pinataRequester                  interfaces.PinataRequesterInterface
	ipfsCheckRequester               interfaces.IpfsCheckRequesterInterface
	agentCollectionRequester         interfaces.AgentCollectionRequesterInterface
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface
	pinMetrics                       metrics_interfaces.PinMetricsInterface
}

func NewPushAgentImageCidToPinata(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	pinataRequester interfaces.PinataRequesterInterface,
	ipfsCheckRequester interfaces.IpfsCheckRequesterInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
	pinMetrics metrics_interfaces.PinMetricsInterface,
) *PushAgentImageCidToPinata {
	return &PushAgentImageCidToPinata{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
		pinMetrics:                       pinMetrics,
	}
}

func (p *PushAgentImageCidToPinata) PushMissingImageCids(ctx context.Context) (err error) {
	defer p.recordSweep(ctx, metrics_interfaces.SweepKindAll, time.Now(), &err)

	for _, chainId := range p.chainsSettings.ChainIds() {
		p.logger.Info("Processing chain", zap.Uint64("chainId", chainId))

		allCollections, err := p.agentCollectionsManagerRequester.GetAllCollectionAddresses(ctx, chainId)
		if err != nil {
			return err
		}

		for _, collectionAddress := range allCollections {
			p.logger.Info("Processing collection",
				zap.Uint64("chainId", chainId),
				zap.String("collectionAddress", collectionAddress.String()),
			)

			tokenIds, err := p.agentCollectionRequester.GetAllTokenIds(ctx, chainId, collectionAddress)
			if err != nil {
				return err
			}

			for _, tokenId := range tokenIds {
				err = p.PushMissingImagesOfAgent(ctx, chainId, collectionAddress, tokenId)
				if err != nil {
					p.logger.Error("Failed to push agent image cid to pinata", zap.Error(err))
					continue
				}

				p.logger.Info("Successfully pushed agent image cid to pinata", zap.String("tokenId", tokenId.String()))
			}
		}
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushMissingImagesOfAgent(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	agentCollectionTokenId big.Int,
) (err error) {
	defer p.recordSweep(ctx, metrics_interfaces.SweepKindAgent, time.Now(), &err)

	cids, err := p.agentCollectionRequester.GetAgentImages(ctx, chainId, agentCollectionAddress, agentCollectionTokenId)
	if err != nil {
		return err
	}

	for _, cid := range cids {
		isUploaded, err := p.pinataRequester.IsCidUploaded(ctx, cid)
		if err != nil {
			p.logger.Error("Failed to check if cid is uploaded", zap.String("cid", cid), zap.Error(err))
			return err
		}

		if *isUploaded {
			p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeAlreadyPinned)
			p.logger.Debug("CID already uploaded to pinata, skipping", zap.String("cid", cid))
			continue
		}

		p.logger.Info("Pushing agent images cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("agentCollectionTokenId", agentCollectionTokenId.Int64()),
		)

		err = p.PushFromCid(ctx, cid)
		if err != nil {
			p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomeFailed)
			p.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			return err
		}

		p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindAgent, metrics_interfaces.PinOutcomePinned)
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushImagesOfMintProposal(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	defer p.recordSweep(ctx, metrics_interfaces.SweepKindMintProposal, time.Now(), &err)

	cids, err := p.agentCollectionRequester.GetMintProposalImages(ctx, chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return err
	}

	for _, cid := range cids {
		p.logger.Info("Pushing mint proposal image cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("proposalId", proposalId.Int64()),
		)

		err = p.PushFromCid(ctx, cid)
		if err != nil {
			p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindMintProposal, metrics_interfaces.PinOutcomeFailed)
			p.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			return err
		}

		p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindMintProposal, metrics_interfaces.PinOutcomePinned)
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushImageOfAgentImageProposal(
	ctx context.Context,
	chainId uint64,
	agentCollectionAddress common.Address,
	proposalId big.Int,
) (err error) {
	defer p.recordSweep(ctx, metrics_interfaces.SweepKindImageProposal, time.Now(), &err)

	cid, err := p.agentCollectionRequester.GetAgentImageProposalImage(ctx, chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return err
	}

	p.logger.Info("Pushing agent image proposal cid to pinata",
		zap.String("cid", *cid),
		zap.String("agentCollectionAddress", agentCollectionAddress.String()),
		zap.Int64("proposalId", proposalId.Int64()),
	)

	err = p.PushFromCid(ctx, *cid)
	if err != nil {
		p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomeFailed)
		p.logger.Error("Failed to push agent image proposal cid to pinata", zap.String("cid", *cid), zap.Error(err))
		return err
	}

	p.pinMetrics.RecordSweepImage(ctx, metrics_interfaces.SweepKindImageProposal, metrics_interfaces.PinOutcomePinned)
	return nil
}

func (p *PushAgentImageCidToPinata) PushFromCid(ctx context.Context, cid string) error {
	addresses, err := p.getCidHostAddresses(ctx, cid)
	if err != nil {
		p.logger.Warn("Failed to get host addresses for cid", zap.String("cid", cid), zap.Error(err))
	}
	if len(addresses) == 0 {
		p.logger.Warn("No host addresses found for cid", zap.String("cid", cid))
	}

	withHostAddresses := len(addresses) > 0

	start := time.Now()
	err = p.pinataRequester.PinCid(ctx, cid, addresses)
	if err == nil {
		p.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomePinned, withHostAddresses, time.Since(start))
		return nil
	}
	p.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomeFailed, withHostAddresses, time.Since(start))

	// Retrying is only worth it when the failed attempt actually carried host addresses to drop.
	if !withHostAddresses {
		return err
	}

	p.logger.Warn("Failed to pin cid to pinata with host addresses, retrying without",
		zap.String("cid", cid),
		zap.Error(err),
	)

	start = time.Now()
	err = p.pinataRequester.PinCid(ctx, cid, nil)
	if err != nil {
		p.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomeFailed, false, time.Since(start))
		return err
	}

	p.pinMetrics.RecordPin(ctx, metrics_interfaces.PinOutcomePinned, false, time.Since(start))
	return nil
}

func (p *PushAgentImageCidToPinata) getCidHostAddresses(ctx context.Context, cid string) ([]string, error) {
	var addresses []string
	var err error

	for attempt := 1; attempt <= hostAddressesMaxRetries; attempt++ {
		addresses, err = p.ipfsCheckRequester.GetMultiAddresses(ctx, cid)
		if err == nil && len(addresses) > 0 {
			p.pinMetrics.RecordHostLookup(ctx, metrics_interfaces.HostLookupOutcomeFound, int64(attempt))
			return addresses, nil
		}

		p.logger.Warn("Failed to get host addresses for cid, retrying...",
			zap.String("cid", cid),
			zap.Int("attempt", attempt),
			zap.Error(err),
		)
	}

	outcome := metrics_interfaces.HostLookupOutcomeEmpty
	if err != nil {
		outcome = metrics_interfaces.HostLookupOutcomeFailed
	}
	p.pinMetrics.RecordHostLookup(ctx, outcome, hostAddressesMaxRetries)

	return nil, err
}

func (p *PushAgentImageCidToPinata) recordSweep(ctx context.Context, kind string, start time.Time, err *error) {
	p.pinMetrics.RecordSweep(ctx, kind, time.Since(start), err != nil && *err != nil)
}
