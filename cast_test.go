package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestCast(t *testing.T) {
	type testData struct {
		Name string
	}
	var s any = &testData{
		Name: "blue",
	}

	t.Run("good case", func(t *testing.T) {
		r := newRecorder()
		v := gt.Cast[*testData](r, s)
		if v.Name != "blue" {
			t.Error("Name is not matched")
		}
		if r.errs > 0 {
			t.Error("should not error, but occurred")
		}
	})

	t.Run("bad case", func(t *testing.T) {
		r := newRecorder()
		v := gt.Cast[testData](r, s)
		if v.Name != "" {
			t.Error("Name is not matched")
		}
		if r.errs == 0 {
			t.Error("should error, but not occurred")
		}
		if r.fails == 0 {
			t.Error("should not fail, but occurred")
		}
	})
}

func TestCast_Nil(t *testing.T) {
	t.Run("value is not nil", func(t *testing.T) {
		r := newRecorder()
		var s any = "test"

		v := gt.Cast[string](r, s)
		if v != "test" {
			t.Error("v is not matched")
		}
		if r.errs > 0 {
			t.Error("v should not be nil")
		}
	})
}
