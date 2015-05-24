package main

import ( "fmt"; "strings" )
import "io/ioutil"
import "log"
import "regexp"

type WordNode struct {
	bits Bitstring
	words map[string]bool
}

func (n WordNode) Weight() uint64 {
	return n.bits.Hamming()
}

func (n WordNode) CombinedWeight(n2 *WordNode) uint64 {
	return n.bits.CombinedHamming(n2.bits)
}

func (n WordNode) String() string {
	return fmt.Sprintf("%b (%d) %v", n.bits, n.Weight(), n.words)
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

func main() {
	nodes := make(map[Bitstring]*WordNode)
	stripperRegEx, err := regexp.Compile("[^A-ZÅÄÖa-zåäö-]+")

	content, err := ioutil.ReadFile("pg135.txt")

	if err == nil {
        words := strings.Fields(string(content))
        for _, word := range words {
        	word = strings.ToUpper(stripperRegEx.ReplaceAllString(word, ""))
        	if (word == "") {
        		continue;
        	}
        	newNode := NewNode(word)

        	if len(nodes) == 0 {
        		nodes[newNode.bits] = newNode
        	} else if nodes[newNode.bits] != nil {
    			nodes[newNode.bits].words[word] = true // We found an existing node with same characters, add the word
    			continue;
    		} else {

	    		found := false
	        	for _, oldNode := range nodes {
	        		if oldNode.bits.IsSubstring(newNode.bits) {
	        			found = true
	        			break; // We found superior node, newNode is a character subset 
	        		}
	        	}
	        	if ! found {
		        	for key, oldNode := range nodes {
		        		if newNode.bits.IsSubstring(oldNode.bits) {
		        			delete(nodes, key)
		        		}
		        	}
		        	nodes[newNode.bits] = newNode
	        	}
	        }
	    }
    } else {
        log.Fatal(err)
    }

    pairs := make(map[uint64] [][]*WordNode)

    var maxWeight uint64
    maxWeight = 0
    for _,node1 := range nodes {
    	for _, node2 := range nodes {
    		nodePair := make([]*WordNode, 2)
    		nodePair[0] = node1
    		nodePair[1] = node2
    		weight := node1.CombinedWeight(node2)

    		if pairs[weight] == nil {
    			pairs[weight] = make([][]*WordNode, 1)
    			pairs[weight][0] = nodePair
    		} else {
    			pairs[weight] = append(pairs[weight], nodePair)
    		}

			if (weight > maxWeight) {			
				maxWeight = weight
			}
    	}
    }

    log.Print(maxWeight)
    for _, pair := range pairs[maxWeight] {
    	log.Print(pair[0].words, pair[1].words)
    }

}
