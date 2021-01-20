package expr

import (
	"fmt"
	"strings"
)

const (
	pathDelimiter = "."
)

func PathBase(path string) string {
	return strings.Split(path, ".")[0]
}

func Set(base TypedValue, path string, val interface{}) error {
	return setter(base, val, strings.Split(path, pathDelimiter)...)
}

func setter(base TypedValue, val interface{}, pp ...string) (err error) {
	if len(pp) == 0 {
		panic("setting value with empty path")
	}

	var (
		path    = strings.Join(pp, pathDelimiter)
		failure = fmt.Errorf("can not set value on %s with path '%s'", base.Type(), path)
		key     = pp[0]
	)

	// descend lower by the path but
	// stop before the last part of the path
	for len(pp) > 1 {
		println(path, key)
		s, is := base.(Selector)
		if !is {
			return failure
		}
		key, pp = pp[0], pp[1:]
		if base, err = s.Select(key); err != nil {
			return err
		}
	}

	key = pp[0]

	// try with field setter first
	// if not a FieldSetter it has to be a Selector
	// that returns TypedValue that we can set
	switch setter := base.(type) {
	case FieldSetter:
		return setter.SetFieldValue(key, val)

	case Selector:
		if base, err = setter.Select(key); err != nil {
			return err
		}

		return base.Set(val)
	}

	return failure
}

func Select(base TypedValue, path string) (TypedValue, error) {
	return selector(base, strings.Split(path, pathDelimiter)...)
}

func selector(base TypedValue, pp ...string) (v TypedValue, err error) {
	if len(pp) == 0 {
		panic("selecting value with empty path")
	}

	var (
		path    = strings.Join(pp, pathDelimiter)
		failure = fmt.Errorf("can not get value from %s with path '%s'", base.Type(), path)
		key     string
	)

	// descend lower by the path but
	// stop before the last part of the path
	for len(pp) > 0 {
		s, is := base.(Selector)
		if !is {
			return nil, failure
		}

		key, pp = pp[0], pp[1:]
		if base, err = s.Select(key); err != nil {
			return nil, err
		}

	}

	return base, nil
}
