package bufferr

import (
	"bytes"
	"fmt"
)

// IntToString fa
func IntToString(values []int) string {
	var buf bytes.Buffer
	buf.WriteByte('[')
	for i, v := range values {
		if i > 0 {
			buf.WriteString(", ")
		}
		fmt.Fprintf(&buf, "%d", v)
	}
	buf.WriteByte(']')
	return buf.String()
}

func comma(s string) string {
	var buf bytes.Buffer
	if len(s) <= 3 {
		return s
	}

	startAt := len(s) % 3
	buf.WriteString(s[:startAt])

	endAt := startAt + 3
	for startAt < len(s) {
		if startAt > 0 {
			buf.WriteByte(',')
		}
		buf.WriteString(s[startAt:endAt])
		startAt += 3
		endAt += 3
	}
	return buf.String()
}
