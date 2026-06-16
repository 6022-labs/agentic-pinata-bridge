package clients_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

type WhenQueryingFileByCidTestingSuite struct {
	sut     *clients.PinataClient
	server  *httptest.Server
	handler http.HandlerFunc
}

func WhenQueryingFileByCidBeforeEach() *WhenQueryingFileByCidTestingSuite {
	suite := &WhenQueryingFileByCidTestingSuite{}

	suite.server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if suite.handler == nil {
			http.NotFound(w, r)
			return
		}
		suite.handler(w, r)
	}))

	suite.sut = clients.NewPinataClient(
		zap.NewNop(),
		http.DefaultClient,
		&settings.PinataSettings{
			ApiKey:  pinataApiKey,
			BaseUrl: suite.server.URL,
		},
	)

	return suite
}

func WhenQueryingFileByCidAfterEach(suite *WhenQueryingFileByCidTestingSuite) {
	suite.server.Close()
}

func TestWhenQueryingFileByCid(t *testing.T) {
	t.Parallel()

	t.Run("Given Pinata returns one matching file", func(t *testing.T) {
		t.Parallel()

		var seenMethod, seenPath, seenQuery, seenAuth string

		initSuite := func(suite *WhenQueryingFileByCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, r *http.Request) {
				seenMethod = r.Method
				seenPath = r.URL.Path
				seenQuery = r.URL.RawQuery
				seenAuth = r.Header.Get("Authorization")

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"data":{"files":[{"cid":"QmHash","name":"file","size":10}],"next_page_token":null}}`))
			}
		}

		t.Run("Should GET the files endpoint with the cid query and auth header", func(t *testing.T) {
			t.Parallel()

			suite := WhenQueryingFileByCidBeforeEach()
			defer WhenQueryingFileByCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.QueryFileByCid("QmHash")

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Len(t, resp.Data.Files, 1)
			assert.Equal(t, "QmHash", resp.Data.Files[0].Cid)

			assert.Equal(t, http.MethodGet, seenMethod)
			assert.Equal(t, clients.QueryFilesEndpoint, seenPath)
			assert.Equal(t, "cid=QmHash", seenQuery)
			assert.Equal(t, "Bearer "+pinataApiKey, seenAuth)
		})
	})

	t.Run("Given Pinata returns a non-200 status", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenQueryingFileByCidTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusUnauthorized)
				_, _ = w.Write([]byte("nope"))
			}
		}

		t.Run("Should return an error", func(t *testing.T) {
			t.Parallel()

			suite := WhenQueryingFileByCidBeforeEach()
			defer WhenQueryingFileByCidAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.QueryFileByCid("QmHash")

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	})
}
