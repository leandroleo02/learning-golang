package sandbox

import (
	"fmt"
	"math"
	"testing"
)

func TestBinaryToDecimal(t *testing.T) {
	decimal := BinaryToDecimal(01001010)

	expected := 12.0

	if decimal != expected {
		t.Fatalf("actual %f != expected 0", decimal)
	}
}

func TestDecimalToBinary(t *testing.T) {
	decimal := DecimalToBinary(12)

	expected := "1100"

	if decimal != expected {
		t.Fatalf("actual %s != expected 0", decimal)
	}
}

func TestBinary(t *testing.T) {
	var x uint8 = 1<<1 | 1<<5
	var y uint8 = 1<<1 | 1<<2

	fmt.Printf("%d => %08b\n", x, x) // "00100010", the set {1, 5}
	fmt.Printf("%d => %08b\n", y, y) // "00000110", the set {1, 2}

	fmt.Printf("%08b\n", x&y)                      // "00000010", the intersection {1}
	fmt.Printf("%08b\n", x|y)                      // "00100110", the union {1, 2, 5}
	fmt.Printf("%08b\n", x^y)                      // "00100100", the symmetric difference {2, 5}
	fmt.Printf("%d => %08b\n\n\n\n\n", x&^y, x&^y) // "00100000", the difference {5}

	for i := uint(0); i < 8; i++ {
		fmt.Printf("%08b (%d)&(1<<%d) %08b => %08b\n", x, x, i, (1 << i), x&(1<<i))
		if x&(1<<i) != 0 { // membership test
			fmt.Println(i) // "1", "5"
		}
	}

	ua := -10
	fmt.Printf("%08b\n", (ua >> 2))
}

func TestFormatter(t *testing.T) {
	o := 0666
	fmt.Printf("%d %[1]o %#[1]o\n", o) // "438 666 0666"

	fmt.Println(math.Exp(1))
	math.IsNaN
}
