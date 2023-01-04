package gt

import (
	"testing"
)

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

// Number provides NumberTest that has not only Value test methods but also large-small comparison methods
func Number[T number](t testing.TB, actual T) NumberTest[T] {
	t.Helper()
	return NumberTest[T]{
		actual: actual,
		t:      t,
	}
}

// N is sugar syntax of Number
func N[T number](t testing.TB, actual T) NumberTest[T] {
	t.Helper()
	return Number(t, actual)
}

// Equal checks if expect equals given actual value.
//
//	n := 2
//	gt.Number(t, n).Equal(2)
func (x NumberTest[T]) Equal(expect T) NumberTest[T] {
	x.t.Helper()
	if x.actual != expect {
		x.t.Error("numbers are not matched\n" + Diff(expect, x.actual))
	}

	return x
}

// NotEqual checks if expect does not equal given actual value.
//
//	n := 5
//	gt.Number(t, n).NotEqual(1) // Pass
//	gt.number(t, n).Equal(5)    // Fail
func (x NumberTest[T]) NotEqual(expect T) NumberTest[T] {
	x.t.Helper()
	if x.actual == expect {
		x.t.Errorf("numbers should not be matched, %v", x.actual)
	}

	return x
}

// Greater checks if actual value is greater than expect
//
//	n := 5
//	gt.Number(t, n).Greater(3) // Pass
//	gt.Number(t, n).Greater(5) // Fail
func (x NumberTest[T]) Greater(expect T) NumberTest[T] {
	x.t.Helper()
	if !(expect < x.actual) {
		x.t.Errorf("got %v, want grater than %v", x.actual, expect)
	}

	return x
}

// GreaterOrEqual checks if actual value is expect or greater
//
//	n := 5
//	gt.Number(t, n).GreaterOrEqual(3) // Pass
//	gt.Number(t, n).GreaterOrEqual(5) // Pass
//	gt.Number(t, n).GreaterOrEqual(6) // Fail
func (x NumberTest[T]) GreaterOrEqual(expect T) NumberTest[T] {
	x.t.Helper()
	if !(expect <= x.actual) {
		x.t.Errorf("got %v, want greater than or equal %v", x.actual, expect)
	}

	return x
}

// Less checks if actual value is less than expect
//
//	n := 5
//	gt.Number(t, n).Less(6) // Pass
//	gt.Number(t, n).Less(5) // Fail
func (x NumberTest[T]) Less(expect T) NumberTest[T] {
	x.t.Helper()
	if !(x.actual < expect) {
		x.t.Errorf("got %v, want less than %v", x.actual, expect)
	}

	return x
}

// LessOrEqual checks if actual value is expect or Less
//
//	n := 5
//	gt.Number(t, n).LessOrEqual(6) // Pass
//	gt.Number(t, n).LessOrEqual(5) // Pass
//	gt.Number(t, n).LessOrEqual(3) // Fail
func (x NumberTest[T]) LessOrEqual(expect T) NumberTest[T] {
	x.t.Helper()
	if !(x.actual <= expect) {
		x.t.Errorf("got %v, want less than or equal %v", x.actual, expect)
	}

	return x
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x NumberTest[T]) Must() NumberTest[T] {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}
