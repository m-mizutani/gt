---
paths:
  - "**/*_test.go"
---

# gt Test Library Rules

This project uses [gt](https://github.com/m-mizutani/gt) for type-safe test assertions.

## Rules

1. **Always use specialized types** - Don't use `gt.Value` or `gt.Bool` when a specialized type exists
2. **Use `Required()` for critical checks** - Stop test early if prerequisite fails
3. **Add descriptions for complex assertions** - Use `Describef()` for context

## Type Selection

| Data Type | Use | NOT |
|-----------|-----|-----|
| `[]T` (slice) | `gt.Array(t, arr).Length(3)` | `gt.Value(t, len(arr)).Equal(3)` |
| `string` | `gt.String(t, s).Contains("x")` | `gt.Bool(t, strings.Contains(s, "x")).True()` |
| `error` | `gt.NoError(t, err)` | `gt.Value(t, err).Nil()` |
| `int`, `float`, etc. | `gt.Number(t, n).Greater(5)` | `gt.Bool(t, n > 5).True()` |
| `map[K]V` | `gt.Map(t, m).HasKey("k")` | manual check with `gt.Bool` |
| `bool` | `gt.Bool(t, b).True()` | `gt.Value(t, b).Equal(true)` |

## Available Methods

**Value** - `gt.Value(t, v)` / `gt.V(t, v)`
- `Equal(T)`, `NotEqual(T)`, `Nil()`, `NotNil()`, `In(...T)`

**Array** - `gt.Array(t, arr)` / `gt.A(t, arr)`
- `Equal([]T)`, `NotEqual([]T)`, `Length(int)`, `Longer(int)`, `Less(int)`
- `Has(T)`, `NotHas(T)`, `Contains([]T)`, `NotContains([]T)`
- `EqualAt(idx, T)`, `NotEqualAt(idx, T)`
- `UnorderedEqual([]T)`, `NotUnorderedEqual([]T)`
- `Any(func(T) bool)`, `All(func(T) bool)`, `Distinct()`
- `At(idx, func(t, T))`, `MatchThen(match, then)`

**Map** - `gt.Map(t, m)` / `gt.M(t, m)`
- `Equal(map[K]V)`, `NotEqual(map[K]V)`, `Length(int)`
- `HasKey(K)`, `NotHasKey(K)`, `HasValue(V)`, `NotHasValue(V)`
- `HasKeyValue(K, V)`, `NotHasKeyValue(K, V)`
- `EqualAt(K, V)`, `NotEqualAt(K, V)`, `At(K, func(t, V))`

**Number** - `gt.Number(t, n)` / `gt.N(t, n)`
- `Equal(T)`, `NotEqual(T)`
- `Greater(T)`, `GreaterOrEqual(T)`, `Less(T)`, `LessOrEqual(T)`

**String** - `gt.String(t, s)` / `gt.S(t, s)`
- `Equal(string)`, `NotEqual(string)`
- `Contains(string)`, `NotContains(string)`
- `ContainsAny(...string)`, `ContainsNone(...string)`
- `HasPrefix(string)`, `NotHasPrefix(string)`
- `HasSuffix(string)`, `NotHasSuffix(string)`
- `Match(pattern)`, `NotMatch(pattern)`
- `IsEmpty()`, `IsNotEmpty()`

**Bool** - `gt.Bool(t, b)` / `gt.B(t, b)`
- `True()`, `False()`

**Error**
- `gt.Error(t, err)` - expects error (fails if nil)
- `gt.NoError(t, err)` - expects no error
- `gt.ErrorAs[T](t, err, callback)` - type assertion
- `.Is(error)`, `.IsNot(error)`, `.Contains(string)`

**File** - `gt.File(t, path)` / `gt.F(t, path)`
- `Exists()`, `NotExists()`
- `String(func(t, string))`, `Reader(func(t, io.Reader))`

**Cast** - `gt.Cast[T](t, v)` / `gt.C[T](t, v)`
- Returns T directly, fails test if type assertion fails

**Return** - `gt.R1(f())`, `gt.R2(f())`, `gt.R3(f())`
- `.NoError(t)` - returns value(s) if no error
- `.Error(t)` - returns ErrorTest if error expected

## Common Patterns

```go
// Error handling with fail-fast
result, err := doSomething()
gt.NoError(t, err).Required()
gt.Value(t, result).Equal(expected)

// Array validation
gt.Array(t, users).
    Length(3).
    All(func(u User) bool { return u.Active })

// Function return testing
data := gt.R1(parseJSON(input)).NoError(t)
gt.String(t, data.Name).Equal("Alice")
```
