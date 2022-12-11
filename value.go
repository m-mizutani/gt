package gt

import (
	"testing"
)

type ValueTest[T any] struct {
	actual T
	t      testing.TB
}

func Value[T any](t testing.TB, actual T) ValueTest[T] {
	t.Helper()
	return ValueTest[T]{
		actual: actual,
		t:      t,
	}
}

// Equal compares Value and expect T.
func (x ValueTest[T]) Equal(expect T) ValueTest[T] {
	x.t.Helper()
	if !EvalCompare(x.actual, expect) {
		x.t.Error("expected equal, but not matched\n" + Diff(expect, x.actual))
	}

	return x
}

func (x ValueTest[T]) NotEqual(expect T) ValueTest[T] {
	x.t.Helper()
	if EvalCompare(x.actual, expect) {
		x.t.Error("expected not equal, but matched\n" + Diff(expect, x.actual))
	}

	return x
}

func (x ValueTest[T]) Nil() ValueTest[T] {
	x.t.Helper()

	if !EvalIsNil(x.actual) {
		x.t.Error("expected nil, but got not nil")
	}

	return x
}

func (x ValueTest[T]) NotNil() ValueTest[T] {
	x.t.Helper()

	if EvalIsNil(x.actual) {
		x.t.Error("expected not nil, but got nil")
	}

	return x
}

// Must check if error has occurred in previous test. If errors in previous test, it immediately stop test by t.Failed().
//
//	name := "Alice"
//	gt.Value(name).Equal("Bob").Must() // Test will stop here
func (x ValueTest[T]) Must() ValueTest[T] {
	x.t.Helper()
	if x.t.Failed() {
		x.t.FailNow()
	}
	return x
}
