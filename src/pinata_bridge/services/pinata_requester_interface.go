package services

type PinataRequesterInterface interface {
	PinCidToPinata(cid string, hostNodes []string) error
}
