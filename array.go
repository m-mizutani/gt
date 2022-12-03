package gt

import "testing"

type ArrayTest[T any] struct {
	actual []T
	t      testing.TB
}

func Array[T comparable](t testing.TB, actual []T) ArrayTest[T] {
	t.Helper()
	return ArrayTest[T]{
		actual: actual,
		t:      t,
	}
}

func (x ArrayTest[T]) Equal(expect []T) ArrayTest[T] {
	x.t.Helper()

	if !EvalCompare(x.actual, expect) {
		x.t.Error("not equal")
		return x
	}

	return x
}

func (x ArrayTest[T]) NotEqual(expect []T) ArrayTest[T] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		x.t.Error("equal")
		return x
	}

	return x
}

func (x ArrayTest[T]) Contain(expect T) ArrayTest[T] {
	x.t.Helper()
	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			return x
		}
	}

	x.t.Error("not contains")
	return x
}

func (x ArrayTest[T]) NotContain(expect T) ArrayTest[T] {
	x.t.Helper()
	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			x.t.Error("contains")
			return x
		}
	}

	return x
}

func (x ArrayTest[T]) Length(expect int) ArrayTest[T] {
	x.t.Helper()
	if len(x.actual) != expect {
		x.t.Error("not contains")
	}
	return x
}

func (x ArrayTest[T]) Required() ArrayTest[T] {
	x.t.Helper()
	if x.t.Failed() {
		x.t.FailNow()
	}
	return x
}
