package occ

import "strings"

// Occurrence - represent element in Occurrences slice
type Occurrence struct {
	Symb        rune
	Occurrences int
}

// Occurrences - represent slice with symblos of word and their occurrences
type Occurrences []Occurrence

// GetOccurrences - returns array of symbols and their occurrences in struct
func GetOccurrences(word string) (Occurrences, []rune) {
	var unsortedOccurrences Occurrences
	var doubles []rune
	for _, v := range word {
		if isUnique(v, doubles) {
			unsortedOccurrences = append(unsortedOccurrences, Occurrence{Symb: v, Occurrences: strings.Count(string(word), string(v))})
			doubles = append(doubles, v)
		}
	}
	occurrencesAreSorted := false
	occurrencesAreSortedInReverse := false
	var occurrencesToReturn Occurrences
	if isSorted(unsortedOccurrences) {
		occurrencesAreSorted = true
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
