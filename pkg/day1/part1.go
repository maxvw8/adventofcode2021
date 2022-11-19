package day1

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"strconv"
)

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
func createWindows(input []int, windowSize int) ([][]int, error) {
	if len(input) < windowSize {
		return nil, fmt.Errorf("windows size %d is lower than input size %d", input, windowSize)
	}
	nWindows := len(input) - windowSize + 1
	windows := make([][]int, nWindows)
	for i := 0; i < nWindows; i++ {
		windows[i] = input[i : i+windowSize]
	}
	return windows, nil
}
func reduce(input [][]int) []int {
	r := make([]int, len(input))
	for i, v := range input {
		for j := 0; j < len(v); j++ {
			r[i] += v[j]
		}
	}
	return r
}
func WindowIncreased(input []int) int {
	w, err := createWindows(input, 3)
	if err != nil {
		fmt.Printf("we screwed up!")
	}
	r := reduce(w)
	return TimesIncreased(r)
}

type sol struct {
	input []int
	file  string
}

func NewPart1(filename string) *sol {
	return &sol{file: filename}
}

func (p *sol) FileName() string {
	return p.file
}

func (p *sol) Load(reader io.Reader) error {
	input, err := utils.ReadFile(reader, strconv.Atoi)
	p.input = input
	return err
}
func (p *sol) solve() int {
	return WindowIncreased(p.input)
}
func (p *sol) Solve() string {
	return fmt.Sprintf("%d", p.solve())
}
