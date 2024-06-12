package gt

import (
	"testing"
	"time"
)

type TimeTest struct {
	actual time.Time
	t      testing.TB
}

// Time provides TimeTest that has basic comparison methods
func Time(t testing.TB, actual time.Time) TimeTest {
	t.Helper()
	return TimeTest{
		actual: actual,
		t:      t,
	}
}

// T is sugar syntax of Time
func T(t testing.TB, actual time.Time) TimeTest {
	t.Helper()
	return Time(t, actual)
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x TimeTest) Must() TimeTest {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}

// Equal check if actual equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x TimeTest) Equal(expect time.Time) TimeTest {
	x.t.Helper()
	if !x.actual.Equal(expect) {
		x.t.Error("values are not matched\n" + Diff(expect, x.actual))
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x TimeTest) NotEqual(expect time.Time) TimeTest {
	x.t.Helper()
	if x.actual.Equal(expect) {
		x.t.Errorf("values should not be matched, %+v", x.actual)
	}

	return x
}

// Before check if actual is before than expect.
func (x TimeTest) Before(expect time.Time) TimeTest {
	x.t.Helper()
	if !x.actual.Before(expect) {
		x.t.Errorf("value should be before than %+v, %+v", expect, x.actual)
	}

	return x
}

// After check if actual is after than expect.
func (x TimeTest) After(expect time.Time) TimeTest {
	x.t.Helper()
	if !x.actual.After(expect) {
		x.t.Errorf("value should be after than %+v, %+v", expect, x.actual)
	}

	return x
}
