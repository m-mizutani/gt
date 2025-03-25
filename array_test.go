package gt_test

import (
	"context"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestArray(t *testing.T) {
	target := []string{"blue", "orange", "red"}

	type testCase struct {
		test func(arr gt.ArrayTest[string])
		pass bool
	}

	testCases := map[string]map[string]testCase{
		"Equal": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Equal([]string{"blue", "orange", "red"})
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Equal([]string{"blue", "yellow", "red"})
				},
				pass: false,
			},
			"fail with nil": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Equal(nil)
				},
				pass: false,
			},
			"fail with bad length (longer)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Equal([]string{"blue", "yellow", "red", "white"})
				},
				pass: false,
			},
			"fail with bad length (shorter)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Equal([]string{"blue", "yellow"})
				},
				pass: false,
			},
		},
		"NotEqual": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqual([]string{"blue", "yellow", "red"})
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqual([]string{"blue", "orange", "red"})
				},
				pass: false,
			},
			"pass with nil": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqual(nil)
				},
				pass: true,
			},
			"pass with bad length (longer)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqual([]string{"blue", "yellow", "red", "white"})
				},
				pass: true,
			},
			"pass with bad length (shorter)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqual([]string{"blue", "yellow"})
				},
				pass: true,
			},
		},
		"EqualAt": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.EqualAt(0, "blue")
				},
				pass: true,
			},
			"fail with not equal": {
				test: func(arr gt.ArrayTest[string]) {
					arr.EqualAt(1, "blue")
				},
				pass: false,
			},
			"fail with out of range (lower)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.EqualAt(-1, "blue")
				},
				pass: false,
			},
			"fail with out of range (upper)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.EqualAt(3, "blue")
				},
				pass: false,
			},
		},
		"NotEqualAt": {
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqualAt(0, "blue")
				},
				pass: false,
			},
			"pass with not equal": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqualAt(1, "blue")
				},
				pass: true,
			},
			"fail with out of range (lower)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqualAt(-1, "blue")
				},
				pass: false,
			},
			"fail with out of range (upper)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotEqualAt(3, "blue")
				},
				pass: false,
			},
		},
		"Contain": {
			"pass (prefix)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Contains([]string{"blue", "orange"})
				},
				pass: true,
			},
			"pass (middle)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Contains([]string{"orange"})
				},
				pass: true,
			},
			"pass (suffix)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Contains([]string{"orange", "red"})
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Contains([]string{"orange", "blue"})
				},
				pass: false,
			},
		},
		"NotContain": {
			"fail (prefix)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotContains([]string{"blue", "orange"})
				},
				pass: false,
			},
			"fail (middle)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotContains([]string{"orange"})
				},
				pass: false,
			},
			"fail (suffix)": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotContains([]string{"orange", "red"})
				},
				pass: false,
			},
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotContains([]string{"orange", "blue"})
				},
				pass: true,
			},
		},
		"Have": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Has("blue")
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Has("yellow")
				},
				pass: false,
			},
		},
		"NotHave": {
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotHas("blue")
				},
				pass: false,
			},
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.NotHas("yellow")
				},
				pass: true,
			},
		},
		"Length": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Length(3)
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Length(4)
				},
				pass: false,
			},
		},
		"Longer": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Longer(2)
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Longer(3)
				},
				pass: false,
			},
		},
		"Less": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Less(4)
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Less(3)
				},
				pass: false,
			},
		},

		"Any": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Any(func(x string) bool {
						return len(x) > 4
					})
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Any(func(x string) bool {
						return len(x) > 6
					})
				},
				pass: false,
			},
		},

		"All": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.All(func(x string) bool {
						return len(x) > 2
					})
				},
				pass: true,
			},
			"fail": {
				test: func(arr gt.ArrayTest[string]) {
					arr.All(func(x string) bool {
						return len(x) > 4
					})
				},
				pass: false,
			},
		},

		"Distinct": {
			"pass": {
				test: func(arr gt.ArrayTest[string]) {
					arr.Distinct()
				},
				pass: true,
			},
		},
	}

	for feature, cases := range testCases {
		t.Run(feature, func(t *testing.T) {
			for title, tc := range cases {
				t.Run(title, func(t *testing.T) {
					r := newRecorder()
					mt := gt.Array(r, target)
					tc.test(mt)

					if tc.pass != (r.errs == 0) {
						t.Errorf("expected: pass=%+v, actual: err=%d", tc.pass, r.errs)
					}
				})
			}
		})
	}
}

func TestArrayDistinct(t *testing.T) {
	target := []string{"blue", "orange", "red", "blue"}

	r := newRecorder()
	mt := gt.Array(r, target)
	mt.Distinct()

	if r.errs == 0 {
		t.Errorf("expected error, but no error")
	}
}

func TestArrayExample1(t *testing.T) {
	data := []int{1, 2, 3}

	gt.Array(t, data).
		Has(1).
		NotHas(4).
		Contains([]int{1, 2}).
		Contains([]int{2, 3}).
		NotContains([]int{1, 3}).
		Equal([]int{1, 2, 3}).
		NotEqual([]int{1, 2, 3, 4}).
		NotEqual([]int{0, 1, 2}).
		Length(3).Longer(2).Less(4)
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

	unorderedUsers := gt.R1(GetUsers(ctx)).NoError(t)

	gt.Array(t, unorderedUsers).
		Has(&user{
			ID:   1000,
			Name: "Alice",
		}).
		Has(&user{
			ID:   1024,
			Name: "Bob",
		}).
		NotHas(&user{
			ID:   9999,
			Name: "TestUser",
		}).
		Length(3)
}

func TestMatchThen(t *testing.T) {
	t.Run("should call then function when match is found", func(t *testing.T) {
		testData := []int{1, 2, 3}
		called := false
		r := newRecorder()
		arrayTest := gt.Array(r, testData)

		arrayTest.MatchThen(func(v int) bool {
			return v == 2
		}, func(t testing.TB, v int) {
			called = true
			if v != 2 {
				t.Errorf("then func received wrong value, got %d, want %d", v, 2)
			}
		})

		if !called {
			t.Errorf("then function was not called")
		}
	})

	t.Run("should report error when no match is found", func(t *testing.T) {
		testData := []int{1, 2, 3}
		called := false
		mockT := &testing.T{}

		arrayTest := gt.Array(mockT, testData)

		arrayTest.MatchThen(func(v int) bool {
			return v == 4
		}, func(t testing.TB, v int) {
			called = true
		})

		if called {
			t.Errorf("then function was called but should not have been")
		}

		if !mockT.Failed() {
			t.Errorf("expected an error when no match is found")
		}
	})
}
