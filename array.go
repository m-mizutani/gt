package gt

import "testing"

type ArrayTest[T any] struct {
	actual []T
	t      testing.TB
}

// Array provides ArrayTest that has not only Value test methods but also array (slice) comparison methods
func Array[T any](t testing.TB, actual []T) ArrayTest[T] {
	t.Helper()
	return ArrayTest[T]{
		actual: actual,
		t:      t,
	}
}

// A is sugar syntax of Array
func A[T any](t testing.TB, actual []T) ArrayTest[T] {
	t.Helper()
	return Array(t, actual)
}

// Equal check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Equal([]int{1, 2, 3, 5}) // Pass
//	gt.Array(t, v).Equal([]int{1, 2, 3, 4}) // Fail
func (x ArrayTest[T]) Equal(expect []T) ArrayTest[T] {
	x.t.Helper()

	if !EvalCompare(x.actual, expect) {
		x.t.Errorf("arrays are not matched\n" + Diff(x.actual, expect))
		return x
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).NotEqual([]int{1, 2, 3, 5}) // Fail
//	gt.Array(t, v).NotEqual([]int{1, 2, 3, 4}) // Pass
func (x ArrayTest[T]) NotEqual(expect []T) ArrayTest[T] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		x.t.Errorf("arrays should not be matched, %+v", x.actual)
		return x
	}

	return x
}

// EqualAt checks if actual[idx] equals expect. If idx is out of range, f is not called and test will trigger error.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).EqualAt(2, 3) // Pass
//	gt.Array(t, v).EqualAt(2, 1) // Fail
//	gt.Array(t, v).EqualAt(2, 5) // Fail by out of range
func (x ArrayTest[T]) EqualAt(idx int, expect T) ArrayTest[T] {
	x.t.Helper()

	if idx < 0 || len(x.actual) <= idx {
		x.t.Errorf("array length is %d, then %d is out of range", len(x.actual), idx)
	} else if !EvalCompare(x.actual[idx], expect) {
		x.t.Errorf("array[%d] is expected %+v, but actual is %+v", idx, expect, x.actual[idx])
	}

	return x
}

// NotEqualAt checks if actual[idx] equals expect. If idx is out of range, f is not called and test will trigger error.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).NotEqualAt(2, 1) // Pass
//	gt.Array(t, v).NotEqualAt(2, 3) // Fail
//	gt.Array(t, v).NotEqualAt(2, 5) // Fail by out of range
func (x ArrayTest[T]) NotEqualAt(idx int, expect T) ArrayTest[T] {
	x.t.Helper()

	if idx < 0 || len(x.actual) <= idx {
		x.t.Errorf("array length is %d, then %d is out of range", len(x.actual), idx)
	} else if EvalCompare(x.actual[idx], expect) {
		x.t.Errorf("array[%d] is not expected %+v, but actual is %+v", idx, expect, x.actual[idx])
	}

	return x
}

func (x ArrayTest[T]) have(expect T) bool {
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
func (x ArrayTest[T]) Have(expect T) ArrayTest[T] {
	x.t.Helper()
	if !x.have(expect) {
		x.t.Errorf("%+v expects to have %+v, but not contains", x.actual, expect)
	}
	return x
}

// NotHave check if actual does not have an expect element
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Have(5)) // Fail
//	gt.Array(t, v).Have(4)) // Pass
func (x ArrayTest[T]) NotHave(expect T) ArrayTest[T] {
	x.t.Helper()
	if x.have(expect) {
		x.t.Errorf("%+v does not expects to have %+v, but contains", x.actual, expect)
	}
	return x
}

