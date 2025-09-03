package main

import (
	"math/rand"
	"testing"
	"time"
)

// 定义测试规模：1亿数据
const maxNumber = 100_000_000

// 生成测试数据：随机生成1000个在范围内的数字
func generateTestNumbers(count int) []int {
	rand.Seed(time.Now().UnixNano())
	nums := make([]int, count)
	for i := 0; i < count; i++ {
		nums[i] = rand.Intn(maxNumber) + 1 // 确保数字在1到maxNumber之间
	}
	return nums
}

// 基准测试：bitmap的内存使用情况
func BenchmarkBitmapMemory(b *testing.B) {
	// 只执行一次初始化，因为我们主要关注内存使用
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 创建能容纳1亿数据的bitmap
		m := newBitmap(int64(maxNumber))

		// 随机设置一些数据（模拟实际使用场景）
		testNums := generateTestNumbers(10000)
		for _, n := range testNums {
			m.Set(n)
		}

		// 执行一些存在性检查操作
		for _, n := range testNums {
			_ = m.Exists(n)
		}
		// <-time.NewTimer(30 * time.Second).C
	}
}

// 基准测试：bitmap的Set操作性能
func BenchmarkBitmapSet(b *testing.B) {
	m := newBitmap(int64(maxNumber))
	testNums := generateTestNumbers(b.N)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m.Set(testNums[i])
	}
}

// 基准测试：bitmap的Exists操作性能
func BenchmarkBitmapExists(b *testing.B) {
	m := newBitmap(int64(maxNumber))
	testNums := generateTestNumbers(b.N)

	// 预先设置数据
	for _, n := range testNums {
		m.Set(n)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = m.Exists(testNums[i])
	}
}

// 对比测试：使用map的内存使用情况
func BenchmarkMapMemory(b *testing.B) {
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 创建map存储数据
		set := make(map[int]bool, 10000)

		// 随机设置一些数据
		testNums := generateTestNumbers(10000)
		for _, n := range testNums {
			set[n] = true
		}

		// 执行一些存在性检查操作
		for _, n := range testNums {
			_ = set[n]
		}
	}
}

// 对比测试：使用切片的内存使用情况（只适用于较小范围，此处仅作对比）
func BenchmarkSliceMemory(b *testing.B) {
	// 注意：对于1亿数据，这个测试会占用大量内存，可能导致OOM
	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		// 创建切片存储数据存在性
		// 警告：这将分配约100MB内存（1亿个bool）
		set := make([]bool, maxNumber+1)

		// 随机设置一些数据
		testNums := generateTestNumbers(10000)
		for _, n := range testNums {
			set[n] = true
		}

		// 执行一些存在性检查操作
		for _, n := range testNums {
			_ = set[n]
		}
	}
}
