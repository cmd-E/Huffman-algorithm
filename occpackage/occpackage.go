package occ

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
	"unicode/utf8"
)

type Occurrence struct {
	Symb        rune
	Occurrences float64
}
type Occurrences []Occurrence

func GetOccurrencesAndUniqueSymbols(word string) (Occurrences, []rune) {
	counts := make(map[rune]int)
	order := []rune{}
	for _, r := range word {
		if counts[r] == 0 {
			order = append(order, r)
		}
		counts[r]++
	}
	values := make(Occurrences, 0, len(order))
	for _, r := range order {
		values = append(values, Occurrence{r, float64(counts[r])})
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].Occurrences < values[j].Occurrences })
	return values, order
}

// Parse reads an occurrence table without depending on CLI flags or filesystem state.
func Parse(reader io.Reader, separator string, probabilities bool) (Occurrences, []rune, error) {
	if separator == "" {
		return nil, nil, fmt.Errorf("separator cannot be empty")
	}
	scanner := bufio.NewScanner(reader)
	scanner.Buffer(make([]byte, 4096), 1024*1024)
	values := Occurrences{}
	order := []rune{}
	seen := make(map[rune]bool)
	line := 0
	for scanner.Scan() {
		line++
		text := strings.TrimSpace(scanner.Text())
		if text == "" || strings.HasPrefix(text, "#") {
			continue
		}
		fields := strings.SplitN(text, separator, 2)
		if len(fields) != 2 || utf8.RuneCountInString(fields[0]) != 1 {
			return nil, nil, fmt.Errorf("line %d: expected one symbol and a weight separated by %q", line, separator)
		}
		symbol, _ := utf8.DecodeRuneInString(fields[0])
		weight, err := strconv.ParseFloat(strings.TrimSpace(fields[1]), 64)
		if err != nil || math.IsNaN(weight) || math.IsInf(weight, 0) || weight <= 0 || (probabilities && weight > 1) {
			return nil, nil, fmt.Errorf("line %d: invalid positive weight", line)
		}
		if seen[symbol] {
			return nil, nil, fmt.Errorf("line %d: duplicate symbol", line)
		}
		seen[symbol] = true
		order = append(order, symbol)
		values = append(values, Occurrence{symbol, weight})
	}
	if err := scanner.Err(); err != nil {
		return nil, nil, fmt.Errorf("read occurrence table: %w", err)
	}
	if len(values) < 2 {
		return nil, nil, fmt.Errorf("at least two distinct symbols required")
	}
	sort.SliceStable(values, func(i, j int) bool { return values[i].Occurrences < values[j].Occurrences })
	return values, order, nil
}
