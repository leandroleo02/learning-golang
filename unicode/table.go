package unicode

import (
	"fmt"
)

// PrintUnicodeTable prints the nth firsts code points
func PrintUnicodeTable(max int32) {
	for i := int32(0); i <= max; i++ {
		fmt.Printf("%v\t%q\n", i, rune(i))
	}
}
