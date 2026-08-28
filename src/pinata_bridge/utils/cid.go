package utils

import (
	"strings"

	go_cid "github.com/ipfs/go-cid"
)

const (
	ipfsSchemePrefix = "ipfs://"
	ipfsPathSegment  = "/ipfs/"
)

// ExtractCid pulls the CID out of an on-chain image value; ok is false when the value carries no CID at all.
func ExtractCid(imageValue string) (cid string, ok bool) {
	candidate := strings.TrimSpace(imageValue)

	// Gateway URLs — and the ipfs://ipfs/<cid> form — carry the CID in the segment right after /ipfs/.
	if index := strings.LastIndex(candidate, ipfsPathSegment); index >= 0 {
		candidate = candidate[index+len(ipfsPathSegment):]
	} else {
		candidate = strings.TrimPrefix(candidate, ipfsSchemePrefix)
	}

	// Whatever follows the CID is a sub-path, query or fragment Pinata cannot pin.
	if index := strings.IndexAny(candidate, "/?#"); index >= 0 {
		candidate = candidate[:index]
	}

	if _, err := go_cid.Decode(candidate); err != nil {
		return "", false
	}

	return candidate, true
}
