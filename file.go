package gt

import (
	"fmt"
	"io"
	"os"
	"testing"
)

type FileTest struct {
	TestMeta
	path string
}

// File provides FileTest that has basic comparison methods
func File(t testing.TB, path string) FileTest {
	t.Helper()
	return FileTest{
		TestMeta: TestMeta{t: t},
		path:     path,
	}
}

// F is sugar syntax of File
func F(t testing.TB, path string) FileTest {
	t.Helper()
	return File(t, path)
}

// Describe sets a description for the test. The description will be displayed when the test fails.
func (x FileTest) Describe(description string) FileTest {
	x.setDesc(description)
	return x
}

// Describef sets a formatted description for the test. The description will be displayed when the test fails.
func (x FileTest) Describef(format string, args ...any) FileTest {
	x.setDescf(format, args...)
	return x
}

// Exists check if file exists
//
//	gt.File(t, "testdata/file.txt").Exists() // Pass
//	gt.File(t, "testdata/no-file.txt").Exists() // Fail
func (x FileTest) Exists() FileTest {
	x.t.Helper()
	if !EvalFileExists(x.path) {
		msg := fmt.Sprintf("file should exist, %s", x.path)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// NotExists check if file does not exist
//
//	gt.File(t, "testdata/file.txt").NotExists() // Fail
//	gt.File(t, "testdata/no-file.txt").NotExists() // Pass
func (x FileTest) NotExists() FileTest {
	x.t.Helper()
	if EvalFileExists(x.path) {
		msg := fmt.Sprintf("file should not exist, %s", x.path)
		x.t.Error(formatErrorMessage(x.description, msg))
	}

	return x
}

// String calls f with file content
//
//	gt.File(t, "testdata/file.txt").String(func(t testing.TB, s string) {
//	   gt.Equal(t, s, "hello")
//	})
func (x FileTest) String(f func(t testing.TB, s string)) FileTest {
	x.t.Helper()
	data, err := os.ReadFile(x.path)
	if err != nil {
		msg := fmt.Sprintf("failed to read file, %s", x.path)
		x.t.Error(formatErrorMessage(x.description, msg))
		return x
	}

	f(x.t, string(data))
	return x
}

// Reader calls f with file content
//
//	gt.File(t, "testdata/file.txt").Reader(func(t testing.TB, r io.Reader) {
//	   // Read file content from r
//	})
func (x FileTest) Reader(f func(testing.TB, io.Reader)) FileTest {
	x.t.Helper()
	r, err := os.Open(x.path)
	if err != nil {
		msg := fmt.Sprintf("failed to open file, %s", x.path)
		x.t.Error(formatErrorMessage(x.description, msg))
		return x
	}
	defer r.Close()

	f(x.t, r)
	return x
}

// Required check if error has occurred in previous test. If errors has been occurred in previous test, it immediately stop test by t.FailNow().
func (x FileTest) Required() FileTest {
	x.requiredWithMeta()
	return x
}
