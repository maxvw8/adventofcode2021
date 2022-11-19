package part1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	p := sol{}
	p.input = []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	actual := p.solve()
	if actual != expected {
		assert.Equal(t, expected, actual)
	}
}
