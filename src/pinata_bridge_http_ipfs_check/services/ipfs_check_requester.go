package services

import (
	"context"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"
	"github.com/samber/lo"
)

type IpfsCheckRequester struct {
	ipfsCheckClient clients.IpfsCheckClientInterface
}

func NewIpfsCheckRequester(ipfsCheckClient clients.IpfsCheckClientInterface) *IpfsCheckRequester {
	return &IpfsCheckRequester{
		ipfsCheckClient: ipfsCheckClient,
	}
}

func (r *IpfsCheckRequester) GetHostNodeIds(ctx context.Context, cid string) ([]string, error) {
	checkResponses, err := r.ipfsCheckClient.Check(ctx, cid)
	if err != nil {
		return nil, err
	}

	hostNodeIds := make([]string, 0)
	for _, provider := range checkResponses {
		// An empty ConnectionMaddrs means ipfs-check never reached the provider; Pinata gains nothing from it.
		if provider.ID == "" || len(provider.ConnectionMaddrs) == 0 {
			continue
		}

		hostNodeIds = append(hostNodeIds, provider.ID)
	}

	// ipfs-check reports a provider once per source that advertised it.
	return lo.Uniq(hostNodeIds), nil
}
