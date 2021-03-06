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
	if userinput.HelpRequested() {
		userinput.PrintHelp()
		os.Exit(0)
	}
	word, customOccurrencesPath := userinput.GetData()
	word = strings.Trim(word, " ")
	customOccurrencesPath = strings.Trim(customOccurrencesPath, " ")
	if word == "" && customOccurrencesPath == "" {
		fmt.Println("No user input provided. Use -h to get help")
		os.Exit(0)
	}
	var occurrences occ.Occurrences
	var uniqueSymbols []rune
	var err error
	if word != "" {
		occurrences, uniqueSymbols = occ.GetOccurrencesAndUniqueSymbols(word)
		if len(uniqueSymbols) < 2 {
			fmt.Println("At least two unique symbols required to encode")
			os.Exit(0)
		}
	} else if customOccurrencesPath != "" {
		occurrences, uniqueSymbols, err = occ.GetOccurrencesAndUniqueSymbolsFromFile(customOccurrencesPath)
		// TODO: Occurrences are need to be sorted
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		fmt.Printf("Code for %c is %s\n", v.Symb, v.Code)
	}
}
