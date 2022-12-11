package gt_test

import (
	"fmt"
	"os"
	"strings"
	"testing"

	"github.com/m-mizutani/gt"
)

func TestDiff(t *testing.T) {
	type testStruct struct {
		Name string
	}

	testCases := map[string]struct {
		wants  []string
		expect any
		actual any
	}{
		/*
			"numbers": {
				wants:  []string{"expect: 2", "actual: 1"},
				expect: 2,
				actual: 1,
			},

			"struct": {
				wants:  []string{`gt_test.testStruct{`},
				expect: testStruct{Name: "blue"},
				actual: testStruct{Name: "orange"},
			},

			"pointer": {
				wants:  []string{"[]int{"},
				expect: []int{1, 2, 3},
				actual: []int{1, 5, 3},
			},
		*/
	}
	for title, tc := range testCases {
		t.Run(title, func(t *testing.T) {

			got := gt.Diff(tc.expect, tc.actual)
			for _, want := range tc.wants {
				if !strings.Contains(got, want) {
					toOutput(got)
					t.Error(strings.Join([]string{
						"",
						"--- want to contain ---",
						want,
						"",
						"--- got ---",
						got,
					}, "\n"))
				}
			}
		})
	}
}

func toOutput(s string) {
	tmp, err := os.CreateTemp("", "*.go")
	if err != nil {
		panic("toOutput: " + err.Error())
	}
	defer tmp.Close()

	tmp.WriteString(s)
	fmt.Fprintln(os.Stderr, "out:", tmp.Name())
}
