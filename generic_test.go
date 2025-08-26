package gt_test

import (
	"errors"
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

func TestExpectError(t *testing.T) {
	testCases := []struct {
		name       string
		expected   bool
		err        error
		shouldFail bool
	}{
		{"expect error and got error", true, errors.New("test error"), false},
		{"expect error but got no error", true, nil, true},
		{"expect no error and got no error", false, nil, false},
		{"expect no error but got error", false, errors.New("test error"), true},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			r := newRecorder()
			gt.ExpectError(r, tc.expected, tc.err)

			if tc.shouldFail && r.errs == 0 {
				t.Errorf("expected test to fail, but it passed")
			}
			if !tc.shouldFail && r.errs > 0 {
				var msg string
				if len(r.msgs) > 0 {
					msg = r.msgs[len(r.msgs)-1]
				}
				t.Errorf("expected test to pass, but it failed: %s", msg)
			}
		})
	}
}

func TestExpectErrorMessages(t *testing.T) {
	checkErrorMessage := func(t *testing.T, r *recorder, expectedMsg string) {
		if r.errs == 0 {
			t.Error("expected test to fail")
		}
		var actualMsg string
		if len(r.msgs) > 0 {
			actualMsg = r.msgs[len(r.msgs)-1]
		}
		if actualMsg != expectedMsg {
			t.Errorf("expected message %q, got %q", expectedMsg, actualMsg)
		}
	}

	t.Run("expected error but got none", func(t *testing.T) {
		r := newRecorder()
		gt.ExpectError(r, true, nil)
		checkErrorMessage(t, r, "expected error, but got no error")
	})

	t.Run("expected no error but got error", func(t *testing.T) {
		r := newRecorder()
		testErr := errors.New("test error")
		gt.ExpectError(r, false, testErr)
		checkErrorMessage(t, r, "expected no error, but got error: test error")
	})
}
