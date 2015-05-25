package main

import "fmt"

type WordNode struct {
	bits  Bitstring
	words map[string] bool
}

func (n WordNode) String() string {
	return fmt.Sprintf("%b (%d) %v", n.bits, n.Weight(), n.words)
}

func (n WordNode) Weight() uint64 {
	return n.bits.Hamming()
}

func (n WordNode) CombinedWeight(n2 *WordNode) uint64 {
	return n.bits.CombinedHamming(n2.bits)
}


func NewNode(word string) *WordNode {
	node := new(WordNode)
	node.words = make(map[string]bool)
	node.words[word] = true

	for _, r := range word {
		// Normalize to zero
		c := r - 65

		// Handle the Röck Döts
		if c == 132 {
			c = 26
		}
		if c == 131 {
			c = 27
		}
		if c == 149 {
			c = 28
		}

		node.bits.SetBit(int(c))
	}
	return node
}
