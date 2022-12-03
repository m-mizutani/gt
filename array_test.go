package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestArray(t *testing.T) {
	mock := newErrCounter()

	d := []int{1, 2, 3}
	gt.Array(mock, d).Equal([]int{1, 2, 3})
	if mock.errCount != 0 {
		t.Error("not errored")
	}
}

func TestArrayExample(t *testing.T) {
	data := []int{1, 2, 3}

	gt.Array(t, data).
		Contain(1).
		NotContain(4).
		Equal([]int{1, 2, 3}).
		NotEqual([]int{1, 2, 3, 4}).
		NotEqual([]int{0, 1, 2}).
		Length(3)
}
