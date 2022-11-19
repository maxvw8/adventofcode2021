package advent

import "io"

type Problem interface {
	Name() string
	FileName() string
	Load(reader io.Reader) error
	Solve() (string, error)
}
