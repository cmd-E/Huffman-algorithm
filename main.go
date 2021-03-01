package main

import (
	"flag"
	"fmt"
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
	} else {
		occurrences, uniqueSymbols = occ.ParseOccurrencesFromFile(customOccurrencesPath)
	}
	// word := "ab"
	// customOccurrencesPath := "C:\\Users\\me\\Desktop\\occs.txt"
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		fmt.Printf("Code for %c is %s\n", v.Symb, v.Code)
	}
}
