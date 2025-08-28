package gt_test

import (
	"errors"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestReturn(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		r := newRecorder()

		doSomeRequest := func() (string, error) {
			return "ok", nil
		}
		gt.V(r, gt.Return1(doSomeRequest()).NoError(r)).Equal("ok").Required()

		if r.errs != 0 {
			t.Error("should not fail")
		}
	})

	t.Run("success", func(t *testing.T) {
		r := newRecorder()

		doSomeRequest := func() (string, error) {
			return "ng", errors.New("test")
		}
		gt.Return1(doSomeRequest()).Error(r).Required()

		if r.errs != 0 {
			t.Error("should not fail")
		}
	})

	t.Run("example", func(t *testing.T) {

		goodFunc := func() (string, error) {
			return "ok", nil
		}
		badFunc := func() (string, error) {
			return "ng", errors.New("test")
		}

		// Check if getting no error and will get value
		gt.Value(t, gt.Return1(goodFunc()).NoError(t)).Equal("ok").Required()

		// Check if getting error
		gt.Return1(badFunc()).Error(t).Required()

	})
}
