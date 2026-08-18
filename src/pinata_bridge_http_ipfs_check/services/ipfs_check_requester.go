package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"
)

type IpfsCheckRequester struct {
	ipfsCheckClient clients.IpfsCheckClientInterface
}

func NewIpfsCheckRequester(ipfsCheckClient clients.IpfsCheckClientInterface) *IpfsCheckRequester {
	return &IpfsCheckRequester{
		ipfsCheckClient: ipfsCheckClient,
	}
}

func (r *IpfsCheckRequester) GetMultiAddresses(ctx context.Context, cid string) ([]string, error) {
	checkResponses, err := r.ipfsCheckClient.Check(ctx, cid)
	if err != nil {
		return nil, err
	}

	multiAddresses := make([]string, 0)
	for _, provider := range checkResponses {
		multiAddresses = append(multiAddresses, provider.ConnectionMaddrs...)
	}

	return multiAddresses, nil
}
