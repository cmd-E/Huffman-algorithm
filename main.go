package main

import (
	"fmt"
	"os"
)

func main() {
	word := os.Args[1]
	var occurense map[rune]int
	for _, v := range word {
		occurense[v]++
	}
	fmt.Println(occurense)
}
