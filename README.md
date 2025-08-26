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

## Test Types Overview

`gt` provides specialized test types for different kinds of data, each with type-safe methods:

| Test Type | Constructor | Purpose | Key Methods |
|-----------|-------------|---------|-------------|
| **Value** | `gt.Value(t, v)` | Generic value testing | `Equal`, `NotEqual`, `Nil`, `NotNil` |
| **Array** | `gt.Array(t, arr)` | Slice/array testing | `Has`, `Contains`, `Length`, `Any`, `All`, `Distinct` |
| **Map** | `gt.Map(t, m)` | Map testing | `HasKey`, `HasValue`, `HasKeyValue`, `EqualAt` |
| **Number** | `gt.Number(t, n)` | Numeric comparisons | `Greater`, `Less`, `GreaterOrEqual`, `LessOrEqual` |
| **String** | `gt.String(t, s)` | String operations | `Contains`, `HasPrefix`, `HasSuffix`, `Match`, `IsEmpty` |
| **Bool** | `gt.Bool(t, b)` | Boolean testing | `True`, `False` |
| **Error** | `gt.Error(t, err)` | Error validation | `Is`, `Contains`, `As` |
| **File** | `gt.File(t, path)` | File system testing | `Exists`, `NotExists`, `String` |
| **Cast** | `gt.Cast[T](t, v)` | Type casting | Type-safe casting with `Nil`/`NotNil` |

### Common Features

All test types support:

- **Fluent Interface**: Method chaining for readable test code
- **Descriptions**: `Describe(msg)` and `Describef(format, args...)` for context
- **Required Pattern**: `Required()` for fail-fast behavior
- **Sugar Syntax**: Short aliases like `gt.A()` for `gt.Array()`, `gt.S()` for `gt.String()`
- **Type Safety**: Compile-time type checking prevents type mismatches

### Example Overview

```go
// All test types can be chained and described
gt.Array(t, users).
    Describe("User validation").
    Required().
    Length(3).
    All(func(u User) bool { return u.ID > 0 })

gt.Map(t, config).
    Describef("Config for env %s", env).
    HasKey("database").
    At("database", func(t testing.TB, db string) {
        gt.String(t, db).HasPrefix("postgres://")
    })
```

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

Accepts array/slice of any type including primitive types and structs. Provides comprehensive testing methods for collections.

```go
colors := []string{"red", "blue", "yellow"}

// Basic equality and length
gt.Array(t, colors).
    Equal([]string{"red", "blue", "yellow"}). // Pass
    Length(3).                                // Pass
    NotEqual([]string{"red", "blue"})         // Pass

// Element checking
gt.Array(t, colors).
    Has("yellow").                            // Pass - contains element
    NotHas("orange").                         // Pass - doesn't contain element
    Contains([]string{"red", "blue"}).        // Pass - contains subsequence
    NotContains([]string{"red", "yellow"})    // Pass - doesn't contain subsequence

// Index-based operations
gt.Array(t, colors).
    EqualAt(0, "red").                        // Pass - check specific index
    NotEqualAt(1, "yellow")                   // Pass

// Length comparisons
gt.Array(t, colors).
    Longer(2).                                // Pass - length > 2
    Less(5)                                   // Pass - length < 5

// Advanced operations
gt.Array(t, colors).
    Distinct().                               // Pass - all elements unique
    Any(func(v string) bool {                 // Pass - at least one matches
        return v == "blue"
    }).
    All(func(v string) bool {                 // Pass - all elements match
        return len(v) > 2
    })

// Test individual elements with callback
gt.Array(t, colors).At(1, func(t testing.TB, v string) {
    gt.String(t, v).Equal("blue")             // Pass
})

// Find and test matching element
users := []User{{"Alice", 25}, {"Bob", 30}}
gt.Array(t, users).MatchThen(
    func(u User) bool { return u.Name == "Bob" },
    func(t testing.TB, u User) {
        gt.Number(t, u.Age).Equal(30)         // Pass
    })

// Sugar syntax
gt.A(t, colors).Has("blue")                   // Same as gt.Array()
```

### Map

Provides type-safe testing for maps with comprehensive key-value operations.

