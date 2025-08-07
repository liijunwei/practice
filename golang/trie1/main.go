package main

import (
	"fmt"
	"strconv"
)

func main() {
	t1 := NewTrie()
	fmt.Printf("t1: %v\n", t1)
	fmt.Printf("t1.root: %v\n", t1.root)
	t1.Insert("hello")
	t1.Insert("hell")
	t1.Insert("hell")
	fmt.Println("search hello", t1.Search("hello"))
	fmt.Println("search hell", t1.Search("hell"))
	fmt.Println("search foo", t1.Search("foo"))
	fmt.Printf("t1.root: %v\n", t1.root)
}

const AlphabetSize = 26

type Node struct {
	children [AlphabetSize]*Node
	isEnd    bool
}

func NewNode() *Node {
	return new(Node)
}

type Trie struct {
	root *Node
}

func NewTrie() *Trie {
	return &Trie{
		root: new(Node),
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, c := range word {
		idx := c - 'a'
		if curr.children[idx] == nil {
			fmt.Println("inserting", strconv.QuoteRune(c))
			curr.children[idx] = NewNode()
		}
		curr = curr.children[idx]
	}
	if !curr.isEnd {
		curr.isEnd = true
		fmt.Println("marking end for", word)
	}
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, c := range word {
		idx := c - 'a'
		if curr.children[idx] == nil {
			return false
		}
		curr = curr.children[idx]
	}
	return curr.isEnd
}
