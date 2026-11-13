# Time

`TimeTest` is a type for testing `time.Time` values. It checks that a time is close to an expected time, which is what a test needs for a timestamp that was stored with reduced precision or taken from the clock a moment earlier.

## Constructor

```go
gt.Time(t testing.TB, actual time.Time) TimeTest
```

## Methods

| Method | Description |
|--------|-------------|
| `Near(expect time.Time, tolerance time.Duration)` | Check if value is less than `tolerance` away from `expect`, earlier or later |
| `Required()` | Stop immediately if previous test failed |
| `Describe(string)` | Set test description |
| `Describef(format, args...)` | Set formatted test description |

## Examples

### Timestamp Stored With Reduced Precision

A database that keeps microseconds returns a time that differs from the one written by less than a microsecond. Compare with a tolerance rather than with equality.

```go
func TestStoredTimestamp(t *testing.T) {
    now := time.Now()
    saved := gt.R1(repo.Save(ctx, record, now)).NoError(t)

    gt.Time(t, saved.CreatedAt).Near(now, time.Millisecond)
}
```

### Both Sides Are Checked

`Near` fails when the value is too early as well as when it is too late. A hand-written `got.Sub(want) < time.Second` passes for every value earlier than `want`; `Near` does not.

```go
func TestNear(t *testing.T) {
    base := time.Date(2024, 5, 1, 12, 0, 0, 0, time.UTC)

    gt.Time(t, base.Add(500*time.Millisecond)).Near(base, time.Second)  // Pass
    gt.Time(t, base.Add(-500*time.Millisecond)).Near(base, time.Second) // Pass
    gt.Time(t, base.Add(-5*time.Minute)).Near(base, time.Second)        // Fail
    gt.Time(t, base.Add(time.Second)).Near(base, time.Second)           // Fail: exactly the tolerance away
}
```

### Failure Message

A failure shows both times and the signed difference:

```
time is not near
expected: 2024-05-01T12:00:00Z (tolerance 1s)
actual:   2024-05-01T11:55:00Z (off by -5m0s)
```

### With a Description

```go
gt.Time(t, run.EndedAt).
    Describe("the run is dated by when it finished").
    Near(finishedAt, time.Second)
```

### A Pointer to a Time

`gt.Time` takes a `time.Time`. Check a pointer for nil first, then dereference it.

```go
gt.Value(t, summary.LastIngestedAt).NotNil().Required()
gt.Time(t, *summary.LastIngestedAt).Near(endedAt, time.Second)
```
