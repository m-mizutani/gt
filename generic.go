package gt

import (
	"fmt"
	"reflect"
	"testing"
)

type baseTest[T any] struct {
	t           testing.TB
	description string
}

// Describe sets a description for the test. The description will be displayed when the test fails.
//
//	gt.Value(t, actual).Describe("User ID should match expected value").Equal(expected)
func (x baseTest[T]) Describe(description string) baseTest[T] {
	x.t.Helper()
	x.description = description
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
//
//	gt.Array(t, items).Describef("Array should contain %d items for user %s", 5, "Alice").Length(5)
func (x baseTest[T]) Describef(format string, args ...any) baseTest[T] {
	x.t.Helper()
	x.description = fmt.Sprintf(format, args...)
	return x
}

func (x baseTest[T]) Required() baseTest[T] {
	x.t.Helper()
	requiredWithDescription(x.t, x.description)
	return x
}

func Equal[T any](t testing.TB, actual T, expected T) baseTest[T] {
	t.Helper()
	if !EvalCompare(actual, expected) {
		t.Error("values should be matched, but not match\n" + Diff(expected, actual))
	}
	return baseTest[T]{t: t}
}

func EQ[T any](t testing.TB, actual T, expected T) baseTest[T] {
	t.Helper()
	return Equal(t, actual, expected)
}

func NotEqual[T any](t testing.TB, actual T, expected T) baseTest[T] {
	t.Helper()
	if EvalCompare(actual, expected) {
		t.Error("values should not be matched, but match\n" + Diff(expected, actual))
	}
	return baseTest[T]{t: t}
}

func NE[T any](t testing.TB, actual T, expected T) baseTest[T] {
	t.Helper()
	return NotEqual(t, actual, expected)
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
func Nil(t testing.TB, actual any) baseTest[any] {
	t.Helper()
	if !isNil(actual) {
		t.Error("value should be nil, but not nil")
	}
	return baseTest[any]{t: t}
}

// NotNil is a helper function to check if a value is not nil.
// It handles nil pointers, interfaces, slices, maps, and channels.
//
//	gt.NotNil(t, 1)
//	gt.NotNil(t, "not nil")
//	gt.NotNil(t, []int{1, 2, 3})
func NotNil(t testing.TB, actual any) baseTest[any] {
	t.Helper()
	if isNil(actual) {
		t.Error("value should not be nil, but nil")
	}
	return baseTest[any]{t: t}
}
