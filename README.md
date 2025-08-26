# gt: Generics based Test library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/m-mizutani/gt.svg)](https://pkg.go.dev/github.com/m-mizutani/gt) [![test](https://github.com/m-mizutani/gt/actions/workflows/test.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/test.yml) [![gosec](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml) [![lint](https://github.com/m-mizutani/gt/actions/workflows/lint.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/lint.yml)

`gt` is test library leveraging Go generics to check variable type in IDE and compiler.

```go
color := "blue"

// gt.Value(t, color).Equal(5) // <- Compile error

gt.Value(t, color).Equal("orange") // <- Fail
gt.Value(t, color).Equal("blue")   // <- Pass
```

```go
colors := ["red", "blue"]

// gt.Array(t, colors).Equal("red")       // <- Compile error
// gt.Array(t, colors).Equal([]int{1, 2}) // <- Compile error

gt.Array(t, colors).Equal([]string{"red", "blue"}) // <- Pass
gt.Array(t, colors).Has("orange")                 // <- Fail
```

## Motivation

Existing test libraries in Go such as [testify](https://github.com/stretchr/testify) strongly support writing unit test. Many test libraries uses `reflect` package to identify and compare variable type and value and test functions of the libraries accept any type by `interface{}` or `any`. However the approach has two problems:

- Variable types mismatch between _expected_ and _actual_ can not be detected before running the test.
- IDE can not support variable completion because types can not be determined before running the test.

On the other hand, Go started to provide [Generics](https://go.dev/doc/tutorial/generics) feature by version 1.18. It can be leveraged to support type completion and checking types before running a test.

## Usage

In many cases, a developer does not care Go generics in using `gt`. However, a developer need to specify generic type (`Value`, `Array`, `Map`, `Error`, etc.) explicitly to use specific test functions for each types.

See [reference](https://pkg.go.dev/github.com/m-mizutani/gt) for more detail.

### Value

Generic test type has a minimum set of test methods.

```go
type user struct {
    Name string
}
u1 := user{Name: "blue"}

// gt.Value(t, u1).Equal(1)                  // Compile error
// gt.Value(t, u1).Equal("blue")             // Compile error
// gt.Value(t, u1).Equal(&user{Name:"blue"}) // Compile error

gt.Value(t, u1).Equal(user{Name:"blue"}) // Pass
```

#### Test Descriptions

All test types support `Describe()` and `Describef()` methods to add context to test failures:

```go
userID := 123
expectedID := 456

// Basic description
gt.Value(t, userID).
    Describe("User ID should match the expected value").
    Equal(expectedID)

// Formatted description
gt.Value(t, userID).
    Describef("User ID should be %d, but got %d", expectedID, userID).
    Equal(expectedID)
```

Error output with description:
```
User ID should be 456, but got 123
values are not matched
actual: 123
expect: 456
```

Error output without description:
```
values are not matched
actual: 123
expect: 456
```

### Number

Accepts only number types: `int`, `uint`, `int64`, `float64`, etc.

```go
var f float64 = 12.5
gt.Number(t, f).
    Equal(12.5).         // Pass
    Greater(12).         // Pass
    Less(10).            // Fail
    GreaterOrEqual(12.5) // Pass
```

### Array

Accepts array of any type not only primitive type but also struct.

```go
colors := []string{"red", "blue", "yellow"}

gt.Array(t, colors).
    Equal([]string{"red", "blue", "yellow"}) // Pass
    Equal([]string{"red", "blue"})           // Fail
    // Equal([]int{1, 2})                    // Compile error
    Contain([]string{"red", "blue"})         // Pass
    Has("yellow")                           // Pass
    Length(3)                                // Pass

gt.Array(t, colors).Required().Has("orange") // Fail and stop test
```

### Map

```go
colorMap := map[string]int{
    "red": 1,
    "yellow": 2,
    "blue": 5,
}

gt.Map(t, colorMap)
    .HasKey("blue")           // Pass
    .HasValue(5)              // Pass
    // .HasValue("red")       // Compile error
    .HasKeyValue("yellow", 2) // Pass

gt.Map(t, colorMap).Required().HasKey("orange") // Fail and stop test
```

### Cast

```go
type user struct {
    Name string
}
var v any = &user{
    Name: "blue",
}

u1 := gt.Cast[user](t, v).NotNil()  // Fail (because v is *user, not user)
gt.Cast[*user](t, v).Nil()          // Fail (because v is not nil)

u2 := gt.Cast[*user](t, v).NotNil() // Pass
gt.Value(t, u2.Name).Equal("blue")       // Pass
```

### Bool

```go
gt.Bool(t, true).True()   // Pass
gt.Bool(t, false).False() // Pass
gt.Bool(t, true).False()  // Fail

// Sugar syntax
gt.True(t, true)          // Pass
gt.False(t, false)        // Pass
```

### String

```go
name := "Alice"
gt.String(t, name).
    Equal("Alice").           // Pass
    IsNotEmpty().             // Pass
    Contains("lic").          // Pass
    HasPrefix("Al").          // Pass
    HasSuffix("ce").          // Pass
    Match(`^A\w+e$`)          // Pass (regex)

// Sugar syntax
gt.S(t, name).Equal("Alice")
```

### Error

Error testing with specialized methods:

```go
err := errors.New("test error")
gt.Error(t, err).
    Is(errors.New("test error")).  // Check error equality
    Contains("test")               // Check error message contains substring

// NoError for functions that should succeed
gt.NoError(t, someFunc()).Required() // Fail fast if error occurs

// ErrorAs for type checking
var customErr *MyCustomError
gt.ErrorAs(t, err, func(e *MyCustomError) {
    gt.Value(t, e.Code).Equal(404)
})
```

### File

File system testing:

```go
gt.File(t, "testdata/file.txt").
    Exists().                      // Check file exists
    String(func(t testing.TB, content string) {
        gt.String(t, content).Contains("expected text")
    })

gt.File(t, "nonexistent.txt").NotExists() // Check file doesn't exist
```

### Return Values

Test function return values with error handling:

```go
// Function returning (value, error)
result := gt.Return1(myFunc()).NoError(t) // Get result if no error
gt.Value(t, result).Equal("expected")

// Function returning (val1, val2, error)
val1, val2 := gt.Return2(myFunc2()).NoError(t)
gt.Value(t, val1).Equal("expected1")
gt.Value(t, val2).Equal("expected2")

// Test error cases
gt.Return1(myFailingFunc()).Error(t).Contains("expected error message")

// Sugar syntax
gt.R1(myFunc()).NoError(t)
gt.R2(myFunc2()).NoError(t)
```

### Nil

```go
gt.Nil(t, nil)
gt.Nil(t, (*int)(nil))
gt.Nil(t, []int(nil))
gt.NotNil(t, "not nil")
```

## License

Apache License 2.0