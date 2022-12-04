package gt_test

import (
	"testing"
)

type counter struct {
	testing.TB

	errs  int
	fails int
}

func newCounter() *counter {
	return &counter{}
}

func (x *counter) Helper() {}

func (x *counter) Error(args ...any) {
	if x.fails == 0 {
		x.errs++
	}
}

func (x *counter) Errorf(fmt string, args ...any) {
	if x.fails == 0 {
		x.errs++
	}
}

func (x *counter) FailNow() {
	x.errs++
}
