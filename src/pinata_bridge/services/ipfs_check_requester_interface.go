package services

type IpfsCheckRequesterInterface interface {
	GetMultiAddresses(cid string) ([]string, error)
}
