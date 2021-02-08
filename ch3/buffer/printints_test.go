package bufferr

import (
	"fmt"
	"testing"
)

func TestBuffer(t *testing.T) {
	values := []int{1, 2, 3, 4}
	s := IntToString(values)
	fmt.Println(s)
}

func TestComma(t *testing.T) {
	s := "12345"
	fmt.Println(comma(s))
}

func TestComma2(t *testing.T) {
	s := "123456789012"
	fmt.Println(comma(s))
}
