package services

import (
	"math/big"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/services/interfaces"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/settings"
	"github.com/ethereum/go-ethereum/common"
	"go.uber.org/zap"
)

type PushAgentImageCidToPinata struct {
	logger                           *zap.Logger
	chainsSettings                   *settings.ChainsSettings
	pinataRequester                  interfaces.PinataRequesterInterface
	ipfsCheckRequester               interfaces.IpfsCheckRequesterInterface
	agentCollectionRequester         interfaces.AgentCollectionRequesterInterface
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface
}

func NewPushAgentImageCidToPinata(
	logger *zap.Logger,
	chainsSettings *settings.ChainsSettings,
	pinataRequester interfaces.PinataRequesterInterface,
	ipfsCheckRequester interfaces.IpfsCheckRequesterInterface,
	agentCollectionRequester interfaces.AgentCollectionRequesterInterface,
	agentCollectionsManagerRequester interfaces.AgentCollectionsManagerRequesterInterface,
) *PushAgentImageCidToPinata {
	return &PushAgentImageCidToPinata{
		logger:                           logger,
		chainsSettings:                   chainsSettings,
		pinataRequester:                  pinataRequester,
		ipfsCheckRequester:               ipfsCheckRequester,
		agentCollectionRequester:         agentCollectionRequester,
		agentCollectionsManagerRequester: agentCollectionsManagerRequester,
	}
}

func (p *PushAgentImageCidToPinata) PushMissingImageCids() error {
	for _, chainId := range p.chainsSettings.ChainIds() {
		p.logger.Info("Processing chain", zap.Uint64("chainId", chainId))

		allCollections, err := p.agentCollectionsManagerRequester.GetAllCollectionAddresses(chainId)
		if err != nil {
			return err
		}

		for _, collectionAddress := range allCollections {
			p.logger.Info("Processing collection",
				zap.Uint64("chainId", chainId),
				zap.String("collectionAddress", collectionAddress.String()),
			)

			tokenIds, err := p.agentCollectionRequester.GetAllTokenIds(chainId, collectionAddress)
			if err != nil {
				return err
			}

			for _, tokenId := range tokenIds {
				err = p.PushMissingImagesOfAgent(chainId, collectionAddress, tokenId)
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

func (p *PushAgentImageCidToPinata) PushMissingImagesOfAgent(chainId uint64, agentCollectionAddress common.Address, agentCollectionTokenId big.Int) error {
	cids, err := p.agentCollectionRequester.GetAgentImages(chainId, agentCollectionAddress, agentCollectionTokenId)
	if err != nil {
		return err
	}

	for _, cid := range cids {
		isUploaded, err := p.pinataRequester.IsCidUploaded(cid)
		if err != nil {
			p.logger.Error("Failed to check if cid is uploaded", zap.String("cid", cid), zap.Error(err))
			return err
		}

		if *isUploaded {
			p.logger.Debug("CID already uploaded to pinata, skipping", zap.String("cid", cid))
			continue
		}

		p.logger.Info("Pushing agent images cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("agentCollectionTokenId", agentCollectionTokenId.Int64()),
		)

		err = p.PushFromCid(cid)
		if err != nil {
			p.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			return err
		}
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushImagesOfMintProposal(chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error {
	cids, err := p.agentCollectionRequester.GetMintProposalImages(chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return err
	}

	for _, cid := range cids {
		p.logger.Info("Pushing mint proposal image cid to pinata",
			zap.String("cid", cid),
			zap.String("agentCollectionAddress", agentCollectionAddress.String()),
			zap.Int64("proposalId", proposalId.Int64()),
		)

		err = p.PushFromCid(cid)
		if err != nil {
			p.logger.Error("Failed to push agent image cid to pinata", zap.String("cid", cid), zap.Error(err))
			return err
		}
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushImageOfAgentImageProposal(chainId uint64, agentCollectionAddress common.Address, proposalId big.Int) error {
	cid, err := p.agentCollectionRequester.GetAgentImageProposalImage(chainId, agentCollectionAddress, proposalId)
	if err != nil {
		return err
	}

	p.logger.Info("Pushing agent image proposal cid to pinata",
		zap.String("cid", *cid),
		zap.String("agentCollectionAddress", agentCollectionAddress.String()),
		zap.Int64("proposalId", proposalId.Int64()),
	)

	err = p.PushFromCid(*cid)
	if err != nil {
		p.logger.Error("Failed to push agent image proposal cid to pinata", zap.String("cid", *cid), zap.Error(err))
		return err
	}

	return nil
}

func (p *PushAgentImageCidToPinata) PushFromCid(cid string) error {
	addresses, err := p.getCidHostAddresses(cid)
	if err != nil {
		p.logger.Warn("Failed to get host addresses for cid", zap.String("cid", cid), zap.Error(err))
	}
	if len(addresses) == 0 {
		p.logger.Warn("No host addresses found for cid", zap.String("cid", cid))
	}

	err = p.pinataRequester.PinCid(cid, addresses)
	if err != nil {
		// Try again without host addresses
		p.logger.Warn("Failed to pin cid to pinata with host addresses, retrying without", zap.String("cid", cid), zap.Error(err))
		err = p.pinataRequester.PinCid(cid, nil)
		if err != nil {
			return err
		}
	}

	return nil
}

func (p *PushAgentImageCidToPinata) getCidHostAddresses(cid string) ([]string, error) {
	var addresses []string
	var err error

	maxRetries := 3
	for attempt := 1; attempt <= maxRetries; attempt++ {
		addresses, err = p.ipfsCheckRequester.GetMultiAddresses(cid)
		if err == nil && len(addresses) > 0 {
			return addresses, nil
		}

		p.logger.Warn("Failed to get host addresses for cid, retrying...", zap.String("cid", cid), zap.Int("attempt", attempt), zap.Error(err))
	}

	return nil, err
}
