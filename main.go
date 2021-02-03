package main

import (
	"fmt"
	"sort"
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

type Pair struct {
	Key   rune
	Value int
}

type PairList []Pair

func (p PairList) Len() int           { return len(p) }
func (p PairList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p PairList) Less(i, j int) bool { return p[i].Value > p[j].Value }

func main() {
	// word := os.Args[1]
	word := "Hello"
	occurrences := make(map[rune]int)
	for _, v := range word {
		occurrences[v]++
	}
	p := make(PairList, len(occurrences))
	i := 0
	for k, v := range occurrences {
		p[i] = Pair{k, v}
		i++
	}
	sort.Sort(p)
	sortedOccurrences := make(map[rune]int)
	for _, k := range p {
		sortedOccurrences[k.Key] = k.Value
	}
	for k, v := range sortedOccurrences {
		fmt.Printf("%v: %v\n", k, v)
	}
	nl := &NodeList{}
	nl.createList(sortedOccurrences)
	nl.displayList()
}

func (n *NodeList) createList(m map[rune]int) {
	for k, v := range m {
		n.insertNode(k, v)
	}
}

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

func getRarest(m map[rune]int) (rune, int, map[rune]int) {
	var smK rune
	var smV int
	for i := 0; i < len(m)-1; i++ {
		// for j := 0; j < len(m)-i-1; j++ {
		// 	if (m)
		// }
	}
	for k, v := range m {
		if v < smV {
			smV = v
			smK = k
			delete(m, k)
		}
	}
	return smK, smV, m
}
