package configurations_test

import (
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/common/host_configurations"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_host/configurations"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_listeners"
	"github.com/knadh/koanf/providers/confmap"
	"github.com/knadh/koanf/v2"
	"go.uber.org/dig"
)

func TestDISmoke(t *testing.T) {
	k := koanf.New(".")
	_ = k.Load(confmap.Provider(map[string]any{
		"host.api_port":                 3000,
		"chains.80002.rpc_http_url":     "http://localhost:8545",
		"chains.80002.rpc_ws_url":       "ws://localhost:8546",
		"pinata.api_key":                "key",
		"pinata.base_url":               "https://api.pinata.cloud",
		"ipfs_check.base_url":           "https://ipfs-check.example",
	}, "."), nil)

	container := configurations.ConfigureDI(k)
	host_configurations.ConfigureLogging(container)

	type params struct {
		dig.In
		Listeners []pinata_bridge_listeners.EventListenerInterface `group:"event_listeners"`
		Subs      []pinata_bridge_listeners.CollectionEventSubscriberInterface `group:"collection_event_subscribers"`
	}
	if err := container.Invoke(func(p params) {
		if len(p.Listeners) != 4 {
			t.Fatalf("expected 4 event listeners, got %d", len(p.Listeners))
		}
		if len(p.Subs) != 3 {
			t.Fatalf("expected 3 collection subscribers, got %d", len(p.Subs))
		}
	}); err != nil {
		t.Fatalf("DI failed: %v", err)
	}
}
