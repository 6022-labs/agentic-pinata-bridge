package interfaces

type IpfsCheckRequesterInterface interface {
	GetMultiAddresses(cid string) ([]string, error)
}
