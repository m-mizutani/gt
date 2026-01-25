# Value

`ValueTest[T]` is a type for general value testing. It provides basic comparison operations for any type.

## Constructor

```go
gt.Value[T any](t testing.TB, actual T) ValueTest[T]
gt.V[T any](t testing.TB, actual T) ValueTest[T]  // Short form
```

## Methods

| Method | Description |
|--------|-------------|
| `Equal(expect T)` | Check if value equals expect |
| `NotEqual(expect T)` | Check if value does not equal expect |
| `Nil()` | Check if value is nil |
| `NotNil()` | Check if value is not nil |
| `In(expects ...T)` | Check if value is one of the expects |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Basic Comparison

```go
type User struct {
    Name string
    Age  int
}

func TestValueEqual(t *testing.T) {
    u1 := User{Name: "Alice", Age: 30}

    // Struct comparison (uses DeepEqual)
    gt.Value(t, u1).Equal(User{Name: "Alice", Age: 30}) // Pass
    gt.Value(t, u1).Equal(User{Name: "Bob", Age: 30})   // Fail

    // Primitive type comparison
    gt.Value(t, "hello").Equal("hello") // Pass
    gt.Value(t, 42).Equal(42)           // Pass
}
```

### NotEqual

```go
func TestValueNotEqual(t *testing.T) {
    color := "blue"
    gt.Value(t, color).NotEqual("red")   // Pass
    gt.Value(t, color).NotEqual("blue")  // Fail
}
```

### Nil / NotNil

```go
func TestValueNil(t *testing.T) {
    var ptr *User
    gt.Value(t, ptr).Nil() // Pass

    ptr = &User{Name: "Alice"}
    gt.Value(t, ptr).NotNil() // Pass
    gt.Value(t, ptr).Nil()    // Fail
}
```

### In

```go
func TestValueIn(t *testing.T) {
    status := "active"
    gt.Value(t, status).In("active", "pending", "closed") // Pass
    gt.Value(t, status).In("deleted", "archived")         // Fail
}
```

### Method Chaining

```go
func TestValueChain(t *testing.T) {
    user := getUser()
    gt.Value(t, user).
        NotNil().
        Required().  // Stop immediately if nil
        Equal(expectedUser)
}
```

### Adding Descriptions

```go
func TestValueDescribe(t *testing.T) {
    userID := 123
    gt.Value(t, userID).
        Describef("User ID for %s", "test case 1").
        Equal(456)

    // Output:
    // User ID for test case 1
    // values are not matched
    // - 456
    // + 123
}
```

## How Comparison Works

`Equal` and `NotEqual` use `reflect.DeepEqual` internally:

- Struct field values are compared recursively
- Slices, maps, and pointers are compared appropriately
- Primitive types behave the same as regular equality comparison

## When to Use Other Types

`ValueTest` is general-purpose, but use specialized types when available:

```go
// Not recommended
gt.Value(t, len(items)).Equal(3)
gt.Value(t, err).Nil()

// Recommended
gt.Array(t, items).Length(3)
gt.NoError(t, err)
```

Benefits of specialized types:
- More intuitive method names
- Additional convenience methods
- More specific error messages
