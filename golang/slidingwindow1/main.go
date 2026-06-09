package main

import (
	"fmt"
	"strings"
	"time"
)

// TCP 滑动窗口的核心作用：流量控制（Flow Control）
//
// 接收方在 TCP 头部放入 "Window Size" 字段，告诉发送方自己还能收多少字节。
// 发送方维护一个发送窗口：
//   - 窗口左边界 = 已发送且已确认的最小字节序号
//   - 窗口右边界 = 左边界 + 接收方通告窗口
//   - 发送方只能发送窗口内的数据
//   - 收到 ACK 后左边界右移 → 窗口"滑动"

// ============================================================
// 发送窗口
// ============================================================

type SendWindow struct {
	data    []byte // 全部待发送数据
	winSize int    // 接收方通告的窗口大小

	sendBase int    // 窗口左边界：已发送但未确认的最小序号
	nextSeq  int    // 下一个要发送的序号
	acked    []bool // 记录每个字节是否已确认
}

func NewSendWindow(data []byte, winSize int) *SendWindow {
	return &SendWindow{
		data:    data,
		winSize: winSize,
		acked:   make([]bool, len(data)),
	}
}

// 窗口右边界
func (sw *SendWindow) end() int {
	e := sw.sendBase + sw.winSize
	if e > len(sw.data) {
		e = len(sw.data)
	}
	return e
}

// 飞行中的字节数（已发送未确认）
func (sw *SendWindow) InFlight() int {
	count := 0
	for i := sw.sendBase; i < sw.nextSeq; i++ {
		if !sw.acked[i] {
			count++
		}
	}
	return count
}

// 可发送的序号列表
func (sw *SendWindow) Usable() []int {
	var seqs []int
	for i := sw.nextSeq; i < sw.end(); i++ {
		if !sw.acked[i] {
			seqs = append(seqs, i)
		}
	}
	return seqs
}

// 发送一个字节
func (sw *SendWindow) Send(seq int) (byte, bool) {
	if seq >= sw.end() || seq < sw.nextSeq || sw.acked[seq] {
		return 0, false
	}
	if seq >= sw.nextSeq {
		sw.nextSeq = seq + 1
	}
	return sw.data[seq], true
}

// 收到累积 ACK——确认 seq 及之前所有未确认字节
func (sw *SendWindow) RecvACK(seq int) {
	if seq < 0 || seq >= len(sw.data) {
		return
	}
	for i := sw.sendBase; i <= seq; i++ {
		sw.acked[i] = true
	}
	// 滑动：左边界跳过已确认的连续前缀
	for sw.sendBase < len(sw.data) && sw.acked[sw.sendBase] {
		sw.sendBase++
	}
}

func (sw *SendWindow) Done() bool {
	return sw.sendBase == len(sw.data)
}

// 可视化
func (sw *SendWindow) String() string {
	var sb strings.Builder
	sb.WriteString("Seq:  ")
	for i := 0; i < len(sw.data); i++ {
		fmt.Fprintf(&sb, " %3d", i)
	}
	sb.WriteString("\nData: ")
	for i := 0; i < len(sw.data); i++ {
		fmt.Fprintf(&sb, "  %c ", rune(sw.data[i]))
	}
	sb.WriteString("\n      ")
	for i := 0; i < len(sw.data); i++ {
		switch {
		case sw.acked[i]:
			sb.WriteString(" ✓ ")
		case i >= sw.sendBase && i < sw.end():
			sb.WriteString(" ▣ ")
		default:
			sb.WriteString(" · ")
		}
	}
	fmt.Fprintf(&sb, "\n窗口:  [%d, %d)  |  飞行中: %d  |  已确认: %d/%d",
		sw.sendBase, sw.end(), sw.InFlight(), sw.sendBase, len(sw.data))
	return sb.String()
}

// ============================================================
// 模拟：接收方逐渐处理数据并返回 ACK
// ============================================================

func main() {
	fmt.Println("===== TCP 滑动窗口 — 流量控制 Demo =====")
	fmt.Println()

	data := []byte("HELLOWORLD") // 10 字节
	winSize := 4                 // 接收方通告窗口大小 = 4
	sw := NewSendWindow(data, winSize)

	fmt.Println("初始状态:")
	fmt.Printf("  发送数据: %q (%d 字节)\n", data, len(data))
	fmt.Printf("  接收方通告窗口: %d 字节\n", winSize)
	fmt.Println("  （接收方说\"我最多还能收 4 字节\"）")
	fmt.Println()
	fmt.Println(sw)
	fmt.Println()

	round := 0
	for !sw.Done() {
		round++
		fmt.Printf("══════ 第 %d 轮 ══════\n", round)

		// 1. 发送方把窗口内可发送的字节全发出去
		usable := sw.Usable()
		if len(usable) == 0 {
			fmt.Println("  ⛔ 窗口已满，无法发送新数据（流量控制生效）")
			fmt.Println("     必须等接收方 ACK 才能继续发")
			fmt.Println()
			time.Sleep(800 * time.Millisecond)
			continue
		}

		for _, seq := range usable {
			b, _ := sw.Send(seq)
			fmt.Printf("  → 发送 Seq=%d ('%c')\n", seq, b)
		}
		fmt.Println()
		fmt.Println(sw)
		fmt.Println()

		// 2. 模拟：RTT + 接收方处理
		time.Sleep(500 * time.Millisecond)

		// 3. 接收方返回累积 ACK——确认已收到的连续前缀
		//    这里假设接收方窗口足够，每个字节都成功收到
		ackSeq := usable[len(usable)-1]
		fmt.Printf("  ← 收到累积 ACK=%d\n", ackSeq)
		sw.RecvACK(ackSeq)
		fmt.Println()
		fmt.Println(sw)
		fmt.Println()

		// 4. 模拟接收方处理速度（偶尔慢一点）
		if round%3 == 0 {
			fmt.Println("  🐢 接收方处理变慢，暂时不消费更多数据...")
			time.Sleep(600 * time.Millisecond)
		}
	}

	fmt.Println("════════════════════════")
	fmt.Println("✅ 传输完成")
	fmt.Println()
	fmt.Println("核心要点:")
	fmt.Println("  · 窗口大小由接收方控制（它告诉发送方\"我只能收这么多\"）")
	fmt.Println("  · 发送方绝不超出窗口——这就是流量控制")
	fmt.Println("  · ACK 到达→窗口左边界右移→腾出空间→可以发更多数据")
	fmt.Println("  · 飞行中的字节 = 已发送未确认的，不能超过窗口大小")
}
