package main

import (
	"testing"

	"github.com/leanovate/gopter"
	"github.com/leanovate/gopter/gen"
	"github.com/leanovate/gopter/prop"
	"github.com/stretchr/testify/assert"
)

// 基本测试
func TestNewBitmap(t *testing.T) {
	t.Run("创建指定大小的bitmap", func(t *testing.T) {
		max := int64(10)
		m := newBitmap(max)
		// 预期长度：10/8 + 1 = 2
		assert.Equal(t, 2, len(m.flags))
	})

	t.Run("空bitmap", func(t *testing.T) {
		m := newBitmap(0)
		assert.Equal(t, 1, len(m.flags)) // 至少有一个字节
	})
}

func TestBitmap_Set_Exists(t *testing.T) {
	m := newBitmap(20)

	t.Run("设置后存在", func(t *testing.T) {
		m.Set(5)
		assert.True(t, m.Exists(5))
	})

	t.Run("未设置则不存在", func(t *testing.T) {
		assert.False(t, m.Exists(6))
	})

	t.Run("设置多个值", func(t *testing.T) {
		m.Set(3)
		m.Set(10)
		m.Set(15)

		assert.True(t, m.Exists(3))
		assert.True(t, m.Exists(10))
		assert.True(t, m.Exists(15))
		assert.False(t, m.Exists(4))
	})

	t.Run("边界值测试", func(t *testing.T) {
		m := newBitmap(8)
		m.Set(8)
		assert.True(t, m.Exists(8))
	})
}

func TestBitmap_Del(t *testing.T) {
	m := newBitmap(20)
	m.Set(5)
	m.Set(10)

	t.Run("删除后不存在", func(t *testing.T) {
		m.Del(5)
		assert.False(t, m.Exists(5))
		assert.True(t, m.Exists(10)) // 其他值不受影响
	})

	t.Run("删除未设置的值", func(t *testing.T) {
		// 应该是安全的，没有副作用
		m.Del(6)
	})
}

func TestLocate(t *testing.T) {
	tests := []struct {
		name     string
		input    int
		expected []int // [arridx, offset]
	}{
		{"n=1", 1, []int{0, 1}},
		{"n=8", 8, []int{1, 0}},
		{"n=9", 9, []int{1, 1}},
		{"n=15", 15, []int{1, 7}},
		{"n=16", 16, []int{2, 0}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			arridx, offset := locate(tt.input)
			assert.Equal(t, tt.expected[0], arridx)
			assert.Equal(t, tt.expected[1], offset)
		})
	}
}

// 属性测试 (Property-based tests)
func TestBitmapProperties(t *testing.T) {
	properties := gopter.NewProperties(nil)

	// 属性1: 任何数字在Set之后Exists应该返回true
	properties.Property("Set后Exists返回true", prop.ForAll(
		func(n int) bool {
			// 生成有效的n值 (1到1000之间)
			if n <= 0 {
				n = -n + 1 // 确保n为正数
			}
			if n > 1000 {
				n = n%1000 + 1
			}

			m := newBitmap(int64(n + 1))
			m.Set(n)
			return m.Exists(n)
		},
		gen.Int(), // 生成任意整数作为输入
	))

	// 属性2: 任何数字在Set然后Del之后Exists应该返回false
	properties.Property("Set然后Del后Exists返回false", prop.ForAll(
		func(n int) bool {
			if n <= 0 {
				n = -n + 1
			}
			if n > 1000 {
				n = n%1000 + 1
			}

			m := newBitmap(int64(n + 1))
			m.Set(n)
			m.Del(n)
			return !m.Exists(n)
		},
		gen.Int(),
	))

	// 属性3: 对一个数字多次Set的效果与一次Set相同
	properties.Property("多次Set与一次Set效果相同", prop.ForAll(
		func(n, k int) bool {
			if n <= 0 {
				n = -n + 1
			}
			if n > 1000 {
				n = n%1000 + 1
			}
			// 确保k是正数
			if k <= 0 {
				k = -k + 1
			}
			if k > 100 {
				k = 100
			}

			m := newBitmap(int64(n + 1))
			// 多次Set
			for i := 0; i < k; i++ {
				m.Set(n)
			}
			return m.Exists(n)
		},
		gen.Int(), gen.Int(),
	))

	// 属性4: 对未设置的数字执行Del操作是安全的
	properties.Property("Del未设置的数字是安全的", prop.ForAll(
		func(n int) bool {
			if n <= 0 {
				n = -n + 1
			}
			if n > 1000 {
				n = n%1000 + 1
			}

			m := newBitmap(int64(n + 1))
			// 尝试删除未设置的值
			m.Del(n)
			// 检查是否仍然不存在
			return !m.Exists(n)
		},
		gen.Int(),
	))

	// 运行所有属性测试
	properties.TestingRun(t)
}

