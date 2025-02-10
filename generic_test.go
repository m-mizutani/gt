package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestNil(t *testing.T) {
	var v *int
	testCases := map[string]struct {
		input any
		pass  bool
	}{
		"nil_pass": {
			input: v,
			pass:  true,
		},
		"nil_pass_pointer": {
			input: (*int)(nil),
			pass:  true,
		},
		"nil_pass_interface": {
			input: (any)(nil),
			pass:  true,
		},
		"nil_pass_slice": {
			input: []int(nil),
			pass:  true,
		},
		"nil_pass_map": {
			input: map[string]int(nil),
			pass:  true,
		},
		"nil_pass_channel": {
			input: (chan int)(nil),
			pass:  true,
		},
		"nil_fail": {
			input: "not nil",
			pass:  false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			r := newRecorder()
			gt.Nil(r, tc.input)
			if tc.pass != (r.errs == 0) {
				t.Errorf("unexpected result for %s: want pass=%v, got pass=%v", name, tc.pass, r.errs == 0)
			}
		})
	}
}

func TestNotNil(t *testing.T) {
	var v *int
	testCases := map[string]struct {
		input any
		pass  bool
	}{
		"not_nil_pass": {
			input: "not nil",
			pass:  true,
		},
		"not_nil_pass_pointer": {
			input: (*int)(nil),
			pass:  false,
		},
		"not_nil_pass_interface": {
			input: (any)(nil),
			pass:  false,
		},
		"not_nil_pass_slice": {
			input: []int(nil),
			pass:  false,
		},
		"not_nil_pass_map": {
			input: map[string]int(nil),
			pass:  false,
		},
		"not_nil_fail": {
			input: v,
			pass:  false,
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			r := newRecorder()
			gt.NotNil(r, tc.input)
			if tc.pass != (r.errs == 0) {
				t.Errorf("unexpected result for %s: want pass=%v, got pass=%v", name, tc.pass, r.errs == 0)
			}
		})
	}
}
