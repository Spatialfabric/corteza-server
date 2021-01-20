package expr

import (
	"bytes"
	"fmt"
	"github.com/cortezaproject/corteza-server/pkg/errors"
	"github.com/spf13/cast"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// Unresolved is a special type that holds value + type it needs to be resolved to
//
// This solves problem with typed value serialization
type Unresolved struct {
	typ   string
	value interface{}
}

// NewUnresolved creates new instance of Unresolved expression type
func NewUnresolved(typ string, val interface{}) (TypedValue, error) {
	return &Unresolved{
		typ:   typ,
		value: UntypedValue(val),
	}, nil
}

// Returns underlying value on Unresolved
func (t Unresolved) Get() interface{} { return t.value }

// Returns type name
func (t Unresolved) Type() string { return t.typ }

// Casts value to interface{}
func (Unresolved) Cast(interface{}) (TypedValue, error) {
	return nil, fmt.Errorf("can not cast to unresolved type")
}

func (t *Unresolved) Set(interface{}) (err error) {
	return fmt.Errorf("can not set on unresolved type")
}

func castAny(val interface{}) (interface{}, error) {
	return val, nil
}

func castDateTime(val interface{}) (out *time.Time, err error) {
	val = UntypedValue(val)
	switch casted := val.(type) {
	case *time.Time:
		return casted, nil
	case time.Time:
		return &casted, nil
	default:
		var c time.Time
		if c, err = cast.ToTimeE(casted); err != nil {
			return nil, err
		}

		return &c, nil
	}
}

func (t *KV) SetFieldValue(key string, val interface{}) error {
	if t.value == nil {
		t.value = make(map[string]string)
	}

	str, err := cast.ToStringE(val)
	t.value[key] = str
	return err
}

func (t *KV) Has(k string) bool {
	_, has := t.value[k]
	return has
}

func (t *KV) Select(k string) (TypedValue, error) {
	if v, has := t.value[k]; has {
		return Must(NewString(v)), nil
	} else {
		return nil, errors.NotFound("no such key '%s'", k)
	}
}

func castKV(val interface{}) (out map[string]string, err error) {
	val = UntypedValue(val)
	switch casted := val.(type) {
	case map[string]string:
		return casted, nil
	default:
		return cast.ToStringMapStringE(casted)
	}
}

func (t *KVV) SetFieldValue(key string, val interface{}) error {
	if t.value == nil {
		t.value = make(map[string][]string)
	}

	str, err := cast.ToStringSliceE(val)
	t.value[key] = str
	return err
}

func castKVV(val interface{}) (out map[string][]string, err error) {
	val = UntypedValue(val)
	switch casted := val.(type) {
	case http.Header:
		return casted, nil
	case url.Values:
		return casted, nil
	default:
		return cast.ToStringMapStringSliceE(casted)
	}
}

func castReader(val interface{}) (out io.Reader, err error) {
	val = UntypedValue(val)

	switch casted := val.(type) {
	case []byte:
		return bytes.NewReader(casted), nil
	case string:
		return strings.NewReader(casted), nil
	case io.Reader:
		return casted, nil
	default:
		return nil, fmt.Errorf("unable to cast %T to io.Reader", val)
	}
}
