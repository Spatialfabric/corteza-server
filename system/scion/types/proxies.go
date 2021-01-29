package types

import "context"

type (
	// Proxy proxies the received driver output to somewhere
	Proxy interface {
		Proxy(ctx context.Context, req *Request, resp *Response) error
	}
)
