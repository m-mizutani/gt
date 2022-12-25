package gt

import "testing"

func Cast[T any](t testing.TB, v any) T {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
	}
	return casted
}

func MustCast[T any](t testing.TB, v any) T {
	casted, ok := v.(T)
	if !ok {
		var a T
		t.Errorf("expected %T, but can not cast", a)
		t.FailNow()
	}
	return casted
}
