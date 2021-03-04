package btll

import (
	"log"

	occ "github.com/cmd-e/huffman-algorithm/occpackage"
)

// TreeNode - represent node in binary tree
type TreeNode struct {
	LeftData       interface{}
	RightData      interface{}
	LeftBranchHas  []rune
	RightBranchHas []rune
	Parent         *TreeNode
	Freq           float64
}

// BinaryTree - represent binary tree
type BinaryTree struct {
	Root  *TreeNode
	Depth int
}

// Encoded struct represents symbol and it's encoded form
type Encoded struct {
	Symb rune
	Code string
}

// CreateTree - creates tree from list
func (bt *BinaryTree) CreateTree(list *NodeList) {
	for list.Length != 1 {
		firstElement := list.getElementBySmallestFrequency()
		secondElement := list.getElementBySmallestFrequency()
		tn := &TreeNode{LeftData: firstElement.Data, RightData: secondElement.Data, Freq: firstElement.Freq + secondElement.Freq}
		if LDNode, ok := tn.LeftData.(*TreeNode); ok {
			tn.LeftData.(*TreeNode).Parent = tn
			tn.LeftBranchHas = getAllChildren(LDNode)
		}
		if RDNode, ok := tn.RightData.(*TreeNode); ok {
			tn.RightData.(*TreeNode).Parent = tn
			tn.RightBranchHas = getAllChildren(RDNode)
		}

		list.insertElementByFrequency(tn)
	}
	var ok bool
	if bt.Root, ok = list.Head.Data.(*TreeNode); !ok {
		log.Fatalf("Error occured can't cast interface to struct")
	}
}

// EncodeSymbols encodes symbols to binary by created tree
func (bt *BinaryTree) EncodeSymbols(symbolsToEncode []rune) []Encoded {
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

func itemExist(arr []rune, toFind rune) bool {
	for _, r := range arr {
		if r == toFind {
			return true
		}
	}
	return false
}

// Node - represent element of NodeList
//  Data - data to be coded
//  Freq - frequency of occurrence
type Node struct {
	Data interface{}
	Freq float64
	Next *Node
	Prev *Node
}

// NodeList - represent list of nodes
type NodeList struct {
	Head   *Node
	Tail   *Node
	Length int
}

// CreateList - creates linked list from user's input
func (n *NodeList) CreateList(o occ.Occurrences) {
	for _, v := range o {
		if n.Length == 0 {
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
		n.Length++
	}
}

func (n *NodeList) insertElementByFrequency(tn *TreeNode) {
	nodeToInsert := &Node{Data: tn, Freq: tn.Freq}
	isInserted := false
	if n.Length == 0 {
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
	n.Length++
}

func (n *NodeList) getElementBySmallestFrequency() Node {
	toReturn := *n.Head
	if n.Head.Next != nil {
		n.Head = n.Head.Next
		n.Head.Prev = nil
	} else {
		n.Head = nil
		n.Tail = nil
	}
	n.Length--
	return toReturn
}
