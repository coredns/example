package example

import (
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"

	"github.com/mholt/caddy"
)

func init() {
	caddy.RegisterPlugin("example", caddy.Plugin{
		ServerType: "dns",
		Action:     setup,
	})
}

func setup(c *caddy.Controller) error {
	c.Next()
	if c.NextArg() {
		return plugin.Error("example", c.ArgErr())
	}

	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		return Example{Next: next}
	})

	return nil
}
