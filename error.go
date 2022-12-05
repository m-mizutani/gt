package gt

import (
	"errors"
	"testing"
)

type ErrorTest struct {
	t      testing.TB
	actual error
}

func Error(t testing.TB, actual error) ErrorTest {
	t.Helper()
	return ErrorTest{
		t:      t,
		actual: actual,
	}
}

func (x ErrorTest) Passed() ErrorTest {
	x.t.Helper()
	if x.actual != nil {
		x.t.Error("expected no error, but got error")
	}
	return x
}

func (x ErrorTest) Failed() ErrorTest {
	x.t.Helper()
	if x.actual == nil {
		x.t.Error("expected error, but got no error")
	}
	return x
}

func (x ErrorTest) Is(expected error) {
	x.t.Helper()
	if !errors.Is(x.actual, expected) {
		x.t.Errorf("expected %T, but got %T", x.actual, expected)
	}
}
