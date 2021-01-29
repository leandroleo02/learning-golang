package sandbox

import (
	"fmt"
	"testing"
)

func TestStrings(t *testing.T) {
	s := "hello, world"
	fmt.Println(len(s))     // "12"
	fmt.Println(s[0], s[7]) // "104 119"  ('h' and 'w')

	fmt.Println(s[:6])

	const GoUsage = `Go is a tool for managing Go source code.

	Usage:
			go command [arguments]
	...`

	fmt.Println(GoUsage)

	fmt.Println("\xe4\xb8\x96\xe7\x95\x8c")
	fmt.Println("\u4e16\u754c")
	fmt.Println("\U00004e16\U0000754c")
	
	fmt.Println("\xe4\xb8\x96")

	fmt.Println("\x21")
}

func TestStringsEncoding(t *testing.T) {
	asianHello := "Hello, 世界"
	fmt.Println(len(asianHello))

	for i, r := range asianHello {
		fmt.Printf("%d\t%q\t%d\n", i, r, r)
	}

	fb := "hello"
	fmt.Println(len(fb))

	fa := "世界"
	fmt.Printf("%q => %d\n", fa, len(fa))
}
