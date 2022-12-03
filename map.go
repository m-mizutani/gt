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
		x.t.Error("not equal")
		return x
	}

	return x
}

func (x MapTest[K, V]) NotEqual(expect map[K]V) MapTest[K, V] {
	x.t.Helper()

	if EvalCompare(x.actual, expect) {
		x.t.Error("equal")
		return x
	}

	return x
}

func (x MapTest[K, V]) HasKey(expect K) MapTest[K, V] {
	if _, ok := x.actual[expect]; !ok {
		x.t.Error("does not have key")
	}

	return x
}

func (x MapTest[K, V]) NotHaveKey(expect K) MapTest[K, V] {
	if _, ok := x.actual[expect]; ok {
		x.t.Error("has key")
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

	x.t.Error("not contains")
	return x
}

func (x MapTest[K, V]) NotContain(expect V) MapTest[K, V] {
	x.t.Helper()
	for i := range x.actual {
		if EvalCompare(x.actual[i], expect) {
			x.t.Error("contains")
			return x
		}
	}

	return x
}

func (x MapTest[K, V]) Length(expect int) MapTest[K, V] {
	x.t.Helper()
	if len(x.actual) != expect {
		x.t.Error("not contains")
	}
	return x
}

func (x MapTest[K, V]) Required() MapTest[K, V] {
	x.t.Helper()
	if x.t.Failed() {
		x.t.FailNow()
	}
	return x
}
