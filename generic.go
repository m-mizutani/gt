package gt

import (
	"fmt"
	"reflect"
	"testing"
)

// TestMeta holds common test metadata including the testing.TB instance and description.
// This struct is embedded in all test types to provide consistent description handling.
type TestMeta struct {
	t           testing.TB
	description string
}

// setDesc sets a plain description for the test
func (m *TestMeta) setDesc(desc string) {
	m.t.Helper()
	m.description = desc
}

// setDescf sets a formatted description for the test
func (m *TestMeta) setDescf(format string, args ...any) {
	m.t.Helper()
	m.description = fmt.Sprintf(format, args...)
}

// requiredWithMeta implements Required() functionality with description support
func (m *TestMeta) requiredWithMeta() {
	m.t.Helper()
	requiredWithDescription(m.t, m.description)
}

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

// ExpectError checks if an error occurrence matches the expectation.
// If expected is true, the test passes when err is not nil.
// If expected is false, the test passes when err is nil.
//
//	// Expect an error
//	err := someFailingFunction()
//	gt.ExpectError(t, true, err)  // Pass if err != nil
//
//	// Expect no error
//	err := someSuccessFunction()
//	gt.ExpectError(t, false, err) // Pass if err == nil
//
//	// Conditional error testing
//	shouldFail := true
//	err := conditionalFunction(shouldFail)
//	gt.ExpectError(t, shouldFail, err)
func ExpectError(t testing.TB, expected bool, err error) {
	t.Helper()

	if expected {
		// Error is expected
		if err == nil {
			t.Error("expected error, but got no error")
		}
	} else {
		// No error is expected
		if err != nil {
			t.Errorf("expected no error, but got error: %v", err)
		}
	}
}
