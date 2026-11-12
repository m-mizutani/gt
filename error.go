package gt

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/m-mizutani/goerr/v2"
)

type ErrorTest struct {
	TestMeta
	actual error
}

// Value provides ErrorTest that is specialized for error testing
func Error(t testing.TB, actual error) ErrorTest {
	t.Helper()
	if actual == nil {
		t.Errorf("expected error, but got no error")
	}
	return ErrorTest{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

// Required checks if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.FailNow().
// Describe sets a description for the test. The description will be displayed when the test fails.
func (x ErrorTest) Describe(description string) ErrorTest {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
func (x ErrorTest) Describef(format string, args ...any) ErrorTest {
	x.setDescf(format, args...)
	return x
}

func (x ErrorTest) Required() ErrorTest {
	x.requiredWithMeta()
	return x
}

// Is checks error object equality by errors.Is() function.
func (x ErrorTest) Is(expected error) {
	x.t.Helper()
	if x.actual != nil && !errors.Is(x.actual, expected) {
		msg := fmt.Sprintf("expected error is %q (%T), but got: %+v", expected, expected, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
}

// IsNot checks error object not-equality by errors.Is() function.
func (x ErrorTest) IsNot(expected error) {
	x.t.Helper()
	if x.actual != nil && errors.Is(x.actual, expected) {
		msg := fmt.Sprintf("expected error is not %q (%T), but got: %+v", expected, expected, x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
}

// HasTag checks that the error carries the goerr tag, by goerr.HasTag(). A tag is given as a
// fmt.Stringer because goerr does not export its tag type; a value returned by goerr.NewTag
// satisfies it. It does nothing for a nil error, which gt.Error has already reported.
//
//	var TagNotFound = goerr.NewTag("not_found")
//	gt.Error(t, err).HasTag(TagNotFound)
func (x ErrorTest) HasTag(tag fmt.Stringer) {
	x.t.Helper()
	if x.actual != nil && !hasTag(x.actual, tag) {
		msg := fmt.Sprintf("expected error has tag %q, but tags are %v: %+v", tag.String(), goerr.Tags(x.actual), x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
}

// NotHasTag checks that the error does not carry the goerr tag, by goerr.HasTag(). It does
// nothing for a nil error, which gt.Error has already reported.
//
//	gt.Error(t, err).NotHasTag(TagNotFound)
func (x ErrorTest) NotHasTag(tag fmt.Stringer) {
	x.t.Helper()
	if x.actual != nil && hasTag(x.actual, tag) {
		msg := fmt.Sprintf("expected error does not have tag %q, but got: %+v", tag.String(), x.actual)
		x.t.Error(formatErrorMessage(x.description, msg))
	}
}

// hasTag rebuilds the tag from its string: goerr keys tags by value, so a tag made from the
// same string is the same tag.
func hasTag(err error, tag fmt.Stringer) bool {
	return goerr.HasTag(err, goerr.NewTag(tag.String()))
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
		msg := fmt.Sprintf("expected error containing %q, but got no error", substr)
		x.t.Error(formatErrorMessage(x.description, msg))
		return
	}
	if msg := x.actual.Error(); !strings.Contains(msg, substr) {
		msgText := fmt.Sprintf("expected error message containing %q, but got %q", substr, msg)
		x.t.Error(formatErrorMessage(x.description, msgText))
	}
}
