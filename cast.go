package gt

import "testing"

// Cast tries type assertion and will error if type assertion fails
//
//	var a any = "hello"
//	b := gt.Cast[string](a) // <- Pass and set "hello" to b
//	c := gt.Cast[int](a)    // <- Fail and set empty to c
func Cast[T any](t testing.TB, v any) T {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
	}
	return casted
}

// MustCast tries type assertion and will error and stop test if type assertion fails
//
//	var a any = "hello"
//	b := gt.MustCast[string](a) // <- Pass and set "hello" to b
//	c := gt.MustCast[int](a)    // <- Fail and stop test
func MustCast[T any](t testing.TB, v any) T {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
		t.FailNow()
	}
	return casted
}
