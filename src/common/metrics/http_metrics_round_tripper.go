package metrics

import (
	"net/http"
	"time"

	metrics_interfaces "github.com/6022-labs/agentic-pinata-bridge/src/common/metrics/interfaces"
)

// HttpMetricsRoundTripper records http.client.* metrics for every outbound call, then delegates to base.
type HttpMetricsRoundTripper struct {
	base    http.RoundTripper
	metrics metrics_interfaces.ExternalHttpMetricsInterface
}

func NewHttpMetricsRoundTripper(
	base http.RoundTripper,
	metrics metrics_interfaces.ExternalHttpMetricsInterface,
) *HttpMetricsRoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &HttpMetricsRoundTripper{base: base, metrics: metrics}
}

func (t *HttpMetricsRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	start := time.Now()

	// Keyed by server, not path: http.route is a server-side attribute and a remote path is unbounded.
	method := req.Method
	serverUrl := req.URL.String()

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		t.metrics.RecordTransportError(req.Context(), method, serverUrl, time.Since(start), err)
		return resp, err
	}

	t.metrics.RecordRequest(req.Context(), method, serverUrl, resp.StatusCode, time.Since(start), resp.ContentLength)
	return resp, nil
}
