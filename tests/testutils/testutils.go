package testutils

import (
	btll "github.com/cmd-e/huffman-algorithm/btAndLinkedList"
	occ "github.com/cmd-e/huffman-algorithm/occpackage"
)

// PrepareNodeList - prepares sample node list for testing based on number of test
func PrepareNodeList(num int) btll.NodeList {
	n := btll.NodeList{}
	o := PrepareOccurrences(num)
	for _, v := range o {
		node := &btll.Node{Data: v.Symb, Freq: v.Occurrences}
		if n.Length == 0 {
			n.Head = node
			n.Tail = node
		} else {
			lastNode := n.Tail
			newNode := node
			lastNode.Next = newNode
			lastNode.Next.Prev = lastNode
			n.Tail = newNode
		}
		n.Length++
	}
	return n
}

// PrepareOccurrences - prepares sample occurrences for testing based on number of test
func PrepareOccurrences(num int) occ.Occurrences {
	occs := occ.Occurrences{}
	switch num {
	case 1:
		occs = append(occs,
			occ.Occurrence{Symb: 'д', Occurrences: 5},
			occ.Occurrence{Symb: 'г', Occurrences: 6},
			occ.Occurrence{Symb: 'в', Occurrences: 6},
			occ.Occurrence{Symb: 'б', Occurrences: 7},
			occ.Occurrence{Symb: 'а', Occurrences: 15})
		break
	case 2:
		occs = append(occs,
			occ.Occurrence{Symb: 'c', Occurrences: 3},
			occ.Occurrence{Symb: 'b', Occurrences: 2},
			occ.Occurrence{Symb: 'a', Occurrences: 1})
		break
	case 3:
		break
	case 4:
		break
	case 5:
		break
	}
	return occs
}

// LinkedListsAreEqual - checks if two node lists are equal
func LinkedListsAreEqual(list1, list2 btll.NodeList) bool {
	if list1.Length != list2.Length || (list1.Head.Data != list2.Head.Data || list1.Tail.Data != list2.Tail.Data) {
		return false
	}
	for list2.Head != nil || list1.Head != nil {
		if list1.Head.Data != list2.Head.Data || list1.Head.Freq != list2.Head.Freq {
			return false
		}
		list2.Head = list2.Head.Next
		list1.Head = list1.Head.Next
	}
	if list1.Head == nil && list2.Head == nil {
		return true
	}
	return false
}

func insertNode(list btll.NodeList, data occ.Occurrence) btll.NodeList {
	node := &btll.Node{Data: data.Symb, Freq: data.Occurrences}
	if list.Length == 0 {
		list.Head = node
		list.Tail = node
	} else {
		lastNode := list.Tail
		newNode := node
		lastNode.Next = newNode
		lastNode.Next.Prev = lastNode
		list.Tail = newNode
	}
	return list
}
