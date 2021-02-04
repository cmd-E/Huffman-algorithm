package main

import (
	"fmt"
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
	Prev *Node
}

// NodeList - represent list of nodes
type NodeList struct {
	Head   *Node
	Tail   *Node
	length int
}

// Occurence - represent element in Occurrences slice
type Occurence struct {
	Symb        rune
	Occurrences int
}

// Occurrences - represent slice with symblos of word and their occurrencea
type Occurrences []Occurence

// Methods for sort.Sort()
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
	// words := os.Args[1:]
	// if len(words) == 0 {
	// 	log.Fatalln("No argument provided")
	// }
	// word := words[0]
	word := "hello"
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
	fmt.Println(occurrences)
	nl.createList(occurrences)
	nl.displayList()
	nl.displayListReverse()
}

func (n *NodeList) createList(o Occurrences) {
	for _, v := range o {
		n.insertNode(v.Symb, v.Occurrences)
	}
}

func (n *NodeList) insertNode(symb rune, freq int) {
	if n.length == 0 {
		node := &Node{Data: symb, Freq: freq}
		n.Head = node
		n.Tail = node
	} else {
		lastNode := n.Tail
		newNode := &Node{Data: symb, Freq: freq}
		lastNode.Next = newNode
		lastNode.Next.Prev = lastNode
		n.Tail = newNode
	}
	n.length++
}

func (n NodeList) displayList() {
	toPrint := n.Head
	for toPrint != nil {
		fmt.Printf("%v -> ", toPrint.Data)
		toPrint = toPrint.Next
	}
	fmt.Println("<nil>")
}

func (n NodeList) displayListReverse() {
	toPrint := n.Tail
	for toPrint != nil {
		fmt.Printf("%v -> ", toPrint.Data)
		toPrint = toPrint.Prev
	}
	fmt.Print("<nil>")
}
