package maps

import (
	"bytes"
	"fmt"
	"testing"
)

func TestDedup(t *testing.T) {
	var buf bytes.Buffer
	buf.WriteString("ola\n")
	buf.WriteString("eu\n")
	buf.WriteString("sou\n")
	buf.WriteString("ola\n")
	buf.WriteString("leandro\n")
	buf.WriteString("ola\n")

	lines, err := dedup(&buf)
	if(err != nil) {
		t.Fail()
	}

	fmt.Println(lines)
}