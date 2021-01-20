package expr

import (
	"fmt"
	"github.com/spf13/cast"
	"time"
)

func SafeIDSet(dst *uint64, val interface{}, pp ...string) error {
	val = UntypedValue(val)

	if len(pp) > 0 {
		return fmt.Errorf("can not set ID with path")
	}

	if tmp, err := cast.ToUint64E(val); err != nil {
		return err
	} else {
		*dst = tmp
	}

	return nil
}

func SafeStringSet(dst *string, val interface{}, pp ...string) error {
	val = UntypedValue(val)

	if len(pp) > 0 {
		return fmt.Errorf("can not set string with path")
	}

	if tmp, err := cast.ToStringE(val); err != nil {
		return err
	} else {
		*dst = tmp
	}

	return nil
}
func SafeHandleSet(dst *string, val interface{}, pp ...string) error {
	val = UntypedValue(val)

	if len(pp) > 0 {
		return fmt.Errorf("can not set string with path")
	}

	if tmp, err := cast.ToStringE(val); err != nil {
		return err
	} else {
		*dst = tmp
	}

	return nil
}

func SafeBooleanSet(dst *bool, val interface{}, pp ...string) error {
	val = UntypedValue(val)

	if len(pp) > 0 {
		return fmt.Errorf("can not set string with path")
	}

	if tmp, err := cast.ToBoolE(val); err != nil {
		return err
	} else {
		*dst = tmp
	}

	return nil
}

func SafeDateTimeSet(dst *time.Time, val interface{}, pp ...string) error {
	val = UntypedValue(val)

	if len(pp) > 0 {
		return fmt.Errorf("can not set time with path")
	}

	if tmp, err := cast.ToTimeE(val); err != nil {
		return err
	} else {
		*dst = tmp
	}

	return nil
}

func SafeKVSet(dst *map[string]string, val interface{}, pp ...string) (err error) {
	val = UntypedValue(val)

	switch len(pp) {
	case 0:
		var aux map[string]string
		switch casted := val.(type) {
		case nil:
			aux = make(map[string]string)
		case *KV:
			aux = casted.value
		default:
			aux, err = cast.ToStringMapStringE(val)
			if err != nil {
				return err
			}
		}

		*dst = aux
	case 1:
		var aux string
		switch casted := val.(type) {
		case *String:
			aux = casted.value
		default:
			aux, err = cast.ToStringE(val)
			if err != nil {
				return err
			}
		}

		(*dst)[pp[0]] = aux
	default:
		return fmt.Errorf("can not set KV with path deeper than 1 level")
	}

	return nil
}
