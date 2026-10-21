package btll

import (
	"container/heap"
	"fmt"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
	"math"
)

type TreeNode struct {
	Symbol      rune
	Leaf        bool
	Left, Right *TreeNode
	Freq        float64
	priority    int
}
type BinaryTree struct {
	Root  *TreeNode
	Depth int
	codes map[rune]string
}
type Encoded struct {
	Symb rune
	Code string
}
type priorityQueue []*TreeNode

func (q priorityQueue) Len() int { return len(q) }
func (q priorityQueue) Less(i, j int) bool {
	if q[i].Freq == q[j].Freq {
		return q[i].priority < q[j].priority
	}
	return q[i].Freq < q[j].Freq
}
func (q priorityQueue) Swap(i, j int)       { q[i], q[j] = q[j], q[i] }
func (q *priorityQueue) Push(x interface{}) { *q = append(*q, x.(*TreeNode)) }
func (q *priorityQueue) Pop() interface{} {
	old := *q
	n := old[len(old)-1]
	*q = old[:len(old)-1]
	return n
}

// CreateTree retains the educational linked-list API without using it in the CLI.
func (bt *BinaryTree) CreateTree(list *NodeList) error {
	if list == nil || list.Length < 2 {
		return fmt.Errorf("at least two symbols required")
	}
	values := make(occ.Occurrences, 0, list.Length)
	for node := list.Head; node != nil; node = node.Next {
		if len(values) >= list.Length {
			return fmt.Errorf("invalid or cyclic list")
		}
		symbol, ok := node.Data.(rune)
		if !ok {
			return fmt.Errorf("invalid list symbol")
		}
		values = append(values, occ.Occurrence{Symb: symbol, Occurrences: node.Freq})
	}
	if len(values) != list.Length {
		return fmt.Errorf("invalid list length")
	}
	return bt.Build(values)
}

// Build constructs a code table directly from typed symbol weights.
func (bt *BinaryTree) Build(values occ.Occurrences) error {
	bt.Root = nil
	bt.codes = nil
	bt.Depth = 0
	if len(values) < 2 {
		return fmt.Errorf("at least two symbols required")
	}
	queue := priorityQueue{}
	seen := make(map[rune]bool)
	for index, value := range values {
		if seen[value.Symb] || value.Occurrences <= 0 || math.IsNaN(value.Occurrences) || math.IsInf(value.Occurrences, 0) {
			return fmt.Errorf("invalid or duplicate symbol weight")
		}
		seen[value.Symb] = true
		queue = append(queue, &TreeNode{Symbol: value.Symb, Leaf: true, Freq: value.Occurrences, priority: index})
	}
	heap.Init(&queue)
	sequence := 0
	for queue.Len() > 1 {
		left := heap.Pop(&queue).(*TreeNode)
		right := heap.Pop(&queue).(*TreeNode)
		sequence++
		weight := left.Freq + right.Freq
		if math.IsInf(weight, 0) {
			return fmt.Errorf("combined weights overflow")
		}
		heap.Push(&queue, &TreeNode{Left: left, Right: right, Freq: weight, priority: -sequence})
	}
	bt.Root = heap.Pop(&queue).(*TreeNode)
	bt.codes = make(map[rune]string)
	var visit func(*TreeNode, string)
	visit = func(node *TreeNode, code string) {
		if node.Leaf {
			bt.codes[node.Symbol] = code
			if len(code) > bt.Depth {
				bt.Depth = len(code)
			}
			return
		}
		visit(node.Left, code+"0")
		visit(node.Right, code+"1")
	}
	visit(bt.Root, "")
	return nil
}
func (bt *BinaryTree) EncodeSymbols(symbols []rune) []Encoded {
	values := make([]Encoded, 0, len(symbols))
	for _, symbol := range symbols {
		if code, ok := bt.codes[symbol]; ok {
			values = append(values, Encoded{symbol, code})
		}
	}
	return values
}

type Node struct {
	Data       interface{}
	Freq       float64
	Next, Prev *Node
}
type NodeList struct {
	Head, Tail *Node
	Length     int
}

func (n *NodeList) CreateList(values occ.Occurrences) {
	n.Head = nil
	n.Tail = nil
	n.Length = 0
	for _, value := range values {
		node := &Node{Data: value.Symb, Freq: value.Occurrences, Prev: n.Tail}
		if n.Tail == nil {
			n.Head = node
		} else {
			n.Tail.Next = node
		}
		n.Tail = node
		n.Length++
	}
}
