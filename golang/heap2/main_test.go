package main

import (
	"math"
	"testing"
)

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

		h := NewMaxHeap()

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
