package gt_test

import (
	"testing"
	"time"

	"github.com/m-mizutani/gt"
)

func TestTime(t *testing.T) {
	testCases := map[string]struct {
		f        func(mock testing.TB)
		errCount int
	}{
		"match time": {
			f: func(mock testing.TB) {
				gt.Time(mock,
					time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)).
					Equal(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC))
			},
			errCount: 0,
		},
		"unmatched time": {
			f: func(mock testing.TB) {
				gt.Time(mock, time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)).
					Equal(time.Date(2020, 1, 2, 3, 4, 5, 1, time.UTC))
			},
			errCount: 1,
		},
		"match time with location": {
			f: func(mock testing.TB) {
				a := time.Date(2020, 1, 1, 18, 4, 5, 0, time.UTC)
				b := time.Date(2020, 1, 2, 3, 4, 5, 0, time.FixedZone("Asia/Tokyo", 9*60*60))
				gt.Time(mock, a).Equal(b)
			},
			errCount: 0,
		},
		"unmatched time with location": {
			f: func(mock testing.TB) {
				a := time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)
				b := time.Date(2020, 1, 2, 3, 4, 5, 0, time.FixedZone("Asia/Tokyo", 8*60*60))
				gt.Time(mock, a).Equal(b)
			},
			errCount: 1,
		},
		"after: success": {
			f: func(mock testing.TB) {
				gt.Time(mock, time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)).
					After(time.Date(2020, 1, 2, 3, 4, 4, 0, time.UTC))
			},
			errCount: 0,
		},
		"after: fail": {
			f: func(mock testing.TB) {
				gt.Time(mock, time.Date(2020, 1, 2, 3, 4, 4, 0, time.UTC)).
					After(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC))
			},
			errCount: 1,
		},
		"before: success": {
			f: func(mock testing.TB) {
				gt.Time(mock, time.Date(2020, 1, 2, 3, 4, 4, 0, time.UTC)).
					Before(time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC))
			},
			errCount: 0,
		},
		"before: fail": {
			f: func(mock testing.TB) {
				gt.Time(mock, time.Date(2020, 1, 2, 3, 4, 5, 0, time.UTC)).
					Before(time.Date(2020, 1, 2, 3, 4, 4, 0, time.UTC))
			},
			errCount: 1,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			r := newRecorder()
			tc.f(r)
			if r.errs != tc.errCount {
				t.Errorf("unexpected error count: %d", r.errs)
			}
		})
	}
}
