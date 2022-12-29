package gt

import "testing"

type CastTest[T any] struct {
	t testing.TB
	v T
}

// Cast tries type assertion and will error if type assertion fails
//
//	var a any = "hello"
//	b := gt.Cast[string](a) // <- Pass and set "hello" to b
//	c := gt.Cast[int](a)    // <- Fail and set empty to c
func Cast[T any](t testing.TB, v any) CastTest[T] {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
	}
	return CastTest[T]{
		t: t,
		v: casted,
	}
}

// MustCast tries type assertion and will error and stop test if type assertion fails
//
//	var a any = "hello"
//	b := gt.MustCast[string](a) // <- Pass and set "hello" to b
//	c := gt.MustCast[int](a)    // <- Fail and stop test
func MustCast[T any](t testing.TB, v any) CastTest[T] {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
		t.FailNow()
	}
	return CastTest[T]{
		t: t,
		v: casted,
	}
}

// Nil will fail if the result of cast is not nil
func (x CastTest[T]) Nil() {
	if !EvalIsNil(x.v) {
		var a T
		x.t.Errorf("expected %T is nil, but actual is not nil", a)
	}
}

// NotNil will fail if the result of cast is nil
func (x CastTest[T]) NotNil() T {
	if EvalIsNil(x.v) {
		var a T
		x.t.Errorf("expected %T is nil, but actual is not nil", a)
	}

	return x.v
}
