package gt_test

import (
	"context"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestArray(t *testing.T) {
	mock := newCounter()

	d := []int{1, 2, 3}
	gt.Array(mock, d).Equal([]int{1, 2, 3})
	if mock.errs != 0 {
		t.Error("not errored")
	}
}

func TestArrayExample1(t *testing.T) {
	data := []int{1, 2, 3}

	gt.Array(t, data).
		Contain(1).
		NotContain(4).
		Equal([]int{1, 2, 3}).
		NotEqual([]int{1, 2, 3, 4}).
		NotEqual([]int{0, 1, 2}).
		Length(3)
}

func TestArrayExample2(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	GetUsers := func(ctx context.Context) ([]*user, error) {
		return []*user{
			{
				ID:   1000,
				Name: "Alice",
			},
			{
				ID:   1024,
				Name: "Bob",
			},
			{
				ID:   1025,
				Name: "Cyno",
			},
		}, nil
	}
	ctx := context.Background()

	unorderedUsers, err := GetUsers(ctx)
	gt.Error(t, err).Passed()
	gt.Array(t, unorderedUsers).
		Contain(&user{
			ID:   1000,
			Name: "Alice",
		}).
		Contain(&user{
			ID:   1024,
			Name: "Bob",
		}).
		NotContain(&user{
			ID:   9999,
			Name: "TestUser",
		}).
		Length(3)
}
