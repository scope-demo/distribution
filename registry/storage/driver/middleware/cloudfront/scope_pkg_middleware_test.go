package middleware

import (
	"os"
	"testing"

	_ "go.undefinedlabs.com/scopeagent/autoinstrument"
	"go.undefinedlabs.com/scopeagent/instrumentation/nethttp"
)

func TestMain(m *testing.M) {
	nethttp.PatchHttpDefaultClient(nethttp.WithPayloadInstrumentation(), nethttp.WithStacktrace())
	os.Exit(m.Run())
}