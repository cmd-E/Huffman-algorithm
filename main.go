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

// Encoded struct represents symbol and it's encoded form
type Encoded struct {
	Symb rune
	Code string
}

// Methods for sort.Sort()
func (o Occurrences) Len() int           { return len(o) }
func (o Occurrences) Swap(i, j int)      { o[i], o[j] = o[j], o[i] }
func (o Occurrences) Less(i, j int) bool { return o[i].Occurrences <= o[j].Occurrences }

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
	// word := "hellow"
	// word := "aaabbccccde"
	word := "ааааааааааааааабббббббввввввггггггддддд"
	var occurrences Occurrences
	var doubles []rune
	log.Printf("word to encode is %s\n", word)
	log.Printf("runes to encode are %v\n", []rune(word))
	for _, v := range word {
		if isUnique(v, doubles) {
			occurrences = append(occurrences, Occurence{Symb: v, Occurrences: strings.Count(string(word), string(v))})
			doubles = append(doubles, v)
		}
	}
	// BUG sorting methods are working not as intended. Check with "ааааааааааааааабббббббввввввггггггддддд"
	sort.Sort(occurrences)
	nl := &NodeList{}
	nl.createList(occurrences)
	bt := &BinaryTree{}
	bt.createTree(nl)
	uniqueSymbols := doubles
	encodedSymbols := bt.encodeSymbols(uniqueSymbols)
	for _, v := range encodedSymbols {
		log.Printf("Code for %c is %s", v.Symb, v.Code)
	}
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
	nodeToInsert := &Node{Data: tn, Freq: tn.Freq}
	isInserted := false
	if n.length == 0 {
		n.Head = nodeToInsert
		n.Tail = nodeToInsert
	} else {
		currentNode := n.Head
		for currentNode != nil {
			if nodeToInsert.Freq <= currentNode.Freq {
				if currentNode == n.Head {
					n.Head = nodeToInsert
					nodeToInsert.Next = currentNode
					currentNode.Prev = nodeToInsert
					isInserted = true
				} else {
					nodeToInsert.Prev = currentNode.Prev
					currentNode.Prev.Next = nodeToInsert
					nodeToInsert.Next = currentNode
					currentNode.Prev = nodeToInsert
					isInserted = true
				}
				break
			}
			currentNode = currentNode.Next
		}
		if !isInserted {
			reserve := n.Tail
			nodeToInsert.Prev = reserve
			reserve.Next = nodeToInsert
			n.Tail = nodeToInsert
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
		if LDNode, ok := tn.LeftData.(*TreeNode); ok {
			tn.LeftData.(*TreeNode).Parent = tn
			tn.LeftBranchHas = getAllChildren(LDNode)
		}
		if RDNode, ok := tn.RightData.(*TreeNode); ok {
			tn.RightData.(*TreeNode).Parent = tn
			tn.RightBranchHas = getAllChildren(RDNode)
		}

		list.insertByFreq(tn)
		// list.displayList()
	}
	log.Println(list.Head.Data.(*TreeNode))
	var ok bool
	if bt.Root, ok = list.Head.Data.(*TreeNode); !ok {
		log.Fatalf("Error occured can't cast interface to struct")
	}
}

func getAllChildren(node *TreeNode) []rune {
	var children []rune
	if node.LeftBranchHas != nil {
		children = append(children, node.LeftBranchHas...)
	}
	if node.RightBranchHas != nil {
		children = append(children, node.RightBranchHas...)
	}
	if _, ok := node.LeftData.(*TreeNode); !ok {
		children = append(children, node.LeftData.(rune))
	}
	if _, ok := node.RightData.(*TreeNode); !ok {
		children = append(children, node.RightData.(rune))
	}
	return children
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

func (bt *BinaryTree) encodeSymbols(symbolsToEncode []rune) []Encoded {
	var encodedVals []Encoded
	for i := 0; i < len(symbolsToEncode); i++ {
		target := symbolsToEncode[i]
		encodedVals = append(encodedVals, Encoded{Symb: target, Code: bt.traverseTree(target)})
	}
	return encodedVals
}

func (bt *BinaryTree) traverseTree(symbol rune) string {
	var strCode []rune
	localRoot := bt.Root
	for {
		if tempRune, ok := localRoot.LeftData.(rune); ok && tempRune == symbol {
			strCode = append(strCode, '0')
			return string(strCode)
		}
		if tempRune, ok := localRoot.RightData.(rune); ok && tempRune == symbol {
			strCode = append(strCode, '1')
			return string(strCode)
		}
		if localRoot.LeftBranchHas != nil && itemExist(localRoot.LeftBranchHas, symbol) {
			localRoot = localRoot.LeftData.(*TreeNode)
			strCode = append(strCode, '0')
			continue
		}
		if localRoot.RightBranchHas != nil && itemExist(localRoot.RightBranchHas, symbol) {
			localRoot = localRoot.RightData.(*TreeNode)
			strCode = append(strCode, '1')
			continue
		}
	}
}

func itemExist(arr []rune, toFind rune) bool {
	for _, r := range arr {
		if r == toFind {
			return true
		}
	}
	return false
}
