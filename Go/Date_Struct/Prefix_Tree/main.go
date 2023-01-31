package main

import "log"

const alphabetSize = 26

type Node struct {
	children [alphabetSize]*Node
	isEnd    bool
}

type Trie struct {
	root *Node
}

func InitTrie() *Trie {
	result := &Trie{root: &Node{}}
	return result
}

func (t *Trie) Insert(w string) {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			currentNode.children[charIndex] = &Node{}
		}
		currentNode = currentNode.children[charIndex]
	}
	currentNode.isEnd = true
}

func (t *Trie) Search(w string) bool {
	wordLenght := len(w)
	currentNode := t.root
	for i := 0; i < wordLenght; i++ {
		charIndex := w[i] - 'a'
		if currentNode.children[charIndex] == nil {
			return false
		}
		currentNode = currentNode.children[charIndex]
	}
	if currentNode.isEnd == true {
		return true
	}
	return false
}

func main() {
	trie := InitTrie()

	toAdd := []string{
		"ninja",
		"janil",
		"lopata",
	}

	for _, m := range toAdd {
		trie.Insert(m)
	}

	log.Println(trie.Search("ninja"))
}
