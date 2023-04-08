package gt_test

import (
	"fmt"
	"strings"
	"testing"

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

func TestDiff(t *testing.T) {
	type testStruct struct {
		Name string
	}

	testCases := map[string]struct {
		wants  []string
		expect any
		actual any
	}{
		"numbers": {
			wants:  []string{"expect: 2", "actual: 1"},
			expect: 2,
			actual: 1,
		},
		"struct": {
			wants:  []string{`gt_test.testStruct{`},
			expect: testStruct{Name: "blue"},
			actual: testStruct{Name: "orange"},
		},
		"pointer": {
			wants:  []string{"[]int{"},
			expect: []int{1, 2, 3},
			actual: []int{1, 5, 3},
		},
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {

			got := gt.Diff(tc.expect, tc.actual)
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					t.Error(strings.Join([]string{
						"",
						"--- want to contain ---",
						want,
						"",
						"--- got ---",
						got,
					}, "\n"))
				}
			}
		})
	}
}
