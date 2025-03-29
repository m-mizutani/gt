package gt

import "testing"

type BoolTest struct {
	t      testing.TB
	actual bool
}

func (x BoolTest) Required() BoolTest {
	x.t.Helper()
	required(x.t)
	return x
}

func Bool(t testing.TB, actual bool) BoolTest {
	return BoolTest{
		t:      t,
		actual: actual,
	}
}

func B(t testing.TB, actual bool) BoolTest {
	return Bool(t, actual)
}

func (x BoolTest) True() BoolTest {
	x.t.Helper()
	if !x.actual {
		x.t.Error("expected true, but false")
	}
	return x
}

func (x BoolTest) False() BoolTest {
	x.t.Helper()
	if x.actual {
		x.t.Error("expected false, but true")
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
