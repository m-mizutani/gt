package basic_test

import (
	"testing"

	"github.com/m-mizutani/gt"
)

func TestBasicUsage(t *testing.T) {
	// Value comparison
	user := struct{ Name string }{"Alice"}
	expectedUser := struct{ Name string }{"Alice"}
	gt.Value(t, user).Equal(expectedUser)

	// Array testing
	items := []string{"apple", "banana", "cherry"}
	gt.Array(t, items).Length(3).Has("banana")

	// Map testing
	config := map[string]int{"port": 8080, "timeout": 30}
	gt.Map(t, config).HasKey("port").EqualAt("port", 8080)

	// String testing
	email := "user@example.com"
	gt.String(t, email).Contains("@").HasSuffix(".com")

	// Number comparison
	count := 42
	gt.Number(t, count).Greater(0).LessOrEqual(100)

	// Error handling
	var err error // nil error
	gt.NoError(t, err).Required()
}

func TestArrayOperations(t *testing.T) {
	users := []struct {
		Name   string
		Active bool
	}{
		{"Alice", true},
		{"Bob", true},
		{"Carol", false},
	}

	gt.Array(t, users).
		Length(3).
		All(func(u struct{ Name string; Active bool }) bool {
			return u.Name != ""
		})
}

func TestFunctionReturns(t *testing.T) {
	// Simulate function that returns (value, error)
	parseData := func(input string) (string, error) {
		return "parsed: " + input, nil
	}

	result := gt.R1(parseData("test")).NoError(t)
	gt.String(t, result).HasPrefix("parsed:")
}

func TestDescriptions(t *testing.T) {
	users := []string{"Alice", "Bob"}

	gt.Array(t, users).
		Describef("User list should have %d members", 2).
		Length(2)
}
