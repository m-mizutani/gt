package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestMap(t *testing.T) {
	m := map[string]int{
		"void":  1,
		"white": 3,
		"blue":  5,
	}

	type testCase struct {
		test func(mt gt.MapTest[string, int])
		pass bool
	}

	testCases := map[string]map[string]testCase{
		"Equal": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.Equal(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  5,
					})
				},
				pass: true,
			},
			"fail_missing_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.Equal(map[string]int{
						"void":  1,
						"white": 3,
					})
				},
				pass: false,
			},
			"fail_additional_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.Equal(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  5,
						"???":   6,
					})
				},
				pass: false,
			},
			"fail_value_not_matched": {
				test: func(mt gt.MapTest[string, int]) {
					mt.Equal(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  666,
					})
				},
				pass: false,
			},
		},
		"NotEqual": {
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqual(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  5,
					})
				},
				pass: false,
			},
			"pass_missing_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqual(map[string]int{
						"void":  1,
						"white": 3,
					})
				},
				pass: true,
			},
			"pass_additional_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqual(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  5,
						"???":   6,
					})
				},
				pass: true,
			},
			"pass_value_not_matched": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqual(map[string]int{
						"void":  1,
						"white": 3,
						"blue":  666,
					})
				},
				pass: true,
			},
		},
		"EqualAt": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.EqualAt("blue", 5)
				},
				pass: true,
			},
			"fail by not equal": {
				test: func(mt gt.MapTest[string, int]) {
					mt.EqualAt("blue", 6)
				},
			},
			"fail by key not found": {
				test: func(mt gt.MapTest[string, int]) {
					mt.EqualAt("orange", 5)
				},
			},
		},
		"NotEqualAt": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqualAt("blue", 6)
				},
				pass: true,
			},
			"fail by equal": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqualAt("blue", 5)
				},
			},
			"fail by key not found": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotEqualAt("orange", 6)
				},
			},
		},

		"HaveKey": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKey("white")
				},
				pass: true,
			},
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKey("orange")
				},
				pass: false,
			},
		},
		"NotHaveKey": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKey("orange")
				},
				pass: true,
			},
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKey("white")
				},
				pass: false,
			},
		},
		"HaveValue": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveValue(5)
				},
				pass: true,
			},
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveValue(0)
				},
				pass: false,
			},
		},
		"NotHaveValue": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveValue(0)
				},
				pass: true,
			},
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveValue(5)
				},
				pass: false,
			},
		},
		"HaveKeyValue": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKeyValue("blue", 5)
				},
				pass: true,
			},
			"fail_no_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKeyValue("orange", 5)
				},
				pass: false,
			},
			"fail_no_value": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKeyValue("blue", 0)
				},
				pass: false,
			},
			"fail_no_key_value": {
				test: func(mt gt.MapTest[string, int]) {
					mt.HaveKeyValue("orange", 0)
				},
				pass: false,
			},
		},
		"NotHaveKeyValue": {
			"fail": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKeyValue("blue", 5)
				},
				pass: false,
			},
			"pass_no_key": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKeyValue("orange", 5)
				},
				pass: true,
			},
			"pass_no_value": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKeyValue("blue", 0)
				},
				pass: true,
			},
			"pass_no_key_value": {
				test: func(mt gt.MapTest[string, int]) {
					mt.NotHaveKeyValue("orange", 0)
				},
				pass: true,
			},
		},
		"Elem": {
			"pass": {
				test: func(mt gt.MapTest[string, int]) {
					mt.At("blue", func(t testing.TB, v int) {
						gt.V(t, v).Equal(5)
					})
				},
				pass: true,
			},
			"fail by not equal": {
				test: func(mt gt.MapTest[string, int]) {
					mt.At("blue", func(t testing.TB, v int) {
						gt.V(t, v).Equal(6)
					})
				},
				pass: false,
			},
			"fail by key not found": {
				test: func(mt gt.MapTest[string, int]) {
					mt.At("orange", func(t testing.TB, v int) {
						gt.V(t, v).Equal(5)
					})
				},
				pass: false,
			},
		},
	}

	for feature, cases := range testCases {
		t.Run(feature, func(t *testing.T) {
			for title, tc := range cases {
				t.Run(title, func(t *testing.T) {
					r := newRecorder()
					mt := gt.M(r, m)
					tc.test(mt)

					if tc.pass != (r.errs == 0) {
						t.Errorf("expected: pass=%+v, actual: err=%d", tc.pass, r.errs)
					}
				})
			}
		})
	}
}
