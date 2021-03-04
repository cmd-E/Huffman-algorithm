package occ

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/cmd-e/huffman-algorithm/userinput"
)

// Occurrence - represent element in Occurrences slice
type Occurrence struct {
	Symb        rune
	Occurrences float64
}

// Occurrences - represent slice with symblos of word and their occurrences
type Occurrences []Occurrence

// GetOccurrences - returns array of symbols and their occurrences in struct
func GetOccurrences(word string) (Occurrences, []rune) {
	var unsortedOccurrences Occurrences
	var doubles []rune
	for _, v := range word {
		if isUnique(v, doubles) {
			unsortedOccurrences = append(unsortedOccurrences, Occurrence{Symb: v, Occurrences: float64(strings.Count(string(word), string(v)))})
			doubles = append(doubles, v)
		}
	}
	occurrencesAreSorted := false
	occurrencesAreSortedInReverse := false
	var occurrencesToReturn Occurrences
	if isSorted(unsortedOccurrences) {
		occurrencesAreSorted = true
		occurrencesToReturn = unsortedOccurrences
	}
	if !occurrencesAreSorted && isSortedInReverse(unsortedOccurrences) {
		occurrencesAreSortedInReverse = true
	}
	if !occurrencesAreSorted && !occurrencesAreSortedInReverse {
		occurrencesToReturn = sortByOccurrences(unsortedOccurrences)
	} else if occurrencesAreSortedInReverse {
		occurrencesToReturn = reverseArr(unsortedOccurrences)
	}
	return occurrencesToReturn, doubles
}

// ParseOccurrencesFromFile - parses file at given path and returns defined occurrences in the file for every defined symbol
func ParseOccurrencesFromFile(path string) (Occurrences, []rune, error) {
	checkForProbability := userinput.ContainsProbabilities()
	file, err := os.Open(path)
	if err != nil {
		return nil, nil, fmt.Errorf("Error occured while parsing file: %s", err.Error())
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var occurs float64
	var probability float64
	var isValid bool
	var occurrences Occurrences
	var uniqueSymbols []rune
	for scanner.Scan() {
		if []rune(strings.Trim(scanner.Text(), " "))[0] == '#' {
			continue
		}
		//TODO let user choose separator for split
		txt := scanner.Text()
		splitted := strings.Split(strings.Trim(txt, " "), "-")
		symb := []rune(splitted[0])[0]
		var entity Occurrence
		if checkForProbability {
			if probability, isValid, err = getProbabilityIfValid(splitted[1]); !isValid && err != nil {
				return nil, nil, fmt.Errorf("Error occurred while parsing file with probabilities: %s", err.Error())
			}
			entity = Occurrence{Symb: symb, Occurrences: probability}
		} else if !checkForProbability {
			if occurs, isValid, err = getOccurrenceIfValid(splitted[1]); !isValid && err != nil {
				return nil, nil, fmt.Errorf("Error occurred while parsing file with probabilities: %s", err.Error())
			}
			entity = Occurrence{Symb: symb, Occurrences: occurs}
		}
		if isUnique(symb, uniqueSymbols) {
			occurrences = append(occurrences, entity)
			uniqueSymbols = append(uniqueSymbols, symb)
		}
	}
	return occurrences, uniqueSymbols, nil
}

func getProbabilityIfValid(suspect string) (float64, bool, error) {
	probability, err := strconv.ParseFloat(suspect, 32)
	if err != nil {
		return 0, false, err
	} else if probability < 0 || probability > 1 {
		return 0, false, nil
	}
	return probability, true, nil
}

func getOccurrenceIfValid(suspect string) (float64, bool, error) {
	occurrence, err := strconv.ParseFloat(suspect, 32)
	if err != nil {
		return 0, false, err
	}
	return occurrence, true, nil
}

func sortByOccurrences(occ Occurrences) Occurrences {
	for i := 1; i < len(occ); i++ {
		key := occ[i]
		j := i - 1
		for j >= 0 && occ[j].Occurrences > key.Occurrences {
			occ[j+1] = occ[j]
			j = j - 1
		}
		occ[j+1] = key
	}
	return occ
}

func isSorted(occ Occurrences) bool {
	for i := 1; i < len(occ); i++ {
		if occ[i-1].Occurrences > occ[i].Occurrences {
			return false
		}
	}
	return true
}

func isSortedInReverse(occ Occurrences) bool {
	for i := 1; i < len(occ); i++ {
		if occ[i-1].Occurrences < occ[i].Occurrences {
			return false
		}
	}
	return true
}

func reverseArr(occ Occurrences) Occurrences {
	for i, j := 0, len(occ)-1; i < j; i, j = i+1, j-1 {
		occ[i], occ[j] = occ[j], occ[i]
	}
	return occ
}

func isUnique(r rune, list []rune) bool {
	for _, v := range list {
		if v == r {
			return false
		}
	}
	return true
}
