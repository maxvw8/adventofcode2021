package day1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPart1(t *testing.T) {

	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	actual := TimesIncreased(input)
	if actual != expected {
		assert.Equal(t, expected, actual)
	}
}
func TestPart2(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 5
	actual, err := WindowIncreased(input)
	assert.NoError(t, err)
	if actual != expected {
		assert.Equal(t, expected, actual)
	}
}
