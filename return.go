package gt

import "testing"

type Return1Test[T1 any] struct {
	r1  T1
	err error
}

func Return1[T1 any](r1 T1, err error) Return1Test[T1] {
	return Return1Test[T1]{
		r1:  r1,
		err: err,
	}
}

func R1[T1 any](r1 T1, err error) Return1Test[T1] {
	return Return1(r1, err)
}

func (x Return1Test[T1]) Error(t testing.TB) {
	t.Helper()
	if x.err == nil {
		t.Errorf("got no error, but should get errored")
	}
}

func (x Return1Test[T1]) NoError(t testing.TB) ValueTest[T1] {
	t.Helper()
	if x.err != nil {
		t.Errorf("got errored, but should not get error")
	}

	return Value(t, x.r1)
}

type Return2Test[T1, T2 any] struct {
	r1  T1
	r2  T2
	err error
}

func Return2[T1, T2 any](r1 T1, r2 T2, err error) Return2Test[T1, T2] {
	return Return2Test[T1, T2]{
		r1:  r1,
		r2:  r2,
		err: err,
	}
}

func R2[T1, T2 any](r1 T1, r2 T2, err error) Return2Test[T1, T2] {
	return Return2(r1, r2, err)
}

func (x Return2Test[T1, T2]) Error(t testing.TB) {
	t.Helper()
	if x.err == nil {
		t.Errorf("got no error, but should get errored")
	}
}

func (x Return2Test[T1, T2]) NoError(t testing.TB) (ValueTest[T1], ValueTest[T2]) {
	t.Helper()
	if x.err != nil {
		t.Errorf("got errored, but should not get error")
	}

	return Value(t, x.r1), Value(t, x.r2)
}
