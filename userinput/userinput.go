package userinput

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"

	occ "github.com/cmd-e/huffman-algorithm/occpackage"
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
	if strings.Trim(filePathToWord, " ") != "" { // -f flag
		return getWordFromFile(filePathToWord), ""
	} else if strings.Trim(word, " ") != "" && strings.Trim(customOccurrencesFilePath, " ") != " " { // -w and -p flags
		if !customOccurrencesAreValid(customOccurrencesFilePath, word) {
			fmt.Println("File with occurences is not valid for this word. You can provide only file and symbols in it will be coded, but not validated with input")
			os.Exit(0)
		}
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

// BUG check with "ab" and file occs.txt
func customOccurrencesAreValid(p, w string) bool {
	runes := []rune(w)
	_, uniqueSymbols := occ.ParseOccurrencesFromFile(p)
	for _, us := range uniqueSymbols {
		for j, suspect := range runes {
			if us == suspect {
				if j != len(runes) {
					runes = append(runes[:j], runes[j+1:]...)
					log.Println("a")

				} else {
					runes = append(runes[:j], runes[j:]...)
					log.Println("b")
				}
				log.Println(runes)
			}
		}
	}
	if len(runes) == 0 {
		return true
	}
	return false
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
