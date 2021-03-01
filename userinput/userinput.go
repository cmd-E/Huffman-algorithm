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
var printHelp bool

// InitFlags - defines flags for user to operate program
func InitFlags() {
	flag.StringVar(&word, "w", "", "Input to encode")
	flag.StringVar(&filePathToWord, "f", "", "file where word is defined (in case it is too long)")
	flag.StringVar(&customOccurrencesFilePath, "p", "", "File where custom occurrences for all symbols in input are defined")
	flag.BoolVar(&printHelp, "h", false, "Print help")
}

// GetData - returns user input and path to file with custom occurrences
func GetData() (string, string) {
	if strings.Trim(filePathToWord, " ") != "" {
		return getWordFromFile(filePathToWord), customOccurrencesFilePath
	}
	return word, customOccurrencesFilePath
}

func getWordFromFile(path string) string {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Error occured while reading file ", err.Error())
		os.Exit(1)
	}
	return string(file)
}

// GetHelp - checks if user requested help
func GetHelp() bool {
	return printHelp
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
