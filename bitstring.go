package main

import "github.com/hideo55/go-popcount"

type Bitstring struct {
	value uint64
}

// Returns true if given bit is on
func (b Bitstring) IsOn(index int) bool {
	return b.value&(1<<uint(index)) != 0
}

// Set bit on
func (b *Bitstring) SetBit(index int) {
	b.value |= 1 << uint(index)
}

// Returns true if given bitstring is a substring
func (b Bitstring) IsSubstring(that Bitstring) bool {
	return b.value&that.value == that.value
}

func (b Bitstring) Hamming() uint64 {
	return popcount.Count(b.value)
}

func (b Bitstring) CombinedHamming(that Bitstring) uint64 {
	return popcount.Count(b.value | that.value)
}
