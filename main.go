package main

import (
	"fmt"
	"sort"
	"strings"
)

// Node - represent element of NodeList
//     Data - data to be coded
//     Freq - frequency of occurrence
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

// TreeNode - represent node in binary tree
type TreeNode struct {
	LeftData  interface{}
	RightData interface{}
	Parent    *TreeNode
	Freq      int
}

// type TreeNode struct {
// 	SubRoot          *TreeNode
// 	Freq             int
// 	LeftSubRoot      *TreeNode
// 	RightSubRoot     *TreeNode
// 	LeftSubRootData  interface{}
// 	RigthSubRootData interface{}
// 	Parent           *TreeNode
// }

// BinaryTree - represent binary tree
type BinaryTree struct {
	Root  *TreeNode
	Depth int
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
func (o Occurrences) Less(i, j int) bool { return o[i].Occurrences < o[j].Occurrences }

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
	// nl.displayList()
	// nl.displayListReverse()
	bt := &BinaryTree{}
	bt.createTree(*nl)
}

// Linked list methods
func (n *NodeList) createList(o Occurrences) {
	for _, v := range o {
		if n.length == 0 {
			node := &Node{Data: v.Symb, Freq: v.Occurrences}
			n.Head = node
			n.Tail = node
		} else {
			lastNode := n.Tail
			newNode := &Node{Data: v.Symb, Freq: v.Occurrences}
			lastNode.Next = newNode
			lastNode.Next.Prev = lastNode
			n.Tail = newNode
		}
		n.length++
	}
}

func (n *NodeList) insertByFreq(tn *TreeNode) {
	node := &Node{Data: tn, Freq: tn.Freq}
	if n.length == 0 {
		n.Head = node
		n.Tail = node
	} else {
		start := n.Head
		for start != nil {
			if start.Next == nil {
				node.Prev = start
				start.Next = node
				n.Tail = start.Next
				break
			}
			if tn.Freq > start.Freq && tn.Freq <= start.Next.Freq {
				reserve := start.Next
				start.Next = &Node{Data: tn, Freq: tn.Freq, Next: reserve, Prev: start}
				break
			}
			start = start.Next
		}
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

// Binary tree methods
var i = 1

func (bt *BinaryTree) createTree(list NodeList) {
	for list.length != 1 {
		firstElement := list.getSmallestFreq()
		secondElement := list.getSmallestFreq()
		fmt.Printf("%d. first: %v, second: %v, linked list length: %d\n", i, firstElement, secondElement, list.length)
		tn := &TreeNode{LeftData: firstElement, RightData: secondElement, Freq: firstElement.Freq + secondElement.Freq}
		list.insertByFreq(tn)
		list.displayList()
		fmt.Println("After new node insertion length is ", list.length)
		i++
	}
}

func (n *NodeList) getSmallestFreq() Node {
	toReturn := *n.Head
	if n.Head.Next != nil {
		n.Head = n.Head.Next
		n.Head.Prev = nil
	} else {
		n.Head = nil
	}
	n.length--
	return toReturn
}
