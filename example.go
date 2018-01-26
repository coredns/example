// The example plugin prints example to stdout on every packet received.
package example

import (
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/plugin"
	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

// Example is an example plugin to show how to write a plugin.
type Example struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface.
func (e Example) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// Somewhat convoluted, as we could have printed "example" here and just call
	// the next plugin - but as an example, show how to wrap a ResponseWriter might be
	// educational.
	pw := NewResponsePrinter(w)
	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

// Name implements the Handler interface.
func (e Example) Name() string { return "example" }

type ResponsePrinter struct {
	dns.ResponseWriter
}

// NewResponsePrinter returns a dns.Msg modifier that print example when a query is received.
func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

// WriteMsg records the status code and calls the underlying ResponseWriter's WriteMsg method.
func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	fmt.Fprintf(out, "example")
	return r.ResponseWriter.WriteMsg(res)
}

var out io.Writer = os.Stdout
