package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := make(map[string]map[string]bool)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, words)
	} else {
		for _, filename := range files {
			f, err := os.Open(filename)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, words)
			f.Close()
		}
	}

	printed := make(map[string]bool)
	for line, appeard := range words {
		if len(appeard) > 1 {
			for name := range words[line] {
				if !printed[name] {
					printed[name] = true
					fmt.Println(name)
				}
			}
		}
	}
}

func countLines(f *os.File, words map[string]map[string]bool) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		word := input.Text()
		if words[word] == nil {
			words[word] = make(map[string]bool)
		}
		words[word][f.Name()] = true
	}
}
