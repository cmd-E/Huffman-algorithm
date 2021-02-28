package main

import (
	"log"
	"os"

	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
)

func main() {
	words := os.Args[1:]
	if len(words) == 0 {
		log.Fatalln("No argument provided")
	}
	word := words[0]
	occurrences, uniqueSymbols := occ.GetOccurrences(word)

	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		log.Printf("Code for %c is %s", v.Symb, v.Code)
	}
}
