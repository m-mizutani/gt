# Map

`MapTest[K, V]` is a type for testing maps. It provides map-specific operations like key/value existence checks and value validation for specific keys.

## Constructor

```go
gt.Map[K comparable, V any](t testing.TB, actual map[K]V) MapTest[K, V]
gt.M[K comparable, V any](t testing.TB, actual map[K]V) MapTest[K, V]  // Short form
```

## Methods

### Comparison Methods

| Method | Description |
|--------|-------------|
| `Equal(expect map[K]V)` | Check if map equals exactly |
| `NotEqual(expect map[K]V)` | Check if map does not equal |
| `EqualAt(key K, expect V)` | Check if value at key equals expect |
| `NotEqualAt(key K, expect V)` | Check if value at key does not equal expect |

### Existence Methods

| Method | Description |
|--------|-------------|
| `HasKey(expect K)` | Check if key exists |
| `NotHasKey(expect K)` | Check if key does not exist |
| `HasValue(expect V)` | Check if value exists |
| `NotHasValue(expect V)` | Check if value does not exist |
| `HasKeyValue(key K, value V)` | Check if key-value pair exists |
| `NotHasKeyValue(key K, value V)` | Check if key-value pair does not exist |

### Other

| Method | Description |
|--------|-------------|
| `Length(expect int)` | Check if map size equals expect |
| `At(key K, f func(t, V))` | Execute callback with value at key |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Basic Comparison

```go
func TestMapEqual(t *testing.T) {
    config := map[string]int{
        "port":    8080,
        "timeout": 30,
    }

    gt.Map(t, config).Equal(map[string]int{
        "port":    8080,
        "timeout": 30,
    }) // Pass

    gt.Map(t, config).NotEqual(map[string]int{
        "port": 3000,
    }) // Pass
}
```

### Key Existence

```go
func TestMapHasKey(t *testing.T) {
    settings := map[string]string{
        "theme":    "dark",
        "language": "en",
    }

    gt.Map(t, settings).HasKey("theme")      // Pass
    gt.Map(t, settings).NotHasKey("version") // Pass
}
```

### Value Existence

```go
func TestMapHasValue(t *testing.T) {
    scores := map[string]int{
        "Alice": 100,
        "Bob":   85,
        "Carol": 90,
    }

    gt.Map(t, scores).HasValue(100)    // Pass
    gt.Map(t, scores).NotHasValue(50)  // Pass
}
```

### Key-Value Pair Check

```go
func TestMapHasKeyValue(t *testing.T) {
    data := map[string]int{
        "count": 42,
        "limit": 100,
    }

    gt.Map(t, data).HasKeyValue("count", 42)       // Pass
    gt.Map(t, data).NotHasKeyValue("count", 0)     // Pass
    gt.Map(t, data).NotHasKeyValue("missing", 42)  // Pass
}
```

### Value at Specific Key

```go
func TestMapEqualAt(t *testing.T) {
    config := map[string]string{
        "host": "localhost",
        "port": "8080",
    }

    gt.Map(t, config).EqualAt("host", "localhost")  // Pass
    gt.Map(t, config).NotEqualAt("port", "3000")    // Pass

    // Accessing non-existent key causes error
    gt.Map(t, config).EqualAt("missing", "value")   // Fail (key not found)
}
```

### Size Check

```go
func TestMapLength(t *testing.T) {
    m := map[int]string{
        1: "one",
        2: "two",
        3: "three",
    }

    gt.Map(t, m).Length(3) // Pass
}
```

### Detailed Testing with Callback

```go
type Config struct {
    Enabled bool
    Value   int
}

func TestMapAt(t *testing.T) {
    configs := map[string]Config{
        "feature_a": {Enabled: true, Value: 100},
        "feature_b": {Enabled: false, Value: 0},
    }

    gt.Map(t, configs).At("feature_a", func(t testing.TB, c Config) {
        gt.Bool(t, c.Enabled).True()
        gt.Number(t, c.Value).Greater(0)
    })
}
```

### Method Chaining

```go
func TestMapChaining(t *testing.T) {
    response := getAPIResponse()

    gt.Map(t, response).
        Describef("API response for user %d", userID).
        HasKey("status").
        Required().
        EqualAt("status", "success").
        HasKey("data")
}
```

## Common Patterns

### Configuration Testing

```go
func TestConfiguration(t *testing.T) {
    config := loadConfig()

    gt.Map(t, config).
        HasKey("database").
        HasKey("server").
        EqualAt("environment", "test")
}
```

### JSON Response Testing

```go
func TestJSONResponse(t *testing.T) {
    var response map[string]any
    json.Unmarshal(body, &response)

    gt.Map(t, response).
        HasKey("id").
        HasKeyValue("status", "ok")
}
```

### Multiple Conditions

```go
func TestMapMultipleConditions(t *testing.T) {
    data := fetchData()

    gt.Map(t, data).
        Length(3).
        HasKey("required_field").
        NotHasKey("deprecated_field").
        EqualAt("version", "2.0")
}
```
