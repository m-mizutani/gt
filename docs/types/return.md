# Return

`Return1Test`, `Return2Test`, `Return3Test` are types for testing function return values (value and error). They allow concise testing of the common Go `(value, error)` return pattern.

## Constructors

```go
// For functions returning 1 value and error
gt.Return1[T1 any](r1 T1, err error) Return1Test[T1]
gt.R1[T1 any](r1 T1, err error) Return1Test[T1]  // Short form

// For functions returning 2 values and error
gt.Return2[T1, T2 any](r1 T1, r2 T2, err error) Return2Test[T1, T2]
gt.R2[T1, T2 any](r1 T1, r2 T2, err error) Return2Test[T1, T2]  // Short form

// For functions returning 3 values and error
gt.Return3[T1, T2, T3 any](r1 T1, r2 T2, r3 T3, err error) Return3Test[T1, T2, T3]
gt.R3[T1, T2, T3 any](r1 T1, r2 T2, r3 T3, err error) Return3Test[T1, T2, T3]  // Short form
```

## Methods

### Return1Test

| Method | Description |
|--------|-------------|
| `NoError(t testing.TB) T1` | Verify no error and return value. Stops test if error exists |
| `Error(t testing.TB) ErrorTest` | Verify error exists and return ErrorTest |

### Return2Test

| Method | Description |
|--------|-------------|
| `NoError(t testing.TB) (T1, T2)` | Verify no error and return 2 values. Stops test if error exists |
| `Error(t testing.TB) ErrorTest` | Verify error exists and return ErrorTest |

### Return3Test

| Method | Description |
|--------|-------------|
| `NoError(t testing.TB) (T1, T2, T3)` | Verify no error and return 3 values. Stops test if error exists |
| `Error(t testing.TB) ErrorTest` | Verify error exists and return ErrorTest |

## Examples

### Return1: 1 Value and Error

```go
func parseJSON(data string) (Config, error) {
    // ...
}

func TestReturn1NoError(t *testing.T) {
    // No error case
    config := gt.R1(parseJSON(`{"name": "test"}`)).NoError(t)
    gt.String(t, config.Name).Equal("test")
}

func TestReturn1Error(t *testing.T) {
    // Expecting error case
    gt.R1(parseJSON(`invalid`)).Error(t).Contains("invalid JSON")
}
```

### Return2: 2 Values and Error

```go
func divmod(a, b int) (quotient, remainder int, err error) {
    if b == 0 {
        return 0, 0, errors.New("division by zero")
    }
    return a / b, a % b, nil
}

func TestReturn2NoError(t *testing.T) {
    q, r := gt.R2(divmod(10, 3)).NoError(t)
    gt.Number(t, q).Equal(3)
    gt.Number(t, r).Equal(1)
}

func TestReturn2Error(t *testing.T) {
    gt.R2(divmod(10, 0)).Error(t).Contains("division by zero")
}
```

### Return3: 3 Values and Error

```go
func processData(input string) (result string, count int, valid bool, err error) {
    // ...
}

func TestReturn3NoError(t *testing.T) {
    result, count, valid := gt.R3(processData("test")).NoError(t)
    gt.String(t, result).IsNotEmpty()
    gt.Number(t, count).Greater(0)
    gt.Bool(t, valid).True()
}
```

### Detailed Error Verification

```go
func TestReturnErrorDetails(t *testing.T) {
    gt.R1(fetchUser("unknown")).
        Error(t).
        Is(ErrNotFound)

    gt.R1(validateInput("")).
        Error(t).
        Contains("required")
}
```

## Common Patterns

### HTTP Client Testing

```go
func TestHTTPClient(t *testing.T) {
    // Success case
    resp := gt.R1(http.Get("http://example.com")).NoError(t)
    defer resp.Body.Close()
    gt.Number(t, resp.StatusCode).Equal(200)
}
```

### File Operations Testing

```go
func TestFileOperations(t *testing.T) {
    // File reading
    data := gt.R1(os.ReadFile("testdata/config.json")).NoError(t)
    gt.Array(t, data).Longer(0)

    // File does not exist
    gt.R1(os.ReadFile("nonexistent.txt")).Error(t)
}
```

### JSON Parsing Testing

```go
func TestJSONParsing(t *testing.T) {
    validJSON := `{"users": [{"name": "Alice"}, {"name": "Bob"}]}`

    var result Response
    gt.NoError(t, json.Unmarshal([]byte(validJSON), &result)).Required()

    gt.Array(t, result.Users).Length(2)
}
```

### Database Operations Testing

```go
func TestDatabaseQuery(t *testing.T) {
    // Success case
    user := gt.R1(db.FindUser(1)).NoError(t)
    gt.String(t, user.Name).Equal("Alice")

    // Not found case
    gt.R1(db.FindUser(999)).Error(t).Is(ErrNotFound)
}
```

## Why Use Return

### Traditional Approach

```go
func TestTraditional(t *testing.T) {
    config, err := parseJSON(input)
    if err != nil {
        t.Fatalf("unexpected error: %v", err)
    }
    if config.Name != "test" {
        t.Errorf("expected name to be 'test', got %s", config.Name)
    }
}
```

### Using Return

```go
func TestWithReturn(t *testing.T) {
    config := gt.R1(parseJSON(input)).NoError(t)
    gt.String(t, config.Name).Equal("test")
}
```

Benefits:
- More concise code
- Error check and value retrieval in one line
- Test automatically stops on `NoError` failure
- Error messages include detailed stack traces

## Notes

### NoError Calls FailNow

When `NoError` detects an error, `FailNow()` is called and the test stops immediately. This is intentional design - it's meaningless to continue subsequent tests when an error exists.

```go
func TestNoErrorStops(t *testing.T) {
    result := gt.R1(failingFunction()).NoError(t)
    // This line is not executed if error occurred
    gt.Value(t, result).Equal(expected)
}
```

### Error Performs nil Check

`Error` fails the test if the error is `nil`:

```go
func TestErrorExpected(t *testing.T) {
    // This test fails if no error occurred
    gt.R1(successfulFunction()).Error(t)  // Fail if no error
}
```
