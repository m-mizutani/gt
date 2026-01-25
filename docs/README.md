# gt Documentation

gt is a type-safe testing library for Go that leverages generics. It provides IDE support and compile-time type checking.

## Table of Contents

### Getting Started

- [Getting Started](./getting-started.md) - Installation and quick start

### Test Type Reference

Detailed usage and method reference for each test type.

| Type | Description | Documentation |
|------|-------------|---------------|
| **Value** | General value testing | [value.md](./types/value.md) |
| **Array** | Slice/array testing | [array.md](./types/array.md) |
| **Map** | Map testing | [map.md](./types/map.md) |
| **Number** | Numeric comparison | [number.md](./types/number.md) |
| **String** | String testing | [string.md](./types/string.md) |
| **Bool** | Boolean testing | [bool.md](./types/bool.md) |
| **Error** | Error testing | [error.md](./types/error.md) |
| **File** | File testing | [file.md](./types/file.md) |
| **Cast** | Type assertion | [cast.md](./types/cast.md) |
| **Return** | Function return testing | [return.md](./types/return.md) |

### Pattern Guide

Common patterns and best practices.

- [Method Chaining](./patterns/method-chaining.md) - Chain methods for fluent tests
- [Required Pattern](./patterns/required.md) - Stop test immediately on failure
- [Describe Pattern](./patterns/describe.md) - Add descriptions to improve error messages

### Claude Code Integration

- [Skill](./skills/SKILL.md) - Claude Code skill for writing tests with gt

## Quick Reference

```go
import "github.com/m-mizutani/gt"

func TestExample(t *testing.T) {
    // Value comparison
    gt.Value(t, actual).Equal(expected)

    // Array testing
    gt.Array(t, items).Length(3).Has(item)

    // Map testing
    gt.Map(t, config).HasKey("key").EqualAt("key", value)

    // Number comparison
    gt.Number(t, count).Greater(0).LessOrEqual(100)

    // String testing
    gt.String(t, email).Contains("@").HasSuffix(".com")

    // Error testing
    gt.NoError(t, err).Required()

    // Function return testing
    result := gt.R1(parseJSON(input)).NoError(t)
}
```

## Short Aliases

All constructors have short forms:

| Full Form | Short Form |
|-----------|------------|
| `gt.Value(t, v)` | `gt.V(t, v)` |
| `gt.Array(t, a)` | `gt.A(t, a)` |
| `gt.Map(t, m)` | `gt.M(t, m)` |
| `gt.Number(t, n)` | `gt.N(t, n)` |
| `gt.String(t, s)` | `gt.S(t, s)` |
| `gt.Bool(t, b)` | `gt.B(t, b)` |
| `gt.File(t, path)` | `gt.F(t, path)` |
| `gt.Cast[T](t, v)` | `gt.C[T](t, v)` |
| `gt.Return1(...)` | `gt.R1(...)` |
| `gt.Return2(...)` | `gt.R2(...)` |
| `gt.Return3(...)` | `gt.R3(...)` |
