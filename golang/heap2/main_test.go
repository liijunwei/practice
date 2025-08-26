package main

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

func TestHeapWithRandomData(t *testing.T) {
	h := NewMaxHeap()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	data := make([]int, 1000000)
	for i := 0; i < 1000000; i++ {
		data[i] = r.Intn(10000)
		h.Insert(data[i])
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
}
