package utils_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge/utils"
	"github.com/stretchr/testify/assert"
)

const (
	testCidV0 = "QmYwAPJzv5CZsnA625s3Xf2nemtYgPpHdWEz79ojWnPbdG"
	testCidV1 = "bafkreidkdrjbtxjtczfhoiqjcmv2fbnbnpx6erhcsxyxthadvyuovkjhpu"
)

func TestWhenExtractingACid(t *testing.T) {
	t.Parallel()

	t.Run("Given the image value is a bare cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the cid", func(t *testing.T) {
			t.Parallel()

			for _, value := range []string{testCidV0, testCidV1, "  " + testCidV1 + "  "} {
				cid, ok := utils.ExtractCid(value)

				assert.True(t, ok, value)
				assert.Contains(t, value, cid)
			}
		})
	})

	t.Run("Given the image value wraps the cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should return the wrapped cid", func(t *testing.T) {
			t.Parallel()

			for _, value := range []string{
				"ipfs://" + testCidV1,
				"ipfs://" + testCidV1 + "/avatar.png",
				"ipfs://ipfs/" + testCidV0,
				"https://ipfs.io/ipfs/" + testCidV0,
				"https://gateway.pinata.cloud/ipfs/" + testCidV1 + "/avatar.png?download=true",
			} {
				cid, ok := utils.ExtractCid(value)

				assert.True(t, ok, value)
				assert.Contains(t, []string{testCidV0, testCidV1}, cid)
			}
		})
	})

	t.Run("Given the image value carries no cid", func(t *testing.T) {
		t.Parallel()

		t.Run("Should report the value as unusable", func(t *testing.T) {
			t.Parallel()

			for _, value := range []string{
				"https://api.dicebear.com/9.x/bottts/svg?seed=RIGI",
				"http://example.com/avatar.png",
				"data:image/svg+xml;base64,PHN2Zz48L3N2Zz4=",
				"test-cid",
				"Qm",
				"ipfs://",
				"",
				"   ",
			} {
				cid, ok := utils.ExtractCid(value)

				assert.False(t, ok, value)
				assert.Empty(t, cid, value)
			}
		})
	})
}
