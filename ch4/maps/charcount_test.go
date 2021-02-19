package maps

import (
	"bytes"
	"testing"
)

func TestCharcount(t *testing.T) {
	var buf bytes.Buffer
	// °  27  Image  15  Image  14  é  13  x  10  ≤  5  ×  5  Image  4  Image  4  Image  3
	buf.WriteRune(27)
	buf.WriteRune('°')
	buf.WriteString("af")
	buf.WriteRune('世')
	buf.WriteRune('界')
	buf.WriteRune('♞')
	buf.WriteRune('␗')
	charcount(&buf)
}
