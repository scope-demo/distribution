package tracinghttp

import (
	"context"
	"go.undefinedlabs.com/scopeagent/env"
	"go.undefinedlabs.com/scopeagent/instrumentation/nethttp"
	"io"
	"net/http"
)

func TracedHTTPTransport() http.RoundTripper {
	return &nethttp.Transport{
		Stacktrace:             true,
		PayloadInstrumentation: true,
	}
}

func TracedHTTPHandler(h http.Handler) http.Handler {
	if env.ScopeDsn.Value != "" {
		h = nethttp.Middleware(h, nethttp.MWPayloadInstrumentation())
	}

	return h
}

func TracedHTTPHandlerFunc(h http.HandlerFunc) http.Handler {
	var handler http.Handler
	handler = h

	if env.ScopeDsn.Value != "" {
		handler = TracedHTTPHandler(h)
	}

	return handler
}

func Get(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := NewRequest(ctx, "GET", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

func Post(ctx context.Context, url, contentType string, body io.Reader) (resp *http.Response, err error) {
	req, err := NewRequest(ctx, "POST", url, body)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

func Head(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := NewRequest(ctx, "HEAD", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

func Delete(ctx context.Context, url string) (resp *http.Response, err error) {
	req, err := NewRequest(ctx, "DELETE", url, nil)
	if err != nil {
		return nil, err
	}
	resp, err = http.DefaultClient.Do(req)
	return resp, err
}

func NewRequest(ctx context.Context, method, url string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)
	return req, nil
}
