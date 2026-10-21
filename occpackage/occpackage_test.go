package occ

import (
	"errors"
	"reflect"
	"strings"
	"testing"
)

type failingReader struct{}

func (failingReader) Read([]byte) (int, error) { return 0, errors.New("read failure") }

func TestParseCustomSeparatorAndProbabilityBoundaries(t *testing.T) {
	values, order, err := Parse(strings.NewReader("# heading\n😀::0.25\nб::1\na::0.5\n"), "::", true)
	if err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(order, []rune{'😀', 'б', 'a'}) || values[0].Symb != '😀' || values[2].Occurrences != 1 {
		t.Fatalf("unexpected result: %v %v", values, order)
	}
	for _, text := range []string{"a::1.01\nb::0.5", "ab::0.5\nb::1", "a::invalid\nb::1", "a::0\nb::1"} {
		if _, _, err := Parse(strings.NewReader(text), "::", true); err == nil {
			t.Fatalf("accepted %q", text)
		}
	}
	if _, _, err := Parse(strings.NewReader("a-1\nb-2"), "", false); err == nil {
		t.Fatal("empty separator accepted")
	}
	if _, _, err := Parse(failingReader{}, "-", false); err == nil || !strings.Contains(err.Error(), "read failure") {
		t.Fatalf("lost read error: %v", err)
	}
}

func TestFrequencyCountingPreservesFirstSeenOrderAndStableTies(t *testing.T) {
	values, order := GetOccurrencesAndUniqueSymbols("😀ба😀б")
	if !reflect.DeepEqual(order, []rune{'😀', 'б', 'а'}) {
		t.Fatal(order)
	}
	if !reflect.DeepEqual(values, Occurrences{{'а', 1}, {'😀', 2}, {'б', 2}}) {
		t.Fatal(values)
	}
	values, order = GetOccurrencesAndUniqueSymbols("")
	if len(values) != 0 || len(order) != 0 {
		t.Fatal("empty input produced symbols")
	}
}
