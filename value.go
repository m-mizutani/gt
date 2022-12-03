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
		x.t.Error("not match")
	}

	return x
}

func (x ValueTest[T]) Nil() ValueTest[T] {
	x.t.Helper()

	if !EvalIsNil(x.actual) {
		x.t.Error("not nil")
	}

	return x
}

func (x ValueTest[T]) NotNil() ValueTest[T] {
	x.t.Helper()

	if EvalIsNil(x.actual) {
		x.t.Error("nil")
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
