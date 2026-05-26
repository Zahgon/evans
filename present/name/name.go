package name

import (
	"reflect"
)

// Presenter is a presenter that formats v into the list of names.
type Presenter struct{}

func indirect(rv reflect.Value) reflect.Value {
	_ = "STUB: not implemented"
	return *new(reflect.Value)
}

// Format formats v into the list of names. v should be a struct type.
// Format tries to find "name" tag from the struct fields and format the first appeared field.
// The struct type is only allowed to have struct, slice or primitive type fields.
func (p *Presenter) Format(v interface{}) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func formatFromStruct(rv reflect.Value) (string, error) { _ = "STUB: not implemented"; return "", nil }

func NewPresenter() *Presenter { _ = "STUB: not implemented"; return nil }
