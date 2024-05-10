package main

import (
	"bufio"
	"fmt"
	"os"
)

type TrieNode struct {
	children    map[rune]*TrieNode
	isEndOfWord bool
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode), isEndOfWord: false}
}

type Trie struct {
	root *TrieNode
}

func NewTrie() *Trie {
	return &Trie{root: NewTrieNode()}
}

// Insert a word into the Trie, making sure the word is in uppercase
func (t *Trie) Insert(word string) {
	currentNode := t.root
	for _, char := range word {
		if _, ok := currentNode.children[char]; !ok {
			currentNode.children[char] = NewTrieNode()
		}
		currentNode = currentNode.children[char]
	}
	currentNode.isEndOfWord = true
}

// Search the Trie for words that start with the given prefix, prefix is assumed to be uppercase
func (t *Trie) SearchPrefix(prefix string) []string {
	currentNode := t.root
	for _, char := range prefix {
		if node, ok := currentNode.children[char]; ok {
			currentNode = node
		} else {
			return nil
		}
	}
	return t.collectWords(currentNode, prefix)
}

// Helper function to collect all words from a given node
func (t *Trie) collectWords(node *TrieNode, prefix string) []string {
	var words []string
	if node.isEndOfWord {
		words = append(words, prefix)
	}
	for char, childNode := range node.children {
		words = append(words, t.collectWords(childNode, prefix+string(char))...)
	}
	return words
}

func main() {
	trie := NewTrie()
	words := []string{"hello", "helium", "hero", "heron", "help", "apple", "aardvark"}
	for _, word := range words {
		trie.Insert(word)
	}

	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter a prefix to get autocomplete suggestions (type 'exit' to quit):")
	for scanner.Scan() {
		input := scanner.Text()
		if input == "exit" {
			break
		}
		suggestions := trie.SearchPrefix(input)
		if suggestions != nil {
			fmt.Println("Autocomplete suggestions:", suggestions)
		} else {
			fmt.Println("No suggestions found.")
		}
		fmt.Println("Enter another prefix (type 'exit' to quit):")
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
}
