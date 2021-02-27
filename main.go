package main

import (
	"log"
	"sort"
	"strings"
)

// Node - represent element of NodeList
//  Data - data to be coded
//  Freq - frequency of occurrence
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
	LeftData       interface{}
	RightData      interface{}
	LeftBranchHas  []rune
	RightBranchHas []rune
	Parent         *TreeNode
	Freq           int
}

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
	// word := "aaabbccccde"
	var occurrences Occurrences
	var doubles []rune
	log.Printf("%v", []rune(word))
	for _, v := range word {
		if isUnique(v, doubles) {
			occurrences = append(occurrences, Occurence{Symb: v, Occurrences: strings.Count(string(word), string(v))})
			doubles = append(doubles, v)
		}
	}
	sort.Sort(occurrences)
	nl := &NodeList{}
	nl.createList(occurrences)
	bt := &BinaryTree{}
	bt.createTree(nl)
	nl.displayList()
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
		for {
			if start.Next == nil {
				if start.Freq > tn.Freq {
					start.Prev = node
					node.Next = start
				} else {
					node.Prev = start
					start.Next = node
					n.Tail = start.Next
				}
				break
			} else if tn.Freq >= start.Freq && tn.Freq <= start.Next.Freq {
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
	log.Printf("Displaying linked list:\n")
	for toPrint != nil {
		log.Printf("%v -> ", toPrint.Data)
		toPrint = toPrint.Next
	}
	log.Println("<nil>")
	log.Printf("Done\n")
}

func (n NodeList) displayListReverse() {
	toPrint := n.Tail
	for toPrint != nil {
		log.Printf("%v -> ", toPrint.Data)
		toPrint = toPrint.Prev
	}
	log.Print("<nil>")
}

func (bt *BinaryTree) createTree(list *NodeList) {
	for list.length != 1 {
		firstElement := list.getSmallestFreq()
		secondElement := list.getSmallestFreq()
		tn := &TreeNode{LeftData: firstElement.Data, RightData: secondElement.Data, Freq: firstElement.Freq + secondElement.Freq}
		if _, ok := tn.LeftData.(*TreeNode); ok {
			tn.LeftData.(*TreeNode).Parent = tn
			/* TODO Gather all children by instruction below
			 * tn.LeftData is *TreeNode. Gather all children from:
			 *    tn.LeftData.(*TreeNode).LeftBranchHas
			 *    tn.LeftData.(*TreeNode).RightBranchHas
			 *    tn.LeftData.(*TreeNode).LeftData (if not *TreeNode)
			 *    tn.LeftData.(*TreeNode).RightData (if not *TreeNode)
			 * assign all this runes to tn.LeftBranchHas
			 */
			// if tn.LeftData.(*TreeNode).RightBranchHas == nil && tn.LeftData.(*TreeNode).LeftBranchHas == nil {
			// 	if tn.LeftData.(*TreeNode).RightData != nil {
			// 		tn.LeftBranchHas = append(tn.LeftBranchHas, tn.LeftData.(*TreeNode).RightData.(rune))
			// 	}
			// 	if tn.LeftData.(*TreeNode).LeftData != nil {
			// 		tn.LeftBranchHas = append(tn.LeftBranchHas, tn.LeftData.(*TreeNode).LeftData.(rune))
			// 	}
			// } else {
			// 	if tn.RightData.(*TreeNode).RightBranchHas != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).RightBranchHas...)
			// 	}
			// 	if tn.RightData.(*TreeNode).LeftBranchHas != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).LeftBranchHas...)
			// 	}
			// }
		} else if _, ok = tn.RightData.(*TreeNode); ok {
			tn.RightData.(*TreeNode).Parent = tn
			/* TODO Gather all children by instruction below
			 * tn.RightData is *TreeNode. Gather all children from:
			 *    tn.RightData.(*TreeNode).LeftBranchHas
			 *    tn.RightData.(*TreeNode).RightBranchHas
			 *    tn.RightData.(*TreeNode).LeftData (if not *TreeNode)
			 *    tn.RightData.(*TreeNode).RightData (if not *TreeNode)
			 * assign all this runes to tn.RightBranchHas
			 */
			// if tn.RightData.(*TreeNode).RightBranchHas == nil && tn.RightData.(*TreeNode).LeftBranchHas == nil {
			// 	if tn.RightData.(*TreeNode).RightData != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).RightData.(rune))
			// 	}
			// 	if tn.RightData.(*TreeNode).LeftData != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).LeftData.(rune))
			// 	}
			// } else {
			// 	if tn.RightData.(*TreeNode).RightBranchHas != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).RightBranchHas...)

			// 	}
			// 	if tn.RightData.(*TreeNode).LeftBranchHas != nil {
			// 		tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).LeftBranchHas...)
			// 	}
			// 	tn.RightBranchHas = append(tn.RightBranchHas, tn.RightData.(*TreeNode).LeftData.(rune))
			// }
		}
		list.insertByFreq(tn)
		list.displayList()
	}
	log.Println(list.Head.Data.(*TreeNode))
	var ok bool
	if bt.Root, ok = list.Head.Data.(*TreeNode); !ok {
		log.Fatalf("Error occured can't cast interface to struct")
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
