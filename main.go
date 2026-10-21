package main

import (
	_ "embed"
	"flag"
	"fmt"
	"os"
	"strings"

	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
	"github.com/cmd-e/huffman-algorithm/userinput"
)

//go:embed help.txt
var helpText string

func init() {
	userinput.InitFlags()
}

func main() {
	flag.Parse()
	if userinput.HelpRequested() {
		fmt.Print(helpText)
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
			fmt.Fprintln(os.Stderr, "At least two unique symbols required to encode")
			os.Exit(1)
		}
	} else if customOccurrencesPath != "" {
		file, openErr := os.Open(customOccurrencesPath)
		if openErr != nil {
			fmt.Fprintln(os.Stderr, openErr)
			os.Exit(1)
		}
		defer file.Close()
		occurrences, uniqueSymbols, err = occ.Parse(file, userinput.GetSeparator(), userinput.ContainsProbabilities())
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	}
	binaryTree := &btll.BinaryTree{}
	if err := binaryTree.Build(occurrences); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		fmt.Printf("Code for %c is %s\n", v.Symb, v.Code)
	}
}