```go
colorMap := map[string]int{
    "red":    1,
    "yellow": 2,
    "blue":   5,
}

// Basic equality and length
gt.Map(t, colorMap).
    Equal(map[string]int{"red": 1, "yellow": 2, "blue": 5}). // Pass
    NotEqual(map[string]int{"red": 1, "yellow": 2}).         // Pass
    Length(3)                                                // Pass

// Key operations
gt.Map(t, colorMap).
    HasKey("blue").                           // Pass - key exists
    NotHasKey("orange").                      // Pass - key doesn't exist
    HasKey("purple")                          // Fail - key doesn't exist

// Value operations
gt.Map(t, colorMap).
    HasValue(5).                              // Pass - value exists
    NotHasValue(10).                          // Pass - value doesn't exist
    HasValue(99)                              // Fail - value doesn't exist

// Key-value pair operations
gt.Map(t, colorMap).
    HasKeyValue("yellow", 2).                 // Pass - pair exists
    NotHasKeyValue("red", 5).                 // Pass - pair doesn't exist
    HasKeyValue("blue", 1)                    // Fail - wrong value for key

// Index-based operations (similar to accessing map[key])
gt.Map(t, colorMap).
    EqualAt("red", 1).                        // Pass - map["red"] == 1
    NotEqualAt("blue", 1).                    // Pass - map["blue"] != 1
    EqualAt("orange", 1)                      // Fail - key doesn't exist

// Test individual values with callback
gt.Map(t, colorMap).At("blue", func(t testing.TB, v int) {
    gt.Number(t, v).Greater(3)                // Pass
})

// Sugar syntax
gt.M(t, colorMap).HasKey("red")               // Same as gt.Map()
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

### ExpectError

Helper function for conditional error testing based on expectations:

```go
// Expect an error to occur
err := someFunctionThatShouldFail()
gt.ExpectError(t, true, err)  // Pass if err != nil

// Expect no error to occur
err := someFunctionThatShouldSucceed()
gt.ExpectError(t, false, err) // Pass if err == nil

// Conditional error testing
shouldFail := true
err := conditionalFunction(shouldFail)
gt.ExpectError(t, shouldFail, err) // Pass based on shouldFail expectation
```

This is particularly useful for:
- Testing functions with conditional error behavior
- Parameterized tests where error expectation varies
- Reducing test code duplication when testing both success and failure cases

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

## Required Pattern (Fail-Fast Testing)

All test types support the `Required()` method which provides fail-fast behavior. When `Required()` is called, it checks if any previous test in the chain has failed. If so, it immediately stops the test execution using `t.FailNow()`, preventing subsequent tests from running.

### Basic Usage

```go
// Test will stop immediately if the first assertion fails
gt.Value(t, result).
    Describe("Critical validation step").
    Required().           // Stop here if previous test failed
    Equal(expected).      // This won't run if Required() triggered
    NotNil()              // This won't run either

// Vs. normal testing - all assertions run even if some fail
gt.Value(t, result).
    Equal(expected).      // Fails but continues
    NotNil()              // Still runs and might also fail
```

### Required with All Test Types

Every test type supports `Required()` for fail-fast behavior:

```go
// Value tests
gt.Value(t, user).Required().NotNil().Equal(expectedUser)

// Array tests
gt.Array(t, items).Required().Length(5).Has(expectedItem)

// Map tests
gt.Map(t, data).Required().HasKey("id").HasValue(123)

// String tests
gt.String(t, name).Required().IsNotEmpty().HasPrefix("user_")

// Number tests
gt.Number(t, count).Required().Greater(0).Less(100)

// Bool tests
gt.Bool(t, isValid).Required().True()

// Error tests
gt.NoError(t, err).Required() // Common pattern - stop if error occurs
gt.Error(t, err).Required().Contains("validation failed")

// File tests
gt.File(t, "config.json").Required().Exists().String(func(t testing.TB, content string) {
    gt.String(t, content).Contains("database")
})
```

### Required with Descriptions

`Required()` works seamlessly with `Describe()` and `Describef()` to provide meaningful error context:

```go
gt.Value(t, response).
    Describef("API response for user %d should be valid", userID).
    Required().           // Will show description if this fails
    NotNil().
    Equal(expectedResponse)

// Error output:
// API response for user 123 should be valid
// Previous test failed
```

### Common Patterns

```go
// Pattern 1: Critical setup validation
config := loadConfig()
gt.Value(t, config).
    Describe("Configuration must be loaded successfully").
    Required().
    NotNil()

// Pattern 2: Function return value validation
result, err := processData(input)
gt.NoError(t, err).Required()  // Stop if error
gt.Value(t, result).Required().NotNil().Equal(expected)

// Pattern 3: Multi-step validation
gt.Array(t, users).
    Describe("User list validation").
    Required().
    Length(expectedCount).
    All(func(u User) bool { return u.ID > 0 })

// Pattern 4: Map validation with early exit
gt.Map(t, apiResponse).
    Describe("API response structure validation").
    Required().
    HasKey("status").
    HasKey("data").
    At("status", func(t testing.TB, status string) {
        gt.String(t, status).Equal("success")
    })
```

### When to Use Required

- **Critical validations**: Use `Required()` when subsequent tests depend on the current assertion
- **Setup validation**: Validate test prerequisites before running main test logic  
- **Error handling**: Stop immediately when errors occur in functions that should succeed
- **Complex test chains**: Prevent cascading failures in multi-step validations
- **Resource validation**: Ensure files, connections, or configurations exist before using them

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