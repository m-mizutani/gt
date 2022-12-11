package gt

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
)

func toComparable(v any) any {
	switch reflect.ValueOf(v).Kind() {
	case reflect.Pointer, reflect.UnsafePointer:
		if p := reflect.ValueOf(v); !p.IsNil() {
			return p.Elem().Interface()
		}
	}

	return v
}

func Diff(expect, actual any) string {
	switch reflect.ValueOf(actual).Kind() {
	case reflect.Pointer, reflect.UnsafePointer,
		reflect.Array, reflect.Slice,
		reflect.Struct, reflect.Map:
		return "diff:\n" + cmp.Diff(expect, actual, cmp.Exporter(func(t reflect.Type) bool { return true }))

	default:
		return strings.Join([]string{
			fmt.Sprintf("expect: %v", expect),
			fmt.Sprintf("actual: %v", actual),
		}, "\n")
	}
}
