package gt_test

import (
	"fmt"

	"github.com/m-mizutani/gt"
)

func ExampleMap() {
	t := newRecorder()

	type user struct {
		ID   int
		Name string
	}
	a := &user{ID: 1, Name: "Alice"}
	b := &user{ID: 2, Name: "Bob"}

	gt.Value(t, a).Equal(b)

	fmt.Println(t.msgs[0])
}

func ExampleMap_nil() {
	t := newRecorder()

	type user struct {
		ID   int
		Name string
	}
	a := &user{ID: 1, Name: "Alice"}

	gt.Value(t, a).Equal(nil)

	fmt.Println(t.msgs[0])
}

// nolint
func ExampleValue_nil() {
	t := newRecorder()

	a := "test"

	gt.Value(t, &a).Nil()

	fmt.Println(t.msgs[0])
}

func ExampleArray() {
	t := newRecorder()

	a := []int{2, 3, 4}
	b := []int{2, 3, 5}

	gt.Value(t, a).Equal(b)

	fmt.Println(t.msgs[0])
}
