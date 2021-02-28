package occ

// Occurrence - represent element in Occurrences slice
type Occurrence struct {
	Symb        rune
	Occurrences int
}

// Occurrences - represent slice with symblos of word and their occurrences
type Occurrences []Occurrence

func SortByOccurrences(occ Occurrences) Occurrences {
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

func IsSorted(occ Occurrences) bool {
	for i := 1; i < len(occ); i++ {
		if occ[i-1].Occurrences > occ[i].Occurrences {
			return false
		}
	}
	return true
}

func IsSortedInReverse(occ Occurrences) bool {
	for i := 1; i < len(occ); i++ {
		if occ[i-1].Occurrences < occ[i].Occurrences {
			return false
		}
	}
	return true
}

func ReverseArr(occ Occurrences) Occurrences {
	for i, j := 0, len(occ)-1; i < j; i, j = i+1, j-1 {
		occ[i], occ[j] = occ[j], occ[i]
	}
	return occ
}

func IsUnique(r rune, list []rune) bool {
	for _, v := range list {
		if v == r {
			return false
		}
	}
	return true
}
