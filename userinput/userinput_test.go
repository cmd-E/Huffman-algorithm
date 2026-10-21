package userinput

import (
	"flag"
	"os"
	"path/filepath"
	"testing"
)

func parseFlags(t *testing.T, args ...string) {
	t.Helper()
	previous := flag.CommandLine
	flag.CommandLine = flag.NewFlagSet("huffman", flag.ContinueOnError)
	t.Cleanup(func() { flag.CommandLine = previous })
	InitFlags()
	if err := flag.CommandLine.Parse(args); err != nil {
		t.Fatal(err)
	}
}

func TestWordTakesPriorityAndFlagAccessors(t *testing.T) {
	parseFlags(t, "-w", "ab", "-f", "not-read", "-p", "not-read", "-s", "::", "-prob", "-h")
	text, table := GetData()
	if text != "ab" || table != "" || !HelpRequested() || !ContainsProbabilities() || GetSeparator() != "::" {
		t.Fatal("flags or priority lost")
	}
}

func TestWordFileAndOccurrenceFallback(t *testing.T) {
	file := filepath.Join(t.TempDir(), "word.txt")
	if err := os.WriteFile(file, []byte("😀abc"), 0600); err != nil {
		t.Fatal(err)
	}
	parseFlags(t, "-w", "   ", "-f", file, "-p", "table.txt")
	text, table := GetData()
	if text != "😀abc" || table != "" {
		t.Fatal(text, table)
	}
	filePathToWord = ""
	text, table = GetData()
	if text != "" || table != "table.txt" {
		t.Fatal(text, table)
	}
}

func TestDefaultFlags(t *testing.T) {
	parseFlags(t)
	text, table := GetData()
	if text != "" || table != "" || HelpRequested() || ContainsProbabilities() || GetSeparator() != "-" {
		t.Fatal("unexpected defaults")
	}
}
