// Package example is a CoreDNS plugin that prints "example" to stdout on every packet received.
//
// It serves as an example CoreDNS plugin with numerous code comments.
package example

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/coredns/coredns/plugin"
	"github.com/coredns/coredns/plugin/metrics"
	clog "github.com/coredns/coredns/plugin/pkg/log"

	"github.com/miekg/dns"
)

// Define log to be a logger with the plugin name in it. This way we can just use log.Info and
// friends to log.
var log = clog.NewWithPlugin("example")

// Example is an example plugin to show how to write a plugin.
type Example struct {
	Next plugin.Handler
}

// ServeDNS implements the plugin.Handler interface. This method gets called when example is used
// in a Server.
func (e Example) ServeDNS(ctx context.Context, w dns.ResponseWriter, r *dns.Msg) (int, error) {
	// This function could be simpler. I.e. just fmt.Println("example") here, but we want to show
	// a slightly more complex example as to make this more interesting.
	// Here we wrap the dns.ResponseWriter in a new ResponseWriter and call the next plugin, when the
	// answer comes back, it will print "example".

	// Debug log that we've have seen the query. This will only be shown when the debug plugin is loaded.
	log.Debug("Received response")

	// Wrap.
	pw := NewResponsePrinter(w)

	// Export metric with the server label set to the current server handling the request.
	requestCount.WithLabelValues(metrics.WithServer(ctx)).Inc()

	// Call next plugin (if any).
	return plugin.NextOrFailure(e.Name(), e.Next, ctx, pw, r)
}

// Name implements the Handler interface.
func (e Example) Name() string { return "example" }

// ResponsePrinter wrap a dns.ResponseWriter and will write example to standard output when WriteMsg is called.
type ResponsePrinter struct {
	dns.ResponseWriter
}

// NewResponsePrinter returns ResponseWriter.
func NewResponsePrinter(w dns.ResponseWriter) *ResponsePrinter {
	return &ResponsePrinter{ResponseWriter: w}
}

// WriteMsg calls the underlying ResponseWriter's WriteMsg method and prints "example" to standard output.
func (r *ResponsePrinter) WriteMsg(res *dns.Msg) error {
	fmt.Fprintln(out, "example")
	return r.ResponseWriter.WriteMsg(res)
}

// Make out a reference to os.Stdout so we can easily overwrite it for testing.
var out io.Writer = os.Stdout
