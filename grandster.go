package main

import ( 
	"strings"
 	"regexp"
 	"os"
 	"fmt"
 	"bufio"
 	"log"
 )

type NodePair struct {
	node1 *WordNode
	node2 *WordNode
	weight uint64
}

func (n NodePair) String() string {
	return fmt.Sprintf("%d %v %v", n.weight, n.node1.words, n.node2.words)
}

func main() {
	inputFile := os.Args[1]
	grandestWordPairs := grandestWordPairs(inputFile)
	for _,pair := range grandestWordPairs {
		fmt.Println(pair)
	}
}

func grandestWordPairs(path string) []NodePair {
	nodes := make(map[Bitstring]*WordNode)
	stripperRegEx, _ := regexp.Compile("[^A-ZÅÄÖa-zåäö-]+")

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

    scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		words := strings.Fields(scanner.Text())
		for _,word := range words {
			word = strings.ToUpper(stripperRegEx.ReplaceAllString(word, ""))
			if word != "" {
				wordToNode(word, nodes)
			}
		}
	}
	fmt.Println(len(nodes))

	return findMaxNodePairs(nodes)
}

// Handle words
func wordToNode(word string, nodes map[Bitstring]*WordNode) {
	newNode := NewNode(word)

	if len(nodes) == 0 {
		nodes[newNode.bits] = newNode
	} else if nodes[newNode.bits] != nil {
		nodes[newNode.bits].words[word] = true // We found an existing node with same characters, add the word
		return
	} else {
		found := false
		for _, oldNode := range nodes {
			if oldNode.bits.IsSubstring(newNode.bits) {
				found = true
				break // We found superior node, newNode is a character subset
			}
		}
		if !found {
			for key, oldNode := range nodes {
				if newNode.bits.IsSubstring(oldNode.bits) {
					delete(nodes, key)
				}
			}
			nodes[newNode.bits] = newNode
		}
	}
}

func findMaxNodePairs(nodes map[Bitstring]*WordNode) []NodePair {	
	pairs := make(map[uint64] []NodePair) 
	foundPairs := make(map[*WordNode] map[*WordNode]bool)

	var maxWeight uint64
	maxWeight = 0
	for _, node1 := range nodes {
		for _, node2 := range nodes {

			weight := node1.CombinedWeight(node2)
			if weight < maxWeight {
			   continue
			}

			nodePair := NodePair {node1: node1, node2: node2, weight: weight}

			if foundPairs[node2] != nil {
				if foundPairs[node2][node1] {
					continue
				} 
			} 
			if foundPairs[node1] == nil {
				foundPairs[node1] = make(map[*WordNode] bool)
			}
			foundPairs[node1][node2] = true

			if pairs[weight] == nil {
				pairs[weight] = make([]NodePair, 1)
				pairs[weight][0] = nodePair
			} else {
				pairs[weight] = append(pairs[weight], nodePair)
			}

			if weight > maxWeight {
				maxWeight = weight
			}
		}
	}

	return pairs[maxWeight]
}

