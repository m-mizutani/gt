package gt_test

import (
	"fmt"
	"strings"
	"testing"
)

type recorder struct {
	testing.TB

	errs  int
	fails int
	msgs  []string
}

func newRecorder() *recorder {
	return &recorder{}
}

func (x *recorder) Helper() {}

func (x *recorder) Error(args ...any) {
	var argv []string
	for _, arg := range args {
		argv = append(argv, fmt.Sprintf("%+v", arg))
	}
	if x.fails == 0 {
		x.msgs = append(x.msgs, strings.Join(argv, " "))
		x.errs++
	}
}

func (x *recorder) Errorf(format string, args ...any) {
	if x.fails == 0 {
		x.msgs = append(x.msgs, fmt.Sprintf(format, args...))
		x.errs++
	}
}

func (x *recorder) FailNow() {
	x.fails++
}

func (x *recorder) Failed() bool {
	return x.errs > 0
}
