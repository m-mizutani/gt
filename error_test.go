package gt_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestError(t *testing.T) {
	err := errors.New("test")
	cnt := newRecorder()
	gt.Error(cnt, err)

	if cnt.errs > 0 {
		t.Error("error test has unexpected result")
	}
}

type testError struct{ N int }

func (x testError) Error() string { return "test error" }

func TestErrorAs(t *testing.T) {
	testErr := testError{N: 5}
	err := fmt.Errorf("run error: %w", testErr)

	var called int
	gt.ErrorAs(t, err, func(tgt *testError) {
		called++
		if tgt.N != 5 {
			t.Errorf("testError.N must be 5, but %+v", tgt.N)
		}
	})
	if called != 1 {
		t.Errorf("callback must be called once, but %+v times", called)
	}
}
