package main

import (
	"log"

	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
)

func main() {
	// words := os.Args[1:]
	// if len(words) == 0 {
	// 	log.Fatalln("No argument provided")
	// }
	// word := words[0]
	// var word string
	// var customOccurrencesFilePath string
	// flag.StringVar(&word, "w", "", "User input to encode")
	// flag.StringVar(&customOccurrencesFilePath, "p", "", "File where custom occurrences for all symbols in word are defined")
	// flag.Parse()
	word := "heee"
	occurrences, uniqueSymbols := occ.GetOccurrences(word)
	// word := "beep boop beer!"
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		log.Printf("Code for %c is %s", v.Symb, v.Code)
	}
}
