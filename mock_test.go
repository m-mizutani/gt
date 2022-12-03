package gt_test

import "testing"

type errCounter struct {
	testing.TB

	errCount  int
	failCount int
}

func newErrCounter() *errCounter {
	return &errCounter{}
}

func (x *errCounter) Helper() {}

func (x *errCounter) Error(args ...any) {
	if x.failCount == 0 {
		x.errCount++
	}
}

func (x *errCounter) FailNow() {
	x.failCount++
}
