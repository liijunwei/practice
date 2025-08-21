package main

import (
	"encoding/json"
	"fmt"
	"strings"
	"unicode"
)

// TODO try
// https://github.com/yanyiwu/gojieba
// https://github.com/huichen/sego
// https://github.com/go-ego/gse
// https://github.com/go-ego/riot

// 倒排索引结构：键为单词，值为包含该单词的文档ID列表
type InvertedIndex map[string][]int

// 文档结构体：包含ID和内容
type Document struct {
	ID      int
	Content string
}

// 初始化一个新的倒排索引
func NewInvertedIndex() InvertedIndex {
	return make(InvertedIndex)
}

// 向索引中添加文档
func (idx InvertedIndex) AddDocument(doc Document) {
	// 提取文档中的所有单词（分词处理）
	words := tokenize(doc.Content)

	// 记录单词在当前文档中是否已添加，避免重复记录
	added := make(map[string]bool)

	// 为每个单词更新索引
	for _, word := range words {
		// 转为小写，使检索不区分大小写
		word = strings.ToLower(word)

		// 避免同一文档中同一单词被多次添加
		if !added[word] {
			idx[word] = append(idx[word], doc.ID)
			added[word] = true
		}
	}
}

// 检索包含指定单词的文档ID列表
func (idx InvertedIndex) Search(word string) []int {
	// 转为小写以匹配索引中的单词
	word = strings.ToLower(word)
	return idx[word]
}

// 分词处理：将文本拆分为单词，并过滤非字母字符
func tokenize(text string) []string {
	// 使用函数式编程将文本转换为符文切片，过滤非字母字符
	f := func(r rune) rune {
		if unicode.IsLetter(r) {
			return r
		}
		return ' ' // 非字母字符替换为空格
	}
	// 替换非字母字符为空格
	cleaned := strings.Map(f, text)

	// 按空格分割为单词
	words := strings.Fields(cleaned)

	return words
}

// 创建文档集合
var documents = []Document{
	{
		ID:      1,
		Content: "Go is a programming language developed at Google.",
	},
	{
		ID:      2,
		Content: "Golang is another name for Go. It's simple and efficient.",
	},
	{
		ID:      3,
		Content: "Programming in Go is fun. Google created Go in 2009.",
	},
	{
		ID:      4,
		Content: "Python is a popular programming language too, but different from Go.",
	},
}

func main() {
	// 创建并构建倒排索引
	index := NewInvertedIndex()
	for _, doc := range documents {
		index.AddDocument(doc)
	}

	// 测试检索功能
	testWords := []string{"go", "google", "programming", "python", "language", "hiiii", "simple and"}
	for _, word := range testWords {
		docs := index.Search(word)
		fmt.Printf("包含单词 '%s' 的文档ID: %v\n", word, dumpjson(findDocs(docs)))
	}
}

func findDocs(ids []int) []Document {
	res := make([]Document, 0, len(ids))
	for _, id := range ids {
		if r := findDoc(id); r.ID > 0 {
			res = append(res, r)
		}
	}
	return res
}

func findDoc(id int) Document {
	for _, doc := range documents {
		if doc.ID == id {
			return doc
		}
	}
	return Document{}
}

func dumpjson(o any) string {
	data, _ := json.Marshal(o)
	return string(data)
}
