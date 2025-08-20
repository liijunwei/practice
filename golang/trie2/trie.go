package main

import (
	"fmt"
	"os"
)

func boom(err error) {
	if err != nil {
		panic(err)
	}
}

type Node struct {
	Children map[rune]*Node
	IsEnd    bool
	ID       int
}

func NewNode(id int) *Node {
	return &Node{
		Children: make(map[rune]*Node),
		IsEnd:    false,
		ID:       id,
	}
}

type Trie struct {
	Root   *Node
	NextID int
}

func NewTrie() *Trie {
	return &Trie{
		Root:   NewNode(0),
		NextID: 1,
	}
}

func (t *Trie) Insert(word string) {
	curr := t.Root
	for _, c := range word {
		if _, ok := curr.Children[c]; !ok {
			curr.Children[c] = NewNode(t.NextID)
			t.NextID++
		}
		curr = curr.Children[c]
	}
	curr.IsEnd = true
}

func (t *Trie) Search(word string) bool {
	curr := t.Root
	for _, c := range word {
		if _, ok := curr.Children[c]; !ok {
			return false
		}
		curr = curr.Children[c]
	}
	return curr.IsEnd
}

// TODO
func (t *Trie) FindPrefix(prefix string) []string {
	result := make([]string, 0, len(samples))
	curr := t.Root

	// 先找到前缀的最后一个节点
	for _, c := range prefix {
		if _, ok := curr.Children[c]; !ok {
			return result // 前缀不存在，返回空结果
		}
		curr = curr.Children[c]
	}

	// 从前缀的最后一个节点开始，收集所有单词
	var dfs func(node *Node, path string)
	dfs = func(node *Node, path string) {
		if node.IsEnd {
			result = append(result, path)
		}

		for char, child := range node.Children {
			dfs(child, path+string(char))
		}
	}

	dfs(curr, prefix)
	return result
}

func (t *Trie) StartsWith(prefix string) bool {
	curr := t.Root
	for _, c := range prefix {
		if _, ok := curr.Children[c]; !ok {
			return false
		}
		curr = curr.Children[c]
	}
	return true
}

func (t *Trie) DumpGraph(filename string) {
	file, err := os.Create(filename)
	boom(err)
	defer file.Close()

	_, err = file.WriteString("digraph Trie {\n")
	boom(err)

	_, err = fmt.Fprintf(file, "    %d [label=\"root\", shape=circle];\n", t.Root.ID)
	boom(err)

	var traverse func(node *Node) error
	traverse = func(node *Node) error {
		for char, child := range node.Children {
			shape := "circle"
			if child.IsEnd {
				shape = "doublecircle"
			}

			label := fmt.Sprintf("%q", string(char)) // whitespace
			_, err := fmt.Fprintf(file, "    %d [label=%s, shape=%s];\n", child.ID, label, shape)
			boom(err)

			_, err = fmt.Fprintf(file, "    %d -> %d [label=%s];\n", node.ID, child.ID, label)
			boom(err)

			err = traverse(child)
			boom(err)
		}
		return nil
	}

	err = traverse(t.Root)
	boom(err)

	_, err = file.WriteString("}\n")
	boom(err)
}
