package gt

import "testing"

type BoolTest struct {
	t      testing.TB
	actual bool
}

func Bool(t *testing.T, actual bool) BoolTest {
	return BoolTest{
		t:      t,
		actual: actual,
	}
}

func B(t *testing.T, actual bool) BoolTest {
	return Bool(t, actual)
}

func (x BoolTest) True() BoolTest {
	if !x.actual {
		x.t.Error("expected true, but false")
	}
	return x
}

func (x BoolTest) False() BoolTest {
	if x.actual {
		x.t.Error("expected false, but true")
	}
	return x
}
