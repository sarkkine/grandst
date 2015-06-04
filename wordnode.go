package main

import "fmt"

type WordNode struct {
	bits  Bitstring
	words map[string] bool
}

func (n WordNode) String() string {
	return fmt.Sprintf("%b (%d) %v", n.bits, n.Weight(), n.words)
}

func (n WordNode) Weight() int {
	return int(n.bits.Hamming())
}

func (n WordNode) CombinedWeight(n2 *WordNode) int {
	return int(n.bits.CombinedHamming(n2.bits))
}


func NewNode(word string) *WordNode {
	node := new(WordNode)
	node.words = make(map[string]bool)
	node.words[word] = true

	for _, r := range word {
		// Normalize to zero
		c := r - 65

		// Handle lé Röck Döts
		if c >= 127 {
			c = c - 99
		}

		node.bits.SetBit(int(c))
	}
	return node
}
