package gt

import "reflect"

var EvalCompare = func(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

var EvalIsNil = func(v any) bool {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Invalid:
		return true

	case reflect.Array, reflect.Slice:
		return value.Len() == 0

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Uintptr, reflect.UnsafePointer:
		return value.IsNil()

	default:
		return false
	}
}
