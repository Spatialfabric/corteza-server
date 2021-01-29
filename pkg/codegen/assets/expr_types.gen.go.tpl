package {{ .Package }}

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
// {{ .Source }}

{{ if .Imports }}
import (
	"context"
	"fmt"
{{- range .Imports }}
  {{ normalizeImport . }}
{{- end }}
{{- if ne .Package "expr" }}
	. "github.com/cortezaproject/corteza-server/pkg/expr"
{{- end }}
)
{{ end }}


var _ = context.Background
var _ = fmt.Errorf


{{ range $exprType, $def := .Types }}
// {{ $exprType }} is an expression type, wrapper for {{ $def.As }} type
type {{ $exprType }} struct{ value {{ $def.As }} }

// New{{ $exprType }} creates new instance of {{ $exprType }} expression type
func New{{ $exprType }}(val interface{}) (*{{ $exprType }}, error) {
	if c, err := {{ $def.CastFn }}(UntypedValue(val)); err != nil {
		return nil, fmt.Errorf("unable to create {{ $exprType }}: %w", err)
	} else {
		return &{{ $exprType }}{value: c}, nil
	}
}


// Return underlying value on {{ $exprType }}
func (t {{ $exprType }}) Get() interface{}                         { return t.value }

// Return type name
func ({{ $exprType }}) Type() string                               { return "{{ $.Prefix }}{{ $exprType }}" }

// Convert value to {{ $def.As }}
func ({{ $exprType }}) Cast(val interface{}) (TypedValue, error) {
	return New{{ $exprType }}(val)
}

// Set updates {{ $exprType }}
func (t *{{ $exprType }}) Set(val interface{}) (error) {
	// Using {{ unexport $exprType "ctor" }} to do the casting for us
	if c, err := {{ $def.CastFn }}(UntypedValue(val)); err != nil {
		return err
	} else {
		t.value = c
	}

	return nil
}

{{ if $def.BuiltInCastFn }}
// {{ $def.CastFn }} arbitrary value to casts {{ $exprType }}
func {{ $def.CastFn }}(val interface{}) ({{ $def.As }}, error) {
	val = UntypedValue(val)

	if val == nil {
		// Creating an empty value
		{{- if hasPtr $def.As }}
		return &{{ removePtr $def.As }}{}, nil
		{{- else }}
		return {{ $def.As }}{}, nil
		{{- end }}
	}

	switch val := val.(type) {
	case {{ $def.As }}:
		return val, nil
	{{ if $def.Struct }}
	case Iterator:
		res := &{{ removePtr $def.As }}{}
		err := val.Each(func(k string, v TypedValue) error {
			return {{ unexport $exprType "assigner" }}(res, k, v)
		})

		if err != nil {
			return nil, err
		}

		return res, nil
	{{ end }}
	}
	return {{ $def.Default }}, fmt.Errorf("unable to cast type %T to {{ removePtr $def.As }}", val)
}
{{ end }}

{{ if $def.Struct }}
// SelectGVal Implements gval.Selector requirements
func (t {{ $exprType }}) SelectGVal(ctx context.Context, k string) (interface{}, error) {
	return {{ unexport $exprType "selector" }}(t.value, k)
}

func (t {{ $exprType }}) Select(k string) (TypedValue, error) {
	return {{ unexport $exprType "selector" }}(t.value, k)
}

func (t {{ $exprType }}) Has(k string) bool {
	switch k {
	{{- range $def.Struct }}
		{{- if .ExprType }}
		case {{ printf "%q" .Name }}{{ if .Alias }}, {{ printf "%q" .Alias }}{{ end }}:
			return true
		{{- end }}
	{{- end }}
	}
	return false
}

// {{ unexport $exprType "selector" }} is field accessor for {{ $def.As }}
func {{ unexport $exprType "selector" }}(res {{ $def.As }}, k string) (TypedValue, error) {
switch k {
{{- range $def.Struct }}
	{{- if .ExprType }}
	case {{ printf "%q" .Name }}{{ if .Alias }}, {{ printf "%q" .Alias }}{{ end }}:
		return {{ export "New" .ExprType }}(res.{{ export .Name }})
	{{- end }}
{{- end }}
}
	return nil, fmt.Errorf("unknown field '%s'", k)
}

// {{ unexport $exprType "assigner" }} is field value setter for {{ $def.As }}
func {{ unexport $exprType "assigner" }}(res {{ $def.As }}, k string, val interface{}) (err error) {
	switch k {
{{- range $def.Struct }}
	{{- if not .Readonly }}
	case {{ printf "%q" .Name }}{{ if .Alias }}, {{ printf "%q" .Alias }}{{ end }}:
		return {{ export "Safe" .ExprType "Set" }}(&res.{{ export .Name }}, val)
	{{- end }}
{{- end }}
	}

	return fmt.Errorf("unknown field '%s'", k)
}
{{ end }}

{{ end }}
