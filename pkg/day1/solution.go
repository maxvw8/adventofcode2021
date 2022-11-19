package day1

import (
	"advent/pkg/utils"
	"fmt"
	"io"
	"strconv"
)

const defaultWindowSize = 3

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
func sumUpWindows(input [][]int) []int {
	r := make([]int, len(input))
	for i, v := range input {
		for j := 0; j < len(v); j++ {
			r[i] += v[j]
		}
	}
	return r
}
func WindowIncreased(input []int) (int, error) {
	w, err := createWindows(input, defaultWindowSize)
	if err != nil {
		return 0, fmt.Errorf("problem while calculating window sum %d", err)
	}
	r := sumUpWindows(w)
	return TimesIncreased(r), nil
}

type solution struct {
	input []int
	file  string
}

func New(filename string) *solution {
	return &solution{file: filename}
}

func (p *solution) FileName() string {
	return p.file
}
func (p *solution) Name() string {
	return "Day 1"
}

func (p *solution) Load(reader io.Reader) error {
	input, err := utils.ReadFile(reader, strconv.Atoi)
	p.input = input
	return err
}
func (p *solution) solve() (int, error) {
	return WindowIncreased(p.input)
}
func (p *solution) Solve() (string, error) {
	res, err := p.solve()
	return fmt.Sprintf("%d", res), err
}
