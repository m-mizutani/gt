package gt

import (
	"errors"
	"testing"
)

type errorTest struct {
	t      testing.TB
	actual error
}

// Value provides errorTest that is specialized for error testing
func Error(t testing.TB, actual error) errorTest {
	t.Helper()
	return errorTest{
		t:      t,
		actual: actual,
	}
}

// Pass checks if error is nil
func (x errorTest) Pass() errorTest {
	x.t.Helper()
	if x.actual != nil {
		x.t.Errorf("expected no error, but got error: %v", x.actual)
	}
	return x
}

// Fail checks if error is not nil
func (x errorTest) Fail() errorTest {
	x.t.Helper()
	if x.actual == nil {
		x.t.Error("expected error, but no error")
	}
	return x
}

// Must checks if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x errorTest) Must() errorTest {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}

// Is checks error object equality by errors.Is() function.
func (x errorTest) Is(expected error) {
	x.t.Helper()
	if !errors.Is(x.actual, expected) {
		x.t.Errorf("expected %T, but not got from %T", expected, x.actual)
	}
}

// IsNot checks error object not-equality by errors.Is() function.
func (x errorTest) IsNot(expected error) {
	x.t.Helper()
	if errors.Is(x.actual, expected) {
		x.t.Errorf("not expected %T, but got from %T", expected, x.actual)
	}
}

// ErrorAs checks error type by errors.As() function. If type check passed, callback will be invoked and given extracted error by errors.As.
func ErrorAs[T any](t testing.TB, actual error, callback func(expect *T)) {
	t.Helper()
	tgt := new(T)
	if errors.As(actual, tgt) {
		callback(tgt)
	} else {
		t.Errorf("expected %T, but got %T", tgt, actual)
	}
}
