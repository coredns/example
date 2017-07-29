// The example middleware prints example to stdout on every packet received.
package example

import (
	"fmt"

	"github.com/coredns/coredns/middleware"

	"github.com/miekg/dns"
	"golang.org/x/net/context"
)

// Example is an example middleware to ...
type Example struct {
	Next middleware.Handler
}

// ServeDNS implements the middleware.Handler interface.
func (e Example) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// Somewhat convoluted, as we could have printed "example" here and just call
	// the next middleware - but as an example, show how to wrap a ResponseWriter might be
	// educational.
	pw := NewResponsePrinter(w)
	return middleware.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
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
	fmt.Println("example")
	return r.ResponseWriter.WriteMsg(res)
}
