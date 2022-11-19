package utils

import (
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func TestParseInt(t *testing.T) {
	tests := []struct {
		Input    string
		Expected []int
		Fn       StringParseFn[int]
	}{
		{"1\n2\n3\n4\n5\n", []int{1, 2, 3, 4, 5}, strconv.Atoi},
		{"1\n2\n3\n4\n\n5\n", []int{1, 2, 3, 4, 5}, strconv.Atoi},
		{"\n5\n4\n3\n2\n\n1\n", []int{5, 4, 3, 2, 1}, strconv.Atoi},
	}

	for _, tc := range tests {
		t.Run(tc.Input, func(t *testing.T) {
			obtained, err := ReadFile(strings.NewReader(tc.Input), tc.Fn)
			AssertNoError(err, t)

			if !reflect.DeepEqual(tc.Expected, obtained) {
				t.Errorf("Obtained %v, but expected %v", obtained, tc.Expected)
			}
		})
	}

}

func AssertNoError(err error, t *testing.T) {
	t.Helper()
	if err != nil {
		t.Errorf("Expected no error but was: %s", err)
	}
}
