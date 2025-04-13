package models

type ExternalPinByHashRequest struct {
	HashToPin      string                 `json:"hashToPin"`
	PinataOptions  ExternalPinataOptions  `json:"pinataOptions"`
	PinataMetadata ExternalPinataMetadata `json:"pinataMetadata"`
}

type ExternalPinataOptions struct {
	GroupId   string   `json:"groupId"`
	HostNodes []string `json:"hostNodes"`
}

type ExternalPinataMetadata struct {
	Name      string            `json:"name"`
	KeyValues map[string]string `json:"keyvalues"`
}
