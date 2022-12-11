package gt_test

import (
	"errors"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestError(t *testing.T) {
	err := errors.New("test")
	cnt := newRecorder()
	gt.Error(cnt, err).Failed()

	if cnt.errs > 0 {
		t.Error("error test has unexpected result")
	}
}
