package use_cases_test

// Fixtures shared by every push use case suite in this package.
const (
	testChainId       uint64 = 80002
	testChainIdString        = "80002"

	testCollectionAddress = "0x0000000000000000000000000000000000000000"

	testValidCid      = "bafkreidkdrjbtxjtczfhoiqjcmv2fbnbnpx6erhcsxyxthadvyuovkjhpu"
	testOtherValidCid = "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"

	// testHostNodeId is a libp2p peer id, the shape pinata's host_nodes expects.
	testHostNodeId = "12D3KooWEyoppNCUx8Yx66oV9fJnriXwCcXwDDUA2kj6vnc6iDEg"

	// testNonCidImage is the dicebear avatar url shape agents were minted with in production.
	testNonCidImage = "https://api.dicebear.com/9.x/bottts/svg?seed=RIGI"
)
