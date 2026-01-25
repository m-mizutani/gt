# String

`StringTest` is a type for testing strings. It provides string-specific operations like substring search, prefix/suffix checking, and regex matching.

## Constructor

```go
gt.String(t testing.TB, actual string) StringTest
gt.S(t testing.TB, actual string) StringTest  // Short form
```

## Methods

### Comparison Methods

| Method | Description |
|--------|-------------|
| `Equal(expect string)` | Check if string equals expect |
| `NotEqual(expect string)` | Check if string does not equal expect |

### Empty String Check

| Method | Description |
|--------|-------------|
| `IsEmpty()` | Check if string is empty |
| `IsNotEmpty()` | Check if string is not empty |

### Substring Check

| Method | Description |
|--------|-------------|
| `Contains(sub string)` | Check if string contains substring |
| `NotContains(sub string)` | Check if string does not contain substring |
| `ContainsAny(substrs ...string)` | Check if string contains any of the substrings |
| `ContainsNone(substrs ...string)` | Check if string contains none of the substrings |

### Prefix/Suffix Check

| Method | Description |
|--------|-------------|
| `HasPrefix(prefix string)` | Check if string starts with prefix |
| `NotHasPrefix(prefix string)` | Check if string does not start with prefix |
| `HasSuffix(suffix string)` | Check if string ends with suffix |
| `NotHasSuffix(suffix string)` | Check if string does not end with suffix |

### Regex Matching

| Method | Description |
|--------|-------------|
| `Match(pattern string)` | Check if string matches regex pattern |
| `NotMatch(pattern string)` | Check if string does not match regex pattern |

### Other

| Method | Description |
|--------|-------------|
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Basic Comparison

```go
func TestStringEqual(t *testing.T) {
    name := "Alice"

    gt.String(t, name).Equal("Alice")     // Pass
    gt.String(t, name).NotEqual("Bob")    // Pass
}
```

### Empty String Check

```go
func TestStringEmpty(t *testing.T) {
    empty := ""
    filled := "hello"

    gt.String(t, empty).IsEmpty()      // Pass
    gt.String(t, filled).IsNotEmpty()  // Pass
}
```

### Substring Search

```go
func TestStringContains(t *testing.T) {
    message := "Hello, World!"

    gt.String(t, message).Contains("World")      // Pass
    gt.String(t, message).NotContains("Goodbye") // Pass
}
```

### Multiple Substring Check

```go
func TestStringContainsMultiple(t *testing.T) {
    log := "ERROR: Connection timeout"

    // Pass if contains any
    gt.String(t, log).ContainsAny("ERROR", "WARN", "INFO") // Pass

    // Pass if contains none
    gt.String(t, log).ContainsNone("DEBUG", "TRACE") // Pass
}
```

### Prefix and Suffix

```go
func TestStringPrefixSuffix(t *testing.T) {
    filename := "document.pdf"

    gt.String(t, filename).HasPrefix("doc")    // Pass
    gt.String(t, filename).HasSuffix(".pdf")   // Pass
    gt.String(t, filename).NotHasPrefix("img") // Pass
}
```

### Regex Matching

```go
func TestStringMatch(t *testing.T) {
    email := "user@example.com"

    // Email pattern
    gt.String(t, email).Match(`^[a-z]+@[a-z]+\.[a-z]+$`) // Pass

    // Check no digits
    gt.String(t, email).NotMatch(`\d`) // Pass
}
```

### Method Chaining

```go
func TestStringChaining(t *testing.T) {
    url := "https://api.example.com/v2/users"

    gt.String(t, url).
        HasPrefix("https://").
        Contains("/v2/").
        HasSuffix("/users")
}
```

### Test with Description

```go
func TestStringWithDescription(t *testing.T) {
    response := getAPIResponse()

    gt.String(t, response.ContentType).
        Describef("Content-Type header for %s", endpoint).
        Contains("application/json")
}
```

## Common Patterns

### Email Validation

```go
func TestEmail(t *testing.T) {
    email := getEmail()

    gt.String(t, email).
        Contains("@").
        Match(`^[^@]+@[^@]+\.[^@]+$`)
}
```

### File Path Validation

```go
func TestFilePath(t *testing.T) {
    path := getFilePath()

    gt.String(t, path).
        HasPrefix("/data/").
        HasSuffix(".json").
        NotContains("..")  // Prevent path traversal
}
```

### Log Message Validation

```go
func TestLogMessage(t *testing.T) {
    log := captureLog()

    gt.String(t, log).
        Contains("[INFO]").
        Contains("request completed").
        NotContains("[ERROR]")
}
```

### URL Validation

```go
func TestURL(t *testing.T) {
    url := buildURL()

    gt.String(t, url).
        HasPrefix("https://").
        Contains("/api/v1/").
        NotContains(" ")  // No spaces
}
```

## Why Use String

```go
// Using gt.String (recommended)
gt.String(t, s).Contains("error")

// Using gt.Bool (not recommended)
gt.Bool(t, strings.Contains(s, "error")).True()
```

Using `String`:
- Makes test intent clear
- Error messages include actual string value
- Chain multiple conditions with method chaining
