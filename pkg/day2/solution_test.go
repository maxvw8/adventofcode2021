package day2

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCommandParsing(t *testing.T) {
	tcs := []struct {
		command  string
		expected transformation
	}{
		{"forward 1", transformation{horizontal: 1, depth: 0}},
		{"up 1", transformation{horizontal: 0, depth: -1}},
		{"down 1", transformation{horizontal: 0, depth: 1}},
	}
	for _, tc := range tcs {
		t.Run(fmt.Sprintf("Transforming %s", tc.command), func(t *testing.T) {
			actual, err := toTransformation(tc.command)
			assert.NoError(t, err)
			assert.Equal(t, tc.expected, actual)
		})
	}
}

func TestPart2(t *testing.T) {
	moves := []transformation{
		{horizontal: 5},
		{depth: 5},
		{horizontal: 8},
		{depth: -3},
		{depth: 8},
		{horizontal: 2},
	}
	pos := reduce(moves)
	assert.Equal(t, transformation{horizontal: 15, depth: 60, aim: 10}, pos)
}
