package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func buildTrie() *Trie {
	// 初始化字典树
	trie := NewTrie()

	trie.Insert("今天")
	trie.Insert("今天心情不错")
	trie.Insert("今天心情不错")
	trie.Insert("🤣☀️🌿🐦🌸🐝✨💧🌈🌼🦋")
	for _, word := range readSamples() {
		trie.Insert(word)
	}
	fmt.Println("trie is ready")

	// 生成可视化图
	// dotfilename := "/tmp/trie.dot"
	// trie.DumpGraph(dotfilename)
	// fmt.Println("dot file generated", dotfilename)

	// svgfilename := "/tmp/trie.svg"
	// err := exec.Command("dot", "-Tsvg", dotfilename, "-o", svgfilename).Run()
	// boom(err)
	// fmt.Println("svg file generated", svgfilename)
	return trie
}

func searchHandler(trie *Trie) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 只接受GET请求
		if r.Method != http.MethodGet {
			http.Error(w, "只支持GET请求", http.StatusMethodNotAllowed)
			return
		}

		// 设置CORS头，允许前端页面访问
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		// 获取prefix参数
		prefix := r.URL.Query().Get("prefix")
		if prefix == "" {
			http.Error(w, "请提供prefix参数", http.StatusBadRequest)
			return
		}

		// 查找匹配的字符串
		matches := trie.FindPrefix(prefix)

		// 返回JSON格式的结果
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]interface{}{
			"prefix":  prefix,
			"matches": matches,
			"count":   len(matches),
		})
	}
}

func main() {
	// 初始化字典树
	trie := buildTrie()

	// 注册处理函数
	http.HandleFunc("/search", searchHandler(trie))

	// 提供静态文件服务
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			http.ServeFile(w, r, "index.html")
		} else {
			http.NotFound(w, r)
		}
	})

	// 启动HTTP服务
	port := ":8080"
	fmt.Printf("服务已启动，监听端口%s\n", port)
	fmt.Printf("可以通过 http://localhost%s 访问前缀搜索页面\n", port)
	fmt.Printf("API接口: http://localhost%s/search?prefix=你的前缀\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
