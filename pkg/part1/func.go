package part1

const MaxStars = 50

func TimesIncreased(input []int) int {
	if len(input) == 0 {
		return 0
	}
	var increaseCounter = 0
	previous := input[0]
	for _, current := range input[1:] {
		if current > previous {
			increaseCounter++
		}
		previous = current
	}
	return increaseCounter
}