// 生成有效的数字（确保在1到max之间）
func validNumber(n int, max int) int {
	if max <= 0 {
		return 1 // 当max无效时，返回默认有效值
	}
	// 将n规范化到1到max之间
	n = ((n % max) + max) % max
	if n == 0 {
		return max
	}
	return n
}

func TestBitmapFixedProperties(t *testing.T) {
	properties := gopter.NewProperties(nil)

	// 属性1: Set后Exists返回true
	properties.Property("Set后Exists返回true", prop.ForAll(
		func(max int64, n int) bool {
			maxVal := int(max)
			if maxVal <= 0 {
				maxVal = 100
			}
			// 确保n是有效数字
			validN := validNumber(n, maxVal)

			m := newBitmap(max)
			m.Set(validN)
			return m.Exists(validN)
		},
		gen.Int64Range(1, 1000),
		gen.Int(),
	))

	// 属性2: Set然后Del后Exists返回false
	properties.Property("Set然后Del后Exists返回false", prop.ForAll(
		func(max int64, n int) bool {
			maxVal := int(max)
			if maxVal <= 0 {
				maxVal = 100
			}
			validN := validNumber(n, maxVal)

			m := newBitmap(max)
			m.Set(validN)
			m.Del(validN)
			return !m.Exists(validN)
		},
		gen.Int64Range(1, 1000),
		gen.Int(),
	))

	// 属性3: Set一个数字不影响其他数字
	properties.Property("Set一个数字不影响其他数字", prop.ForAll(
		func(max int64, n int, other int) bool {
			maxVal := int(max)
			if maxVal <= 0 {
				maxVal = 100
			}
			// 确保n和other都是有效且不同的数字
			validN := validNumber(n, maxVal)
			validOther := validNumber(other, maxVal)

			// 如果两个数字相同，手动调整other
			if validOther == validN {
				validOther = validNumber(other+1, maxVal)
			}

			m := newBitmap(max)
			initial := m.Exists(validOther)
			m.Set(validN)
			return m.Exists(validOther) == initial
		},
		gen.Int64Range(1, 1000),
		gen.Int(),
		gen.Int(),
	))

	// 属性4: 交替Set和Del应正确切换状态
	properties.Property("交替Set和Del应正确切换状态", prop.ForAll(
		func(max int64, n int, k int) bool {
			maxVal := int(max)
			if maxVal <= 0 {
				maxVal = 100
			}
			validN := validNumber(n, maxVal)

			// 确保k是合理的正数
			if k <= 0 {
				k = 5
			} else if k > 100 {
				k = 100
			}

			m := newBitmap(max)
			state := false

			for i := 0; i < k; i++ {
				if i%2 == 0 {
					m.Set(validN)
					state = true
				} else {
					m.Del(validN)
					state = false
				}
				if m.Exists(validN) != state {
					return false
				}
			}
			return true
		},
		gen.Int64Range(1, 1000),
		gen.Int(),
		gen.Int(),
	))

	// 运行所有测试
	properties.TestingRun(t)
}
