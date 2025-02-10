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

func Nil(t testing.TB, actual any) {
	t.Helper()
	if !isNil(actual) {
		t.Error("value should be nil, but not nil")
	}
}

func NotNil(t testing.TB, actual any) {
	t.Helper()
	if isNil(actual) {
		t.Error("value should not be nil, but nil")
	}
}
