package chord

import (
	"github.com/coredns/caddy"
	"github.com/coredns/coredns/core/dnsserver"
	"github.com/coredns/coredns/plugin"
)

func init() { plugin.Register("chord", setup) }

func setup(c *caddy.Controller) error {
	c.Next() // 'chord'
	if c.NextArg() {
		return plugin.Error("chord", c.ArgErr())
	}

	chord := Chord{}
	dnsserver.GetConfig(c).AddPlugin(func(next plugin.Handler) plugin.Handler {
		chord.Next = next
		return chord
	})

	return nil
}
