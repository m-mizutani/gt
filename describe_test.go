package gt_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

// Test basic Describe functionality with ValueTest
func TestValueTest_BasicDescribe(t *testing.T) {
	// Test that Describe methods exist and can be called
	result := gt.Value(t, "test").Describe("Test description")
	if result == (gt.ValueTest[string]{}) {
		t.Error("Describe() should return a ValueTest instance")
	}
}

func TestValueTest_BasicDescribef(t *testing.T) {
	// Test that Describef methods exist and can be called
	result := gt.Value(t, 42).Describef("Test number %d", 42)
	if result == (gt.ValueTest[int]{}) {
		t.Error("Describef() should return a ValueTest instance")
	}
}

// Test basic Describe functionality with ArrayTest
func TestArrayTest_BasicDescribe(t *testing.T) {
	// Test that Describe methods exist and can be called
	arr := []int{1, 2, 3}
	result := gt.Array(t, arr).Describe("Test array description")
	// Just test that we can call the method without error - result should not be nil
	_ = result
}

func TestArrayTest_BasicDescribef(t *testing.T) {
	// Test that Describef methods exist and can be called
	arr := []string{"a", "b"}
	result := gt.Array(t, arr).Describef("Array with %d items", len(arr))
	// Just test that we can call the method without error - result should not be nil
	_ = result
}

// Test method chaining works
func TestMethodChaining_WithDescribe(t *testing.T) {
	// This should not fail - just testing that chaining works
	gt.Value(t, "test").
		Describe("Method chaining test").
		Equal("test") // This should pass

	// Test with array
	gt.Array(t, []int{1, 2, 3}).
		Describe("Array method chaining").
		Length(3) // This should pass
}

// Test that successful operations don't cause issues
func TestSuccessfulOperations_WithDescribe(t *testing.T) {
	// These should all pass without any errors
	gt.Value(t, 42).Describe("Number test").Equal(42)
	gt.Array(t, []int{1, 2, 3}).Describef("Array of %d elements", 3).Length(3)
	
	// Test with different types
	gt.Value(t, "hello").Describe("String test").Equal("hello")
	gt.Value(t, true).Describef("Boolean test: %v", true).Equal(true)
}

// Test Required method chaining with Describe
func TestRequired_MethodChaining(t *testing.T) {
	// Test that Required() can be chained with Describe()
	// This should pass since no previous test failed
	gt.Value(t, "test").
		Describe("Required test").
		Required().
		Equal("test")
}

// Test Describe with long text (basic test for functionality)
func TestLongDescription_Basic(t *testing.T) {
	longDescription := "This is a very long description that contains more than 80 characters to test how the Describe functionality handles longer text inputs that might need special formatting in error messages."
	
	// This should pass, just testing that long descriptions don't break anything
	gt.Value(t, "value").
		Describe(longDescription).
		Equal("value")
}