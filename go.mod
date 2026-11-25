module github.com/6022-labs/agentic-pinata-bridge

go 1.27.4

require (
	github.com/avast/retry-go/v5 v5.1.1
	github.com/ethereum/go-ethereum v1.17.10
	github.com/gofiber/contrib/v3/otel v1.3.5
	github.com/gofiber/contrib/v3/zap v1.0.19
	github.com/gofiber/fiber/v3 v3.7.0
	github.com/ipfs/go-cid v0.6.2
	github.com/joho/godotenv v1.5.1
	github.com/knadh/koanf/parsers/json v1.0.1
	github.com/knadh/koanf/providers/confmap v1.0.1
	github.com/knadh/koanf/providers/file v1.2.1
	github.com/knadh/koanf/v2 v2.4.1
	github.com/oapi-codegen/runtime v1.9.1
	github.com/samber/lo v1.53.0
	github.com/stretchr/testify v1.12.1
	go.opentelemetry.io/contrib/bridges/otelzap v0.21.0
	go.opentelemetry.io/contrib/instrumentation/net/http/otelhttp v0.72.0
	go.opentelemetry.io/otel v1.47.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploggrpc v0.23.0
	go.opentelemetry.io/otel/exporters/otlp/otlplog/otlploghttp v0.23.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc v1.47.0
	go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetrichttp v1.47.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.47.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp v1.47.0
	go.opentelemetry.io/otel/log v1.47.0
	go.opentelemetry.io/otel/metric v1.47.0
	go.opentelemetry.io/otel/sdk v1.47.0
	go.opentelemetry.io/otel/sdk/log v1.47.0
	go.opentelemetry.io/otel/sdk/metric v1.47.0
	go.opentelemetry.io/otel/trace v1.47.0
	go.uber.org/dig v1.19.0
	go.uber.org/mock v0.6.0
	go.uber.org/multierr v1.11.0
	go.uber.org/zap v1.28.0
	google.golang.org/grpc v1.85.0
)

require (
	github.com/Microsoft/go-winio v0.6.4 // indirect
	github.com/ProjectZKM/Ziren/crates/go-runtime/zkvm_runtime v0.0.0-20261119055430-c9988a098d92 // indirect
	github.com/apapsch/go-jsonmerge/v2 v2.0.0 // indirect
	github.com/bits-and-blooms/bitset v1.30.0 // indirect
	github.com/cenkalti/backoff/v5 v5.0.3 // indirect
	github.com/cespare/xxhash/v2 v2.3.0 // indirect
	github.com/clipperhouse/uax29/v2 v2.7.0 // indirect
	github.com/consensys/gnark-crypto v0.22.0 // indirect
	github.com/crate-crypto/go-eth-kzg v1.5.0 // indirect
	github.com/deckarep/golang-set/v2 v2.9.0 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.4.1 // indirect
	github.com/ethereum/c-kzg-4844/v2 v2.1.8 // indirect
	github.com/felixge/httpsnoop v1.1.0 // indirect
	github.com/fjl/jsonw v0.1.0 // indirect
	github.com/fsnotify/fsnotify v1.10.1 // indirect
	github.com/go-logr/logr v1.4.4 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/go-viper/mapstructure/v2 v2.5.0 // indirect
	github.com/gofiber/schema v1.10.0 // indirect
	github.com/gofiber/utils/v2 v2.6.5 // indirect
	github.com/google/uuid v1.6.0 // indirect
	github.com/gorilla/websocket v1.5.3 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.34.0 // indirect
	github.com/holiman/uint256 v1.3.2 // indirect
	github.com/klauspost/compress v1.20.5 // indirect
	github.com/klauspost/cpuid/v2 v2.4.2 // indirect
	github.com/knadh/koanf/maps v0.1.3 // indirect
	github.com/mattn/go-colorable v0.1.16 // indirect
	github.com/mattn/go-isatty v0.0.24 // indirect
	github.com/mattn/go-runewidth v0.0.28 // indirect
	github.com/minio/sha256-simd v1.0.1 // indirect
	github.com/mitchellh/copystructure v1.2.0 // indirect
	github.com/mitchellh/reflectwalk v1.0.2 // indirect
	github.com/molecule-man/go-brrr v1.3.1 // indirect
	github.com/mr-tron/base58 v1.3.0 // indirect
	github.com/multiformats/go-base32 v0.1.0 // indirect
	github.com/multiformats/go-base36 v0.2.0 // indirect
	github.com/multiformats/go-multibase v0.3.0 // indirect
	github.com/multiformats/go-multihash v0.2.3 // indirect
	github.com/multiformats/go-varint v0.1.0 // indirect
	github.com/philhofer/fwd v1.2.1 // indirect
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/spaolacci/murmur3 v1.1.0 // indirect
	github.com/supranational/blst v0.3.17 // indirect
	github.com/tinylib/msgp v1.6.5 // indirect
	github.com/tklauser/go-sysconf v0.4.0 // indirect
	github.com/tklauser/numcpus v0.12.0 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasthttp v1.79.0 // indirect
	github.com/yusufpapurcu/wmi v1.2.4 // indirect
	go.mongodb.org/mongo-driver v1.17.10 // indirect
	go.opentelemetry.io/auto/sdk v1.2.2 // indirect
	go.opentelemetry.io/contrib v1.47.0 // indirect
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.47.0 // indirect
	go.opentelemetry.io/proto/otlp v1.11.1 // indirect
	go.yaml.in/yaml/v3 v3.0.6 // indirect
	golang.org/x/crypto v0.59.0 // indirect
	golang.org/x/net v0.63.0 // indirect
	golang.org/x/sync v0.25.0 // indirect
	golang.org/x/sys v0.50.0 // indirect
	golang.org/x/text v0.44.0 // indirect
	google.golang.org/genproto/googleapis/api v0.0.0-20261118221215-c8830f628b3f // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20261118221215-c8830f628b3f // indirect
	google.golang.org/protobuf v1.36.12 // indirect
	lukechampine.com/blake3 v1.5.1 // indirect
)
