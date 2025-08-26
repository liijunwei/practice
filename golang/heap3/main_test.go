package main

import (
	"fmt"
	"math"
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

func TestLevelOrder(t *testing.T) {
	h := NewMaxHeap(10)

	// 测试空堆
	result := h.LevelOrder()
	require.Empty(t, result)

	// 插入一些元素
	values := []int{10, 5, 15, 3, 7, 12, 20}
	for _, v := range values {
		h.Insert(v)
	}

	// 获取层序遍历结果
	levelOrder := h.LevelOrder()

	// 验证结果长度
	require.Equal(t, len(values), len(levelOrder))

	// 验证堆性质：父节点应该大于等于子节点
	for i := 1; i <= len(levelOrder)/2; i++ {
		parent := levelOrder[i-1]

		// 检查左子节点
		leftChild := 2 * i
		if leftChild <= len(levelOrder) {
			require.GreaterOrEqual(t, parent, levelOrder[leftChild-1],
				"父节点 %d 应该大于等于左子节点 %d", parent, levelOrder[leftChild-1])
		}

		// 检查右子节点
		rightChild := 2*i + 1
		if rightChild <= len(levelOrder) {
			require.GreaterOrEqual(t, parent, levelOrder[rightChild-1],
				"父节点 %d 应该大于等于右子节点 %d", parent, levelOrder[rightChild-1])
		}
	}

	// 验证根节点是最大值
	maxVal := levelOrder[0]
	for _, v := range levelOrder {
		require.LessOrEqual(t, v, maxVal, "根节点应该是最大值")
	}
}

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
		//dumpgraph(heap, fmt.Sprintf("/tmp/heap%d.dot", size))
		return heap.Size() == size
	}

	err := quick.Check(checkHeapSize, &quick.Config{
		MaxCount: 10_0000,
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

// go test -fuzz=FuzzHeap .
// go test -fuzz=FuzzHeap -fuzztime=30s .
// go test -fuzz=FuzzHeap -fuzztime=10000000x .
func FuzzHeap(f *testing.F) {
	// 添加一些种子测试用例
	f.Add([]byte{1, 2, 3, 4, 5})
	f.Add([]byte{100, 50, 75, 25, 90})
	f.Add([]byte{})

	f.Fuzz(func(t *testing.T, data []byte) {
		if len(data) == 0 {
			return
		}

		h := NewMaxHeap(len(data))

		// 将字节数据转换为整数并插入堆
		for _, b := range data {
			h.Insert(int(b))
		}

		// 验证删除时是否始终弹出最大值
		prev := math.MaxInt32
		for !h.Empty() {
			current := h.DelMax()
			if current > prev {
				t.Errorf("堆序错误：%d > %d", current, prev)
			}
			prev = current
		}
	})
}
