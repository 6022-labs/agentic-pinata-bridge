package clients_test

import (
	"context"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/models"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
)

const pinataApiKey = "secret-key"

type WhenPinningByHashTestingSuite struct {
	sut     *clients.PinataClient
	server  *httptest.Server
	handler http.HandlerFunc
}

func WhenPinningByHashBeforeEach() *WhenPinningByHashTestingSuite {
	suite := &WhenPinningByHashTestingSuite{}

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

func WhenPinningByHashAfterEach(suite *WhenPinningByHashTestingSuite) {
	suite.server.Close()
}

func TestWhenPinningByHash(t *testing.T) {
	t.Parallel()

	t.Run("Given Pinata accepts the pin request", func(t *testing.T) {
		t.Parallel()

		var seenMethod, seenPath, seenAuth, seenContentType string
		var seenBody []byte

		initSuite := func(suite *WhenPinningByHashTestingSuite) {
			suite.handler = func(w http.ResponseWriter, r *http.Request) {
				seenMethod = r.Method
				seenPath = r.URL.Path
				seenAuth = r.Header.Get("Authorization")
				seenContentType = r.Header.Get("Content-Type")
				seenBody, _ = io.ReadAll(r.Body)

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"id":"pin-1","ipfsHash":"QmHash","status":"prechecking","name":"file"}`))
			}
		}

		t.Run("Should POST to the pin endpoint with auth header and decode the response", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningByHashBeforeEach()
			defer WhenPinningByHashAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.PinByHash(
				context.Background(),
				&models.ExternalPinByHashRequest{HashToPin: "QmHash"},
			)

			assert.NoError(t, err)
			assert.NotNil(t, resp)
			assert.Equal(t, "pin-1", resp.Id)
			assert.Equal(t, "QmHash", resp.IpfsHash)

			assert.Equal(t, http.MethodPost, seenMethod)
			assert.Equal(t, clients.PinByCidEndpoint, seenPath)
			assert.Equal(t, "Bearer "+pinataApiKey, seenAuth)
			assert.Equal(t, "application/json", seenContentType)
			assert.Contains(t, string(seenBody), `"hashToPin":"QmHash"`)
		})
	})

	t.Run("Given Pinata returns a non-200 status", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningByHashTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.WriteHeader(http.StatusInternalServerError)
				_, _ = w.Write([]byte("boom"))
			}
		}

		t.Run("Should return an error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningByHashBeforeEach()
			defer WhenPinningByHashAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.PinByHash(
				context.Background(),
				&models.ExternalPinByHashRequest{HashToPin: "QmHash"},
			)

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	})

	t.Run("Given Pinata returns a malformed body", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningByHashTestingSuite) {
			suite.handler = func(w http.ResponseWriter, _ *http.Request) {
				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte("not-json"))
			}
		}

		t.Run("Should return a decode error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningByHashBeforeEach()
			defer WhenPinningByHashAfterEach(suite)
			initSuite(suite)

			resp, err := suite.sut.PinByHash(
				context.Background(),
				&models.ExternalPinByHashRequest{HashToPin: "QmHash"},
			)

			assert.Error(t, err)
			assert.Nil(t, resp)
		})
	})
}
