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

func (x ValueTest[T]) Equal(expect T) ValueTest[T] {
	x.t.Helper()
	if !EvalCompare(x.actual, expect) {
		x.t.Error("expected equal, but not matched")
	}

	return x
}

func (x ValueTest[T]) NotEqual(expect T) ValueTest[T] {
	x.t.Helper()
	if EvalCompare(x.actual, expect) {
		x.t.Error("expected not equal, but matched")
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

func (x ValueTest[T]) Required() ValueTest[T] {
	x.t.Helper()
	if x.t.Failed() {
		x.t.FailNow()
	}
	return x
}
