package gt

import "testing"

type MapTest[K comparable, V any] struct {
	actual map[K]V
	t      testing.TB
}

func Map[K comparable, V any](t testing.TB, actual map[K]V) MapTest[K, V] {
	t.Helper()
	return MapTest[K, V]{
		actual: actual,
		t:      t,
	}
}

func (x MapTest[K, V]) Equal(expect map[K]V) MapTest[K, V] {
	x.t.Helper()

	if !EvalCompare(x.actual, expect) {
		x.t.Error("expected equals, but not matched")
		return x
	}

	return x
}

func (x MapTest[K, V]) NotEqual(expect map[K]V) MapTest[K, V] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		x.t.Error("expected not equals, but matched")
		return x
	}

	return x
}

func (x MapTest[K, V]) HasKey(expect K) MapTest[K, V] {
	x.t.Helper()

	if _, ok := x.actual[expect]; !ok {
		x.t.Error("expected contain the key, but not got")
	}

	return x
}

func (x MapTest[K, V]) NotHaveKey(expect K) MapTest[K, V] {
	x.t.Helper()

	if _, ok := x.actual[expect]; ok {
		x.t.Error("expected not contain the key, but got")
	}

	return x
}

func (x MapTest[K, V]) Contain(expect V) MapTest[K, V] {
	x.t.Helper()

	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			return x
		}
	}

	x.t.Error("expected contain the value, but not got")
	return x
}

func (x MapTest[K, V]) NotContain(expect V) MapTest[K, V] {
	x.t.Helper()

	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			x.t.Error("expected not contain, but got the value")
			break
		}
	}

	return x
}

func (x MapTest[K, V]) Length(expect int) MapTest[K, V] {
	x.t.Helper()
	if len(x.actual) != expect {
		x.t.Error("got non expected length")
	}
	return x
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.Failed().
func (x MapTest[K, V]) Must() MapTest[K, V] {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}
