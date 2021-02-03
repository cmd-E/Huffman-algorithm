package main

import (
	"os"
)

func main() {
	word := os.Args[1]
	occurences := make(map[rune]int)
	for _, v := range word {
		occurences[v]++
	}

}
