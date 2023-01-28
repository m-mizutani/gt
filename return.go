package gt

import "testing"

type Return1Test[T1 any] struct {
	r1  T1
	err error
}

// Return1 creates a test for returned variables (one value and one error)
//
//	f := func() (string, error) {
//		return "ok", nil
//	}
//	gt.Return1(f()).Error()               // Fail
//	gt.Return1(f()).NoError().Equal("ok") // Pass
func Return1[T1 any](r1 T1, err error) Return1Test[T1] {
	return Return1Test[T1]{
		r1:  r1,
		err: err,
	}
}

func R1[T1 any](r1 T1, err error) Return1Test[T1] {
	return Return1(r1, err)
}

// Error check if the function returned error. If error is nil, it will fail. If error is not nil, it provides ErrorTest
func (x Return1Test[T1]) Error(t testing.TB) ErrorTest {
	t.Helper()
	if x.err == nil {
		t.Errorf("got no error, but should get errored")
	}
	return Error(t, x.err)
}

// NoError check if the function returned no error. If error is not nil, it will fail. If error is nil, it provides 1st returned value.
func (x Return1Test[T1]) NoError(t testing.TB) T1 {
	t.Helper()
	if x.err != nil {
		t.Errorf("got errored, but should not get error")
	}

	return x.r1
}

type Return2Test[T1, T2 any] struct {
	r1  T1
	r2  T2
	err error
}

// Return2 creates a test for returned variables (two values and one error)
//
//	f := func() (string, int, error) {
//		return "ok", 1, nil
//	}
//	gt.Return2(f()).Error()           // Fail
//	s, i := gt.Return1(f()).NoError() // Pass
//
//	s.Equal("ok") // Pass
//	i.Equal(1)    // Pass
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

// Error check if the function returned error. If error is nil, it will fail. If error is not nil, it provides ErrorTest
func (x Return2Test[T1, T2]) Error(t testing.TB) ErrorTest {
	t.Helper()
	if x.err == nil {
		t.Errorf("got no error, but should get errored")
	}
	return Error(t, x.err)
}

// NoError check if the function returned no error. If error is not nil, it will fail. If error is nil, it provides 1st and 2nd returned value.
func (x Return2Test[T1, T2]) NoError(t testing.TB) (T1, T2) {
	t.Helper()
	if x.err != nil {
		t.Errorf("got errored, but should not get error")
	}

	return x.r1, x.r2
}

type Return3Test[T1, T2, T3 any] struct {
	r1  T1
	r2  T2
	r3  T3
	err error
}

// Return3 creates a test for returned variables (three values and one error)
//
//	f := func() (string, int, bool, error) {
//		return "ok", 1, true, nil
//	}
//	gt.Return3(f()).Error()           // Fail
//	s, i, b := gt.Return1(f()).NoError() // Pass
//
//	s.Equal("ok") // Pass
//	i.Equal(1)    // Pass
//	b.Equal(true) // Pass
func Return3[T1, T2, T3 any](r1 T1, r2 T2, r3 T3, err error) Return3Test[T1, T2, T3] {
	return Return3Test[T1, T2, T3]{
		r1:  r1,
		r2:  r2,
		r3:  r3,
		err: err,
	}
}

func R3[T1, T2, T3 any](r1 T1, r2 T2, r3 T3, err error) Return3Test[T1, T2, T3] {
	return Return3(r1, r2, r3, err)
}

// Error check if the function returned error. If error is nil, it will fail. If error is not nil, it provides ErrorTest
func (x Return3Test[T1, T2, T3]) Error(t testing.TB) ErrorTest {
	t.Helper()
	if x.err == nil {
		t.Errorf("got no error, but should get errored")
	}
	return Error(t, x.err)
}

// NoError check if the function returned no error. If error is not nil, it will fail. If error is nil, it provides 1st, 2nd and 3rd returned value.
func (x Return3Test[T1, T2, T3]) NoError(t testing.TB) (T1, T2, T3) {
	t.Helper()
	if x.err != nil {
		t.Errorf("got errored, but should not get error")
	}

	return x.r1, x.r2, x.r3
}
