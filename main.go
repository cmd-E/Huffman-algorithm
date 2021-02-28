package main

import (
	"log"
	"os"
	"strings"

	btll "github.com/cmd-e/huffman-algo/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algo/occpackage"
)

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		log.Fatalln("No argument provided")
	}
	word := words[0]
	// word := "hellow"
	// word := "aaabbccccde"
	// word := "ааааааааааааааабббббббввввввггггггддддд"
	// word := "beep boop beer!"
	var unsortedOccurrences occ.Occurrences
	var doubles []rune
	for _, v := range word {
		if occ.IsUnique(v, doubles) {
			unsortedOccurrences = append(unsortedOccurrences, occ.Occurrence{Symb: v, Occurrences: strings.Count(string(word), string(v))})
			doubles = append(doubles, v)
		}
	}
	occurrencesAreSorted := false
	occurrencesAreSortedInReverse := false
	var occurrences occ.Occurrences
	if occ.IsSorted(unsortedOccurrences) {
		occurrencesAreSorted = true
	}
	if !occurrencesAreSorted && occ.IsSortedInReverse(unsortedOccurrences) {
		occurrencesAreSortedInReverse = true
	}
	if !occurrencesAreSorted && !occurrencesAreSortedInReverse {
		occurrences = occ.SortByOccurrences(unsortedOccurrences)
	} else if occurrencesAreSortedInReverse {
		occurrences = occ.ReverseArr(unsortedOccurrences)
	}
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	uniqueSymbols := doubles
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		log.Printf("Code for %c is %s", v.Symb, v.Code)
	}
}
