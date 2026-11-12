package gt

import (
	"fmt"
	"testing"
	"time"
)

type TimeTest struct {
	TestMeta
	actual time.Time
}

// Time provides TimeTest that is specialized for time.Time testing
func Time(t testing.TB, actual time.Time) TimeTest {
	t.Helper()
	return TimeTest{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

// Describe sets a description for the test. The description will be displayed when the test fails.
//
//	gt.Time(t, got.CreatedAt).Describe("CreatedAt should be the time it was stored").Near(now, time.Second)
func (x TimeTest) Describe(description string) TimeTest {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
//
//	gt.Time(t, got.CreatedAt).Describef("CreatedAt of record %d", id).Near(now, time.Second)
func (x TimeTest) Describef(format string, args ...any) TimeTest {
	x.setDescf(format, args...)
	return x
}

// Near checks that actual is less than tolerance away from expect, earlier or later.
//
//	now := time.Now()
//	gt.Time(t, now.Add(500*time.Millisecond)).Near(now, time.Second) // Pass
//	gt.Time(t, now.Add(-2*time.Second)).Near(now, time.Second)       // Fail
func (x TimeTest) Near(expect time.Time, tolerance time.Duration) TimeTest {
	x.t.Helper()
	if diff := x.actual.Sub(expect); diff.Abs() >= tolerance {
		msg := fmt.Sprintf("time is not near\nexpected: %s (tolerance %s)\nactual:   %s (off by %s)",
			expect.Format(time.RFC3339Nano), tolerance, x.actual.Format(time.RFC3339Nano), diff)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
	return x
}

// Required check if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.Failed().
func (x TimeTest) Required() TimeTest {
	x.requiredWithMeta()
	return x
}
