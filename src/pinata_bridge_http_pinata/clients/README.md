# pinata_bridge_http_pinata clients

`oapi_client.gen.go` is the generated HTTP transport this adapter uses to reach Pinata.
It is produced by `oapi-codegen` (pinned `v2.5.1`) from `pinata-api-v3.yaml`, Pinata's
own OpenAPI spec — vendored here so generation never depends on `docs.pinata.cloud`.

`oapi_codegen.yaml` restricts generation to the two operations this service consumes
(`pinByCid`, `listFiles`). The full spec does not generate: `/groups/{network}` declares
a positional path parameter it never defines, which `oapi-codegen` rejects.

Do not edit `oapi_client.gen.go` by hand. To regenerate:

```sh
sh src/pinata_bridge_http_pinata/clients/generate_client.sh                 # from the vendored spec
sh src/pinata_bridge_http_pinata/clients/generate_client.sh --refresh-spec  # pull upstream first
```

Then refresh the mocks with `deepsearch-mockgen -S ./src -O ./tests -A -P`. It writes one
identical copy per interface it finds in the generated file; keep
`tests/pinata_bridge_http_pinata_mocks/clients_mocks/client_with_responses_mock.go`
and delete the `client_mock.go` / `http_request_doer_mock.go` duplicates.

`../services/pinata_requester.go` wraps the generated `ClientWithResponses` and maps
Pinata's wire shapes to the `PinataRequesterInterface` port the domain depends on.
