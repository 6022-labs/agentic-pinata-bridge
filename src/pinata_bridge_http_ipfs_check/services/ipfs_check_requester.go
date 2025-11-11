package services

import "github.com/6022-labs/agentic-ai-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"

type IpfsCheckRequester struct {
	ipfsCheckClient clients.IpfsCheckClientInterface
}

func NewIpfsCheckRequester(ipfsCheckClient clients.IpfsCheckClientInterface) *IpfsCheckRequester {
	return &IpfsCheckRequester{
		ipfsCheckClient: ipfsCheckClient,
	}
}

func (r *IpfsCheckRequester) GetMultiAddresses(cid string) ([]string, error) {
	checkResponses, err := r.ipfsCheckClient.Check(cid)
	if err != nil {
		return nil, err
	}

	multiAddresses := make([]string, 0)
	for _, provider := range checkResponses {
		multiAddresses = append(multiAddresses, provider.ConnectionMaddrs...)
	}

	return multiAddresses, nil
}
