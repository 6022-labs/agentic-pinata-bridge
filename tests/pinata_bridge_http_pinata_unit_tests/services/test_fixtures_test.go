package services_test

import "time"

// Fixtures shared by every pinata requester suite in this package.
const (
	testCid      = "bafkreidkdrjbtxjtczfhoiqjcmv2fbnbnpx6erhcsxyxthadvyuovkjhpu"
	testHostNode = "/ip4/127.0.0.1/tcp/4001"

	testRequestTimeout = 50 * time.Millisecond
)
