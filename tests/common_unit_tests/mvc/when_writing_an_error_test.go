package mvc_test

import (
	"errors"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	apperrors "github.com/6022-labs/agentic-pinata-bridge/src/common/errors"
	"github.com/6022-labs/agentic-pinata-bridge/src/common/mvc"
	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest/observer"
)

type WhenWritingAnErrorTestingSuite struct {
	sut  *mvc.ErrorWriter
	logs *observer.ObservedLogs
}

func WhenWritingAnErrorBeforeEach(t *testing.T) *WhenWritingAnErrorTestingSuite {
	core, logs := observer.New(zapcore.DebugLevel)
	return &WhenWritingAnErrorTestingSuite{
		sut:  mvc.NewErrorWriter(zap.New(core)),
		logs: logs,
	}
}

func (s *WhenWritingAnErrorTestingSuite) write(t *testing.T, err error) (int, string) {
	app := fiber.New()
	app.Get("/boom", func(ctx fiber.Ctx) error {
		return s.sut.Write(ctx, err)
	})

	resp, testErr := app.Test(httptest.NewRequest(http.MethodGet, "/boom", nil))
	assert.NoError(t, testErr)
	body, readErr := io.ReadAll(resp.Body)
	assert.NoError(t, readErr)
	assert.NoError(t, resp.Body.Close())
	return resp.StatusCode, string(body)
}

func TestWhenWritingAnError(t *testing.T) {
	t.Parallel()

	t.Run("Given a ValidationError", func(t *testing.T) {
		t.Parallel()

		sourceErr := apperrors.NewValidationError("chainId", "chainId is invalid")

		t.Run("Should respond 400 with the field and message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusBadRequest, status)
			assert.JSONEq(t, `{"field":"chainId","message":"chainId is invalid"}`, body)
		})

		t.Run("Should log a warning with the request", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.WarnLevel, logs[0].Level)
			assert.Equal(t, "request rejected", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, http.MethodGet, fields["method"])
			assert.Equal(t, "/boom", fields["path"])
			assert.Equal(t, int64(http.StatusBadRequest), fields["status"])
			assert.Equal(t, `{"field":"chainId","message":"chainId is invalid"}`, fields["response"])
		})
	})

	t.Run("Given a BadRequestError", func(t *testing.T) {
		t.Parallel()

		sourceErr := apperrors.NewBadRequestError("invalid_params", "failed to decode: chainId")

		t.Run("Should respond 400 with the code and message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusBadRequest, status)
			assert.JSONEq(t, `{"code":"invalid_params","message":"failed to decode: chainId"}`, body)
		})

		t.Run("Should log a warning with the code and response", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.WarnLevel, logs[0].Level)
			assert.Equal(t, "request rejected", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, "invalid_params", fields["code"])
			assert.Equal(t, `{"code":"invalid_params","message":"failed to decode: chainId"}`, fields["response"])
		})
	})

	t.Run("Given an UnavailableError", func(t *testing.T) {
		t.Parallel()

		sourceErr := apperrors.NewUnavailableError(
			"collections_read_failed",
			"upstream request failed",
			errors.New("dial https://rpc.example/secret-key"),
		)

		t.Run("Should respond 502 with the code and message only", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusBadGateway, status)
			assert.JSONEq(t, `{"code":"collections_read_failed","message":"upstream request failed"}`, body)
		})

		t.Run("Should log an error with the cause", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.ErrorLevel, logs[0].Level)
			assert.Equal(t, "request failed", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, "collections_read_failed", fields["code"])
			assert.Equal(t, int64(http.StatusBadGateway), fields["status"])
			assert.Equal(t, "dial https://rpc.example/secret-key", fields["cause"])
		})
	})

	t.Run("Given an InternalError", func(t *testing.T) {
		t.Parallel()

		sourceErr := apperrors.NewInternalError("pin_failed", errors.New("dial https://rpc.example/secret-key"))

		t.Run("Should respond 500 with the code and the generic message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusInternalServerError, status)
			assert.JSONEq(t, `{"code":"pin_failed","message":"internal server error"}`, body)
		})

		t.Run("Should log an error with the cause", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.ErrorLevel, logs[0].Level)
			assert.Equal(t, "request failed", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, "pin_failed", fields["code"])
			assert.Equal(t, http.MethodGet, fields["method"])
			assert.Equal(t, "/boom", fields["path"])
			assert.Equal(t, int64(http.StatusInternalServerError), fields["status"])
			assert.Equal(t, "dial https://rpc.example/secret-key", fields["cause"])
		})
	})

	t.Run("Given a fiber error", func(t *testing.T) {
		t.Parallel()

		sourceErr := fiber.ErrRequestEntityTooLarge

		t.Run("Should respond with its status, a snake_case code and its message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusRequestEntityTooLarge, status)
			assert.JSONEq(t, `{"code":"request_entity_too_large","message":"Request Entity Too Large"}`, body)
		})

		t.Run("Should log a warning with the code", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.WarnLevel, logs[0].Level)
			assert.Equal(t, "request rejected", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, "request_entity_too_large", fields["code"])
			assert.Equal(t, int64(http.StatusRequestEntityTooLarge), fields["status"])
		})
	})

	t.Run("Given a fiber not found error", func(t *testing.T) {
		t.Parallel()

		sourceErr := fiber.ErrNotFound

		t.Run("Should respond 404 with not_found and its message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusNotFound, status)
			assert.JSONEq(t, `{"code":"not_found","message":"Not Found"}`, body)
		})

		t.Run("Should not log", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			assert.Equal(t, 0, suite.logs.Len())
		})
	})

	t.Run("Given an unknown error", func(t *testing.T) {
		t.Parallel()

		sourceErr := errors.New("dial https://rpc.example/secret-key")

		t.Run("Should respond 500 with internal_error and the generic message", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			status, body := suite.write(t, sourceErr)

			assert.Equal(t, http.StatusInternalServerError, status)
			assert.JSONEq(t, `{"code":"internal_error","message":"internal server error"}`, body)
		})

		t.Run("Should log an error with the cause", func(t *testing.T) {
			t.Parallel()

			suite := WhenWritingAnErrorBeforeEach(t)

			suite.write(t, sourceErr)

			logs := suite.logs.All()
			assert.Len(t, logs, 1)
			assert.Equal(t, zapcore.ErrorLevel, logs[0].Level)
			assert.Equal(t, "request failed", logs[0].Message)
			fields := logs[0].ContextMap()
			assert.Equal(t, "internal_error", fields["code"])
			assert.Equal(t, "dial https://rpc.example/secret-key", fields["cause"])
		})
	})
}
