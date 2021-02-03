package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// Node - represent element of NodeList
//     Data - data to be coded
//     Freq - frequency of occurrence
//     Next - pointer to next element
type Node struct {
	Data interface{}
	Freq int
	Next *Node
}

// NodeList - represent list of nodes
type NodeList struct {
	Node *Node
}

// Occurence - represent element in Occurrences slice
type Occurence struct {
	Symb        rune
	Occurrences int
}

// Occurrences - represent slice with symblos of word and their occurrencea
type Occurrences []Occurence

func (o Occurrences) Len() int           { return len(o) }
func (o Occurrences) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o Occurrences) Less(i, j int) bool { return o[i].Occurrences > o[j].Occurrences }

func isUnique(r rune, list []rune) bool {
	for _, v := range list {
		if v == r {
			return false
		}
	}
	return true
}

func main() {
	word := os.Args[1]
	var occurrences Occurrences
	var doubles []rune
	for _, v := range word {
		if isUnique(v, doubles) {
			occurrences = append(occurrences, Occurence{Symb: v, Occurrences: strings.Count(string(word), string(v))})
			doubles = append(doubles, v)
		}
	}
	sort.Sort(occurrences)
	nl := &NodeList{}
	nl.createList(occurrences)
	nl.displayList()
}

func (n *NodeList) createList(o Occurrences) {
	for _, v := range o {
		n.insertNode(v.Symb, v.Occurrences)
	}
}

// Method is broken
func (n *NodeList) insertNode(symb rune, freq int) {
	if n.Node == nil {
		n.Node = &Node{Data: symb, Freq: freq}
		return
	}
	for n.Node.Next != nil {
		n.Node = n.Node.Next
	}
	n.Node.Next = &Node{Data: symb, Freq: freq}
}

func (n *NodeList) displayList() {
	for n.Node.Next != nil {
		fmt.Printf("%v -> ", n.Node.Data)
		n.Node = n.Node.Next
	}
	fmt.Printf("%v", n.Node.Data)
}
