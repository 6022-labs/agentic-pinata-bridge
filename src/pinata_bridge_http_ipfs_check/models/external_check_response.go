package models

type ExternalCheckResponse struct {
	ID                       string   `json:"ID"`
	ConnectionError          string   `json:"ConnectionError"`
	Addrs                    []string `json:"Addrs"`
	ConnectionMaddrs         []string `json:"ConnectionMaddrs"`
	DataAvailableOverBitswap struct {
		Duration  int64  `json:"Duration"`
		Found     bool   `json:"Found"`
		Responded bool   `json:"Responded"`
		Error     string `json:"Error"`
	} `json:"DataAvailableOverBitswap"`
	Source string `json:"Source"`
}
