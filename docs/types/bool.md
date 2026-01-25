# Bool

`BoolTest` is a type for testing boolean values. It provides simple `True()` / `False()` methods.

## Constructor

```go
gt.Bool(t testing.TB, actual bool) BoolTest
gt.B(t testing.TB, actual bool) BoolTest  // Short form
```

### Convenience Functions

```go
gt.True(t testing.TB, actual bool) BoolTest   // Equivalent to gt.Bool(t, actual).True()
gt.False(t testing.TB, actual bool) BoolTest  // Equivalent to gt.Bool(t, actual).False()
```

## Methods

| Method | Description |
|--------|-------------|
| `True()` | Check if value is true |
| `False()` | Check if value is false |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Basic Usage

```go
func TestBool(t *testing.T) {
    isActive := true
    isDeleted := false

    gt.Bool(t, isActive).True()   // Pass
    gt.Bool(t, isDeleted).False() // Pass
}
```

### Using Convenience Functions

```go
func TestBoolShorthand(t *testing.T) {
    gt.True(t, isValid())   // Pass if isValid() returns true
    gt.False(t, hasError()) // Pass if hasError() returns false
}
```

### Test with Description

```go
func TestBoolWithDescription(t *testing.T) {
    user := getUser()

    gt.Bool(t, user.IsAdmin).
        Describef("Admin status for user %s", user.Name).
        True()
}
```

### Method Chaining with Required

```go
func TestBoolRequired(t *testing.T) {
    config := loadConfig()

    gt.Bool(t, config.IsValid).
        True().
        Required()  // Stop immediately if false

    // Only executed if IsValid is true
    processConfig(config)
}
```

## When to Use Bool

### Appropriate Usage

When testing boolean values directly:

```go
gt.Bool(t, user.IsActive).True()
gt.Bool(t, order.IsCanceled).False()
```

### Avoid Using Bool For

Use specialized types when available:

```go
// Not recommended
gt.Bool(t, len(items) == 3).True()
gt.Bool(t, strings.Contains(s, "error")).True()
gt.Bool(t, err == nil).True()
gt.Bool(t, n > 5).True()

// Recommended
gt.Array(t, items).Length(3)
gt.String(t, s).Contains("error")
gt.NoError(t, err)
gt.Number(t, n).Greater(5)
```

Benefits of specialized types:
- More specific error messages
- Additional methods available
- Clearer test intent

## Error Messages

```go
gt.Bool(t, false).True()
// Output: expected true, but false

gt.Bool(t, true).False()
// Output: expected false, but true
```

With description:

```go
gt.Bool(t, user.IsVerified).
    Describef("User %s verification status", user.Email).
    True()
// Output:
// User user@example.com verification status
// expected true, but false
```
