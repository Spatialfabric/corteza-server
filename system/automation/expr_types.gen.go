package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// system/automation/expr_types.yaml

import (
	"context"
	"fmt"
	. "github.com/cortezaproject/corteza-server/pkg/expr"
	"github.com/cortezaproject/corteza-server/system/types"
)

var _ = context.Background
var _ = fmt.Errorf

// Role is an expression type, wrapper for *types.Role type
type Role struct{ value *types.Role }

// NewRole creates new instance of Role expression type
func NewRole(val interface{}) (*Role, error) {
	if c, err := roleCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create Role: %w", err)
	} else {
		return &Role{value: c}, nil
	}
}

// Return underlying value on Role
func (t Role) Get() interface{} { return t.value }

// Return type name
func (Role) Type() string { return "Role" }

// Convert value to *types.Role
func (Role) Cast(val interface{}) (TypedValue, error) {
	return NewRole(val)
}

// Set updates Role
func (t *Role) Set(val interface{}) error {
	// Using roleCtor to do the casting for us
	if c, err := roleCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// roleCast arbitrary value to casts Role
func roleCast(val interface{}) (*types.Role, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.Role{}, nil
	}

	switch val := val.(type) {
	case *types.Role:
		return val, nil

	case Iterator:
		res := &types.Role{}
		err := val.Each(func(k string, v TypedValue) error {
			return roleAssigner(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.Role", val)
}

// SelectGVal Implements gval.Selector requirements
func (t Role) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return roleSelector(t.value, k)
}

func (t Role) Select(k string) (TypedValue, error) {
	return roleSelector(t.value, k)
}

func (t Role) Has(k string) bool {
	switch k {
	case "ID":
		return true
	case "name":
		return true
	case "handle":
		return true
	case "labels":
		return true
	case "createdAt":
		return true
	case "updatedAt":
		return true
	case "archivedAt":
		return true
	case "deletedAt":
		return true
	}
	return false
}

// roleSelector is field accessor for *types.Role
func roleSelector(res *types.Role, k string) (TypedValue, error) {
	switch k {
	case "ID":
		return NewID(res.ID)
	case "name":
		return NewString(res.Name)
	case "handle":
		return NewHandle(res.Handle)
	case "labels":
		return NewKV(res.Labels)
	case "createdAt":
		return NewDateTime(res.CreatedAt)
	case "updatedAt":
		return NewDateTime(res.UpdatedAt)
	case "archivedAt":
		return NewDateTime(res.ArchivedAt)
	case "deletedAt":
		return NewDateTime(res.DeletedAt)
	}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// roleAssigner is field value setter for *types.Role
func roleAssigner(res *types.Role, k string, val interface{}) (err error) {
	switch k {
	case "name":
		return SafeStringSet(&res.Name, val)
	case "handle":
		return SafeHandleSet(&res.Handle, val)
	case "labels":
		return SafeKVSet(&res.Labels, val)
	case "createdAt":
		return SafeDateTimeSet(&res.CreatedAt, val)
	}

	return fmt.Errorf("unknown field '%s'", k)
}

// User is an expression type, wrapper for *types.User type
type User struct{ value *types.User }

// NewUser creates new instance of User expression type
func NewUser(val interface{}) (*User, error) {
	if c, err := userCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create User: %w", err)
	} else {
		return &User{value: c}, nil
	}
}

// Return underlying value on User
func (t User) Get() interface{} { return t.value }

// Return type name
func (User) Type() string { return "User" }

// Convert value to *types.User
func (User) Cast(val interface{}) (TypedValue, error) {
	return NewUser(val)
}

// Set updates User
func (t *User) Set(val interface{}) error {
	// Using userCtor to do the casting for us
	if c, err := userCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// userCast arbitrary value to casts User
func userCast(val interface{}) (*types.User, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.User{}, nil
	}

	switch val := val.(type) {
	case *types.User:
		return val, nil

	case Iterator:
		res := &types.User{}
		err := val.Each(func(k string, v TypedValue) error {
			return userAssigner(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.User", val)
}

// SelectGVal Implements gval.Selector requirements
func (t User) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return userSelector(t.value, k)
}

func (t User) Select(k string) (TypedValue, error) {
	return userSelector(t.value, k)
}

func (t User) Has(k string) bool {
	switch k {
	case "ID":
		return true
	case "username":
		return true
	case "email":
		return true
	case "name":
		return true
	case "handle":
		return true
	case "emailConfirmed":
		return true
	case "labels":
		return true
	case "createdAt":
		return true
	case "updatedAt":
		return true
	case "suspendedAt":
		return true
	case "deletedAt":
		return true
	}
	return false
}

// userSelector is field accessor for *types.User
func userSelector(res *types.User, k string) (TypedValue, error) {
	switch k {
	case "ID":
		return NewID(res.ID)
	case "username":
		return NewString(res.Username)
	case "email":
		return NewString(res.Email)
	case "name":
		return NewString(res.Name)
	case "handle":
		return NewHandle(res.Handle)
	case "emailConfirmed":
		return NewBoolean(res.EmailConfirmed)
	case "labels":
		return NewKV(res.Labels)
	case "createdAt":
		return NewDateTime(res.CreatedAt)
	case "updatedAt":
		return NewDateTime(res.UpdatedAt)
	case "suspendedAt":
		return NewDateTime(res.SuspendedAt)
	case "deletedAt":
		return NewDateTime(res.DeletedAt)
	}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// userAssigner is field value setter for *types.User
func userAssigner(res *types.User, k string, val interface{}) (err error) {
	switch k {
	case "username":
		return SafeStringSet(&res.Username, val)
	case "email":
		return SafeStringSet(&res.Email, val)
	case "name":
		return SafeStringSet(&res.Name, val)
	case "handle":
		return SafeHandleSet(&res.Handle, val)
	case "emailConfirmed":
		return SafeBooleanSet(&res.EmailConfirmed, val)
	case "labels":
		return SafeKVSet(&res.Labels, val)
	case "createdAt":
		return SafeDateTimeSet(&res.CreatedAt, val)
	}

	return fmt.Errorf("unknown field '%s'", k)
}
