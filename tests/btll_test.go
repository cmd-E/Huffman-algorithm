package test

import (
	"testing"

	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
	"github.com/cmd-e/huffman-algorithm/tests/testutils"
)

func TestEncodeSymbols(t *testing.T) {
	testCases := []struct {
		num      int
		input    string
		expected []btll.Encoded
	}{
		{1, "beep boop beer!", []btll.Encoded{
			{Symb: 'b', Code: "00"},
			{Symb: 'e', Code: "11"},
			{Symb: 'p', Code: "101"},
			{Symb: ' ', Code: "010"},
			{Symb: 'o', Code: "011"},
			{Symb: 'r', Code: "1000"},
			{Symb: '!', Code: "1001"},
		}},
		{2, "ab", []btll.Encoded{
			{Symb: 'a', Code: "0"},
			{Symb: 'b', Code: "1"},
		}},
		{3, "abbbbbb", []btll.Encoded{
			{Symb: 'a', Code: "0"},
			{Symb: 'b', Code: "1"},
		}},
	}
	for _, tcase := range testCases {
		occurrences, uniqueSymbols := occ.GetOccurrencesAndUniqueSymbols(tcase.input)
		nodeList := &btll.NodeList{}
		nodeList.CreateList(occurrences)
		binaryTree := &btll.BinaryTree{}
		binaryTree.CreateTree(nodeList)
		encodedSymbols := binaryTree.EncodeSymbols(uniqueSymbols)
		if len(encodedSymbols) != len(tcase.expected) {
			t.Errorf("Test number %d. Input: %s. Expected: %v, got %v", tcase.num, tcase.input, tcase.expected, encodedSymbols)
		}
		for i := 0; i < len(encodedSymbols); i++ {
			suspect := encodedSymbols[i]
			if suspect != tcase.expected[i] {
				t.Errorf("Test number %d. Input: %s. Expected: %v, got %v", tcase.num, tcase.input, tcase.expected, encodedSymbols)
			}
		}
	}
}

func TestCreateList(t *testing.T) {
	testCases := []struct {
		num         int
		occurrences occ.Occurrences
		expected    btll.NodeList
	}{
		{1, testutils.PrepareOccurrences(1), testutils.PrepareNodeList(1)},
		{2, testutils.PrepareOccurrences(2), testutils.PrepareNodeList(2)},
	}
	for _, tcase := range testCases {
		n := btll.NodeList{}
		n.CreateList(tcase.occurrences)
		if !testutils.LinkedListsAreEqual(n, tcase.expected) {
			t.Errorf("Test number %d. Expected: %v, got %v", tcase.num, tcase.expected, n)
		}
	}
}
