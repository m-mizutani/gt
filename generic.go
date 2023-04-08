package gt

import (
	"testing"
)

func Equal[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	if !EvalCompare(actual, expected) {
		t.Error("values should be matched, but not match\n" + Diff(expected, actual))
	}
}

func NotEqual[T any](t testing.TB, actual T, expected T) {
	t.Helper()
	if EvalCompare(actual, expected) {
		t.Error("values should not be matched, but match\n" + Diff(expected, actual))
	}
}
