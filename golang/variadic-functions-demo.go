// 可变参数函数
package main

import "fmt"

func sum(desc string, nums ...int) {
	fmt.Print(desc, " ", nums, " ")
	total := 0

	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {
	sum("12", 1, 2)
	sum("123", 1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum("1234", nums...)
}
