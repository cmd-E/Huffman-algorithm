package userinput

import (
	"flag"
	"fmt"
)

var word string
var customOccurrencesFilePath string
var printHelp bool

// InitFlags - defines flags for user to operate program
func InitFlags() {
	flag.StringVar(&word, "w", "", "User input to encode")
	flag.StringVar(&customOccurrencesFilePath, "p", "", "File where custom occurrences for all symbols in word are defined")
	flag.BoolVar(&printHelp, "h", false, "Print help to user")
}

// GetData - returns user input and path to file with custom occurrences
func GetData() (string, string) {
	return word, customOccurrencesFilePath
}

// GetHelp - checks if user requested help
func GetHelp() bool {
	return printHelp
}

// PrintHelp - prints help if user is asking for it
func PrintHelp() {
	// TODO print help from file
	fmt.Println("Help is on the way!")
}
