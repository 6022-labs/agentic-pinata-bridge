package models

type ExternalPinByHashResponse struct {
	Id       string `json:"id"`
	IpfsHash string `json:"ipfsHash"`
	Status   string `json:"status"`
	Name     string `json:"name"`
}
