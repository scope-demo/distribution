package context

import (
	"testing"

	"github.com/docker/distribution/testutil/tracing"
)

func TestVersionContext(t *testing.T) {
	ctx := tracing.GetContext(t)

	if GetVersion(ctx) != "" {
		t.Fatalf("context should not yet have a version")
	}

	expected := "2.1-whatever"
	ctx = WithVersion(ctx, expected)
	version := GetVersion(ctx)

	if version != expected {
		t.Fatalf("version was not set: %q != %q", version, expected)
	}
}
