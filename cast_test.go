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
		v := gt.Cast[*testData](r, s).NotNil()
		if v.Name != "blue" {
			t.Error("Name is not matched")
		}
		if r.errs > 0 {
			t.Error("should not error, but occurred")
		}
	})

	t.Run("bad case", func(t *testing.T) {
		r := newRecorder()
		v := gt.Cast[testData](r, s).NotNil()
		if v.Name != "" {
			t.Error("Name is not matched")
		}
		if r.errs == 0 {
			t.Error("should error, but not occurred")
		}
		if r.fails > 0 {
			t.Error("should not fail, but occurred")
		}
	})

	t.Run("must case", func(t *testing.T) {
		r := newRecorder()
		v := gt.MustCast[testData](r, s).NotNil()
		if v.Name != "" {
			t.Error("Name is not matched")
		}
		if r.errs == 0 {
			t.Error("should error, but not occurred")
		}
		if r.fails == 0 {
			t.Error("should fail, but not occurred")
		}
	})
}

func TestCast_Nil(t *testing.T) {
	t.Run("value is not nil", func(t *testing.T) {
		r := newRecorder()
		var s any = "test"

		v := gt.MustCast[string](r, s).NotNil()
		if v != "test" {
			t.Error("v is not matched")
		}
		if r.errs > 0 {
			t.Error("v should not be nil")
		}
	})
}
