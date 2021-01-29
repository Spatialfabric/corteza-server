package automation

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/system/types"
)

type (
	userService interface {
		FindByID(ctx context.Context, userID uint64) (*types.User, error)
		FindByHandle(ctx context.Context, handle string) (*types.User, error)
		FindByEmail(ctx context.Context, email string) (*types.User, error)
		Find(ctx context.Context, filter types.UserFilter) (set types.UserSet, f types.UserFilter, err error)

		Create(ctx context.Context, user *types.User) (*types.User, error)
		Update(ctx context.Context, user *types.User) (*types.User, error)

		Delete(ctx context.Context, id uint64) error
		Suspend(ctx context.Context, id uint64) error
		Unsuspend(ctx context.Context, id uint64) error
		Undelete(ctx context.Context, id uint64) error
	}

	usersHandler struct {
		reg  usersHandlerRegistry
		uSvc userService
	}
)

func UsersHandler(reg usersHandlerRegistry, uSvc userService) *usersHandler {
	h := &usersHandler{
		reg:  reg,
		uSvc: uSvc,
	}

	h.register()
	return h
}

func (h usersHandler) lookup(ctx context.Context, args *usersLookupArgs) (results *usersLookupResults, err error) {
	results = &usersLookupResults{}
	results.User, err = lookupUser(ctx, h.uSvc, args.queryID, args.queryEmail, args.queryHandle)
	return
}

func (h usersHandler) create(ctx context.Context, args *usersCreateArgs) (results *usersCreateResults, err error) {
	results = &usersCreateResults{}
	results.User, err = h.uSvc.Create(ctx, args.User)
	return
}

func (h usersHandler) update(ctx context.Context, args *usersUpdateArgs) (results *usersUpdateResults, err error) {
	results = &usersUpdateResults{}
	results.User, err = h.uSvc.Update(ctx, args.User)
	return
}

func (h usersHandler) delete(ctx context.Context, args *usersDeleteArgs) error {
	if id, err := getUserID(ctx, h.uSvc, args.queryID, args.queryEmail, args.queryHandle); err != nil {
		return err
	} else {
		return h.uSvc.Delete(ctx, id)
	}
}

func (h usersHandler) recover(ctx context.Context, args *usersRecoverArgs) error {
	if id, err := getUserID(ctx, h.uSvc, args.queryID, args.queryEmail, args.queryHandle); err != nil {
		return err
	} else {
		return h.uSvc.Undelete(ctx, id)
	}
}

func (h usersHandler) suspend(ctx context.Context, args *usersSuspendArgs) error {
	if id, err := getUserID(ctx, h.uSvc, args.queryID, args.queryEmail, args.queryHandle); err != nil {
		return err
	} else {
		return h.uSvc.Suspend(ctx, id)
	}
}

func (h usersHandler) unsuspend(ctx context.Context, args *usersUnsuspendArgs) error {
	if id, err := getUserID(ctx, h.uSvc, args.queryID, args.queryEmail, args.queryHandle); err != nil {
		return err
	} else {
		return h.uSvc.Unsuspend(ctx, id)
	}
}

func getUserID(ctx context.Context, svc userService, ID uint64, email, handle string) (uint64, error) {
	if ID > 0 {
		return ID, nil
	}

	user, err := lookupUser(ctx, svc, ID, email, handle)
	if err != nil {
		return 0, err
	}

	return user.ID, nil

}

func lookupUser(ctx context.Context, svc userService, ID uint64, email, handle string) (*types.User, error) {
	switch {
	case ID > 0:
		return svc.FindByID(ctx, ID)
	case len(email) > 0:
		return svc.FindByEmail(ctx, email)
	case len(handle) > 0:
		return svc.FindByHandle(ctx, handle)
	}

	return nil, fmt.Errorf("empty lookup query params")
}
