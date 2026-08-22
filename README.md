# Agentic Pinata Bridge

[![Go Build](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-build.yml/badge.svg)](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-build.yml)
[![Unit Tests](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-unit-tests.yml/badge.svg)](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-unit-tests.yml)
[![Go Lint](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-lint.yml/badge.svg)](https://github.com/6022-labs/agentic-pinata-bridge/actions/workflows/go-lint.yml)
[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)

Agent images live on IPFS as bare CIDs stored on-chain. IPFS only keeps what
someone pins, so an agent's avatar disappears the moment the last node holding
it goes away. This service watches the agent collections deployed by
[`agentic-smart-contracts`](https://github.com/6022-labs/agentic-smart-contracts)
and pins every image CID to [Pinata](https://pinata.cloud), so the images stay
reachable regardless of who else is hosting them.

Part of the [6022 agentic ecosystem](https://docs.agentic.6022.io).

## How it works

Two independent roles, each switchable — run both in one process, or split them
across deployments.

**Listeners** (`HOST__USE_LISTENERS`, default on) follow the chain. The bridge
subscribes to `CollectionCreated` on the `AgentCollectionsManager` of every
configured chain, and to each collection it learns about:

| Event | What gets pinned |
|---|---|
| `Minted` | the new agent's images |
| `MintProposalCreated` | the proposed agent's images, before approval |
| `AgentImageProposalCreated` | the proposed replacement image |

Collections discovered at runtime are picked up without a restart — the
per-collection subscriptions are added as `CollectionCreated` fires.

**Admin sweeps** (`HOST__USE_API`, default on) cover backfill and repair — a
missed event, a Pinata outage, a chain added after the fact. This is an operator
interface, not a public API: it is meant to be reachable by whoever runs the
bridge and by nobody else. See [Exposure](#exposure).

| Route | Scope |
|---|---|
| `POST /push_missing_image_cids` | every chain, every collection, every agent |
| `POST /push_missing_images_of_agent/:chainId/:agentCollectionAddress/:agentCollectionTokenId` | one agent |
| `POST /push_images_of_mint_proposal/:chainId/:agentCollectionAddress/:mintProposalId` | one mint proposal |
| `GET /health` | liveness |

Both paths converge on the same pinner, and pinning by hash is idempotent — the
agent-scoped paths additionally ask Pinata what it already holds and skip those
CIDs. Before each pin the bridge resolves the CID's current host addresses
through an [ipfs-check](https://github.com/ipfs/ipfs-check) endpoint and hands
them to Pinata, so the fetch starts from a known provider instead of a DHT walk.
If that pin fails, it retries once without the addresses.

## Layout

- `src/pinata_bridge` — core domain (pinner, use cases, event handlers).
- `src/pinata_bridge_blockchain` — per-chain RPC clients, contract readers, event subscriptions.
- `src/pinata_bridge_listeners` — listener lifecycle, subscription refresh on new collections.
- `src/pinata_bridge_http_pinata` — Pinata API client.
- `src/pinata_bridge_http_ipfs_check` — ipfs-check client (CID host address lookup).
- `src/pinata_bridge_mvc` — Fiber controllers.
- `src/pinata_bridge_host` — entrypoint, DI wiring, server config.
- `src/common` — telemetry bootstrap, shared metrics and middlewares.
- `grafana/` — the observability stack (Alloy, Mimir, Loki, Tempo, dashboards).

## Run

```sh
cp src/pinata_bridge_host/.env.example src/pinata_bridge_host/.env
# fill the RPC URLs and PINATA__API_KEY
docker compose up
```

- **Bridge admin API:** http://localhost:3000
- **Grafana:** http://localhost:3001

Without Docker, configure `appsettings.json` instead and run the host directly:

```sh
cp src/pinata_bridge_host/appsettings.example.json src/pinata_bridge_host/appsettings.json
cd src/pinata_bridge_host && go run .
```

## Configuration

Every setting is reachable from `appsettings.json` or from the environment —
double underscores map to the JSON nesting (`PINATA__API_KEY` is
`pinata.api_key`). See `src/pinata_bridge_host/.env.example` for the full list.

| Group | What it covers |
|---|---|
| `HOST__*` | `API_PORT`, `LISTEN_ADDRESS` (empty binds every interface), and the `USE_API` / `USE_LISTENERS` role switches |
| `CHAINS__<chain_id>__RPC_HTTP_URL` / `RPC_WS_URL` | one pair per chain to watch; the WS URL carries the event subscriptions |
| `PINATA__*` | API base URL and key |
| `IPFS_CHECK__*` | ipfs-check base URL |
| `TELEMETRY__*` | OTLP export — endpoint, environment, credentials |

`AgentCollectionsManager` addresses are not operator config: they are baked into
the build in `src/pinata_bridge_blockchain/settings/agent_collections_managers.json`
(Polygon and Polygon Amoy today). A chain listed there but absent from `CHAINS__*`
is simply skipped.

## Observability

The host exports metrics, logs and traces over OTLP/gRPC to Grafana Alloy, which
fans them out to Mimir, Loki and Tempo. Dashboards for pin throughput, sweep
duration, chain events, RPC and HTTP are provisioned under `grafana/dashboards/`.
The OTLP ingest is unauthenticated by default; `OTLP_AUTH_*` and `OTLP_TLS_*`
turn that on.

## Exposure

**The HTTP API is an admin interface and must not be reachable from the public
internet.** It carries no authentication and no authorization — none is planned,
because the deployment is expected to keep it private. CORS is open only so an
operator's own tooling can call it from a browser; read it as a symptom of "this
never faces the internet", not as an invitation.

Run it on a private network, a VPN, or behind a reverse proxy that does the
authentication. Nothing here is dangerous in the sense of leaking or destroying
data — the sweeps are idempotent and only ever pin CIDs already published
on-chain — but an open endpoint lets anyone burn your Pinata quota at will.

Two levers: `HOST__USE_API=false` runs listeners only and opens no port at all,
and `HOST__LISTEN_ADDRESS` binds the API to one interface instead of all of them.

## Development

```sh
make build   # go build ./src/...
make test    # unit tests
make lint    # golangci-lint v2
make mocks   # regenerate gomock mocks after an interface change
```

`make help` lists the rest.

## License

MIT — see [LICENSE](LICENSE).
