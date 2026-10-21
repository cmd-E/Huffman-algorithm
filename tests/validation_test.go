package test

import (
	bt "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
	"strings"
	"testing"
)

func TestOccurrenceErrors(t *testing.T) {
	for _, input := range []string{"", "a-1\n", "-1\n", "a-NaN\nb-1", "a-Inf\nb-1", "a-0\nb-1", "a-1\na-2", "a-1\nb-" + strings.Repeat("1", 1024*1024)} {
		if _, _, err := occ.Parse(strings.NewReader(input), "-", false); err == nil {
			t.Errorf("accepted invalid input %q", input[:minLength(len(input), 30)])
		}
	}
	if _, _, err := occ.Parse(strings.NewReader("\n # comment\na-1\nb-2\n"), "-", false); err != nil {
		t.Fatal(err)
	}
	if err := new(bt.BinaryTree).CreateTree(&bt.NodeList{}); err == nil {
		t.Fatal("empty tree accepted")
	}
}
func minLength(a, b int) int {
	if a < b {
		return a
	}
	return b
}
func TestUnicodePrefixFree(t *testing.T) {
	values, symbols := occ.GetOccurrencesAndUniqueSymbols("😀😀абвг")
	list := new(bt.NodeList)
	list.CreateList(values)
	tree := new(bt.BinaryTree)
	if err := tree.CreateTree(list); err != nil {
		t.Fatal(err)
	}
	codes := tree.EncodeSymbols(symbols)
	for i, a := range codes {
		for j, b := range codes {
			if i != j && strings.HasPrefix(a.Code, b.Code) {
				t.Fatal("not prefix free")
			}
		}
	}
}
