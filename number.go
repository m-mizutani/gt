package gt

import "testing"

type number interface {
	int | uint |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

type NumberTest[T number] struct {
	actual T
	t      testing.TB
}

func Number[T number](t testing.TB, actual T) NumberTest[T] {
	t.Helper()
	return NumberTest[T]{
		actual: actual,
		t:      t,
	}
}

// Equal checks if expect equals given actual value.
//
//	n := 2
//	gt.Number(n).Equal(2)
func (x NumberTest[T]) Equal(expect T) NumberTest[T] {
	x.t.Helper()
	if x.actual != expect {
		x.t.Error("expected equal, but not matched\n" + Diff(expect, x.actual))
	}

	return x
}

// NotEqual checks if expect does not equal given actual value.
//
//	n := 2
//	gt.Number(n).NotEqual(1)
func (x NumberTest[T]) NotEqual(expect T) NumberTest[T] {
	x.t.Helper()
	if x.actual == expect {
		x.t.Error("expected not equal, but matched\n" + Diff(expect, x.actual))
	}

	return x
}

func (x NumberTest[T]) GreaterThan(expect T) NumberTest[T] {
	x.t.Helper()
	if !(expect < x.actual) {
		x.t.Error("expected greater, but not greater")
	}

	return x
}

func (x NumberTest[T]) LessThan(expect T) NumberTest[T] {
	x.t.Helper()
	if !(x.actual < expect) {
		x.t.Error("expected less, but not less")
	}

	return x
}

// Must check if error has occurred in previous test. If errors in previous test, it immediately stop test by t.Failed().
//
//	age := 42
//	gt.Number(age).LessThan(18).Must() // Test will stop here
func (x NumberTest[T]) Must() NumberTest[T] {
	x.t.Helper()
	if x.t.Failed() {
		x.t.FailNow()
	}
	return x
}
