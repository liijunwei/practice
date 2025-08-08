package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	trie := NewTrie()

	trie.Insert("hello")
	trie.Insert("hell")
	trie.Insert("hi")
	trie.Insert("今天")
	trie.Insert("今天心情不错")
	trie.Insert("🤣")

	dotfilename := "/tmp/trie.dot"
	trie.DumpGraph(dotfilename)
	fmt.Println("dot file generated", dotfilename)

	svgfilename := "/tmp/trie.svg"
	err := exec.Command("dot", "-Tsvg", dotfilename, "-o", svgfilename).Run()
	boom(err)
	fmt.Println("svg file generated", svgfilename)
	err = exec.Command("open", svgfilename).Run()
	boom(err)
}

func boom(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	children map[rune]*Node
	isEnd    bool
	id       int
}

func NewNode(id int) *Node {
	return &Node{
		children: make(map[rune]*Node),
		isEnd:    false,
		id:       id,
	}
}

type Trie struct {
	root   *Node
	nextID int
}

func NewTrie() *Trie {
	return &Trie{
		root:   NewNode(0),
		nextID: 1,
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			curr.children[c] = NewNode(t.nextID)
			t.nextID++
		}
		curr = curr.children[c]
	}
	curr.isEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for _, c := range word {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return curr.isEnd
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.root
	for _, c := range prefix {
		if _, ok := curr.children[c]; !ok {
			return false
		}
		curr = curr.children[c]
	}
	return true
}

func (t *Trie) DumpGraph(filename string) {
	file, err := os.Create(filename)
	boom(err)
	defer file.Close()

	_, err = file.WriteString("digraph Trie {\n")
	boom(err)

	_, err = fmt.Fprintf(file, "    %d [label=\"root\", shape=circle];\n", t.root.id)
	boom(err)

	var traverse func(node *Node) error
	traverse = func(node *Node) error {
		for char, child := range node.children {
			shape := "circle"
			if child.isEnd {
				shape = "doublecircle"
			}

			_, err := fmt.Fprintf(file, "    %d [label=%s, shape=%s];\n", child.id, string(char), shape)
			boom(err)

			_, err = fmt.Fprintf(file, "    %d -> %d;\n", node.id, child.id)
			boom(err)

			err = traverse(child)
			boom(err)
		}
		return nil
	}

	err = traverse(t.root)
	boom(err)

	_, err = file.WriteString("}\n")
	boom(err)
}
