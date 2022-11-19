package day2

import (
	"advent/pkg/utils"
	"fmt"
	"io"
)

const forward = "forward"
const down = "down"
const up = "up"

type transformation struct {
	depth      int
	horizontal int
	aim        int
}

type solution struct {
	moves []transformation
}

func New() *solution {
	return &solution{}
}

func (p *solution) FileName() string {
	return "2.txt"
}
func (p *solution) Name() string {
	return "Day 2"
}

func (p *solution) Load(reader io.Reader) error {
	moves, err := utils.ReadFile(reader, toTransformation)
	p.moves = moves
	return err
}
func reduce(moves []transformation) transformation {
	r := transformation{depth: 0, horizontal: 0}
	for _, m := range moves {
		r = r.apply(m)
	}
	return r
}

func (t transformation) apply(move transformation) transformation {
	return transformation{depth: t.depth + move.depth, horizontal: t.horizontal + move.horizontal}
}
func toTransformation(line string) (transformation, error) {
	var command string
	var value int
	var move = transformation{}
	fmt.Sscanf(line, "%s %d", &command, &value)
	switch command {
	case forward:
		move.horizontal = value
	case up:
		move.depth = -value
	case down:
		move.depth = value
	default:
		return transformation{}, fmt.Errorf("unrecognized command '%s'", command)
	}

	return move, nil
}

func (p *solution) Solve() (string, error) {
	pos := reduce(p.moves)
	val := pos.depth * pos.horizontal
	return fmt.Sprintf("%d", val), nil
}
