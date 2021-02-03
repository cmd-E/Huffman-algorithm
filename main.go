package main

import (
	"os"
)

func main() {
	word := os.Args[1]
	occurrences := make(map[rune]int)
	for _, v := range word {
		occurrences[v]++
	}

}
