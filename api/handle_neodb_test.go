package api

import (
	"net/http"
)

// The NeoDB fork gates the actor and inbox endpoints on a NeoDB user
// agent. Upstream tests send Go's default user agent, so inject one into
// every test request that does not set its own, instead of editing the
// upstream test files.
func init() {
	http.DefaultTransport = neodbUATransport{inner: http.DefaultTransport}
}

type neodbUATransport struct {
	inner http.RoundTripper
}

func (t neodbUATransport) RoundTrip(req *http.Request) (*http.Response, error) {
	if req.Header.Get("User-Agent") == "" {
		req = req.Clone(req.Context())
		req.Header.Set("User-Agent", "NeoDB/1.0")
	}
	return t.inner.RoundTrip(req)
}
