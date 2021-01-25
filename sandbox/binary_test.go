package sandbox

import (
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	decimal := BinaryToDecimal(1100)

	expected := 12.0

	if decimal != expected {
		t.Fatalf("actual %f != expected 0", decimal)
	}
}
