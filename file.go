package gt

import (
	"io"
	"os"
	"testing"
)

type FileTest struct {
	path string
	t    testing.TB
}

// File provides FileTest that has basic comparison methods
func File(t testing.TB, path string) FileTest {
	t.Helper()
	return FileTest{
		path: path,
		t:    t,
	}
}

// F is sugar syntax of File
func F(t testing.TB, path string) FileTest {
	t.Helper()
	return File(t, path)
}

// Exists check if file exists
//
//	gt.File(t, "testdata/file.txt").Exists() // Pass
//	gt.File(t, "testdata/no-file.txt").Exists() // Fail
func (x FileTest) Exists() FileTest {
	x.t.Helper()
	if !EvalFileExists(x.path) {
		x.t.Errorf("file should exist, %s", x.path)
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
		x.t.Errorf("file should not exist, %s", x.path)
	}

	return x
}

// String calls f with file content
//
//	gt.File(t, "testdata/file.txt").String(func(s StringTest) {
//	   s.Equal("hello")
//	})
func (x FileTest) String(f func(StringTest)) FileTest {
	x.t.Helper()
	data, err := os.ReadFile(x.path)
	if err != nil {
		x.t.Errorf("failed to read file, %s", x.path)
		return x
	}

	f(String(x.t, string(data)))
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
		x.t.Errorf("failed to open file, %s", x.path)
		return x
	}
	defer r.Close()

	f(x.t, r)
	return x
}
