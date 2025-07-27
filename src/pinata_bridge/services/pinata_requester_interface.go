package services

type PinataRequesterInterface interface {
	PinCid(cid string, hostNodes []string) error
	IsCidUploaded(cid string) (*bool, error)
}