func (x ArrayTest[T]) contain(expect []T) bool {
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
func (x ArrayTest[T]) Contain(expect []T) ArrayTest[T] {
	x.t.Helper()
	if !x.contain(expect) {
		x.t.Errorf("%+v expects to have %+v, but not contains", x.actual, expect)
	}
	return x
}

// NotContain check if actual does not have expect as sub sequence.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).NotContain([]int{1, 2, 3})) // Fail
//	gt.Array(t, v).NotContain([]int{1, 2, 5})) // Pass
func (x ArrayTest[T]) NotContain(expect []T) ArrayTest[T] {
	x.t.Helper()
	if x.contain(expect) {
		x.t.Errorf("%+v expects to have %+v, but not contains", x.actual, expect)
	}
	return x
}

// Length checks if element number of actual array is expect.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Length(4) // Pass
func (x ArrayTest[T]) Length(expect int) ArrayTest[T] {
	x.t.Helper()
	if len(x.actual) != expect {
		x.t.Errorf("array length is expected to be %d, but actual is %d", expect, len(x.actual))
	}
	return x
}

// Longer checks if array length is longer than expect.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Longer(3) // Pass
//	gt.Array(t, v).Longer(4) // Fail
func (x ArrayTest[T]) Longer(expect int) ArrayTest[T] {
	x.t.Helper()
	if !(expect < len(x.actual)) {
		x.t.Errorf("array length is expected to be longer than %d, but actual is %d", expect, len(x.actual))
	}
	return x
}

// Shorter checks if array length is shorter than expect.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Shorter(5) // Pass
//	gt.Array(t, v).Shorter(4) // Fail
func (x ArrayTest[T]) Shorter(expect int) ArrayTest[T] {
	x.t.Helper()
	if !(len(x.actual) < expect) {
		x.t.Errorf("array length is expected to be shorter than %d, but actual is %d", expect, len(x.actual))
	}
	return x
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x ArrayTest[T]) Must() ArrayTest[T] {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}

// At calls f with testing.TB and idx th elements in the array. If idx is out of range, f is not called and test will trigger error.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).At(2, func(t testing.TB, v int) {
//		gt.Value(t, v).Equal(3) // Pass
//	})
func (x ArrayTest[T]) At(idx int, f func(t testing.TB, v T)) ArrayTest[T] {
	x.t.Helper()

	if idx < 0 || len(x.actual) <= idx {
		x.t.Errorf("array length is %d, then %d is out of range", len(x.actual), idx)
	} else {
		f(x.t, x.actual[idx])
	}

	return x
}

// Any calls f with testing.TB and each elements in the array. If f returns true, Any returns immediately and test will pass. If f returns false for all elements, Any will trigger error.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).Any(func(v int) bool {
//	    return v == 3
//	}) // Pass
//	gt.Array(t, v).Any(func(v int) bool {
//	    return v == 4
//	}) // Fail
func (x ArrayTest[T]) Any(f func(v T) bool) ArrayTest[T] {
	x.t.Helper()

	for i := range x.actual {
		if f(x.actual[i]) {
			return x
		}
	}
	x.t.Errorf("no matched elements in array")

	return x
}

// All calls f with testing.TB and each elements in the array. If f returns false, All returns immediately and test will trigger error. If f returns true for all elements, All will pass.
//
//	v := []int{1, 2, 3, 5}
//	gt.Array(t, v).All(func(v int) bool {
//	    return v < 6
//	}) // Pass
//	gt.Array(t, v).All(func(v int) bool {
//	    return v < 4
//	}) // Fail
func (x ArrayTest[T]) All(f func(v T) bool) ArrayTest[T] {
	x.t.Helper()

	for i := range x.actual {
		if !f(x.actual[i]) {
			x.t.Errorf("unmatched element found in array: %+v", x.actual[i])
			return x
		}
	}
	return x
}

// Distinct checks if all elements in the array are distinct. If not, test will trigger error.
//
//	gt.Array(t, []int{1, 2, 3, 5}).Distinct() // Pass
//	gt.Array(t, []int{1, 2, 3, 2}).Distinct() // Fail
func (x ArrayTest[T]) Distinct() ArrayTest[T] {
	x.t.Helper()

	for i := range x.actual {
		for j := i + 1; j < len(x.actual); j++ {
			if EvalCompare(x.actual[i], x.actual[j]) {
				x.t.Errorf("array[%d] and array[%d] are not distinct (%+v)", i, j, x.actual[i])
				return x
			}
		}
	}

	return x
}

// MatchThen calls then function with testing.TB and first element in the array that match with match. If no element matches, MatchThen will trigger error.
//
//	v := []struct{
//	    Name string
//	    Age int
//	}{
//	    {"Alice", 20},
//	    {"Bob", 21},
//	    {"Carol", 22},
//	}
//	gt.Array(t, v).MatchThen(func(v struct{Name string, Age int}) bool {
//	    return v.Name == "Bob"
//	}, func(t testing.TB, v struct{Name string, Age int}) {
//	    gt.Value(t, v.Age).Equal(21) // Pass
//	})
//
//	gt.Array(t, v).MatchThen(func(v struct{Name string, Age int}) bool {
//	    return v.Name == "Dave"
//	}, func(t testing.TB, v struct{Name string, Age int}) {
//	    gt.Value(t, v.Age).Equal(21) // Fail
//	})
func (x ArrayTest[T]) MatchThen(match func(v T) bool, then func(t testing.TB, v T)) ArrayTest[T] {
	x.t.Helper()

	for i := range x.actual {
		if match(x.actual[i]) {
			then(x.t, x.actual[i])
			return x
		}
	}

	x.t.Errorf("no matched elements in array")
	return x
}
