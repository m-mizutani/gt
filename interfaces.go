package gt

import "testing"

type errorWithFail struct {
	testing.TB
}

func newErrorWithFail(t testing.TB) *errorWithFail {
	t.Helper()
	return &errorWithFail{
		TB: t,
	}
}

func (x *errorWithFail) Error(args ...any) {
	x.TB.Helper()
	x.TB.Error(args...)
	x.TB.FailNow()
}

func (x *errorWithFail) Errorf(format string, args ...any) {
	x.TB.Helper()
	x.TB.Errorf(format, args...)
	x.TB.FailNow()
}
