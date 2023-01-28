package gt

import "testing"

// Cast tries type assertion and will error and stop test if type assertion fails
//
//	var a any = "hello"
//	b := gt.Cast[string](a) // <- Pass and set "hello" to b
//	c := gt.Cast[int](a)    // <- Fail and stop test
func Cast[T any](t testing.TB, v any) T {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
		t.FailNow()
	}
	return casted
}

func C[T any](t testing.TB, v any) T {
	return Cast[T](t, v)
}
