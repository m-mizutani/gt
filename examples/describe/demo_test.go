package main

import (
	"testing"

	"github.com/m-mizutani/gt"
)

// Demo tests that show error message formatting
// Uncomment t.Skip() lines to see actual error messages

func TestDemo_ValueWithDescription(t *testing.T) {
	// Demo test - this will fail to show error message format

	gt.Value(t, "actual").
		Describe("User ID should match the expected value from the database").
		Equal("expected")
}

func TestDemo_ValueWithDescribef(t *testing.T) {
	// Demo test - this will fail to show error message format

	userID := 123
	expectedID := 456
	userName := "Alice"

	gt.Value(t, userID).
		Describef("User ID should be %d for user %s, but got different value", expectedID, userName).
		Equal(expectedID)
}

func TestDemo_ArrayWithLongDescription(t *testing.T) {
	// Demo test - this will fail to show error message format

	items := []string{"apple", "banana"}

	gt.Array(t, items).
		Describe("Array should contain all expected fruits including orange, grape, and pineapple for the complete fruit collection that was supposed to be delivered to the customer").
		Length(5)
}

func TestDemo_ArrayWithoutDescription(t *testing.T) {
	// Demo test - this will fail to show error message format

	items := []string{"apple", "banana"}
	gt.Array(t, items).Length(5)
}

func TestDemo_ValueWithoutDescription(t *testing.T) {
	// Demo test - this will fail to show error message format without description

	gt.Value(t, "actual").Equal("expected")
}

func TestDemo_NumberWithoutDescription(t *testing.T) {
	// Demo test - this will fail to show error message format without description

	gt.Number(t, 42).Equal(100)
}

func TestDemo_StringWithoutDescription(t *testing.T) {
	// Demo test - this will fail to show error message format without description

	gt.String(t, "hello").Equal("world")
}

func TestDemo_BoolWithoutDescription(t *testing.T) {
	// Demo test - this will fail to show error message format without description

	gt.Bool(t, true).False()
}

func TestDemo_MethodChaining(t *testing.T) {
	// Demo test - this will fail to show error message format

	// This will fail first
	gt.Value(t, "initial").
		Describe("Initial validation must pass for subsequent tests").
		Equal("different")

	// This Required() call will show the description in the failure message
	gt.Value(t, "next").
		Describe("This test requires the previous validation to succeed").
		Required().
		Equal("next")
}
