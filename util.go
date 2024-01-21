package gt

import (
	"fmt"
	"os"
	"reflect"
	"strings"

	"github.com/google/go-cmp/cmp"
)

// EvalCompare is a function to check if actual value equals with expected value. A developer can replace EvalCompare with own evaluation function if needed.
//
//	EvalCompare(1, 2) == false
//	EvalCompare("abc", "axc") == false
//	EvalCompare([]int{1, 2, 3}, []int{1, 2, 3}) == true
var EvalCompare = func(a, b any) bool {
	return reflect.DeepEqual(a, b)
}

// EvalIsNil is a function to check if actual value is nil. A developer can replace EvalIsNil with own nil check function if needed.
//
//	EvalIsNil(1) == false
//	EvalIsNil("abc") == false
//	EvalIsNil(nil) == true
//	EvalIsNil([]int{}) == true
var EvalIsNil = func(v any) bool {
	value := reflect.ValueOf(v)

	switch value.Kind() {
	case reflect.Invalid:
		return true

	case reflect.Array, reflect.Slice:
		return value.Len() == 0

	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Pointer, reflect.Uintptr, reflect.UnsafePointer:
		return value.IsNil()

	default:
		return false
	}
}

// EvalFileExists is a function to check if file exists. A developer can replace EvalFileExists with own file existence check function if needed.
//
//	EvalFileExists("testdata/file.txt") == true
//	EvalFileExists("testdata/no-file.txt") == false
var EvalFileExists = func(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

var Diff = func(expect, actual any) string {
	switch reflect.ValueOf(actual).Kind() {
	case reflect.Pointer, reflect.UnsafePointer,
		reflect.Array, reflect.Slice,
		reflect.Struct, reflect.Map:
		return "diff:\n" + cmp.Diff(expect, actual, cmp.Exporter(func(t reflect.Type) bool { return true }))

	default:
		return strings.Join([]string{
			fmt.Sprintf("actual: %+v", actual),
			fmt.Sprintf("expect: %+v", expect),
		}, "\n")
	}
}

var DumpError = func(err error) string {
	return err.Error()
}
