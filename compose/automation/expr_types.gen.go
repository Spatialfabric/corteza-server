package automation

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// compose/automation/expr_types.yaml

import (
	"context"
	"fmt"
	"github.com/cortezaproject/corteza-server/compose/types"
	. "github.com/cortezaproject/corteza-server/pkg/expr"
)

var _ = context.Background
var _ = fmt.Errorf

// Module is an expression type, wrapper for *types.Module type
type Module struct{ value *types.Module }

// NewModule creates new instance of Module expression type
func NewModule(val interface{}) (*Module, error) {
	if c, err := moduleCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create Module: %w", err)
	} else {
		return &Module{value: c}, nil
	}
}

// Return underlying value on Module
func (t Module) Get() interface{} { return t.value }

// Return type name
func (Module) Type() string { return "ComposeModule" }

// Convert value to *types.Module
func (Module) Cast(val interface{}) (TypedValue, error) {
	return NewModule(val)
}

// Set updates Module
func (t *Module) Set(val interface{}) error {
	// Using moduleCtor to do the casting for us
	if c, err := moduleCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// moduleCast arbitrary value to casts Module
func moduleCast(val interface{}) (*types.Module, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.Module{}, nil
	}

	switch val := val.(type) {
	case *types.Module:
		return val, nil

	case Iterator:
		res := &types.Module{}
		err := val.Each(func(k string, v TypedValue) error {
			return moduleAssigner(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.Module", val)
}

// SelectGVal Implements gval.Selector requirements
func (t Module) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return moduleSelector(t.value, k)
}

func (t Module) Select(k string) (TypedValue, error) {
	return moduleSelector(t.value, k)
}

func (t Module) Has(k string) bool {
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
	case "deletedAt":
		return true
	}
	return false
}

// moduleSelector is field accessor for *types.Module
func moduleSelector(res *types.Module, k string) (TypedValue, error) {
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
	case "deletedAt":
		return NewDateTime(res.DeletedAt)
	}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// moduleAssigner is field value setter for *types.Module
func moduleAssigner(res *types.Module, k string, val interface{}) (err error) {
	switch k {
	case "ID":
		return SafeIDSet(&res.ID, val)
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

// Namespace is an expression type, wrapper for *types.Namespace type
type Namespace struct{ value *types.Namespace }

// NewNamespace creates new instance of Namespace expression type
func NewNamespace(val interface{}) (*Namespace, error) {
	if c, err := namespaceCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create Namespace: %w", err)
	} else {
		return &Namespace{value: c}, nil
	}
}

// Return underlying value on Namespace
func (t Namespace) Get() interface{} { return t.value }

// Return type name
func (Namespace) Type() string { return "ComposeNamespace" }

// Convert value to *types.Namespace
func (Namespace) Cast(val interface{}) (TypedValue, error) {
	return NewNamespace(val)
}

// Set updates Namespace
func (t *Namespace) Set(val interface{}) error {
	// Using namespaceCtor to do the casting for us
	if c, err := namespaceCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// namespaceCast arbitrary value to casts Namespace
func namespaceCast(val interface{}) (*types.Namespace, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.Namespace{}, nil
	}

	switch val := val.(type) {
	case *types.Namespace:
		return val, nil

	case Iterator:
		res := &types.Namespace{}
		err := val.Each(func(k string, v TypedValue) error {
			return namespaceAssigner(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.Namespace", val)
}

// SelectGVal Implements gval.Selector requirements
func (t Namespace) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return namespaceSelector(t.value, k)
}

func (t Namespace) Select(k string) (TypedValue, error) {
	return namespaceSelector(t.value, k)
}

func (t Namespace) Has(k string) bool {
	switch k {
	case "ID":
		return true
	case "name":
		return true
	case "slug", "handle":
		return true
	case "labels":
		return true
	case "createdAt":
		return true
	case "updatedAt":
		return true
	case "deletedAt":
		return true
	}
	return false
}

// namespaceSelector is field accessor for *types.Namespace
func namespaceSelector(res *types.Namespace, k string) (TypedValue, error) {
	switch k {
	case "ID":
		return NewID(res.ID)
	case "name":
		return NewString(res.Name)
	case "slug", "handle":
		return NewHandle(res.Slug)
	case "labels":
		return NewKV(res.Labels)
	case "createdAt":
		return NewDateTime(res.CreatedAt)
	case "updatedAt":
		return NewDateTime(res.UpdatedAt)
	case "deletedAt":
		return NewDateTime(res.DeletedAt)
	}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// namespaceAssigner is field value setter for *types.Namespace
func namespaceAssigner(res *types.Namespace, k string, val interface{}) (err error) {
	switch k {
	case "ID":
		return SafeIDSet(&res.ID, val)
	case "name":
		return SafeStringSet(&res.Name, val)
	case "slug", "handle":
		return SafeHandleSet(&res.Slug, val)
	case "labels":
		return SafeKVSet(&res.Labels, val)
	case "createdAt":
		return SafeDateTimeSet(&res.CreatedAt, val)
	}

	return fmt.Errorf("unknown field '%s'", k)
}

// Record is an expression type, wrapper for *types.Record type
type Record struct{ value *types.Record }

// NewRecord creates new instance of Record expression type
func NewRecord(val interface{}) (*Record, error) {
	if c, err := recordCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create Record: %w", err)
	} else {
		return &Record{value: c}, nil
	}
}

// Return underlying value on Record
func (t Record) Get() interface{} { return t.value }

// Return type name
func (Record) Type() string { return "ComposeRecord" }

// Convert value to *types.Record
func (Record) Cast(val interface{}) (TypedValue, error) {
	return NewRecord(val)
}

// Set updates Record
func (t *Record) Set(val interface{}) error {
	// Using recordCtor to do the casting for us
	if c, err := recordCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// recordCast arbitrary value to casts Record
func recordCast(val interface{}) (*types.Record, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.Record{}, nil
	}

	switch val := val.(type) {
	case *types.Record:
		return val, nil

	case Iterator:
		res := &types.Record{}
		err := val.Each(func(k string, v TypedValue) error {
			return recordAssigner(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.Record", val)
}

// SelectGVal Implements gval.Selector requirements
func (t Record) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return recordSelector(t.value, k)
}

func (t Record) Select(k string) (TypedValue, error) {
	return recordSelector(t.value, k)
}

func (t Record) Has(k string) bool {
	switch k {
	case "ID":
		return true
	case "labels":
		return true
	case "createdAt":
		return true
	case "updatedAt":
		return true
	case "deletedAt":
		return true
	}
	return false
}

// recordSelector is field accessor for *types.Record
func recordSelector(res *types.Record, k string) (TypedValue, error) {
	switch k {
	case "ID":
		return NewID(res.ID)
	case "labels":
		return NewKV(res.Labels)
	case "createdAt":
		return NewDateTime(res.CreatedAt)
	case "updatedAt":
		return NewDateTime(res.UpdatedAt)
	case "deletedAt":
		return NewDateTime(res.DeletedAt)
	}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// recordAssigner is field value setter for *types.Record
func recordAssigner(res *types.Record, k string, val interface{}) (err error) {
	switch k {
	case "ID":
		return SafeIDSet(&res.ID, val)
	case "labels":
		return SafeKVSet(&res.Labels, val)
	case "createdAt":
		return SafeDateTimeSet(&res.CreatedAt, val)
	}

	return fmt.Errorf("unknown field '%s'", k)
}

// RecordValueErrorSet is an expression type, wrapper for *types.RecordValueErrorSet type
type RecordValueErrorSet struct{ value *types.RecordValueErrorSet }

// NewRecordValueErrorSet creates new instance of RecordValueErrorSet expression type
func NewRecordValueErrorSet(val interface{}) (*RecordValueErrorSet, error) {
	if c, err := recordValueErrorSetCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create RecordValueErrorSet: %w", err)
	} else {
		return &RecordValueErrorSet{value: c}, nil
	}
}

// Return underlying value on RecordValueErrorSet
func (t RecordValueErrorSet) Get() interface{} { return t.value }

// Return type name
func (RecordValueErrorSet) Type() string { return "ComposeRecordValueErrorSet" }

// Convert value to *types.RecordValueErrorSet
func (RecordValueErrorSet) Cast(val interface{}) (TypedValue, error) {
	return NewRecordValueErrorSet(val)
}

// Set updates RecordValueErrorSet
func (t *RecordValueErrorSet) Set(val interface{}) error {
	// Using recordValueErrorSetCtor to do the casting for us
	if c, err := recordValueErrorSetCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// recordValueErrorSetCast arbitrary value to casts RecordValueErrorSet
func recordValueErrorSetCast(val interface{}) (*types.RecordValueErrorSet, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return &types.RecordValueErrorSet{}, nil
	}

	switch val := val.(type) {
	case *types.RecordValueErrorSet:
		return val, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.RecordValueErrorSet", val)
}

// RecordValues is an expression type, wrapper for types.RecordValueSet type
type RecordValues struct{ value types.RecordValueSet }

// NewRecordValues creates new instance of RecordValues expression type
func NewRecordValues(val interface{}) (*RecordValues, error) {
	if c, err := recordValuesCast(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create RecordValues: %w", err)
	} else {
		return &RecordValues{value: c}, nil
	}
}

// Return underlying value on RecordValues
func (t RecordValues) Get() interface{} { return t.value }

// Return type name
func (RecordValues) Type() string { return "ComposeRecordValues" }

// Convert value to types.RecordValueSet
func (RecordValues) Cast(val interface{}) (TypedValue, error) {
	return NewRecordValues(val)
}

// Set updates RecordValues
func (t *RecordValues) Set(val interface{}) error {
	// Using recordValuesCtor to do the casting for us
	if c, err := recordValuesCast(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

// recordValuesCast arbitrary value to casts RecordValues
func recordValuesCast(val interface{}) (types.RecordValueSet, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		return types.RecordValueSet{}, nil
	}

	switch val := val.(type) {
	case types.RecordValueSet:
		return val, nil

	}
	return nil, fmt.Errorf("unable to cast type %T to types.RecordValueSet", val)
}
