package main

import (
	"fmt"
	"time"
)

const (
	statusHold    = "hold"
	statusReady   = "ready"
	statusRunning = "running"
	statusDone    = "done"
)

// state machine:
//     from hold    to ready    : task_ready
//     from ready   to running  : task_start
//     from running to done     : task_done

type Graph map[*WorkNode][]*WorkNode

// TODO Q: need to add mutex to WorkNode?
type WorkNode struct {
	Name     string
	Status   string
	Duration int64
}

// check whether dependencies are all done
// mark to ready if yes
func (n *WorkNode) TaskReady(graph *Graph) bool {
	dependencies := (*graph)[n]

	for _, node := range dependencies {
		// fmt.Println("checking dependencies status:", node.Name, node.Status)
		if node.Status != statusDone {
			return false
		}
	}

	n.Status = statusReady
	fmt.Printf("%s is ready...\n", n.Name)

	return true
}

// start the task if the node is ready
func (n *WorkNode) TaskStart() {
	if n.Status == statusReady {
		n.Status = statusRunning
		fmt.Printf("%s is started...\n", n.Name)

		return
	}

	fmt.Printf("%s is not ready\n", n.Name)
}

// do work if it's already running
func (n *WorkNode) TaskDone() {
	if n.Status == statusRunning {
		time.Sleep(time.Duration(n.Duration) * time.Second) // work consumes time
		n.Status = statusDone
		fmt.Printf("%s is done...\n", n.Name)

		return
	}

	fmt.Printf("%s is not running\n", n.Name)
}

func (n *WorkNode) IsDone() bool {
	return n.Status == statusDone
}

func main() {
	defer timer("all tasks")()

	t1 := &WorkNode{
		Name:     "task1",
		Status:   statusHold,
		Duration: 1,
	}

	t2 := &WorkNode{
		Name:     "task2",
		Status:   statusHold,
		Duration: 2,
	}

	t3 := &WorkNode{
		Name:     "task3",
		Status:   statusHold,
		Duration: 1,
	}

	t4 := &WorkNode{
		Name:     "task4",
		Status:   statusHold,
		Duration: 1,
	}

	var taskGraph = Graph{
		t2: {t1},
		t3: {t1},
		t4: {t3},
	}

	for {
		// TODO make this non-blocking
		// TODO make trigger task running on `task_ready` event

		run(t1, &taskGraph)
		go run(t2, &taskGraph)
		run(t3, &taskGraph)
		run(t4, &taskGraph)

		// TODO dfs iterate through all notes
		if t1.IsDone() || t2.IsDone() || t3.IsDone() || t4.IsDone() {
			fmt.Println("all work done")
			break
		}
	}
}

func run(n *WorkNode, g *Graph) {
	defer timer(n.Name)()

	if n.TaskReady(g) {
		n.TaskStart()
		n.TaskDone()
	}
}

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
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
