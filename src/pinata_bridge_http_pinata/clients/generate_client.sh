#!/usr/bin/env sh
set -eu

# Regenerates the Pinata HTTP client from the vendored Pinata V3 OpenAPI spec.
#
# The spec is vendored on purpose: generation must never depend on docs.pinata.cloud
# being reachable. Pass --refresh-spec to pull the current upstream spec first, then
# review the diff before committing.

usage() {
	echo "Usage: generate_client.sh [--refresh-spec]" >&2
}

CLIENTS_DIR="$(CDPATH= cd -- "$(dirname "$0")" && pwd)"
SPEC_URL="https://docs.pinata.cloud/pinata-api-v3.yaml"
OAPI_CODEGEN_VERSION="v2.5.1"
REFRESH_SPEC="false"

while [ $# -gt 0 ]; do
	case "$1" in
	--refresh-spec)
		REFRESH_SPEC="true"
		shift
		;;
	-h | --help)
		usage
		exit 0
		;;
	*)
		echo "unknown argument: $1" >&2
		usage
		exit 1
		;;
	esac
done

if [ "${REFRESH_SPEC}" = "true" ]; then
	echo "Fetching ${SPEC_URL}"
	curl -fsSL "${SPEC_URL}" -o "${CLIENTS_DIR}/pinata-api-v3.yaml"
fi

echo "Generating pinata client with oapi-codegen ${OAPI_CODEGEN_VERSION}"
(
	cd "${CLIENTS_DIR}"
	go run "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen@${OAPI_CODEGEN_VERSION}" \
		-config oapi_codegen.yaml \
		pinata-api-v3.yaml
)

echo "Done"
