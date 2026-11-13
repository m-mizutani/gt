package gt_test

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/m-mizutani/goerr/v2"
	"github.com/m-mizutani/gt"
)

func TestError(t *testing.T) {
	err := errors.New("test")
	cnt := newRecorder()
	gt.Error(cnt, err)

	if cnt.errs > 0 {
		t.Error("error test has unexpected result")
	}
}

type testError struct{ N int }

func (x testError) Error() string { return "test error" }

func TestErrorAs(t *testing.T) {
	testErr := testError{N: 5}
	err := fmt.Errorf("run error: %w", testErr)

	var called int
	gt.ErrorAs(t, err, func(tgt *testError) {
		called++
		if tgt.N != 5 {
			t.Errorf("testError.N must be 5, but %+v", tgt.N)
		}
	})
	if called != 1 {
		t.Errorf("callback must be called once, but %+v times", called)
	}
}

func TestErrorContains(t *testing.T) {
	t.Run("error contains expected string", func(t *testing.T) {
		err := errors.New("test error message")
		cnt := newRecorder()
		gt.Error(cnt, err).Contains("error message")
		if cnt.errs > 0 {
			t.Error("error test has unexpected result")
		}
	})

	t.Run("error does not contain expected string", func(t *testing.T) {
		err := errors.New("test error message")
		cnt := newRecorder()
		gt.Error(cnt, err).Contains("unexpected message")
		if cnt.errs == 0 {
			t.Error("error test should report error when message does not contain expected string")
		}
	})

	t.Run("nil error", func(t *testing.T) {
		cnt := newRecorder()
		// Error()関数自体がnilエラーの場合にエラーを報告するため、
		// 初期エラー数を記録
		initialErrs := cnt.errs
		gt.Error(cnt, nil).Contains("any message")
		// エラー数が増加していることを確認
		if cnt.errs <= initialErrs {
			t.Error("error test should report additional error for Contains check")
		}
	})
}

func TestErrorIsShowsTheErrors(t *testing.T) {
	expected := errors.New("expected sentinel")

	t.Run("matched error reports nothing", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, fmt.Errorf("wrapped: %w", expected)).Is(expected)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("unmatched error reports both the expected and the actual error", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, errors.New("something else failed")).Is(expected)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		for _, want := range []string{"expected sentinel", "something else failed"} {
			if !strings.Contains(cnt.msgs[0], want) {
				t.Errorf("message must contain %q, but got %q", want, cnt.msgs[0])
			}
		}
	})

	t.Run("values attached by goerr are shown", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, goerr.New("lookup failed", goerr.V("user_id", "u-123"))).Is(expected)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		if !strings.Contains(cnt.msgs[0], "u-123") {
			t.Errorf("message must contain the attached value, but got %q", cnt.msgs[0])
		}
	})

	t.Run("nil error is reported once by gt.Error only", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, nil).Is(expected)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d: %v", cnt.errs, cnt.msgs)
		}
	})
}

func TestErrorIsNotShowsTheError(t *testing.T) {
	unexpected := errors.New("unexpected sentinel")

	t.Run("unmatched error reports nothing", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, errors.New("something else failed")).IsNot(unexpected)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("matched error reports the actual error", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, fmt.Errorf("while saving: %w", unexpected)).IsNot(unexpected)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		if !strings.Contains(cnt.msgs[0], "while saving") {
			t.Errorf("message must contain the actual error, but got %q", cnt.msgs[0])
		}
	})

	t.Run("nil error is reported once by gt.Error only", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, nil).IsNot(unexpected)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d: %v", cnt.errs, cnt.msgs)
		}
	})
}

func TestErrorHasTag(t *testing.T) {
	tagNotFound := goerr.NewTag("not_found")
	tagInternal := goerr.NewTag("internal")

	t.Run("tagged error reports nothing", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, goerr.New("no such user", goerr.T(tagNotFound))).HasTag(tagNotFound)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("tag of a wrapped error is found", func(t *testing.T) {
		cnt := newRecorder()
		inner := goerr.New("no such user", goerr.T(tagNotFound))
		gt.Error(cnt, goerr.Wrap(inner, "lookup failed")).HasTag(tagNotFound)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("tag of an error joined into goerr.Errors is found", func(t *testing.T) {
		cnt := newRecorder()
		joined := goerr.Join(errors.New("plain"), goerr.New("no such user", goerr.T(tagNotFound)))
		gt.Error(cnt, joined).HasTag(tagNotFound)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("another tag reports the expected tag, the actual tags and the error", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, goerr.New("database is down", goerr.T(tagInternal))).HasTag(tagNotFound)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		for _, want := range []string{"not_found", "internal", "database is down"} {
			if !strings.Contains(cnt.msgs[0], want) {
				t.Errorf("message must contain %q, but got %q", want, cnt.msgs[0])
			}
		}
	})

	t.Run("an error that is not a goerr error has no tag", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, errors.New("plain")).HasTag(tagNotFound)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d", cnt.errs)
		}
	})

	t.Run("nil error is reported once by gt.Error only", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, nil).HasTag(tagNotFound)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d: %v", cnt.errs, cnt.msgs)
		}
	})
}

func TestErrorNotHasTag(t *testing.T) {
	tagNotFound := goerr.NewTag("not_found")
	tagInternal := goerr.NewTag("internal")

	t.Run("error without the tag reports nothing", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, goerr.New("database is down", goerr.T(tagInternal))).NotHasTag(tagNotFound)
		gt.Error(cnt, errors.New("plain")).NotHasTag(tagNotFound)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("tagged error reports the tag and the error", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, goerr.New("no such user", goerr.T(tagNotFound))).NotHasTag(tagNotFound)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		for _, want := range []string{"not_found", "no such user"} {
			if !strings.Contains(cnt.msgs[0], want) {
				t.Errorf("message must contain %q, but got %q", want, cnt.msgs[0])
			}
		}
	})

	t.Run("nil error is reported once by gt.Error only", func(t *testing.T) {
		cnt := newRecorder()
		gt.Error(cnt, nil).NotHasTag(tagNotFound)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d: %v", cnt.errs, cnt.msgs)
		}
	})
}

func TestErrorTagDescription(t *testing.T) {
	cnt := newRecorder()
	gt.Error(cnt, errors.New("plain")).Describe("lookup of a missing user").HasTag(goerr.NewTag("not_found"))
	if cnt.errs != 1 {
		t.Fatalf("one error must be reported, but got %d", cnt.errs)
	}
	if !strings.HasPrefix(cnt.msgs[0], "lookup of a missing user\n") {
		t.Errorf("message must start with the description, but got %q", cnt.msgs[0])
	}
}
