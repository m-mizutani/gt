package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestNumberTest_Int(t *testing.T) {
	testCases := map[string]struct {
		actual int
		check  func(n gt.NumberTest[int])
		pass   bool
	}{
		"Equal_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.Equal(10) },
			pass:   true,
		},
		"Equal_fail": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.Equal(11) },
			pass:   false,
		},
		"NotEqual_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.NotEqual(11) },
			pass:   true,
		},
		"NotEqual_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[int]) { n.NotEqual(11) },
			pass:   false,
		},
		"Greater_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.Greater(9) },
			pass:   true,
		},
		"Greater_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[int]) { n.Greater(11) },
			pass:   false,
		},
		"GreaterOrEqual_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.GreaterOrEqual(10) },
			pass:   true,
		},
		"GreaterOrEqual_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[int]) { n.GreaterOrEqual(12) },
			pass:   false,
		},
		"Less_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[int]) { n.Less(11) },
			pass:   true,
		},
		"Less_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[int]) { n.Less(11) },
			pass:   false,
		},
		"LessOrEqual_pass": {
			actual: 5,
			check:  func(n gt.NumberTest[int]) { n.LessOrEqual(5) },
			pass:   true,
		},
		"LessOrEqual_fail": {
			actual: 5,
			check:  func(n gt.NumberTest[int]) { n.LessOrEqual(4) },
			pass:   false,
		},
	}

	pass := func(v bool) string {
		if v {
			return "pass"
		} else {
			return "fail"
		}
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			r := newRecorder()
			n := gt.N(r, tc.actual)
			tc.check(n)
			if tc.pass != (r.errs == 0) {
				t.Errorf("not expected result (expect: %v / actual: %v)", pass(tc.pass), pass(r.errs == 0))
			}
		})
	}
}

func TestNumberTest_Float(t *testing.T) {
	testCases := map[string]struct {
		actual float64
		check  func(n gt.NumberTest[float64])
		pass   bool
	}{
		"Equal_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.Equal(10) },
			pass:   true,
		},
		"Equal_fail": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.Equal(11) },
			pass:   false,
		},
		"NotEqual_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.NotEqual(11) },
			pass:   true,
		},
		"NotEqual_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[float64]) { n.NotEqual(11) },
			pass:   false,
		},
		"Greater_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.Greater(9) },
			pass:   true,
		},
		"Greater_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[float64]) { n.Greater(11) },
			pass:   false,
		},
		"GreaterOrEqual_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.GreaterOrEqual(10) },
			pass:   true,
		},
		"GreaterOrEqual_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[float64]) { n.GreaterOrEqual(12) },
			pass:   false,
		},
		"Less_pass": {
			actual: 10,
			check:  func(n gt.NumberTest[float64]) { n.Less(11) },
			pass:   true,
		},
		"Less_fail": {
			actual: 11,
			check:  func(n gt.NumberTest[float64]) { n.Less(11) },
			pass:   false,
		},
		"LessOrEqual_pass": {
			actual: 5,
			check:  func(n gt.NumberTest[float64]) { n.LessOrEqual(5) },
			pass:   true,
		},
		"LessOrEqual_fail": {
			actual: 5,
			check:  func(n gt.NumberTest[float64]) { n.LessOrEqual(4) },
			pass:   false,
		},
	}

	pass := func(v bool) string {
		if v {
			return "pass"
		} else {
			return "fail"
		}
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			r := newRecorder()
			n := gt.N(r, tc.actual)
			tc.check(n)
			if tc.pass != (r.errs == 0) {
				t.Errorf("not expected result (expect: %v / actual: %v)", pass(tc.pass), pass(r.errs == 0))
			}
		})
	}
}
