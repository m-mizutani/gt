# Getting Started

## Installation

```bash
go get github.com/m-mizutani/gt
```

## Basic Usage

### Value Comparison

```go
package example_test

import (
    "testing"
    "github.com/m-mizutani/gt"
)

func TestValue(t *testing.T) {
    color := "blue"
    gt.Value(t, color).Equal("blue")   // Pass
    gt.Value(t, color).Equal("orange") // Fail
    // gt.Value(t, color).Equal(5)     // Compile error!
}
```

gt uses Go generics, so comparing values of different types results in a compile-time error.

### Use Specialized Types

gt provides specialized types for different use cases. Choose the appropriate type for more intuitive tests.

```go
func TestSpecializedTypes(t *testing.T) {
    // Array testing
    items := []string{"apple", "banana", "cherry"}
    gt.Array(t, items).Length(3).Has("banana")

    // Map testing
    config := map[string]int{"port": 8080, "timeout": 30}
    gt.Map(t, config).HasKey("port").EqualAt("port", 8080)

    // Number testing
    count := 42
    gt.Number(t, count).Greater(0).LessOrEqual(100)

    // String testing
    email := "user@example.com"
    gt.String(t, email).Contains("@").HasSuffix(".com")

    // Boolean testing
    isActive := true
    gt.Bool(t, isActive).True()
}
```

### Error Handling

```go
func TestError(t *testing.T) {
    result, err := doSomething()

    // Check no error and stop immediately if there is one
    gt.NoError(t, err).Required()

    // Verify result
    gt.Value(t, result).Equal(expected)
}

func TestExpectedError(t *testing.T) {
    _, err := doSomethingThatFails()

    // Check that an error occurred
    gt.Error(t, err).Contains("invalid input")
}
```

### Testing Function Return Values

```go
func TestFunctionReturn(t *testing.T) {
    // Test (value, error) returning function
    result := gt.R1(parseJSON(input)).NoError(t)
    gt.String(t, result.Name).Equal("Alice")

    // Test expected error case
    gt.R1(parseJSON(badInput)).Error(t).Contains("invalid")
}
```

## Why gt?

### 1. Type Safety

Traditional test libraries (like testify) use `interface{}`/`any`:

```go
// testify - error not caught until runtime
assert.Equal(t, "hello", 123)  // Compiles but fails at runtime
```

gt uses generics:

```go
// gt - error caught at compile time
gt.Value(t, "hello").Equal(123)  // Compile error!
```

### 2. IDE Support

With type information, IDE autocompletion works accurately. Available methods are clearly displayed.

### 3. Clear Error Messages

```go
gt.Array(t, users).
    Describef("Users for tenant %s", tenantID).
    Length(5)

// Output:
// Users for tenant abc123
// array length is expected to be 5, but actual is 3
```

## Next Steps

- [Test Type Reference](./README.md#test-type-reference) - Detailed documentation for each type
- [Pattern Guide](./README.md#pattern-guide) - Common patterns
