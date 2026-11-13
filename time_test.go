package gt_test

import (
	"strings"
	"testing"
	"time"

	"github.com/m-mizutani/gt"
)

func TestTimeNear(t *testing.T) {
	base := time.Date(2024, 5, 1, 12, 0, 0, 0, time.UTC)

	t.Run("within tolerance on either side reports nothing", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base).Near(base, time.Second)
		gt.Time(cnt, base.Add(999*time.Millisecond)).Near(base, time.Second)
		gt.Time(cnt, base.Add(-999*time.Millisecond)).Near(base, time.Second)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})

	t.Run("exactly the tolerance away is not near", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(time.Second)).Near(base, time.Second)
		if cnt.errs != 1 {
			t.Errorf("one error must be reported, but got %d", cnt.errs)
		}
	})

	t.Run("later than tolerance reports both times and the difference", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(5*time.Minute)).Near(base, time.Second)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		for _, want := range []string{"2024-05-01T12:00:00Z", "2024-05-01T12:05:00Z", "5m0s", "1s"} {
			if !strings.Contains(cnt.msgs[0], want) {
				t.Errorf("message must contain %q, but got %q", want, cnt.msgs[0])
			}
		}
	})

	t.Run("earlier than tolerance is not near", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(-5*time.Minute)).Near(base, time.Second)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		if !strings.Contains(cnt.msgs[0], "-5m0s") {
			t.Errorf("message must contain the signed difference, but got %q", cnt.msgs[0])
		}
	})

	t.Run("the same instant in another location is near", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.In(time.FixedZone("JST", 9*60*60))).Near(base, time.Millisecond)
		if cnt.errs != 0 {
			t.Errorf("no error must be reported, but got %v", cnt.msgs)
		}
	})
}

func TestTimeDescribe(t *testing.T) {
	base := time.Date(2024, 5, 1, 12, 0, 0, 0, time.UTC)

	t.Run("description leads the failure message", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(time.Hour)).Describe("created_at of the stored record").Near(base, time.Second)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		if !strings.HasPrefix(cnt.msgs[0], "created_at of the stored record\n") {
			t.Errorf("message must start with the description, but got %q", cnt.msgs[0])
		}
	})

	t.Run("formatted description leads the failure message", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(time.Hour)).Describef("created_at of record %d", 7).Near(base, time.Second)
		if cnt.errs != 1 {
			t.Fatalf("one error must be reported, but got %d", cnt.errs)
		}
		if !strings.HasPrefix(cnt.msgs[0], "created_at of record 7\n") {
			t.Errorf("message must start with the description, but got %q", cnt.msgs[0])
		}
	})
}

func TestTimeRequired(t *testing.T) {
	base := time.Date(2024, 5, 1, 12, 0, 0, 0, time.UTC)

	t.Run("stops the test after a failure", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base.Add(time.Hour)).Near(base, time.Second).Required()
		if cnt.fails != 1 {
			t.Errorf("FailNow must be called once, but got %d", cnt.fails)
		}
	})

	t.Run("does not stop the test without a failure", func(t *testing.T) {
		cnt := newRecorder()
		gt.Time(cnt, base).Near(base, time.Second).Required()
		if cnt.fails != 0 {
			t.Errorf("FailNow must not be called, but got %d", cnt.fails)
		}
	})
}
