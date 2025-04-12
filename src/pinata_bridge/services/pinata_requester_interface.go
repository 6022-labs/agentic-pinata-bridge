package services

type PinataRequesterInterface interface {
	PinCidToPinata(cid string) error
}
