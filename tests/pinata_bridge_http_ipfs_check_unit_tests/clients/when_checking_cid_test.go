package clients_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_ipfs_check/settings"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type WhenCheckingCidTestingSuite struct {
	sut     *clients.IpfsCheckClient
	server  *httptest.Server
	handler http.HandlerFunc
}

func WhenCheckingCidBeforeEach() *WhenCheckingCidTestingSuite {
	suite := &WhenCheckingCidTestingSuite{}

	suite.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if suite.handler == nil {
			http.NotFound(w, r)
			return
		}
		suite.handler(w, r)
	}))

	suite.sut = clients.NewIpfsCheckClient(
		zap.NewNop(),
		http.DefaultClient,
		&settings.IpfsCheckSettings{BaseUrl: suite.server.URL},
	)

	return suite
}

func WhenCheckingCidAfterEach(suite *WhenCheckingCidTestingSuite) {
	suite.server.Close()
}

func TestWhenCheckingCid(t *testing.T) {
	t.Parallel()

	t.Run("Given the ipfs-check service returns providers", func(t *testing.T) {
		t.Parallel()

		var seenMethod, seenPath, seenQuery, seenContentType string

		initSuite := func(suite *WhenCheckingCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, r *http.Request) {
				seenMethod = r.Method
				seenPath = r.URL.Path
				seenQuery = r.URL.RawQuery
				seenContentType = r.Header.Get("Content-Type")

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`[{"ID":"peer-1","ConnectionMaddrs":["/ip4/127.0.0.1/tcp/4001"]}]`))
			}
		}

		t.Run("Should POST to the check endpoint with the cid query and decode the response", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidBeforeEach()
			defer WhenCheckingCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.Check(context.Background(), "QmHash")

			assert.NoError(t, err)
			assert.Len(t, resp, 1)
			assert.Equal(t, "peer-1", resp[0].ID)
			assert.Equal(t, []string{"/ip4/127.0.0.1/tcp/4001"}, resp[0].ConnectionMaddrs)

			assert.Equal(t, http.MethodPost, seenMethod)
			assert.Equal(t, clients.CHECK_ENDPOINT, seenPath)
			assert.Equal(t, "cid=QmHash", seenQuery)
			assert.Equal(t, "application/json", seenContentType)
		})
	})

	t.Run("Given the ipfs-check service returns a non-200 status", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusBadGateway)
				_, _ = w.Write([]byte("upstream error"))
			}
		}

		t.Run("Should return an error", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidBeforeEach()
			defer WhenCheckingCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.Check(context.Background(), "QmHash")

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	})

	t.Run("Given the ipfs-check service answers 200 with a diagnostic object", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"MutableResolution":{"Error":"DNSLink lookup failed"}}`))
			}
		}

		t.Run("Should return an error naming the body", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidBeforeEach()
			defer WhenCheckingCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.Check(context.Background(), "QmHash")

			assert.Nil(t, resp)
			assert.ErrorContains(t, err, "MutableResolution")
		})
	})

	t.Run("Given the ipfs-check service returns a malformed body", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenCheckingCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte("not-json"))
			}
		}

		t.Run("Should return a decode error", func(t *testing.T) {
			t.Parallel()

			suite := WhenCheckingCidBeforeEach()
			defer WhenCheckingCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.Check(context.Background(), "QmHash")

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	})
}
