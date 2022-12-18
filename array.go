package gt

import "testing"

type arrayTest[T any] struct {
	actual []T
	t      testing.TB
}

// Array provides arrayTest that has not only Value test methods but also array (slice) comparison methods
func Array[T comparable](t testing.TB, actual []T) arrayTest[T] {
	t.Helper()
	return arrayTest[T]{
		actual: actual,
		t:      t,
	}
}

// A is sugar syntax of Array
func A[T comparable](t testing.TB, actual []T) arrayTest[T] {
	t.Helper()
	return Array(t, actual)
}

// Equal check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Equal([]int{1, 2, 3, 5}) // Pass
//	gt.Array(t, v).Equal([]int{1, 2, 3, 4}) // Fail
func (x arrayTest[T]) Equal(expect []T) arrayTest[T] {
	x.t.Helper()

	if !EvalCompare(x.actual, expect) {
		x.t.Error("not equal")
		return x
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).NotEqual([]int{1, 2, 3, 5}) // Fail
//	gt.Array(t, v).NotEqual([]int{1, 2, 3, 4}) // Pass
func (x arrayTest[T]) NotEqual(expect []T) arrayTest[T] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		x.t.Error("equal")
		return x
	}

	return x
}

func (x arrayTest[T]) have(expect T) bool {
	x.t.Helper()

	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			return true
		}
	}

	return false
}

// Have check if actual has an expect element
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Have(5)) // Pass
//	gt.Array(t, v).Have(4)) // Fail
func (x arrayTest[T]) Have(expect T) arrayTest[T] {
	x.t.Helper()
	if !x.have(expect) {
		x.t.Errorf("%v expects to have %v, but not contains", x.actual, expect)
	}
	return x
}

// NotHave check if actual does not have an expect element
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Have(5)) // Fail
//	gt.Array(t, v).Have(4)) // Pass
func (x arrayTest[T]) NotHave(expect T) arrayTest[T] {
	x.t.Helper()
	if x.have(expect) {
		x.t.Errorf("%v does not expects to have %v, but contains", x.actual, expect)
	}
	return x
}

func (x arrayTest[T]) contain(expect []T) bool {
	x.t.Helper()

	check := func(i int) bool {
		for j := range expect {
			if i+j >= len(x.actual) || !EvalCompare(x.actual[i+j], expect[j]) {
				return false
			}
		}

		return true
	}

	for i := range x.actual {
		if check(i) {
			return true
		}
	}

	return false
}

// Contain check if actual has expect as sub sequence.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Contain([]int{1, 2, 3})) // Pass
//	gt.Array(t, v).Contain([]int{1, 2, 5})) // Fail
func (x arrayTest[T]) Contain(expect []T) arrayTest[T] {
	x.t.Helper()
	if !x.contain(expect) {
		x.t.Errorf("%v expects to have %v, but not contains", x.actual, expect)
	}
	return x
}

// NotContain check if actual does not have expect as sub sequence.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).NotContain([]int{1, 2, 3})) // Fail
//	gt.Array(t, v).NotContain([]int{1, 2, 5})) // Pass
func (x arrayTest[T]) NotContain(expect []T) arrayTest[T] {
	x.t.Helper()
	if x.contain(expect) {
		x.t.Errorf("%v expects to have %v, but not contains", x.actual, expect)
	}
	return x
}

// Length checks if element number of actual array is expect.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Length(4) // Pass
func (x arrayTest[T]) Length(expect int) arrayTest[T] {
	x.t.Helper()
	if len(x.actual) != expect {
		x.t.Error("not contains")
	}
	return x
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x arrayTest[T]) Must() arrayTest[T] {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}
