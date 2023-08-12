package main

import (
	"fmt"
)

const (
	statusHold     = "hold"
	statusReady    = "ready"
	statusRunning  = "running"
	statusFinished = "finished"
)

func main() {
	t1 := WorkNode{
		Name:   "task1",
		Status: statusHold,
	}

	t2 := WorkNode{
		Name:   "task2",
		Status: statusHold,
	}

	t3 := WorkNode{
		Name:   "task3",
		Status: statusHold,
	}

	t4 := WorkNode{
		Name:   "task4",
		Status: statusHold,
	}

	var taskGraph = map[WorkNode][]WorkNode{
		t2: {t1},
		t3: {t1},
		t4: {t3},
	}

	fmt.Println(taskGraph)
}

// 定义每个任务的输入输出
// 任务可以组合，根据任务的依赖关系生成图
// 对图做剪支优化，合并重复任务(optional for now)
// 对图进行拓扑排序(optional for now)
// 从没有输入依赖 或者 依赖已经满足的节点开始执行，每个任务执行结束检查是否能触发下游任务开始执行

// 主要思路：
// 框架(scheduler) 知道：
// 		图的样子
// 		每个任务的函数签名
// 		知道每个任务的当前状态

type WorkNode struct {
	Name   string
	Status string
}
