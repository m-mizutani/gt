package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestString(t *testing.T) {
	t.Run("Equal", func(t *testing.T) {
		gt.S(t, "hello").Equal("hello")
		gt.S(t, "world").Equal("world")
	})

	t.Run("NotEqual", func(t *testing.T) {
		gt.S(t, "hello").NotEqual("world")
		gt.S(t, "world").NotEqual("hello")
	})

	t.Run("IsEmpty", func(t *testing.T) {
		gt.S(t, "").IsEmpty()
	})

	t.Run("IsNotEmpty", func(t *testing.T) {
		gt.S(t, "hello").IsNotEmpty()
	})

	t.Run("Contains", func(t *testing.T) {
		gt.S(t, "hello, world").Contains("hello")
	})

	t.Run("NotContains", func(t *testing.T) {
		gt.S(t, "hello, world").NotContains("goodbye")
	})

	t.Run("HasPrefix", func(t *testing.T) {
		gt.S(t, "hello, world").HasPrefix("hello")
	})

	t.Run("NotHasPrefix", func(t *testing.T) {
		gt.S(t, "hello, world").NotHasPrefix("goodbye")
	})

	t.Run("HasSuffix", func(t *testing.T) {
		gt.S(t, "hello, world").HasSuffix("world")
	})

	t.Run("NotHasSuffix", func(t *testing.T) {
		gt.S(t, "hello, world").NotHasSuffix("goodbye")
	})

	t.Run("Match", func(t *testing.T) {
		gt.S(t, "hello, world").Match("^hello.*")
	})
}
