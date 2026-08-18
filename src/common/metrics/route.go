// Package metrics holds small helpers shared by adapter-level metrics code.
package metrics

import (
	"regexp"
	"strings"
)

// uuidSegment matches a UUID path segment.
var uuidSegment = regexp.MustCompile(`^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`)

// longHexSegment matches a long hex blob (tx hashes, addresses, CIDs-as-hex).
var longHexSegment = regexp.MustCompile(`^(0x)?[0-9a-fA-F]{16,}$`)

// TemplatizeRoute replaces id-like segments (UUID/hex/number) with {id} to bound http.route cardinality.
func TemplatizeRoute(path string) string {
	if path == "" {
		return path
	}
	segments := strings.Split(path, "/")
	for i, seg := range segments {
		switch {
		case seg == "":
			continue
		case uuidSegment.MatchString(seg), longHexSegment.MatchString(seg), isAllDigits(seg):
			segments[i] = "{id}"
		}
	}
	return strings.Join(segments, "/")
}

func isAllDigits(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return s != ""
}
