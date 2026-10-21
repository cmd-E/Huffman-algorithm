package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"testing"
)

func TestCLI(t *testing.T) {
	directory := t.TempDir()
	binary := filepath.Join(directory, "huffman")
	if runtime.GOOS == "windows" {
		binary += ".exe"
	}
	if output, err := exec.Command("go", "build", "-o", binary, ".").CombinedOutput(); err != nil {
		t.Fatalf("build: %v\n%s", err, output)
	}
	input := filepath.Join(directory, "input.txt")
	if err := os.WriteFile(input, []byte("abbb"), 0600); err != nil {
		t.Fatal(err)
	}
	table := filepath.Join(directory, "table.txt")
	if err := os.WriteFile(table, []byte("a::0.25\nb::0.75\n"), 0600); err != nil {
		t.Fatal(err)
	}
	cases := []struct {
		name     string
		args     []string
		success  bool
		contains string
	}{
		{"help", []string{"-h"}, true, "-w"},
		{"empty", nil, true, "No user input"},
		{"word", []string{"-w", "abbb"}, true, "Code for a is 0"},
		{"file", []string{"-f", input}, true, "Code for b is 1"},
		{"table", []string{"-p", table, "-s", "::", "-prob"}, true, "Code for b is 1"},
		{"one-symbol", []string{"-w", "aaaa"}, false, "two unique"},
		{"missing-file", []string{"-f", filepath.Join(directory, "missing")}, false, "Error"},
		{"bad-table", []string{"-p", input}, false, "line 1"},
	}
	for _, item := range cases {
		t.Run(item.name, func(t *testing.T) {
			command := exec.Command(binary, item.args...)
			command.Dir = directory // Help must work without a help.txt in the working directory.
			output, err := command.CombinedOutput()
			if (err == nil) != item.success || !strings.Contains(string(output), item.contains) {
				t.Fatalf("err=%v output=%s", err, output)
			}
		})
	}
}
