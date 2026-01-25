# Method Chaining

All gt test types support method chaining. This allows you to fluently chain multiple verifications together.

## Basic Concept

Each method returns its own type, allowing you to call methods consecutively:

```go
gt.Array(t, items).
    Length(3).
    Has(expectedItem).
    All(func(item Item) bool { return item.Active })
```

## Practical Examples

### Array Chaining

```go
func TestArrayChaining(t *testing.T) {
    users := getUsers()

    gt.Array(t, users).
        Length(5).                              // Has 5 elements
        Has(admin).                             // Contains admin user
        All(func(u User) bool { return u.ID > 0 })  // All IDs are positive
}
```

### String Chaining

```go
func TestStringChaining(t *testing.T) {
    email := getEmail()

    gt.String(t, email).
        IsNotEmpty().
        Contains("@").
        HasSuffix(".com").
        NotContains(" ")  // No spaces
}
```

### Map Chaining

```go
func TestMapChaining(t *testing.T) {
    config := loadConfig()

    gt.Map(t, config).
        HasKey("database").
        HasKey("server").
        NotHasKey("deprecated_field").
        EqualAt("environment", "production")
}
```

### Number Chaining

```go
func TestNumberChaining(t *testing.T) {
    score := calculateScore()

    gt.Number(t, score).
        GreaterOrEqual(0).   // 0 or more
        LessOrEqual(100)     // 100 or less
}
```

## Required within Chains

Insert `Required()` in the chain to stop immediately if tests up to that point have failed:

```go
func TestChainWithRequired(t *testing.T) {
    data := fetchData()

    gt.Array(t, data).
        Length(3).
        Required().  // Stop here if length is not 3
        EqualAt(0, firstExpected).
        EqualAt(1, secondExpected).
        EqualAt(2, thirdExpected)
}
```

In this example, if `Length(3)` fails, `EqualAt` methods are not executed. This prevents potential out-of-bounds errors from index access.

## Describe within Chains

Descriptions are typically set at the beginning of the chain:

```go
func TestChainWithDescribe(t *testing.T) {
    response := callAPI(endpoint)

    gt.Map(t, response).
        Describef("API response from %s", endpoint).
        HasKey("status").
        EqualAt("status", "success").
        HasKey("data")
}
```

The description is included in failure messages, making it easier to identify which test failed.

## Splitting Chains

For complex tests, you can also split chains:

```go
func TestSplitChain(t *testing.T) {
    user := getUser()

    // First verify existence
    userTest := gt.Value(t, user).NotNil().Required()

    // Detailed verification
    gt.String(t, user.Name).IsNotEmpty()
    gt.String(t, user.Email).Contains("@")
    gt.Bool(t, user.Active).True()
}
```

## Combining with Table-Driven Tests

```go
func TestTableDriven(t *testing.T) {
    testCases := []struct {
        name     string
        input    string
        wantLen  int
        contains string
    }{
        {"empty", "", 0, ""},
        {"hello", "hello world", 11, "world"},
        {"numbers", "12345", 5, "3"},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := process(tc.input)

            // Verify multiple conditions with method chaining
            if tc.wantLen > 0 {
                gt.String(t, result).
                    Describef("Result for input %q", tc.input).
                    Contains(tc.contains)
            }
        })
    }
}
```

## Best Practices

### 1. Chain in Logical Order

```go
// Good: existence check → basic verification → detailed verification
gt.Array(t, items).
    Length(3).       // First check length
    Required().      // Stop if length is wrong
    Has(important).  // Check important element exists
    All(validator)   // Verify all elements
```

### 2. Place Required at Appropriate Points

Use when prerequisites not being met makes subsequent tests meaningless:

```go
gt.Value(t, user).
    NotNil().
    Required()  // No point in continuing if nil

gt.String(t, user.Name).Equal("Alice")
```

### 3. Put Description First

```go
gt.String(t, result).
    Describef("Processing result for user %s", userID).  // First
    IsNotEmpty().
    Contains("success")
```

### 4. Line Breaks for Readability

```go
// Good: one method per line
gt.Array(t, users).
    Length(10).
    Has(admin).
    All(isActive)

// Avoid: cramming into one line
gt.Array(t, users).Length(10).Has(admin).All(isActive)
```
