package utils

import (
	"bufio"
	"io"
	"strings"
)

// StringParseFn defines function which receives text and try to parse to a type.
// It returns the converted type or an error if the convertion was not possible
type StringParseFn[T any] func(string) (T, error)

func ReadFile[T any](r io.Reader, parser func(string) (T, error)) ([]T, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanLines)
	var result []T
	var line string
	for scanner.Scan() {
		line = strings.TrimSpace(scanner.Text())
		if len(line) == 0 {
			continue
		}
		x, err := parser(line)
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}
