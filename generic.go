package gt

import (
	"reflect"
	"testing"
)

func Equal[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	if !EvalCompare(actual, expected) {
		t.Error("values should be matched, but not match\n" + Diff(expected, actual))
	}
}

func EQ[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	Equal(t, actual, expected)
}

func NotEqual[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	if EvalCompare(actual, expected) {
		t.Error("values should not be matched, but match\n" + Diff(expected, actual))
	}
}

func NE[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	NotEqual(t, actual, expected)
}

func isNil(v any) bool {
	if v == nil {
		return true
	}

	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Ptr, reflect.Interface, reflect.Slice, reflect.Map, reflect.Chan:
		return rv.IsNil()
	default:
		return false
	}
}

// Nil is a helper function to check if a value is nil.
// It handles nil pointers, interfaces, slices, maps, and channels.
//
//	gt.Nil(t, nil)
//	gt.Nil(t, (*int)(nil))
//	gt.Nil(t, []int(nil))
//	gt.Nil(t, map[string]int(nil))
//	gt.Nil(t, chan int(nil))
func Nil(t testing.TB, actual any) {
	t.Helper()
	if !isNil(actual) {
		t.Error("value should be nil, but not nil")
	}
}

// NotNil is a helper function to check if a value is not nil.
// It handles nil pointers, interfaces, slices, maps, and channels.
//
//	gt.NotNil(t, 1)
//	gt.NotNil(t, "not nil")
//	gt.NotNil(t, []int{1, 2, 3})
func NotNil(t testing.TB, actual any) {
	t.Helper()
	if isNil(actual) {
		t.Error("value should not be nil, but nil")
	}
}
