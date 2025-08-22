package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"testing/quick"

	"github.com/stretchr/testify/require"
)

func TestFoo(t *testing.T) {
	t.Skip()
	_ = make([]int, 5911001488492326024)
}

// TODO add more property tests

func TestHeapProperty1(t *testing.T) {
	skipCount := 0
	passCount := 0

	checkHeapSize := func(size int) bool {
		if size <= 0 || size >= 10000000 {
			skipCount += 1
			return true
		}
		passCount += 1
		heap := NewMaxHeap(size)
		for range size {
			heap.Insert(rand.Intn(100))
		}
		return heap.Size() == size
	}

	err := quick.Check(checkHeapSize, &quick.Config{
		MaxCount: 100000,
		Values:   valueGen,
	})

	fmt.Println("skip ratio:", float64(skipCount)/(float64(passCount)+float64(skipCount)))
	require.NoError(t, err)
}

func valueGen(v []reflect.Value, r *rand.Rand) {
	// 1. 验证参数数量（测试函数应只有 1 个参数）
	if len(v) != 1 {
		panic("测试函数必须只有 1 个参数")
	}

	// 2. 定义生成范围
	minSize := 1
	maxSize := 10000

	// 3. 初始化参数为 int 类型的 reflect.Value（避免 zero Value）
	// 注意：必须先设置类型，否则 SetInt 会失败
	v[0] = reflect.New(reflect.TypeOf(int(0))).Elem()

	// 4. 生成 [minSize, maxSize] 范围内的随机 int 值
	rangeSize := maxSize - minSize + 1        // 范围大小（包含边界）
	randomSize := minSize + r.Intn(rangeSize) // 生成随机值

	// 5. 将生成的值赋给参数
	v[0].SetInt(int64(randomSize))
}
