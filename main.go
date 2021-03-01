package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
	"github.com/cmd-e/huffman-algorithm/userinput"
)

func init() {
	userinput.InitFlags()
}

func main() {
	flag.Parse()
	if userinput.GetHelp() {
		userinput.PrintHelp()
		os.Exit(0)
	}
	word, customOccurrencesPath := userinput.GetData()
	if strings.Trim(word, " ") == "" {
		fmt.Println("No user input provided. Use -h to get help")
		os.Exit(0)
	}
	var occurrences occ.Occurrences
	var uniqueSymbols []rune
	if strings.Trim(customOccurrencesPath, " ") == "" {
		occurrences, uniqueSymbols = occ.GetOccurrences(word)
	}
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		log.Printf("Code for %c is %s", v.Symb, v.Code)
	}
}
