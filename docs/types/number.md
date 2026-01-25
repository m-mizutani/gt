# Number

`NumberTest[T]` is a type for testing numeric values. It provides equality comparison plus comparison methods for greater/less than.

## Constructor

```go
gt.Number[T number](t testing.TB, actual T) NumberTest[T]
gt.N[T number](t testing.TB, actual T) NumberTest[T]  // Short form
```

The `number` type constraint accepts:
- `int`, `int8`, `int16`, `int32`, `int64`
- `uint`, `uint8`, `uint16`, `uint32`, `uint64`, `uintptr`
- `float32`, `float64`

## Methods

| Method | Description |
|--------|-------------|
| `Equal(expect T)` | Check if value equals expect |
| `NotEqual(expect T)` | Check if value does not equal expect |
| `Greater(expect T)` | Check if value is greater than expect |
| `GreaterOrEqual(expect T)` | Check if value is greater than or equal to expect |
| `Less(expect T)` | Check if value is less than expect |
| `LessOrEqual(expect T)` | Check if value is less than or equal to expect |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Equality Comparison

```go
func TestNumberEqual(t *testing.T) {
    count := 42

    gt.Number(t, count).Equal(42)    // Pass
    gt.Number(t, count).NotEqual(0)  // Pass
}
```

### Greater/Less Comparison

```go
func TestNumberComparison(t *testing.T) {
    score := 85

    gt.Number(t, score).Greater(80)        // Pass (85 > 80)
    gt.Number(t, score).GreaterOrEqual(85) // Pass (85 >= 85)
    gt.Number(t, score).Less(90)           // Pass (85 < 90)
    gt.Number(t, score).LessOrEqual(85)    // Pass (85 <= 85)
}
```

### Range Testing

```go
func TestNumberRange(t *testing.T) {
    percentage := 75.5

    // Check value is within 0-100 range
    gt.Number(t, percentage).
        GreaterOrEqual(0).
        LessOrEqual(100)
}
```

### Integer and Floating Point Types

```go
func TestNumberTypes(t *testing.T) {
    // Works with various numeric types
    gt.Number(t, int8(10)).Greater(5)
    gt.Number(t, int64(1000000)).Less(2000000)
    gt.Number(t, uint32(100)).GreaterOrEqual(100)
    gt.Number(t, float64(3.14)).Less(4.0)
}
```

### Method Chaining

```go
func TestNumberChaining(t *testing.T) {
    value := calculateScore()

    gt.Number(t, value).
        Describef("Score for user %s", userID).
        Greater(0).
        Required().  // Stop if <= 0
        LessOrEqual(100)
}
```

## Why Use Number

### Improved Readability

```go
// Using gt.Number (recommended)
gt.Number(t, score).Greater(80)

// Using gt.Bool (not recommended)
gt.Bool(t, score > 80).True()
```

Using `Number`:
- Makes test intent clear
- Provides more specific error messages

### Error Message Difference

```go
// With gt.Number
gt.Number(t, 50).Greater(80)
// Output: got 50, want greater than 80

// With gt.Bool
gt.Bool(t, 50 > 80).True()
// Output: expected true, but false
```

## Common Patterns

### HTTP Status Code Testing

```go
func TestHTTPStatus(t *testing.T) {
    resp := doRequest()

    gt.Number(t, resp.StatusCode).
        GreaterOrEqual(200).
        Less(300)
}
```

### Array Length Boundary Testing

```go
func TestArrayBounds(t *testing.T) {
    items := getItems()

    gt.Number(t, len(items)).
        Greater(0).       // Not empty
        LessOrEqual(100)  // Within limit
}
```

### Percentage Testing

```go
func TestPercentage(t *testing.T) {
    progress := calculateProgress()

    gt.Number(t, progress).
        GreaterOrEqual(0).
        LessOrEqual(100)
}
```
