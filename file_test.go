package gt_test

import (
	"bytes"
	"io"
	"os"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestFileExists(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Create a FileTest instance
	ft := gt.File(t, file.Name())

	// Test that the file exists
	ft.Exists()
}

func TestFileNotExists(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	os.Remove(file.Name())

	// Create a FileTest instance
	ft := gt.File(t, file.Name())

	// Test that the file does not exist
	ft.NotExists()
}

func TestFileString(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write some content to the file
	content := []byte("hello")
	if _, err := file.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a FileTest instance
	ft := gt.File(t, file.Name())

	// Test the file content using the String method
	ft.String(func(st gt.StringTest) {
		st.Equal("hello")
	})
}

func TestFileReader(t *testing.T) {
	// Create a temporary file for testing
	file, err := os.CreateTemp("", "testfile")
	if err != nil {
		t.Fatal(err)
	}
	defer os.Remove(file.Name())

	// Write some content to the file
	content := []byte("hello")
	if _, err := file.Write(content); err != nil {
		t.Fatal(err)
	}
	if err := file.Close(); err != nil {
		t.Fatal(err)
	}

	// Create a FileTest instance
	ft := gt.File(t, file.Name())

	// Test the file content using the Reader method
	ft.Reader(func(tb testing.TB, r io.Reader) {
		// Read the file content from the reader
		buf := new(bytes.Buffer)
		if _, err := buf.ReadFrom(r); err != nil {
			tb.Fatal(err)
		}

		// Compare the file content
		if buf.String() != "hello" {
			tb.Errorf("Expected 'hello', got '%s'", buf.String())
		}
	})
}
