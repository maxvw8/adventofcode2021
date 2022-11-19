package part1

import "testing"

func Test(t *testing.T) {
	input := []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}
	expected := 7
	obtained := TimesIncreased(input)
	if obtained != expected {
		t.Errorf("expected %d increased, but got %d", expected, obtained)
	}
}
