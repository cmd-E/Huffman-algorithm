package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
	"unicode"

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
	if word != "" && customOccurrencesPath != "" {
		// -e or -d flag is passed
		occurrences, uniqueSymbols, err = occ.GetOccurrencesAndUniqueSymbolsFromFile(customOccurrencesPath)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(1)
		}
	} else if word != "" {
		occurrences, uniqueSymbols = occ.GetOccurrencesAndUniqueSymbols(word)
		if len(uniqueSymbols) < 2 {
			fmt.Println("At least two unique symbols required to encode")
			os.Exit(0)
		}
	}
	if !occ.EnoughDataToEncode([]rune(word), uniqueSymbols) {
		fmt.Println("Not enough data to encode input with this tree")
		os.Exit(0)
	}
	nodeList := &btll.NodeList{}
	nodeList.CreateList(occurrences)
	binaryTree := &btll.BinaryTree{}
	binaryTree.CreateTree(nodeList)
	encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		fmt.Printf("Code for %c is %s\n", v.Symb, v.Code)
	}
	var code []rune
	for _, symbol := range word {
		symbol = unicode.ToLower(symbol)
		found := false
		for _, encodedSymbol := range encodedSymbols {
			if unicode.ToLower(encodedSymbol.Symb) == symbol {
				found = true
				code = append(code, []rune(encodedSymbol.Code)...)
				break
			}
		}
		if !found {
			fmt.Printf("Dude wtf it's not supposed to happen at all. You don't have code for symbol [%c], how did you even pass all checks?\n", symbol)
			os.Exit(66)
		}
	}
	fmt.Printf("Word: [%s]. Code: [%s]", word, string(code))
}
