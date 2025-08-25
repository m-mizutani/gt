package gt

import (
	"fmt"
	"testing"
)

type MapTest[K comparable, V any] struct {
	TestMeta
	actual map[K]V
}

// Map provides MapTest that has not only Value test methods but also key-value test
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).HasKey("blue").HasValue(5)
func Map[K comparable, V any](t testing.TB, actual map[K]V) MapTest[K, V] {
	t.Helper()
	return MapTest[K, V]{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

// M is sugar syntax of Map
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.M(t, m).HasKey("blue").HasValue(5)
func M[K comparable, V any](t testing.TB, actual map[K]V) MapTest[K, V] {
	t.Helper()
	return Map(t, actual)
}

// Describe sets a description for the test. The description will be displayed when the test fails.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).Describe("Map should contain expected key-value pairs").HasKey("blue")
func (x MapTest[K, V]) Describe(description string) MapTest[K, V] {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).Describef("Map should contain key %s with value %d", "blue", 5).HasKeyValue("blue", 5)
func (x MapTest[K, V]) Describef(format string, args ...any) MapTest[K, V] {
	x.setDescf(format, args...)
	return x
}

// Equal checks if expect completely equals given actual map.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).Equal(map[string]int{"blue": 5}) // <- Pass
//	gt.Map(t, m).Equal(map[string]int{"blue": 0}) // <- Fail
func (x MapTest[K, V]) Equal(expect map[K]V) MapTest[K, V] {
	x.t.Helper()

	if !EvalCompare(x.actual, expect) {
		msg := "maps are not matched\n" + Diff(expect, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
		return x
	}

	return x
}

// NotEqual checks if expect does not equal given actual map.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).Equal(map[string]int{ // <- Pass
//		"blue": 0,
//	})
//	gt.Map(t, m).Equal(map[string]int{ // <- Pass
//		"blue": 5,
//		"orange": 9,
//	})
func (x MapTest[K, V]) NotEqual(expect map[K]V) MapTest[K, V] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		msg := fmt.Sprintf("maps should not be matched, %+v", x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
		return x
	}

	return x
}

// EqualAt checks if actual[key] equals expect. If key is not found, test will fail.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).NotEqualAt("blue", 5)   // Pass
//	gt.Map(t, m).NotEqualAt("blue", 1)   // Fail
//	gt.Map(t, m).NotEqualAt("orange", 5) // Fail by key not found
func (x MapTest[K, V]) EqualAt(key K, expect V) MapTest[K, V] {
	x.t.Helper()

	if v, ok := x.actual[key]; !ok {
		msg := fmt.Sprintf("key '%+v' is not found in the map", key)
		x.t.Error(formatErrorMessage(x.description, msg))
	} else if !EvalCompare(v, expect) {
		msg := fmt.Sprintf("map[%+v] is expected %+v, but actual is %+v", key, expect, v)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotEqualAt checks if actual[key] equals expect. If key is not found, test will fail.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).NotEqualAt("blue", 1)   // Pass
//	gt.Map(t, m).NotEqualAt("blue", 5)   // Fail
//	gt.Map(t, m).NotEqualAt("orange", 5) // Fail by key not found
func (x MapTest[K, V]) NotEqualAt(key K, expect V) MapTest[K, V] {
	x.t.Helper()

	if v, ok := x.actual[key]; !ok {
		msg := fmt.Sprintf("key '%+v' is not found in the map", key)
		x.t.Error(formatErrorMessage(x.description, msg))
	} else if EvalCompare(v, expect) {
		msg := fmt.Sprintf("map[%+v] is expected %+v, but actual is %+v", key, expect, v)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// HasKey checks if the map has expect key.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).HasKey("blue")   // <- pass
//	gt.Map(t, m).HasKey("orange") // <- fail
func (x MapTest[K, V]) HasKey(expect K) MapTest[K, V] {
	x.t.Helper()

	if _, ok := x.actual[expect]; !ok {
		msg := fmt.Sprintf("expected to contain the key '%+v', but not got", expect)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotHasKey checks if the map does not have expect key.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).NotHasKey("orange") // <- pass
//	gt.Map(t, m).NotHasKey("blue")   // <- fail
func (x MapTest[K, V]) NotHasKey(expect K) MapTest[K, V] {
	x.t.Helper()

	if _, ok := x.actual[expect]; ok {
		msg := "expected not to contain the key, but got"
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// HasValue checks if the map has expect key.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).HasValue(5) // <- pass
//	gt.Map(t, m).HasValue(7) // <- fail
func (x MapTest[K, V]) HasValue(expect V) MapTest[K, V] {
	x.t.Helper()

	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			return x
		}
	}

	msg := fmt.Sprintf("expected to contain the value '%+v', but not got", expect)
	x.t.Error(formatErrorMessage(x.description, msg))
	return x
}

// NotHasValue checks if the map has expect key.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).NotHasValue(5) // <- fail
//	gt.Map(t, m).NotHasValue(7) // <- pass
func (x MapTest[K, V]) NotHasValue(expect V) MapTest[K, V] {
	x.t.Helper()

	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			msg := "expected not contain, but got the value"
			x.t.Error(formatErrorMessage(x.description, msg))
			break
		}
	}

	return x
}

func (x MapTest[K, V]) hasKeyValue(expectKey K, expectValue V) bool {
	x.t.Helper()

	for k := range x.actual {
		if EvalCompare(k, expectKey) && EvalCompare(x.actual[k], expectValue) {
			return true
		}
	}

	return false
}

// HasKeyValue checks if the map has expect a pair of key and value.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).HasKeyValue("blue", 5)   // <- pass
//	gt.Map(t, m).HasKeyValue("blue", 0)   // <- fail
//	gt.Map(t, m).HasKeyValue("orange", 5) // <- fail
func (x MapTest[K, V]) HasKeyValue(expectKey K, expectValue V) MapTest[K, V] {
	x.t.Helper()

	if !x.hasKeyValue(expectKey, expectValue) {
		msg := fmt.Sprintf("expected to contain (%+v, %+v), but not contain", expectKey, expectValue)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
	return x
}

// NotHasKeyValue checks if the map does not have expect a pair of key and value.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).NotHasKeyValue("blue", 5)   // <- fail
//	gt.Map(t, m).NotHasKeyValue("blue", 0)   // <- pass
//	gt.Map(t, m).NotHasKeyValue("orange", 5) // <- pass
func (x MapTest[K, V]) NotHasKeyValue(expectKey K, expectValue V) MapTest[K, V] {
	x.t.Helper()

	if x.hasKeyValue(expectKey, expectValue) {
		msg := fmt.Sprintf("expected not to contain (%+v, %+v), but contained", expectKey, expectValue)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// Length checks number of a pair of keys.
//
//	m := map[string]int{
//		"blue": 5,
//		"orange: 0,
//	}
//	gt.Map(t, m).Length(2) // <- pass
//	gt.Map(t, m).Length(0) // <- pass
func (x MapTest[K, V]) Length(expect int) MapTest[K, V] {
	x.t.Helper()
	if len(x.actual) != expect {
		msg := fmt.Sprintf("map length is expected to be %d, but actual is %d", expect, len(x.actual))
		x.t.Error(formatErrorMessage(x.description, msg))
	}
	return x
}

// Required check if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.Failed().
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).Required().
//		HasKey("blue", 0).      // <- fail
//		HasKey("blue", 5).      // <- will not be tested
//	gt.Map(t, m).HasKey("blue") // <- will not be tested
func (x MapTest[K, V]) Required() MapTest[K, V] {
	x.requiredWithMeta()
	return x
}

// At calls f with testing.TB and idx th elements in the array. If idx is out of range, f is not called and test will trigger error.
//
//	m := map[string]int{
//		"blue": 5,
//	}
//	gt.Map(t, m).At("blue", func(t testing.TB, v int) {
//		gt.Value(t, v).Equal(5) // <- pass
//	})
func (x MapTest[K, V]) At(key K, f func(t testing.TB, v V)) MapTest[K, V] {
	x.t.Helper()

	if v, ok := x.actual[key]; !ok {
		msg := fmt.Sprintf("key '%+v' is not found in the map", key)
		x.t.Error(formatErrorMessage(x.description, msg))
	} else {
		f(x.t, v)
	}

	return x
}
