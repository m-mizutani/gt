package gt

import "testing"

type BoolTest struct {
	TestMeta
	actual bool
}

func (x BoolTest) Required() BoolTest {
	x.requiredWithMeta()
	return x
}

func Bool(t testing.TB, actual bool) BoolTest {
	t.Helper()
	return BoolTest{
		TestMeta: TestMeta{t: t},
		actual:   actual,
	}
}

func B(t testing.TB, actual bool) BoolTest {
	return Bool(t, actual)
}

// Describe sets a description for the test. The description will be displayed when the test fails.
func (x BoolTest) Describe(description string) BoolTest {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
func (x BoolTest) Describef(format string, args ...any) BoolTest {
	x.setDescf(format, args...)
	return x
}

func (x BoolTest) True() BoolTest {
	x.t.Helper()
	if !x.actual {
		msg := "expected true, but false"
		x.t.Error(formatErrorMessage(x.description, msg))
	}
	return x
}

func (x BoolTest) False() BoolTest {
	x.t.Helper()
	if x.actual {
		msg := "expected false, but true"
		x.t.Error(formatErrorMessage(x.description, msg))
	}
	return x
}

func True(t testing.TB, actual bool) BoolTest {
	t.Helper()
	return Bool(t, actual).True()
}

func False(t testing.TB, actual bool) BoolTest {
	t.Helper()
	return Bool(t, actual).False()
}
