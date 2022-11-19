package part1

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

type sol struct {
	input []int
	file  string
}

func New(filename string) *sol {
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
	return TimesIncreased(p.input)
}
func (p *sol) Solve() string {
	return fmt.Sprintf("%d", p.solve())
}
