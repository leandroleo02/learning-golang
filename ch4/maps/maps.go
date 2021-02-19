package maps

import (
	"bufio"
	"io"
)

func dedup(r io.Reader) ([]string, error) {
	seen := make(map[string]bool) // a set of strings
	input := bufio.NewScanner(r)
	for input.Scan() {
		line := input.Text()
		if !seen[line] {
			seen[line] = true
		}
	}

	if err := input.Err(); err != nil {
		return nil, err	
	}

	lines := make([]string, 0, len(seen))
	for line := range seen {
		lines = append(lines, line)
	}
	return lines, nil
}