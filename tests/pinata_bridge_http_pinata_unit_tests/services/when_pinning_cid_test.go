package services_test

import (
	"context"
	"io"
	"net/http"
	"net/url"
	"testing"
	"testing/synctest"
	"time"

	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/clients"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/services"
	"github.com/6022-labs/agentic-pinata-bridge/src/pinata_bridge_http_pinata/settings"
	"github.com/6022-labs/agentic-pinata-bridge/tests/pinata_bridge_http_pinata_mocks/clients_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
)

type WhenPinningCidTestingSuite struct {
	sut    *services.PinataRequester
	client *clients_mocks.MockClientWithResponsesInterface
}

func WhenPinningCidBeforeEach(t *testing.T) *WhenPinningCidTestingSuite {
	mockController := gomock.NewController(t)
	client := clients_mocks.NewMockClientWithResponsesInterface(mockController)

	return &WhenPinningCidTestingSuite{
		sut:    services.NewPinataRequester(client, &settings.PinataSettings{RequestTimeout: testRequestTimeout}),
		client: client,
	}
}

func pinByCidResponseWith(statusCode int, body string) *clients.PinByCidResponse {
	return &clients.PinByCidResponse{
		Body:         []byte(body),
		HTTPResponse: &http.Response{StatusCode: statusCode},
	}
}

func rateLimitedPinResponse(retryAfter string) *clients.PinByCidResponse {
	response := pinByCidResponseWith(http.StatusTooManyRequests, `{}`)
	response.HTTPResponse.Header = http.Header{}
	if retryAfter != "" {
		response.HTTPResponse.Header.Set("Retry-After", retryAfter)
	}
	return response
}

func pinBodyWithHostNode() clients.PinByCidJSONRequestBody {
	return clients.PinByCidJSONRequestBody{Cid: testCid, HostNodes: &[]string{testHostNode}}
}

func TestWhenPinningCid(t *testing.T) {
	t.Parallel()

	t.Run("Given host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					_ context.Context,
					body clients.PinByCidJSONRequestBody,
					_ ...clients.RequestEditorFn,
				) (*clients.PinByCidResponse, error) {
					assert.Equal(t, testCid, body.Cid)
					assert.NotNil(t, body.HostNodes)
					assert.Equal(t, []string{testHostNode}, *body.HostNodes)

					return pinByCidResponseWith(http.StatusOK, `{}`), nil
				},
			)
		}

		t.Run("Should pin the cid carrying the host nodes", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

			assert.NoError(t, err)
		})
	})

	t.Run("Given no host nodes are provided", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					_ context.Context,
					body clients.PinByCidJSONRequestBody,
					_ ...clients.RequestEditorFn,
				) (*clients.PinByCidResponse, error) {
					assert.Equal(t, testCid, body.Cid)
					assert.Nil(t, body.HostNodes)

					return pinByCidResponseWith(http.StatusOK, `{}`), nil
				},
			)
		}

		t.Run("Should pin the cid without host nodes", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.NoError(t, err)
		})
	})

	t.Run("Given pinata rejects the cid", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).
				Return(pinByCidResponseWith(http.StatusBadRequest, `{"error":"Invalid IPFS hash"}`), nil)
		}

		t.Run("Should return an error carrying the status and body", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.ErrorContains(t, err, "400")
			assert.ErrorContains(t, err, "Invalid IPFS hash")
		})
	})

	t.Run("Given the transport fails", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).Return(nil, assert.AnError)
		}

		t.Run("Should propagate the error", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.ErrorIs(t, err, assert.AnError)
		})
	})

	t.Run("Given pinata never answers", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), gomock.Any()).DoAndReturn(
				func(
					ctx context.Context,
					_ clients.PinByCidJSONRequestBody,
					_ ...clients.RequestEditorFn,
				) (*clients.PinByCidResponse, error) {
					<-ctx.Done()

					return nil, ctx.Err()
				},
			)
		}

		t.Run("Should give up after the request timeout", func(t *testing.T) {
			t.Parallel()

			suite := WhenPinningCidBeforeEach(t)
			initSuite(suite)

			err := suite.sut.PinCid(context.Background(), testCid, nil)

			assert.ErrorIs(t, err, context.DeadlineExceeded)
		})
	})

	t.Run("Given pinata is rate limiting then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse("3"), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry with the same host nodes after the delay pinata asked for", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, 3*time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata asks to wait exactly ten seconds then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse("10"), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry after the delay pinata asked for", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, 10*time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata asks to wait until a date then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse("Sat, 01 Jan 2000 00:00:05 GMT"), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry at that date", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, 5*time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata is rate limiting without a delay then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse(""), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry after one second", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata is rate limiting with a delay that is not valid then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse("soon"), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry after one second", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata is rate limiting with a date in the past then accepts", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			gomock.InOrder(
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(rateLimitedPinResponse("Fri, 31 Dec 1999 00:00:00 GMT"), nil),
				suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
					Return(pinByCidResponseWith(http.StatusOK, `{}`), nil),
			)
		}

		t.Run("Should retry after one second", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.NoError(t, err)
				assert.Equal(t, time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata keeps rate limiting", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
				Return(rateLimitedPinResponse("1"), nil).
				Times(3)
		}

		t.Run("Should return the error after three attempts", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.ErrorContains(t, err, "429")
				assert.Equal(t, 2*time.Second, time.Since(start))
			})
		})
	})

	t.Run("Given pinata asks to wait longer than ten seconds", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
				Return(rateLimitedPinResponse("11"), nil).
				Times(1)
		}

		t.Run("Should return the error without waiting", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.ErrorContains(t, err, "429")
				assert.Equal(t, time.Duration(0), time.Since(start))
			})
		})
	})

	t.Run("Given pinata asks to wait longer than 64 bits can hold", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
				Return(rateLimitedPinResponse("99999999999999999999999"), nil).
				Times(1)
		}

		t.Run("Should return the error without waiting", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.ErrorContains(t, err, "429")
				assert.Equal(t, time.Duration(0), time.Since(start))
			})
		})
	})

	t.Run("Given pinata fails with a server error", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
				Return(pinByCidResponseWith(http.StatusBadGateway, `{}`), nil).
				Times(1)
		}

		t.Run("Should return the error without retrying", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.ErrorContains(t, err, "502")
				assert.Equal(t, time.Duration(0), time.Since(start))
			})
		})
	})

	t.Run("Given the connection drops", func(t *testing.T) {
		t.Parallel()

		initSuite := func(suite *WhenPinningCidTestingSuite) {
			suite.client.EXPECT().PinByCidWithResponse(gomock.Any(), pinBodyWithHostNode()).
				Return(nil, &url.Error{Op: "Post", Err: io.EOF}).
				Times(1)
		}

		t.Run("Should return the error without retrying", func(t *testing.T) {
			t.Parallel()

			synctest.Test(t, func(t *testing.T) {
				suite := WhenPinningCidBeforeEach(t)
				initSuite(suite)
				start := time.Now()

				err := suite.sut.PinCid(context.Background(), testCid, []string{testHostNode})

				assert.ErrorIs(t, err, io.EOF)
				assert.Equal(t, time.Duration(0), time.Since(start))
			})
		})
	})
}
