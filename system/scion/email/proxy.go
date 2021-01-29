package driver

import (
	"context"
	"errors"

	"github.com/cortezaproject/corteza-server/system/scion/types"
)

type (
	proxy struct {
		s ProxySettings
	}

	ProxySettings struct {
		Subject    string
		Body       string
		Recipients []string
		CC         []string
		BCC        []string
	}
)

// Proxy returns a new email proxy instance
func Proxy(s ProxySettings) types.Proxy {
	return &proxy{
		s: s,
	}
}

func (p *proxy) Proxy(ctx context.Context, req *types.Request, resp *types.Response) error {
	return errors.New("not implemented")
}
