package gt

import (
	"fmt"
	"regexp"
	"strings"
	"testing"
)

type StringTest struct {
	TestMeta
	actual string
}

// String provides StringTest that has basic comparison methods
func String(t testing.TB, actual string) StringTest {
	t.Helper()
	return StringTest{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

// S is sugar syntax of String
func S(t testing.TB, actual string) StringTest {
	t.Helper()
	return String(t, actual)
}

// Describe sets a description for the test. The description will be displayed when the test fails.
//
//	gt.String(t, actual).Describe("Username should match expected value").Equal(expected)
func (x StringTest) Describe(description string) StringTest {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
//
//	gt.String(t, actual).Describef("Username should match %s for user %s", expected, "Alice").Equal(expected)
func (x StringTest) Describef(format string, args ...any) StringTest {
	x.setDescf(format, args...)
	return x
}

// Required check if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.FailNow().
func (x StringTest) Required() StringTest {
	x.requiredWithMeta()
	return x
}

// Equal check if actual equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x StringTest) Equal(expect string) StringTest {
	x.t.Helper()
	if !EvalCompare(x.actual, expect) {
		msg := "values are not matched\n" + Diff(expect, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotEqual check if actual does not equals with expect. Default evaluation function uses reflect.DeepEqual.
func (x StringTest) NotEqual(expect string) StringTest {
	x.t.Helper()
	if EvalCompare(x.actual, expect) {
		msg := fmt.Sprintf("values should not be matched, %+v", x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// IsEmpty check if actual is empty.
func (x StringTest) IsEmpty() StringTest {
	x.t.Helper()
	if len(x.actual) > 0 {
		msg := fmt.Sprintf("value should be empty, %+v", x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// IsNotEmpty check if actual is not empty.
func (x StringTest) IsNotEmpty() StringTest {
	x.t.Helper()
	if len(x.actual) == 0 {
		msg := "value should not be empty"
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// Contains check if actual contains expected.
func (x StringTest) Contains(sub string) StringTest {
	x.t.Helper()
	if !strings.Contains(x.actual, sub) {
		msg := fmt.Sprintf("value should contain %+v, %+v", sub, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotContains check if actual does not contain expected.
func (x StringTest) NotContains(sub string) StringTest {
	x.t.Helper()
	if strings.Contains(x.actual, sub) {
		msg := fmt.Sprintf("value should not contain %+v, %+v", sub, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// ContainsAny checks if actual contains any of the expected substrings.
// The test passes if actual contains at least one of the provided substrings.
func (x StringTest) ContainsAny(substrs ...string) StringTest {
	x.t.Helper()

	for _, sub := range substrs {
		if strings.Contains(x.actual, sub) {
			return x
		}
	}

	msg := fmt.Sprintf("value should contain any of %+v, but got: %+v", substrs, x.actual)
	x.t.Error(formatErrorMessage(x.description, msg))

	return x
}

// ContainsNone checks if actual contains none of the expected substrings.
// The test passes if actual does not contain any of the provided substrings.
func (x StringTest) ContainsNone(substrs ...string) StringTest {
	x.t.Helper()

	for _, sub := range substrs {
		if strings.Contains(x.actual, sub) {
			msg := fmt.Sprintf("value should not contain any of %+v, but contains %+v in: %+v", substrs, sub, x.actual)
			x.t.Error(formatErrorMessage(x.description, msg))
			return x
		}
	}

	return x
}

// HasPrefix check if actual has prefix expected.
func (x StringTest) HasPrefix(prefix string) StringTest {
	x.t.Helper()
	if !strings.HasPrefix(x.actual, prefix) {
		msg := fmt.Sprintf("value should have prefix %+v, %+v", prefix, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotHasPrefix check if actual does not have prefix expected.
func (x StringTest) NotHasPrefix(prefix string) StringTest {
	x.t.Helper()
	if strings.HasPrefix(x.actual, prefix) {
		msg := fmt.Sprintf("value should not have prefix %+v, %+v", prefix, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// HasSuffix check if actual has suffix expected.
func (x StringTest) HasSuffix(suffix string) StringTest {
	x.t.Helper()
	if !strings.HasSuffix(x.actual, suffix) {
		msg := fmt.Sprintf("value should have suffix %+v, %+v", suffix, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotHasSuffix check if actual does not have suffix expected.
func (x StringTest) NotHasSuffix(suffix string) StringTest {
	x.t.Helper()
	if strings.HasSuffix(x.actual, suffix) {
		msg := fmt.Sprintf("value should not have suffix %+v, %+v", suffix, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// Match check if actual matches with expected regular expression.
func (x StringTest) Match(pattern string) StringTest {
	x.t.Helper()
	if !x.match(pattern) {
		msg := fmt.Sprintf("value should match '%+v', %+v", pattern, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotMatch check if actual matches with expected regular expression.
func (x StringTest) NotMatch(pattern string) StringTest {
	x.t.Helper()
	if x.match(pattern) {
		msg := fmt.Sprintf("value should not match '%+v', %+v", pattern, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

func (x StringTest) match(pattern string) bool {
	x.t.Helper()
	ptn, err := regexp.Compile(pattern)
	if err != nil {
		msg := fmt.Sprintf("invalid pattern, %+v", pattern)
		x.t.Error(formatErrorMessage(x.description, msg))
		x.t.FailNow()
		return false
	}

	return ptn.MatchString(x.actual)
}
