package main
import (
       "testing"
       "strings"
       "github.com/stretchr/testify/assert"
  )

var testWords = []struct {
    word1 string
    word2 string
    weight1 int
    weight2 int
    combinedWeight int
    isSubstring bool
}{	
	{
		word1: "ABCDEFGHIJKLMNOPQRSTUVWXYZÄÖ",
		word2: "KABANA",
		weight1: 28,
		weight2: 4,
		combinedWeight: 28,
		isSubstring: true,
	},
	{
		word1: "AAAAAAAAAAAAAAAA",
		word2: "BBBBBBBBBBBBBBBBBBB",
		weight1: 1,
		weight2: 1,
		combinedWeight: 2,
		isSubstring: false,
	},
	{
		word1: "LÉÉT",
		word2: "LEET",
		weight1: 3,
		weight2: 3,
		combinedWeight: 4,
		isSubstring: false,
	},
	{
		word1: "VYÖRYTYSTÄ",
		word2: "YRJÖNTYVI",
		weight1: 7,
		weight2: 8,
		combinedWeight: 10,
		isSubstring: false,
	},
}


func TestNodePairFunctions(t *testing.T) {
     for _, test := range testWords {
     	 n1 := NewNode(strings.ToUpper(test.word1))
	 n2 := NewNode(strings.ToUpper(test.word2))
	 assert.Equal(t, test.weight1, n1.Weight())
	 assert.Equal(t, test.weight2, n2.Weight())
	 assert.Equal(t, test.combinedWeight, n1.CombinedWeight(n2))
	 assert.Equal(t, test.isSubstring, n1.bits.IsSubstring(n2.bits))
     }
}