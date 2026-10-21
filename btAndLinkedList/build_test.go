package btll

import (
	"math"
	"reflect"
	"strings"
	"testing"

	occ "github.com/cmd-e/huffman-algorithm/occpackage"
)

func TestBuildRejectsInvalidWeightsAndResetsState(t *testing.T) {
	valid := occ.Occurrences{{Symb: 'a', Occurrences: 1}, {Symb: 'b', Occurrences: 2}}
	cases := []occ.Occurrences{nil, valid[:1], {{Symb: 'a', Occurrences: 1}, {Symb: 'a', Occurrences: 2}},
		{{Symb: 'a', Occurrences: 0}, {Symb: 'b', Occurrences: 1}}, {{Symb: 'a', Occurrences: -1}, {Symb: 'b', Occurrences: 1}},
		{{Symb: 'a', Occurrences: math.NaN()}, {Symb: 'b', Occurrences: 1}},
		{{Symb: 'a', Occurrences: math.Inf(1)}, {Symb: 'b', Occurrences: 1}},
		{{Symb: 'a', Occurrences: math.MaxFloat64}, {Symb: 'b', Occurrences: math.MaxFloat64}}}
	for i, values := range cases {
		t.Run(string(rune('A'+i)), func(t *testing.T) {
			tree := &BinaryTree{}
			if err := tree.Build(valid); err != nil {
				t.Fatal(err)
			}
			if err := tree.Build(values); err == nil {
				t.Fatal("invalid tree accepted")
			}
			if tree.Root != nil || tree.Depth != 0 || len(tree.EncodeSymbols([]rune{'a'})) != 0 {
				t.Fatal("previous tree leaked after failed rebuild")
			}
		})
	}
}

func TestCreateTreeRejectsBrokenLists(t *testing.T) {
	cycle := &Node{Data: 'a', Freq: 1}
	cycle.Next = cycle
	for _, list := range []*NodeList{nil, {}, {Length: 2}, {Length: 2, Head: &Node{Data: "not a rune"}}, {Length: 2, Head: cycle}} {
		if err := new(BinaryTree).CreateTree(list); err == nil {
			t.Fatal("broken list accepted")
		}
	}
}

func TestCodesAreDeterministicPrefixFreeAndWeighted(t *testing.T) {
	values, symbols := occ.GetOccurrencesAndUniqueSymbols("aaaaaaaabbbbcc😀")
	var previous []Encoded
	for iteration := 0; iteration < 10; iteration++ {
		tree := new(BinaryTree)
		if err := tree.Build(values); err != nil {
			t.Fatal(err)
		}
		codes := tree.EncodeSymbols(symbols)
		if len(codes) != len(symbols) || tree.Depth < 1 {
			t.Fatal("missing codes")
		}
		if previous != nil && !reflect.DeepEqual(previous, codes) {
			t.Fatal("unstable codes")
		}
		previous = codes
		for i, first := range codes {
			for j, second := range codes {
				if i != j && strings.HasPrefix(first.Code, second.Code) {
					t.Fatal("prefix collision")
				}
			}
		}
		if len(codes[0].Code) > len(codes[len(codes)-1].Code) {
			t.Fatal("most frequent symbol has longer code")
		}
		if len(tree.EncodeSymbols([]rune{'?'})) != 0 {
			t.Fatal("unknown symbol encoded")
		}
	}
}

func TestCreateListReplacesOldContentsAndLinksBothDirections(t *testing.T) {
	list := new(NodeList)
	list.CreateList(occ.Occurrences{{Symb: 'a', Occurrences: 1}, {Symb: 'b', Occurrences: 2}, {Symb: 'c', Occurrences: 3}})
	if list.Length != 3 || list.Head.Prev != nil || list.Tail.Next != nil || list.Head.Next.Prev != list.Head || list.Tail.Prev.Next != list.Tail {
		t.Fatal("invalid list links")
	}
	list.CreateList(nil)
	if list.Length != 0 || list.Head != nil || list.Tail != nil {
		t.Fatal("old contents retained")
	}
}
