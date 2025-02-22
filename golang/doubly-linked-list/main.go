package main

import (
	"container/list"
	"encoding/json"
	"fmt"
)

func main() {
	l1 := list.New()
	l1.PushFront(1)
	l1.PushFront(2)
	l1.PushFront(3)
	print(l1)

	l2 := list.New()
	l2.PushBack(1)
	l2.PushBack(2)
	l2.PushBack(3)
	print(l2)

	// e := ll.Front()
	// ll.MoveToBack(e)
	// fmt.Println("Back", ll.Back())

	// print(ll)
}

func print(ll *list.List) {
	result := make([]any, 0)

	for e := ll.Front(); e != nil; e = e.Next() {
		result = append(result, e.Value)
	}

	info := make(map[string]any)
	info["list"] = result
	info["length"] = ll.Len()
	info["front"] = ll.Front().Value
	info["back"] = ll.Back().Value

	data, err := json.Marshal(info)
	boom(err)

	fmt.Println(string(data))
}

func boom(e error) {
	if e != nil {
		panic(e)
	}
}
