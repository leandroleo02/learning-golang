package types

import (
	"fmt"
	"testing"
)

func TestToCelsius(t *testing.T) {
	f := Fahrenheit{temperature: 32}
	c := f.ToCelsius()

	expected := Celsius{temperature: 0}

	fmt.Println(c.String()) // "100°C"
	fmt.Printf("%v\n", c)   // "100°C"; no need to call String explicitly
	fmt.Printf("%s\n", c)   // "100°C"
	fmt.Println(c)          // "100°C"
	fmt.Printf("%g\n", c)   // "100"; does not call String
	// fmt.Println(float64(c)) // "100"; does not call String

	if c != expected {
		t.Fatalf("actual %f != expected 0", c)
	}
}
