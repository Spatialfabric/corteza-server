package automation

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	roleService interface {
		FindByID(ctx context.Context, roleID uint64) (*types.Role, error)
		FindByHandle(ctx context.Context, handle string) (*types.Role, error)
		Find(ctx context.Context, filter types.RoleFilter) (set types.RoleSet, f types.RoleFilter, err error)

		Create(ctx context.Context, role *types.Role) (*types.Role, error)
		Update(ctx context.Context, role *types.Role) (*types.Role, error)

		Delete(ctx context.Context, id uint64) error
		Archive(ctx context.Context, id uint64) error
		Unarchive(ctx context.Context, id uint64) error
		Undelete(ctx context.Context, id uint64) error
	}

	rolesHandler struct {
		reg  rolesHandlerRegistry
		rSvc roleService
	}
)

func RolesHandler(reg rolesHandlerRegistry, rSvc roleService) *rolesHandler {
	h := &rolesHandler{
		reg:  reg,
		rSvc: rSvc,
	}

	h.register()
	return h
}

func (h rolesHandler) lookup(ctx context.Context, args *rolesLookupArgs) (results *rolesLookupResults, err error) {
	results = &rolesLookupResults{}
	results.Role, err = lookupRole(ctx, h.rSvc, args.queryID, args.queryHandle)
	return
}

func (h rolesHandler) create(ctx context.Context, args *rolesCreateArgs) (results *rolesCreateResults, err error) {
	results = &rolesCreateResults{}
	results.Role, err = h.rSvc.Create(ctx, args.Role)
	return
}

func (h rolesHandler) update(ctx context.Context, args *rolesUpdateArgs) (results *rolesUpdateResults, err error) {
	results = &rolesUpdateResults{}
	results.Role, err = h.rSvc.Update(ctx, args.Role)
	return
}

func (h rolesHandler) delete(ctx context.Context, args *rolesDeleteArgs) error {
	if id, err := getRoleID(ctx, h.rSvc, args.queryID, args.queryHandle); err != nil {
		return err
	} else {
		return h.rSvc.Delete(ctx, id)
	}
}

func (h rolesHandler) recover(ctx context.Context, args *rolesRecoverArgs) error {
	if id, err := getRoleID(ctx, h.rSvc, args.queryID, args.queryHandle); err != nil {
		return err
	} else {
		return h.rSvc.Undelete(ctx, id)
	}
}

func (h rolesHandler) archive(ctx context.Context, args *rolesArchiveArgs) error {
	if id, err := getRoleID(ctx, h.rSvc, args.queryID, args.queryHandle); err != nil {
		return err
	} else {
		return h.rSvc.Archive(ctx, id)
	}
}

func (h rolesHandler) unarchive(ctx context.Context, args *rolesUnarchiveArgs) error {
	if id, err := getRoleID(ctx, h.rSvc, args.queryID, args.queryHandle); err != nil {
		return err
	} else {
		return h.rSvc.Unarchive(ctx, id)
	}
}

func getRoleID(ctx context.Context, svc roleService, ID uint64, handle string) (uint64, error) {
	if ID > 0 {
		return ID, nil
	}

	role, err := lookupRole(ctx, svc, ID, handle)
	if err != nil {
		return 0, err
	}

	return role.ID, nil

}

func lookupRole(ctx context.Context, svc roleService, ID uint64, handle string) (*types.Role, error) {
	switch {
	case ID > 0:
		return svc.FindByID(ctx, ID)
	case len(handle) > 0:
		return svc.FindByHandle(ctx, handle)
	}

	return nil, fmt.Errorf("empty lookup query params")
}
