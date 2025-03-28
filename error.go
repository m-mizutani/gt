package gt

import (
	"errors"
	"strings"
	"testing"
)

type ErrorTest struct {
	t      testing.TB
	actual error
}

// Value provides ErrorTest that is specialized for error testing
func Error(t testing.TB, actual error) ErrorTest {
	t.Helper()
	if actual == nil {
		t.Errorf("expected error, but got no error")
	}
	return ErrorTest{
		t:      t,
		actual: actual,
	}
}

// Required checks if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x ErrorTest) Required() ErrorTest {
	x.t.Helper()
	if x.actual == nil {
		x.t.FailNow()
	}
	return x
}

// Is checks error object equality by errors.Is() function.
func (x ErrorTest) Is(expected error) {
	x.t.Helper()
	if x.actual != nil && !errors.Is(x.actual, expected) {
		x.t.Errorf("expected %T, but not got from %T", expected, x.actual)
	}
}

// IsNot checks error object not-equality by errors.Is() function.
func (x ErrorTest) IsNot(expected error) {
	x.t.Helper()
	if x.actual != nil && errors.Is(x.actual, expected) {
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

type NoErrorTest struct {
	t      testing.TB
	actual error
}

// NoError checks if error does not occur.
func NoError(t testing.TB, actual error) NoErrorTest {
	t.Helper()
	if actual != nil {
		t.Errorf("expected no error, but got %+v", actual)
	}
	return NoErrorTest{
		t:      t,
		actual: actual,
	}
}

func (x NoErrorTest) Required() {
	x.t.Helper()
	if x.actual != nil {
		x.t.FailNow()
	}
}

// Contains checks if the error message contains the expected substring.
func (x ErrorTest) Contains(substr string) {
	x.t.Helper()
	if x.actual == nil {
		x.t.Errorf("expected error containing %q, but got no error", substr)
		return
	}
	if msg := x.actual.Error(); !strings.Contains(msg, substr) {
		x.t.Errorf("expected error message containing %q, but got %q", substr, msg)
	}
}
