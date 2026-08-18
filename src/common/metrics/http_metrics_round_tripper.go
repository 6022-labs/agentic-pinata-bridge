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

	// Templatize so id-bearing paths don't explode http.route cardinality.
	method := req.Method
	route := TemplatizeRoute(req.URL.Path)

	resp, err := t.base.RoundTrip(req)
	if err != nil {
		t.metrics.RecordTransportError(req.Context(), method, route, time.Since(start))
		return resp, err
	}

	t.metrics.RecordRequest(req.Context(), method, route, resp.StatusCode, time.Since(start), resp.ContentLength)
	return resp, nil
}
