package gt_test

import (
	"fmt"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestNumberTest(t *testing.T) {
	numberTest[int](t)
	numberTest[uint](t)
	numberTest[int8](t)
	numberTest[int16](t)
	numberTest[int32](t)
	numberTest[int64](t)
	numberTest[uint8](t)
	numberTest[uint16](t)
	numberTest[uint32](t)
	numberTest[uint64](t)
	numberTest[float32](t)
	numberTest[float64](t)
}

type number interface {
	int | uint |
		int8 | int16 | int32 | int64 |
		uint8 | uint16 | uint32 | uint64 |
		float32 | float64
}

func numberTest[T number](t *testing.T) {
	var v T
	t.Run(fmt.Sprintf("%T", v), func(t *testing.T) {
		testCases := map[string]struct {
			actual T
			check  func(n gt.NumberTest[T])
			pass   bool
		}{
			"Equal_pass": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.Equal(10) },
				pass:   true,
			},
			"Equal_fail": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.Equal(11) },
				pass:   false,
			},
			"NotEqual_pass": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.NotEqual(11) },
				pass:   true,
			},
			"NotEqual_fail": {
				actual: 11,
				check:  func(n gt.NumberTest[T]) { n.NotEqual(11) },
				pass:   false,
			},
			"Greater_pass": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.Greater(9) },
				pass:   true,
			},
			"Greater_fail": {
				actual: 11,
				check:  func(n gt.NumberTest[T]) { n.Greater(11) },
				pass:   false,
			},
			"GreaterOrEqual_pass": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.GreaterOrEqual(10) },
				pass:   true,
			},
			"GreaterOrEqual_fail": {
				actual: 11,
				check:  func(n gt.NumberTest[T]) { n.GreaterOrEqual(12) },
				pass:   false,
			},
			"Less_pass": {
				actual: 10,
				check:  func(n gt.NumberTest[T]) { n.Less(11) },
				pass:   true,
			},
			"Less_fail": {
				actual: 11,
				check:  func(n gt.NumberTest[T]) { n.Less(11) },
				pass:   false,
			},
			"LessOrEqual_pass": {
				actual: 5,
				check:  func(n gt.NumberTest[T]) { n.LessOrEqual(5) },
				pass:   true,
			},
			"LessOrEqual_fail": {
				actual: 5,
				check:  func(n gt.NumberTest[T]) { n.LessOrEqual(4) },
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
					t.Errorf("not expected result (expect: %+v / actual: %+v)", pass(tc.pass), pass(r.errs == 0))
				}
			})
		}
	})
}
