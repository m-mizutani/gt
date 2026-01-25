# gt: Generics based Test library for Go

[![Go Reference](https://pkg.go.dev/badge/github.com/m-mizutani/gt.svg)](https://pkg.go.dev/github.com/m-mizutani/gt) [![test](https://github.com/m-mizutani/gt/actions/workflows/test.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/test.yml) [![gosec](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/gosec.yml) [![lint](https://github.com/m-mizutani/gt/actions/workflows/lint.yml/badge.svg)](https://github.com/m-mizutani/gt/actions/workflows/lint.yml)

Type-safe test assertions for Go with IDE support and compile-time type checking.

```go
color := "blue"
gt.Value(t, color).Equal("blue")   // Pass
gt.Value(t, color).Equal("orange") // Fail
// gt.Value(t, color).Equal(5)     // Compile error!
```

## Installation

```bash
go get github.com/m-mizutani/gt
```

## Quick Start

```go
import "github.com/m-mizutani/gt"

func TestExample(t *testing.T) {
    // Value comparison
    gt.Value(t, user).Equal(expectedUser)

    // Array testing
    gt.Array(t, items).Length(3).Has(expectedItem)

    // Map testing
    gt.Map(t, config).HasKey("database").EqualAt("port", 5432)

    // String testing
    gt.String(t, email).Contains("@").HasSuffix(".com")

    // Number comparison
    gt.Number(t, count).Greater(0).LessOrEqual(100)

    // Error handling
    gt.NoError(t, err).Required()  // Fail fast if error
}
```

## Available Types

| Type | Constructor | Use for |
|------|-------------|---------|
| **Value** | `gt.Value(t, v)` | Any value (structs, primitives) |
| **Array** | `gt.Array(t, arr)` | Slices and arrays |
| **Map** | `gt.Map(t, m)` | Maps |
| **Number** | `gt.Number(t, n)` | Numeric comparisons (>, <, etc.) |
| **String** | `gt.String(t, s)` | String operations |
| **Bool** | `gt.Bool(t, b)` | Boolean checks |
| **Error** | `gt.Error(t, err)` | Error validation |
| **File** | `gt.File(t, path)` | File system checks |

Each type has a short alias: `gt.V()`, `gt.A()`, `gt.M()`, `gt.N()`, `gt.S()`, `gt.B()`, `gt.F()`

See [GoDoc](https://pkg.go.dev/github.com/m-mizutani/gt) for all available methods.

## Key Features

### Method Chaining

```go
gt.Array(t, users).
    Length(3).
    Has(admin).
    All(func(u User) bool { return u.Active })
```

### Fail Fast with Required()

```go
gt.NoError(t, err).Required()           // Stop immediately if error
gt.Value(t, config).Required().NotNil() // Stop if nil
```

### Descriptive Error Messages

```go
gt.Array(t, users).
    Describef("Users for tenant %s", tenantID).
    Length(5)

// Output:
// Users for tenant abc123
// array length is expected to be 5, but actual is 3
```

### Function Return Values

```go
// Handle (value, error) returns
result := gt.R1(parseJSON(input)).NoError(t)
gt.Value(t, result.Name).Equal("Alice")

// Test expected errors
gt.R1(parseJSON(badInput)).Error(t).Contains("invalid")
```

## Best Practices

Use the specialized type that matches your data:

```go
// Good - specialized types give better errors and methods
gt.Array(t, arr).Length(3)
gt.String(t, s).Contains("x")
gt.NoError(t, err)
gt.Number(t, n).Greater(5)

// Avoid - loses specialized functionality
gt.Value(t, len(arr)).Equal(3)
gt.Bool(t, strings.Contains(s, "x")).True()
gt.Value(t, err).Nil()
gt.Bool(t, n > 5).True()
```

## Why gt?

Traditional test libraries like [testify](https://github.com/stretchr/testify) use `interface{}`/`any`:

```go
// testify - no compile-time type checking
assert.Equal(t, "hello", 123)  // Compiles but fails at runtime
```

gt uses [Go Generics](https://go.dev/doc/tutorial/generics) (1.18+):

```go
// gt - caught at compile time
gt.Value(t, "hello").Equal(123)  // Compile error!
```

Benefits:
- Type mismatches caught before running tests
- Full IDE autocompletion
- Better error messages

## License

Apache License 2.0
