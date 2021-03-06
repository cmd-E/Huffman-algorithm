package userinput

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var word string
var customOccurrencesFilePath string
var filePathToWord string
var treatOccurrencesAsProbabilities bool
var printHelp bool
var customSeparator string

// InitFlags - defines flags for user to operate program
func InitFlags() {
	flag.StringVar(&word, "w", "", "Input to encode")
	flag.StringVar(&filePathToWord, "f", "", "file where word is defined")
	flag.StringVar(&customOccurrencesFilePath, "p", "", "File where custom occurrences for all symbols are defined")
	flag.BoolVar(&printHelp, "h", false, "Print help")
	flag.BoolVar(&treatOccurrencesAsProbabilities, "prob", false, "available if -p is defined. Occurrences for symbols are treated as possibilities")
	flag.StringVar(&customSeparator, "s", "-", "available if -p is defined. Changes default separator (-) on user's one")
}

// GetData - returns user input and path to file with custom occurrences
// w -> f -> p (prob) | h
func GetData() (string, string) {
	if strings.Trim(word, " ") != "" {
		return word, "" // -w flag
	} else if strings.Trim(filePathToWord, " ") != "" {
		return getWordFromFile(filePathToWord), "" // -f flag
	}
	return "", customOccurrencesFilePath // -p flag
}

func getWordFromFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error occured while reading file ", err.Error())
		os.Exit(1)
	}
	return string(file)
}

// HelpRequested - checks if user requested help
func HelpRequested() bool {
	return printHelp
}

// ContainsProbabilities - checks if prob flag is checked. If is checked numbers in file -p are treated as floats in range of [0..1]
func ContainsProbabilities() bool {
	return treatOccurrencesAsProbabilities
}

// GetSeparator - returns separator for file when -p flag is defined
func GetSeparator() string {
	return customSeparator
}

// PrintHelp - prints help if user is asking for it
func PrintHelp() {
	helpFile, err := ioutil.ReadFile("help.txt")
	if err != nil {
		fmt.Println("There was an error while attempting to read file text: ", err.Error())
		return
	}
	fmt.Println(string(helpFile))
}
