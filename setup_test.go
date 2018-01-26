package example

import (
	"testing"

	"github.com/mholt/caddy"
)

func TestSetupWhoami(t *testing.T) {
	c := caddy.NewTestController("dns", `example`)
	if err := setup(c); err != nil {
		t.Fatalf("Expected no errors, but got: %v", err)
	}

	c = caddy.NewTestController("dns", `example more`)
	if err := setup(c); err == nil {
		t.Fatalf("Expected errors, but got: %v", err)
	}
}
