package gt_test

import (
	"context"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestArray(t *testing.T) {
	mock := newRecorder()

	d := []int{1, 2, 3}
	gt.Array(mock, d).Equal([]int{1, 2, 3})
	if mock.errs != 0 {
		t.Error("not errored")
	}
}

func TestArrayExample1(t *testing.T) {
	data := []int{1, 2, 3}

	gt.A(t, data).
		Have(1).
		NotHave(4).
		Contain([]int{1, 2}).
		Contain([]int{2, 3}).
		NotContain([]int{1, 3}).
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
	gt.Error(t, err).Pass().Must()
	gt.Array(t, unorderedUsers).
		Have(&user{
			ID:   1000,
			Name: "Alice",
		}).
		Have(&user{
			ID:   1024,
			Name: "Bob",
		}).
		NotHave(&user{
			ID:   9999,
			Name: "TestUser",
		}).
		Length(3)
}
