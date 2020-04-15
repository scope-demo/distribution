package auth

import (
	_ "go.undefinedlabs.com/scopeagent/autoinstrument"
	"go.undefinedlabs.com/scopeagent/instrumentation/nethttp"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	nethttp.PatchHttpDefaultClient(nethttp.WithPayloadInstrumentation(), nethttp.WithStacktrace())
	os.Exit(m.Run())
}
