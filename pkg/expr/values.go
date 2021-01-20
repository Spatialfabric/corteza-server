package expr

import (
	"reflect"
)

type (
	Type interface {
		Type() string
		Cast(interface{}) (TypedValue, error)
	}

	TypedValue interface {
		Type
		Setter
		Get() interface{}
	}

	typedValueWrap struct {
		Value interface{} `json:"@value"`
		Type  string      `json:"@type"`
	}

	Setter interface {
		Set(interface{}) error
	}

	Selector interface {
		Has(k string) bool
		Select(k string) (TypedValue, error)
	}

	FieldSetter interface {
		SetFieldValue(string, interface{}) error
	}

	Iterator interface {
		Each(func(k string, v TypedValue) error) error
	}

	Decoder interface {
		Decode(reflect.Value) error
	}

	Dict interface {
		Dict() map[string]interface{}
	}
)

func UntypedValue(val interface{}) interface{} {
	if tv, is := val.(TypedValue); is {
		return tv.Get()
	}

	return val
}

func Must(v TypedValue, err error) TypedValue {
	if err != nil {
		panic(err)
	}
	return v
}
