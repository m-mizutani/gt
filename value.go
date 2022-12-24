package gt

import (
	"testing"
)

type ValueTest[T any] struct {
	actual T
	t      testing.TB
}

// Value provides ValueTest that has basic comparison methods
func Value[T any](t testing.TB, actual T) ValueTest[T] {
	t.Helper()
	return ValueTest[T]{
		actual: actual,
		t:      t,
	}
}

// V is sugar syntax of Value
func V[T any](t testing.TB, actual T) ValueTest[T] {
	t.Helper()
	return Value(t, actual)
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
		x.t.Error("expected equal, but not matched\n" + Diff(expect, x.actual))
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
		x.t.Error("expected not equal, but matched\n" + Diff(expect, x.actual))
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
		x.t.Errorf("expected nil, but got %v (%T)", x.actual, x.actual)
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
		x.t.Error("expected not nil, but got nil")
	}

	return x
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.Failed().
//
//	name := "Alice"
//	gt.Value(name).Equal("Bob").Must() // Test will stop here
func (x ValueTest[T]) Must() ValueTest[T] {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}
