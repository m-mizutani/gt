# Required Pattern

`Required()` is a method that stops the test immediately if the previous test has failed. This implements the "fail-fast" pattern and prevents meaningless chains of errors.

## Basic Usage

```go
func TestRequired(t *testing.T) {
    user, err := getUser(id)

    gt.NoError(t, err).Required()  // Stop immediately if error

    // Only executed if no error
    gt.String(t, user.Name).Equal("Alice")
}
```

## Why Use Required

### Problem: Meaningless Error Chains

Without `Required()`, subsequent tests run even when prerequisites are not met:

```go
func TestWithoutRequired(t *testing.T) {
    user, err := getUser(id)

    gt.NoError(t, err)  // Continues even if error
    // user might be nil

    gt.String(t, user.Name).Equal("Alice")  // nil pointer panic!
}
```

### Solution: Early Exit with Required

```go
func TestWithRequired(t *testing.T) {
    user, err := getUser(id)

    gt.NoError(t, err).Required()  // Stop immediately if error

    // user is guaranteed to not be nil
    gt.String(t, user.Name).Equal("Alice")
}
```

## Usage Patterns

### Error Check

```go
config, err := loadConfig()
gt.NoError(t, err).Required()

// Only continue if config loaded successfully
gt.String(t, config.Name).IsNotEmpty()
```

### Nil Check

```go
result := findUser(query)
gt.Value(t, result).NotNil().Required()

// result is guaranteed to not be nil
gt.String(t, result.Email).Contains("@")
```

### Array Length Check

```go
items := getItems()
gt.Array(t, items).Length(3).Required()

// Length is guaranteed to be 3
gt.Value(t, items[0]).Equal(first)
gt.Value(t, items[1]).Equal(second)
gt.Value(t, items[2]).Equal(third)
```

### Map Key Existence Check

```go
data := getData()
gt.Map(t, data).HasKey("config").Required()

// "config" key is guaranteed to exist
config := data["config"]
processConfig(config)
```

## Required within Method Chains

`Required()` can be inserted in the middle of a chain:

```go
gt.Array(t, users).
    Describef("Users from API").
    Length(10).
    Required().      // Stop if length is not 10
    Has(admin).      // Check admin exists
    Required().      // Stop if no admin
    All(isValid)     // Verify all users
```

## NoError and Required

`NoError` is a special case. The combination `NoError(t, err).Required()` is very common:

```go
// Pattern 1: Explicit
result, err := doSomething()
gt.NoError(t, err).Required()

// Using return helpers is implicitly Required
result := gt.R1(doSomething()).NoError(t)  // Auto-stops on error
```

The `NoError()` method of `R1`, `R2`, `R3` internally calls `FailNow()`, so it implicitly behaves the same as Required.

## When to Use Required

### 1. When Subsequent Tests Depend on Prerequisites

```go
file, err := os.Open(path)
gt.NoError(t, err).Required()
defer file.Close()

// Only if file opened successfully
content, err := io.ReadAll(file)
gt.NoError(t, err)
```

### 2. To Prevent nil Pointer Access

```go
user := findUser(id)
gt.Value(t, user).NotNil().Required()

// Safe to access user.Name
gt.String(t, user.Name).Equal("Alice")
```

### 3. Before Array Index Access

```go
items := getItems()
gt.Array(t, items).Longer(0).Required()

// Safe to access items[0]
gt.Value(t, items[0]).Equal(expected)
```

## When Not to Use Required

### Independent Tests

Not needed when tests have no dependencies:

```go
func TestIndependent(t *testing.T) {
    gt.String(t, getName()).Equal("Alice")
    gt.Number(t, getAge()).Greater(0)   // Does not depend on above
    gt.Bool(t, isActive()).True()       // Does not depend on above
}
```

### End of Test

Not needed for the last assertion in a test:

```go
func TestLast(t *testing.T) {
    result := calculate()
    gt.Number(t, result).Equal(42)  // Required not needed
}
```

## Internal Behavior

`Required()` uses `t.Failed()` and `t.FailNow()` internally:

1. Check if previous test failed with `t.Failed()`
2. If failed, call `t.FailNow()` to stop test immediately
3. If not failed, continue as normal

```go
func (x ValueTest[T]) Required() ValueTest[T] {
    if x.t.Failed() {
        x.t.FailNow()
    }
    return x
}
```
