package arrays

import (
	"crypto/sha256"
	"fmt"
	"testing"
)

func TestReverse(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	Reverse(s)

	fmt.Println(s)
}

func TestRotateWithReverse(t *testing.T) {
	s := []int{0, 1, 2, 3, 4, 5}
	Reverse(s[:2])
	fmt.Println(s)
	
	Reverse(s[2:])
	fmt.Println(s)
	
	Reverse(s)
	fmt.Println(s)
}

func TestSha(t *testing.T) {
	sha256.Sum256([]byte("x"))
}