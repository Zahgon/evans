// Package table provides a table like formatting.
package table

import (
	"reflect"
)

// Presenter is a presenter that formats v with table layout.
type Presenter struct{}

func indirect(rv reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

func indirectType(rt reflect.Type) reflect.Type {
	_ = "STUB: not implemented"
	return *new(reflect.Type)
}

// Format formats v with table layout. v should be a struct type that has a slice field. Format extract the slice field
// and display them. See test cases for example.
func (p *Presenter) Format(v interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func findSlice(rv reflect.Value) (_ reflect.Value, ok bool) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), false
}

func processStructKeys(rt reflect.Type) []string { _ = "STUB: not implemented"; return nil }

func processStructValues(rv reflect.Value) []string { _ = "STUB: not implemented"; return nil }

func NewPresenter() *Presenter { _ = "STUB: not implemented"; return nil }
