package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestValueEqual(t *testing.T) {
	testCases := map[string]struct {
		f        func(mock testing.TB)
		errCount int
	}{
		"match nil": {
			f: func(mock testing.TB) {
				type s struct{}
				var p *s
				gt.Value(mock, p).Equal(nil)
			},
			errCount: 0,
		},
		"unmatched pointer": {
			f: func(mock testing.TB) {
				type s struct{}
				gt.Value(mock, &s{}).Equal(nil)
			},
			errCount: 1,
		},
		"match struct (ptr)": {
			f: func(mock testing.TB) {
				type s struct {
					a string
					b int
				}

				gt.Value(mock, &s{a: "x", b: 1}).Equal(&s{a: "x", b: 1})
			},
			errCount: 0,
		},
		"not match struct (ptr)": {
			f: func(mock testing.TB) {
				type s struct {
					a string
					b int
				}

				gt.Value(mock, &s{a: "x", b: 2}).Equal(&s{a: "x", b: 1})
			},
			errCount: 1,
		},
		"match number": {
			f: func(mock testing.TB) {
				gt.Value(mock, 1).Equal(1)
			},
			errCount: 0,
		},
		"unmatched number": {
			f: func(mock testing.TB) {
				gt.Value(mock, 1).Equal(2)
			},
			errCount: 1,
		},
		"match string": {
			f: func(mock testing.TB) {
				gt.Value(mock, "abc").Equal("abc")
			},
			errCount: 0,
		},
		"unmatched string": {
			f: func(mock testing.TB) {
				gt.Value(mock, "abc").Equal("xyz")
			},
			errCount: 1,
		},
		"matched []byte": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte("abc")).Equal([]byte("abc"))
			},
			errCount: 0,
		},
		"unmatched []byte": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte("abc")).Equal([]byte("xyz"))
			},
			errCount: 1,
		},
		"[]byte length 0 and greater than 0": {
			f: func(mock testing.TB) {
				gt.Value(mock, []byte{}).Equal([]byte("a"))
			},
			errCount: 1,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			cnt := newErrCounter()
			tc.f(cnt)

			if cnt.errCount != tc.errCount {
				t.Errorf("error count is not match: expected %d, but actual %d", tc.errCount, cnt.errCount)
			}
			if cnt.errCount != tc.errCount {
				t.Errorf("error count is not match: expected %d, but actual %d", tc.errCount, cnt.errCount)
			}
		})
	}
}

func TestValueNil(t *testing.T) {
	testCases := map[string]struct {
		value any
		isNil bool
		isErr bool
	}{
		"match nil": {
			value: nil,
			isNil: true,
			isErr: false,
		},
		"unmatched not nil": {
			value: nil,
			isNil: false,
			isErr: true,
		},
		"string is not Nil": {
			value: "a",
			isNil: true,
			isErr: true,
		},
		"int is not Nil": {
			value: 0,
			isNil: true,
			isErr: true,
		},
		"float is not Nil": {
			value: 1.23,
			isNil: true,
			isErr: true,
		},
		"struct is not Nil": {
			value: struct{}{},
			isNil: true,
			isErr: true,
		},
		"struct ptr is not Nil": {
			value: &struct{}{},
			isNil: true,
			isErr: true,
		},
		"function is not Nil": {
			value: func() {},
			isNil: true,
			isErr: true,
		},
		"chain is not Nil": {
			value: make(chan bool),
			isNil: true,
			isErr: true,
		},
		"slice is not Nil": {
			value: []int{1, 2, 3},
			isNil: true,
			isErr: true,
		},
		"empty slice is Nil": {
			value: []int{},
			isNil: true,
			isErr: false,
		},
	}

	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {
			cnt := newErrCounter()

			v := gt.Value(cnt, tc.value)

			if tc.isNil {
				v.Nil()
			} else {
				v.NotNil()
			}

			if (cnt.errCount > 0) != tc.isErr {
				t.Errorf("Expected isErr: %v, but actual: %v (%T)", tc.isErr, cnt.errCount, tc.value)
			}
		})
	}
}
