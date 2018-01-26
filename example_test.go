package example

import (
	"bytes"
	"testing"

	"github.com/coredns/coredns/plugin/pkg/dnstest"
	"github.com/coredns/coredns/plugin/test"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

func TestExample(t *testing.T) {
	ex := Example{Next: test.ErrorHandler()}

	b := &bytes.Buffer{}
	out = b

	ctx := context.TODO()
	r := new(dns.Msg)
	r.SetQuestion("example.org.", dns.TypeA)
	rec := dnstest.NewRecorder(&test.ResponseWriter{})

	ex.ServeDNS(ctx, rec, r)
	if x := b.String(); x != "example" {
		t.Errorf("Failed to print 'example', got %s", x)
	}
}
