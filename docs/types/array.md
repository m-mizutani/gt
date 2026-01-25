# Array

`ArrayTest[T]` is a type for testing slices and arrays. It provides array-specific operations like element existence checks, length validation, and ordering tests.

## Constructor

```go
gt.Array[T any](t testing.TB, actual []T) ArrayTest[T]
gt.A[T any](t testing.TB, actual []T) ArrayTest[T]  // Short form
```

## Methods

### Comparison Methods

| Method | Description |
|--------|-------------|
| `Equal(expect []T)` | Check if array equals exactly (including order) |
| `NotEqual(expect []T)` | Check if array does not equal |
| `UnorderedEqual(expect []T)` | Check if arrays have same elements regardless of order |
| `NotUnorderedEqual(expect []T)` | Check if arrays don't have same elements |

### Length Methods

| Method | Description |
|--------|-------------|
| `Length(expect int)` | Check if array length equals expect |
| `Longer(expect int)` | Check if array length is greater than expect |
| `Less(expect int)` | Check if array length is less than expect |

### Element Methods

| Method | Description |
|--------|-------------|
| `Has(expect T)` | Check if array contains element |
| `NotHas(expect T)` | Check if array does not contain element |
| `Contains(expect []T)` | Check if array contains subsequence |
| `NotContains(expect []T)` | Check if array does not contain subsequence |
| `EqualAt(idx int, expect T)` | Check if element at index equals expect |
| `NotEqualAt(idx int, expect T)` | Check if element at index does not equal expect |
| `Distinct()` | Check if all elements are unique |

### Callback Methods

| Method | Description |
|--------|-------------|
| `At(idx int, f func(t, v T))` | Execute callback with element at index |
| `Any(f func(v T) bool)` | Check if any element returns true |
| `All(f func(v T) bool)` | Check if all elements return true |
| `MatchThen(match func(T) bool, then func(t, T))` | Execute callback with first matching element |

### Other

| Method | Description |
|--------|-------------|
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Basic Comparison

```go
func TestArrayEqual(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    gt.Array(t, numbers).Equal([]int{1, 2, 3, 4, 5}) // Pass
    gt.Array(t, numbers).Equal([]int{1, 2, 3})       // Fail
}
```

### Unordered Comparison

```go
func TestArrayUnorderedEqual(t *testing.T) {
    items := []string{"apple", "banana", "cherry"}

    // Pass even with different order
    gt.Array(t, items).UnorderedEqual([]string{"cherry", "apple", "banana"}) // Pass

    // Duplicate elements are considered
    gt.Array(t, []int{1, 1, 2}).UnorderedEqual([]int{2, 1, 1}) // Pass
    gt.Array(t, []int{1, 1, 2}).UnorderedEqual([]int{1, 2, 2}) // Fail
}
```

### Length Validation

```go
func TestArrayLength(t *testing.T) {
    items := []string{"a", "b", "c", "d"}

    gt.Array(t, items).Length(4)  // Pass
    gt.Array(t, items).Longer(3)  // Pass (4 > 3)
    gt.Array(t, items).Less(5)    // Pass (4 < 5)
}
```

### Element Existence

```go
func TestArrayHas(t *testing.T) {
    fruits := []string{"apple", "banana", "cherry"}

    gt.Array(t, fruits).Has("banana")    // Pass
    gt.Array(t, fruits).NotHas("orange") // Pass
}
```

### Subsequence Check

```go
func TestArrayContains(t *testing.T) {
    numbers := []int{1, 2, 3, 4, 5}

    // Check for contiguous subsequence
    gt.Array(t, numbers).Contains([]int{2, 3, 4}) // Pass
    gt.Array(t, numbers).Contains([]int{1, 3, 5}) // Fail (not contiguous)
}
```

### Index Access

```go
func TestArrayAt(t *testing.T) {
    items := []string{"first", "second", "third"}

    gt.Array(t, items).EqualAt(0, "first")   // Pass
    gt.Array(t, items).EqualAt(1, "second")  // Pass
    gt.Array(t, items).NotEqualAt(2, "last") // Pass

    // Callback form
    gt.Array(t, items).At(1, func(t testing.TB, v string) {
        gt.String(t, v).Contains("sec")
    })
}
```

### Condition-based Validation

```go
type User struct {
    Name   string
    Active bool
}

func TestArrayConditions(t *testing.T) {
    users := []User{
        {Name: "Alice", Active: true},
        {Name: "Bob", Active: true},
        {Name: "Carol", Active: false},
    }

    // Any element matches condition
    gt.Array(t, users).Any(func(u User) bool {
        return u.Name == "Bob"
    }) // Pass

    // All elements match condition
    gt.Array(t, users).All(func(u User) bool {
        return u.Name != ""
    }) // Pass

    // Test matching element further
    gt.Array(t, users).MatchThen(
        func(u User) bool { return u.Name == "Carol" },
        func(t testing.TB, u User) {
            gt.Bool(t, u.Active).False()
        },
    )
}
```

### Uniqueness Check

```go
func TestArrayDistinct(t *testing.T) {
    unique := []int{1, 2, 3, 4, 5}
    gt.Array(t, unique).Distinct() // Pass

    duplicates := []int{1, 2, 3, 2, 5}
    gt.Array(t, duplicates).Distinct() // Fail
}
```

### Method Chaining

```go
func TestArrayChaining(t *testing.T) {
    users := getUsers()

    gt.Array(t, users).
        Describef("Users from API").
        Length(3).
        Required().  // Stop if length doesn't match
        Has(admin).
        All(func(u User) bool { return u.Active })
}
```

## Common Patterns

### Testing Only Slice Length

```go
// Recommended
gt.Array(t, items).Length(3)

// Not recommended
gt.Value(t, len(items)).Equal(3)
```

### Testing Empty Slice

```go
gt.Array(t, items).Length(0)
```

### Testing Content Regardless of Order

```go
gt.Array(t, actual).UnorderedEqual(expected)
```
