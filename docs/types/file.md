# File

`FileTest` is a type for testing file system operations. It provides file existence checking and content reading capabilities.

## Constructor

```go
gt.File(t testing.TB, path string) FileTest
gt.F(t testing.TB, path string) FileTest  // Short form
```

## Methods

| Method | Description |
|--------|-------------|
| `Exists()` | Verify file exists |
| `NotExists()` | Verify file does not exist |
| `String(f func(t, string))` | Pass file content as string to callback |
| `Reader(f func(t, io.Reader))` | Pass file Reader to callback |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### File Existence Check

```go
func TestFileExists(t *testing.T) {
    // Verify file exists
    gt.File(t, "testdata/config.json").Exists()

    // Verify file does not exist
    gt.File(t, "testdata/deleted.txt").NotExists()
}
```

### Verify File Content as String

```go
func TestFileContent(t *testing.T) {
    gt.File(t, "testdata/message.txt").String(func(t testing.TB, content string) {
        gt.String(t, content).Contains("Hello")
        gt.String(t, content).HasSuffix("\n")
    })
}
```

### Process File Content as Reader

```go
func TestFileReader(t *testing.T) {
    gt.File(t, "testdata/data.json").Reader(func(t testing.TB, r io.Reader) {
        var data Config
        err := json.NewDecoder(r).Decode(&data)
        gt.NoError(t, err)
        gt.String(t, data.Name).Equal("test")
    })
}
```

### Method Chaining

```go
func TestFileChaining(t *testing.T) {
    gt.File(t, "testdata/config.yaml").
        Describef("Config file for %s environment", env).
        Exists().
        Required().  // Stop immediately if not exists
        String(func(t testing.TB, content string) {
            gt.String(t, content).Contains("database:")
        })
}
```

## Common Patterns

### Verifying Generated Files

```go
func TestGeneratedFile(t *testing.T) {
    // Generate file
    err := generateReport("output/report.txt")
    gt.NoError(t, err).Required()

    // Verify generated file
    gt.File(t, "output/report.txt").
        Exists().
        String(func(t testing.TB, content string) {
            gt.String(t, content).Contains("Summary")
            gt.String(t, content).HasPrefix("Report:")
        })
}
```

### Verifying Temp File Cleanup

```go
func TestTempFileCleanup(t *testing.T) {
    tmpFile := createTempFile()
    defer cleanup()

    // Exists before processing
    gt.File(t, tmpFile).Exists()

    // Does not exist after cleanup
    cleanup()
    gt.File(t, tmpFile).NotExists()
}
```

### Verifying Config Files

```go
func TestConfigFile(t *testing.T) {
    gt.File(t, "config.json").
        Exists().
        Reader(func(t testing.TB, r io.Reader) {
            var config Config
            err := json.NewDecoder(r).Decode(&config)
            gt.NoError(t, err)

            gt.Number(t, config.Port).Greater(0)
            gt.String(t, config.Host).IsNotEmpty()
        })
}
```

## Notes

### File Paths

When using relative paths, they are relative to the current directory when running tests. Typically, this is the package directory.

```go
// Using testdata directory is common
gt.File(t, "testdata/fixture.json").Exists()
```

### Error Handling

`String()` and `Reader()` fail the test if file reading fails. If you're unsure whether the file exists, verify with `Exists()` first.

```go
gt.File(t, path).
    Exists().
    Required().  // Stop here if not exists
    String(func(t testing.TB, content string) {
        // Only executed if file exists
    })
```
