package gt

import (
	"fmt"
	"testing"
)

type ValueTest[T any] struct {
	TestMeta
	actual T
}

// Value provides ValueTest that has basic comparison methods
func Value[T any](t testing.TB, actual T) ValueTest[T] {
	t.Helper()
	return ValueTest[T]{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

// V is sugar syntax of Value
func V[T any](t testing.TB, actual T) ValueTest[T] {
	t.Helper()
	return Value(t, actual)
}

// Describe sets a description for the test. The description will be displayed when the test fails.
//
//	gt.Value(t, actual).Describe("User ID should match expected value").Equal(expected)
func (x ValueTest[T]) Describe(description string) ValueTest[T] {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
//
//	gt.Value(t, user.ID).Describef("User ID should be %d for user %s", 123, user.Name).Equal(123)
func (x ValueTest[T]) Describef(format string, args ...any) ValueTest[T] {
	x.setDescf(format, args...)
	return x
}

// Equal check if actual equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	type user struct {
//	  Name string
//	}
//	u1 := user{Name: "blue"}
//	gt.Value(t, u1).Equal(user{Name: "blue"}) // Pass
//	gt.Value(t, u1).Equal(user{Name: "orange"}) // Fail
func (x ValueTest[T]) Equal(expect T) ValueTest[T] {
	x.t.Helper()
	if !EvalCompare(x.actual, expect) {
		msg := "values are not matched\n" + Diff(expect, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	type user struct {
//	  Name string
//	}
//	u1 := user{Name: "blue"}
//	gt.Value(t, u1).NotEqual(user{Name: "blue"})   // Fail
//	gt.Value(t, u1).NotEqual(user{Name: "orange"}) // Pass
func (x ValueTest[T]) NotEqual(expect T) ValueTest[T] {
	x.t.Helper()
	if EvalCompare(x.actual, expect) {
		msg := fmt.Sprintf("values should not be matched, %+v", x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// Nil checks if actual is nil.
//
//	var u *User
//	gt.Value(t, u).Nil() // Pass
//	u = &User{}
//	gt.Value(t, u).Nil() // Fail
func (x ValueTest[T]) Nil() ValueTest[T] {
	x.t.Helper()

	if !EvalIsNil(x.actual) {
		msg := fmt.Sprintf("expected nil, but got %+v (%T)", x.actual, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotNil checks if actual is not nil.
//
//	var u *User
//	gt.Value(t, u).Nil() // Fail
//	u = &User{}
//	gt.Value(t, u).Nil() // Pass
func (x ValueTest[T]) NotNil() ValueTest[T] {
	x.t.Helper()

	if EvalIsNil(x.actual) {
		msg := "expected not nil, but got nil"
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// In checks actual is in expects. Default evaluation function uses reflect.DeepEqual.
func (x ValueTest[T]) In(expects ...T) ValueTest[T] {
	x.t.Helper()

	for i := range expects {
		if EvalCompare(x.actual, expects[i]) {
			return x
		}
	}

	msg := fmt.Sprintf("values should be in %+v, but not found %+v", expects, x.actual)
	x.t.Error(formatErrorMessage(x.description, msg))
	return x
}

// Required check if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.Failed().
//
//	name := "Alice"
//	gt.Value(t, name).Equal("Bob").Required() // Test will stop here
func (x ValueTest[T]) Required() ValueTest[T] {
	x.requiredWithMeta()
	return x
}
