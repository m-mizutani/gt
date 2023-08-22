package gt

import (
	"regexp"
	"strings"
	"testing"
)

type StringTest struct {
	actual string
	t      testing.TB
}

// String provides StringTest that has basic comparison methods
func String(t testing.TB, actual string) StringTest {
	t.Helper()
	return StringTest{
		actual: actual,
		t:      t,
	}
}

// S is sugar syntax of String
func S(t testing.TB, actual string) StringTest {
	t.Helper()
	return String(t, actual)
}

// Must check if error has occurred in previous test. If errors will occur in following test, it immediately stop test by t.FailNow().
func (x StringTest) Must() StringTest {
	x.t.Helper()
	x.t = newErrorWithFail(x.t)
	return x
}

// Equal check if actual equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x StringTest) Equal(expect string) StringTest {
	x.t.Helper()
	if !EvalCompare(x.actual, expect) {
		x.t.Error("values are not matched\n" + Diff(expect, x.actual))
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x StringTest) NotEqual(expect string) StringTest {
	x.t.Helper()
	if EvalCompare(x.actual, expect) {
		x.t.Errorf("values should not be matched, %+v", x.actual)
	}

	return x
}

// IsEmpty check if actual is empty.
func (x StringTest) IsEmpty() StringTest {
	x.t.Helper()
	if len(x.actual) > 0 {
		x.t.Errorf("value should be empty, %+v", x.actual)
	}

	return x
}

// IsNotEmpty check if actual is not empty.
func (x StringTest) IsNotEmpty() StringTest {
	x.t.Helper()
	if len(x.actual) == 0 {
		x.t.Errorf("value should not be empty")
	}

	return x
}

// Contains check if actual contains expected.
func (x StringTest) Contains(sub string) StringTest {
	x.t.Helper()
	if !strings.Contains(x.actual, sub) {
		x.t.Errorf("value should contain %+v, %+v", sub, x.actual)
	}

	return x
}

// NotContains check if actual does not contain expected.
func (x StringTest) NotContains(sub string) StringTest {
	x.t.Helper()
	if strings.Contains(x.actual, sub) {
		x.t.Errorf("value should not contain %+v, %+v", sub, x.actual)
	}

	return x
}

// HasPrefix check if actual has prefix expected.
func (x StringTest) HasPrefix(prefix string) StringTest {
	x.t.Helper()
	if !strings.HasPrefix(x.actual, prefix) {
		x.t.Errorf("value should have prefix %+v, %+v", prefix, x.actual)
	}

	return x
}

// NotHasPrefix check if actual does not have prefix expected.
func (x StringTest) NotHasPrefix(prefix string) StringTest {
	x.t.Helper()
	if strings.HasPrefix(x.actual, prefix) {
		x.t.Errorf("value should not have prefix %+v, %+v", prefix, x.actual)
	}

	return x
}

// HasSuffix check if actual has suffix expected.
func (x StringTest) HasSuffix(suffix string) StringTest {
	x.t.Helper()
	if !strings.HasSuffix(x.actual, suffix) {
		x.t.Errorf("value should have suffix %+v, %+v", suffix, x.actual)
	}

	return x
}

// NotHasSuffix check if actual does not have suffix expected.
func (x StringTest) NotHasSuffix(suffix string) StringTest {
	x.t.Helper()
	if strings.HasSuffix(x.actual, suffix) {
		x.t.Errorf("value should not have suffix %+v, %+v", suffix, x.actual)
	}

	return x
}

// Match check if actual matches with expected regular expression.
func (x StringTest) Match(pattern string) StringTest {
	x.t.Helper()
	if !x.match(pattern) {
		x.t.Errorf("value should match '%+v', %+v", pattern, x.actual)
	}

	return x
}

// NotMatch check if actual matches with expected regular expression.
func (x StringTest) NotMatch(pattern string) StringTest {
	x.t.Helper()
	if x.match(pattern) {
		x.t.Errorf("value should match '%+v', %+v", pattern, x.actual)
	}

	return x
}

func (x StringTest) match(pattern string) bool {
	x.t.Helper()
	ptn, err := regexp.Compile(pattern)
	if err != nil {
		x.t.Errorf("invalid pattern, %+v", pattern)
		x.t.FailNow()
		return false
	}

	return ptn.MatchString(x.actual)
}
