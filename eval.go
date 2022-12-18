package gt

import "reflect"

// EvalCompare is a function to check if actual value equals with expected value. A developer can replace EvalCompare with own evaluation function if needed.
//
//	EvalCompare(1, 2) == false
//	EvalCompare("abc", "axc") == false
//	EvalCompare([]int{1, 2, 3}, []int{1, 2, 3}) == true
var EvalCompare = func(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// EvalIsNil is a function to check if actual value is nil. A developer can replace EvalIsNil with own nil check function if needed.
//
//	EvalIsNil(1) == false
//	EvalIsNil("abc") == false
//	EvalIsNil(nil) == true
//	EvalIsNil([]int{}) == true
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
