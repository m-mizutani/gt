# Cast

`Cast` is a function for safely performing type assertions in tests. If the type assertion fails, the test is immediately failed and stopped.

## Function Signature

```go
gt.Cast[T any](t testing.TB, v any) T
gt.C[T any](t testing.TB, v any) T  // Short form
```

## Behavior

1. Attempts to assert `v` to type `T`
2. If successful, returns the casted value
3. If failed, marks test as error and stops immediately with `FailNow()`

## Examples

### Basic Usage

```go
func TestCast(t *testing.T) {
    var data any = "hello"

    // Type assertion succeeds
    s := gt.Cast[string](t, data)
    gt.String(t, s).Equal("hello")

    // Type assertion fails (test stops)
    // n := gt.Cast[int](t, data)  // Fail
}
```

### Converting from Interface

```go
type Animal interface {
    Speak() string
}

type Dog struct {
    Name string
}

func (d Dog) Speak() string {
    return "Woof!"
}

func TestInterfaceCast(t *testing.T) {
    var animal Animal = Dog{Name: "Buddy"}

    dog := gt.Cast[Dog](t, animal)
    gt.String(t, dog.Name).Equal("Buddy")
}
```

### Type Conversion of JSON Decode Results

```go
func TestJSONDecode(t *testing.T) {
    var result any
    json.Unmarshal([]byte(`{"name": "Alice", "age": 30}`), &result)

    m := gt.Cast[map[string]any](t, result)
    name := gt.Cast[string](t, m["name"])
    age := gt.Cast[float64](t, m["age"])  // JSON numbers are float64

    gt.String(t, name).Equal("Alice")
    gt.Number(t, age).Equal(30.0)
}
```

### Type Conversion of API Responses

```go
func TestAPIResponse(t *testing.T) {
    response := callAPI()  // returns any

    data := gt.Cast[map[string]any](t, response)
    users := gt.Cast[[]any](t, data["users"])

    gt.Array(t, users).Length(3)

    firstUser := gt.Cast[map[string]any](t, users[0])
    name := gt.Cast[string](t, firstUser["name"])
    gt.String(t, name).Equal("Alice")
}
```

### Using Short Form

```go
func TestCastShorthand(t *testing.T) {
    var v any = 42.0

    // Short form
    n := gt.C[float64](t, v)
    gt.Number(t, n).Equal(42.0)
}
```

## Error Messages

When type assertion fails:

```go
var v any = "hello"
gt.Cast[int](t, v)
// Output: expected int, but can not cast
```

## Why Use Cast

### Safe Type Conversion

Regular type assertions either return two values or may panic:

```go
// Two-value form (safe but verbose)
s, ok := v.(string)
if !ok {
    t.Fatal("expected string")
}

// Single-value form (may panic on failure)
s := v.(string)  // Dangerous
```

Using `Cast`:

```go
s := gt.Cast[string](t, v)  // Safely stops test on failure
```

### Test Readability

You can continue writing code assuming the type assertion succeeded:

```go
func TestComplexData(t *testing.T) {
    data := fetchData()

    result := gt.Cast[map[string]any](t, data)
    items := gt.Cast[[]any](t, result["items"])
    first := gt.Cast[map[string]any](t, items[0])
    name := gt.Cast[string](t, first["name"])

    gt.String(t, name).Equal("expected")
}
```

## Notes

### Casting to Pointer Types

Casting to pointer types is also possible:

```go
var v any = &User{Name: "Alice"}

user := gt.Cast[*User](t, v)
gt.String(t, user.Name).Equal("Alice")
```

### Handling nil

`nil` fails type assertion to most types:

```go
var v any = nil

// This fails
gt.Cast[string](t, v)  // Fail

// Casting to nil-acceptable pointer type
var p any = (*User)(nil)
user := gt.Cast[*User](t, p)  // Pass, user is nil
gt.Value(t, user).Nil()
```
