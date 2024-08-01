package main

import (
	"context"
	"fmt"
	"net/netip"
	"time"

	"github.com/AdguardTeam/dnsproxy/upstream"
)

type Resolver struct {
	resolvers upstream.ParallelResolver
	timeout   time.Duration
}

func NewResolver(addresses []string, timeout time.Duration) (*Resolver, error) {
	resolvers := make([]upstream.Resolver, 0, len(addresses))
	opts := &upstream.Options{
		Timeout: timeout,
	}
	for _, addr := range addresses {
		u, err := upstream.AddressToUpstream(addr, opts)
		if err != nil {
			return nil, fmt.Errorf("unable to construct upstream resolver from string %q: %w",
				addr, err)
		}
		resolvers = append(resolvers, &upstream.UpstreamResolver{Upstream: u})
	}
	return &Resolver{
		resolvers: resolvers,
		timeout:   timeout,
	}, nil
}

func (r *Resolver) LookupNetIP(ctx context.Context, network string, host string) (addrs []netip.Addr, err error) {
	return r.resolvers.LookupNetIP(ctx, network, host)
}
