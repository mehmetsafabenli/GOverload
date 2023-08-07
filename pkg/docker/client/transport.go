package client

import (
	"crypto/tls"
	"net/http"
)

func resolveTLSConfig(transport http.RoundTripper) *tls.Config {
	if transport == nil {
		return nil
	}

	if t, ok := transport.(*http.Transport); ok {
		return t.TLSClientConfig
	}
	return nil
}
