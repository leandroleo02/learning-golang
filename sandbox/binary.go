package sandbox

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

// BinaryToDecimal converts binary numbers to decimal
func BinaryToDecimal(number int) float64 {
	ns := fmt.Sprint(number)
 	var result float64
	for i, n := range ns {
		nn, err := strconv.Atoi(string(n))
		aff(err)
		nf := float64(nn)
		p := float64((len(ns) - 1) - i)
		result += (nf * math.Pow(2, p))
	}
	return result
}

// DecimalToBinary converts decimal to binary
func DecimalToBinary(number int64) string {
	return strconv.FormatInt(number, 2)
}

func reverseString(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}

func aff(err error) {
	if err != nil {
		fmt.Println("Error")
		os.Exit(1)
	}
}
