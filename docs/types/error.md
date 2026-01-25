# Error

gt provides multiple functions and types for testing errors. It handles both cases where an error is expected and cases where no error should occur.

## Constructors and Functions

### When Error is Expected

```go
gt.Error(t testing.TB, actual error) ErrorTest
```

Verifies that an error has occurred. The test fails if `actual` is `nil`.

### When No Error is Expected

```go
gt.NoError(t testing.TB, actual error) NoErrorTest
```

Verifies that no error has occurred. The test fails if `actual` is not `nil`.

### Error Check with Type Assertion

```go
gt.ErrorAs[T any](t testing.TB, actual error, callback func(expect *T))
```

Checks the error type using `errors.As` and executes the callback if matched.

## ErrorTest Methods

| Method | Description |
|--------|-------------|
| `Is(expected error)` | Verify error matches using `errors.Is` |
| `IsNot(expected error)` | Verify error does not match using `errors.Is` |
| `Contains(substr string)` | Verify error message contains substring |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## NoErrorTest Methods

| Method | Description |
|--------|-------------|
| `Required()` | Stop test immediately if error exists |

## Examples

### Verify No Error

```go
func TestNoError(t *testing.T) {
    result, err := doSomething()

    gt.NoError(t, err)
    gt.Value(t, result).Equal(expected)
}
```

### Verify No Error with Fail-Fast

```go
func TestNoErrorRequired(t *testing.T) {
    config, err := loadConfig()
    gt.NoError(t, err).Required()  // Stop immediately if error

    // Only executed if no error
    gt.String(t, config.Name).IsNotEmpty()
}
```

### Verify Error Occurs

```go
func TestError(t *testing.T) {
    _, err := parseInvalidInput()

    gt.Error(t, err)  // Verify error occurred
}
```

### Verify Error Message Content

```go
func TestErrorContains(t *testing.T) {
    _, err := validateInput("")

    gt.Error(t, err).Contains("input is required")
}
```

### Check for Specific Error

```go
var ErrNotFound = errors.New("not found")

func TestErrorIs(t *testing.T) {
    err := findUser("unknown")

    gt.Error(t, err).Is(ErrNotFound)
}

func TestErrorIsNot(t *testing.T) {
    err := findUser("unknown")

    gt.Error(t, err).IsNot(ErrPermissionDenied)
}
```

### Check Error Type

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func TestErrorAs(t *testing.T) {
    err := validate(input)

    gt.ErrorAs[ValidationError](t, err, func(e *ValidationError) {
        gt.String(t, e.Field).Equal("email")
        gt.String(t, e.Message).Contains("invalid")
    })
}
```

### Testing Wrapped Errors

```go
func TestWrappedError(t *testing.T) {
    err := process()
    // err = fmt.Errorf("process failed: %w", ErrNotFound)

    // errors.Is detects wrapped errors
    gt.Error(t, err).Is(ErrNotFound)
}
```

## Common Patterns

### Testing Function Return Values

```go
func TestFunctionReturn(t *testing.T) {
    // No error case
    result := gt.R1(parseJSON(validInput)).NoError(t)
    gt.String(t, result.Name).Equal("Alice")

    // Error case
    gt.R1(parseJSON(invalidInput)).Error(t).Contains("invalid JSON")
}
```

### Testing Multiple Error Conditions

```go
func TestMultipleErrorConditions(t *testing.T) {
    testCases := []struct {
        input   string
        wantErr string
    }{
        {"", "input is required"},
        {"ab", "minimum length is 3"},
        {"invalid@", "invalid format"},
    }

    for _, tc := range testCases {
        t.Run(tc.input, func(t *testing.T) {
            err := validate(tc.input)
            gt.Error(t, err).Contains(tc.wantErr)
        })
    }
}
```

### Don't Test nil Error with Value

```go
// Not recommended
gt.Value(t, err).Nil()

// Recommended
gt.NoError(t, err)
```

Using `NoError`:
- Makes test intent clear
- Error message includes error details
- Can stop immediately with `Required()`
