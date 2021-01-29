package driver

import (
	"context"
	"errors"

	"github.com/cortezaproject/corteza-server/system/scion/types"
)

type (
	driver struct{}
)

// Driver returns a new openoffice driver instance
func Driver() types.Driver {
	return &driver{}
}

func (d *driver) Render(ctx context.Context, r *types.Request) (rsp *types.Response, err error) {
	return nil, errors.New("not implemented")
}

func (d *driver) CanRender(tt types.TemplateType) bool {
	return true
}
func (d *driver) CanProduce(dt types.DocumentType) bool {
	return true
}
