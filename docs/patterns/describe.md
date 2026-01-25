# Describe Pattern

`Describe()` and `Describef()` are methods that add descriptions to tests. When a test fails, the description is included in the error message, making it easier to identify which test failed.

## Basic Usage

### Describe

Set a fixed description:

```go
gt.Value(t, result).
    Describe("API response status").
    Equal(200)
```

### Describef

Set a formatted description:

```go
gt.Value(t, user.Age).
    Describef("Age for user %s (ID: %d)", user.Name, user.ID).
    GreaterOrEqual(0)
```

## Failure Output Examples

Without description:

```go
gt.Value(t, result).Equal(expected)
// Output:
// values are not matched
// - expected
// + actual
```

With description:

```go
gt.Value(t, result).
    Describef("Processing result for user %s", userID).
    Equal(expected)
// Output:
// Processing result for user alice123
// values are not matched
// - expected
// + actual
```

## Usage Patterns

### In Table-Driven Tests

```go
func TestTableDriven(t *testing.T) {
    testCases := []struct {
        name   string
        input  int
        expect int
    }{
        {"positive", 5, 10},
        {"negative", -5, -10},
        {"zero", 0, 0},
    }

    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            result := double(tc.input)
            gt.Number(t, result).
                Describef("double(%d)", tc.input).
                Equal(tc.expect)
        })
    }
}
```

### Distinguishing Multiple Assertions

```go
func TestMultipleAssertions(t *testing.T) {
    user := getUser()

    gt.String(t, user.Name).
        Describe("User name").
        IsNotEmpty()

    gt.String(t, user.Email).
        Describe("User email").
        Contains("@")

    gt.Bool(t, user.Active).
        Describe("User active status").
        True()
}
```

### Including Context Information

```go
func TestWithContext(t *testing.T) {
    for _, env := range []string{"dev", "staging", "prod"} {
        config := loadConfig(env)

        gt.Map(t, config).
            Describef("Config for %s environment", env).
            HasKey("database").
            HasKey("cache")
    }
}
```

### In API Tests

```go
func TestAPI(t *testing.T) {
    endpoints := []string{"/users", "/posts", "/comments"}

    for _, endpoint := range endpoints {
        resp := callAPI(endpoint)

        gt.Number(t, resp.StatusCode).
            Describef("Status code for %s", endpoint).
            Equal(200)

        gt.String(t, resp.ContentType).
            Describef("Content-Type for %s", endpoint).
            Contains("application/json")
    }
}
```

## Position in Method Chains

Descriptions are typically set at the beginning of the chain:

```go
// Recommended: description first
gt.Array(t, items).
    Describef("Items for category %s", category).
    Length(10).
    Has(expectedItem)

// Works but harder to read
gt.Array(t, items).
    Length(10).
    Describef("Items for category %s", category).  // In the middle
    Has(expectedItem)
```

## Usage with Multiple Test Types

Available for all test types:

```go
// Value
gt.Value(t, user).Describe("User object").NotNil()

// Array
gt.Array(t, items).Describef("Items (count: %d)", len(items)).Length(5)

// Map
gt.Map(t, config).Describe("Application config").HasKey("database")

// Number
gt.Number(t, score).Describef("Score for %s", player).Greater(0)

// String
gt.String(t, message).Describe("Error message").Contains("failed")

// Bool
gt.Bool(t, isValid).Describe("Validation result").True()

// Error
gt.Error(t, err).Describe("Expected validation error").Contains("invalid")

// File
gt.File(t, path).Describef("Config file at %s", path).Exists()
```

## Best Practices

### 1. Be Clear About What You're Testing

```go
// Good
gt.String(t, result).
    Describe("JSON parsing result").
    Contains("success")

// Bad (vague)
gt.String(t, result).
    Describe("result").
    Contains("success")
```

### 2. Include Dynamic Values

```go
// Good
gt.Number(t, count).
    Describef("Item count for user %s in category %s", userID, category).
    Equal(expected)

// Bad (hard to debug)
gt.Number(t, count).
    Describe("Item count").
    Equal(expected)
```

### 3. Don't Make It Too Long

```go
// Good
gt.Value(t, order).
    Describef("Order #%d", order.ID).
    NotNil()

// Bad (too long)
gt.Value(t, order).
    Describef("Order with ID %d created at %s by user %s in region %s with items %v",
        order.ID, order.CreatedAt, order.UserID, order.Region, order.Items).
    NotNil()
```

### 4. Distinguish from t.Run

When using subtests, leverage both `t.Run` name and `Describe`:

```go
func TestUsers(t *testing.T) {
    for _, user := range users {
        t.Run(user.Name, func(t *testing.T) {
            // t.Run name + Describe for details
            gt.String(t, user.Email).
                Describe("Email format").
                Contains("@")

            gt.Number(t, user.Age).
                Describe("Age validation").
                GreaterOrEqual(0)
        })
    }
}
```
