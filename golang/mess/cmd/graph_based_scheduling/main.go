package main

import (
	"fmt"
	"strings"
	"time"
)

// TODO add tests
// TODO what's the relationship between this and toposort?
// Q: how to see the result of single thread execution?

const (
	statusHold    = "hold" // initial state
	statusReady   = "ready"
	statusRunning = "running"
	statusDone    = "done"
)

// state machine:
//     from hold    to ready    : TaskReady
//     from ready   to running  : TaskStart
//     from running to done     : TaskDone

// TODO cycle detection
type DiGraph map[*WorkNode]([]*WorkNode)

func (g DiGraph) Transpose() DiGraph {
	transposeGraph := DiGraph{}

	for current, dependencies := range g {
		for _, dep := range dependencies {
			if _, ok := transposeGraph[dep]; ok {
				transposeGraph[dep] = append(transposeGraph[dep], current)
				continue
			}

			transposeGraph[dep] = []*WorkNode{current}
		}
	}

	return transposeGraph
}

func (g DiGraph) String() string {
	var sb strings.Builder

	for current, dependencies := range g {
		depNames := make([]string, len(dependencies))
		for _, dep := range dependencies {
			depNames = append(depNames, dep.ID)
		}

		sb.WriteString(fmt.Sprintf("%s depends on %s\n", current.ID, depNames))
	}

	return sb.String()
}

func (g DiGraph) Nodes() map[*WorkNode]struct{} {
	nodes := make(map[*WorkNode]struct{})

	for current, dependencies := range g {
		nodes[current] = struct{}{}

		for _, dep := range dependencies {
			nodes[dep] = struct{}{}
		}
	}

	return nodes
}

func (g DiGraph) FinalNodes() []*WorkNode {
	nodes := g.Nodes()
	finals := make([]*WorkNode, 0, len(nodes))

	for node := range nodes {
		if node.OutDegree(g) == 0 {
			finals = append(finals, node)
		}
	}

	names := make([]string, 0, len(finals))

	for _, node := range finals {
		names = append(names, node.ID)
	}

	fmt.Printf("waiting for final nodes: %s\n", names)

	return finals
}

type WorkNodeOption func(*WorkNode)

// TODO Q: need to add mutex to WorkNode?
type WorkNode struct {
	ID       string
	Status   string
	Duration int64
	done     chan struct{}
}

func Duration(d int64) WorkNodeOption {
	return func(node *WorkNode) {
		node.Duration = d
	}
}

func NewWorkNode(name string, options ...WorkNodeOption) *WorkNode {
	node := &WorkNode{
		ID:       name,
		Status:   statusHold,
		Duration: 1,
		done:     make(chan struct{}),
	}

	for _, o := range options {
		o(node)
	}

	return node
}

// check whether dependencies are all done
// mark to ready if yes
func (n *WorkNode) TaskReady(graph DiGraph) bool {
	dependencies := graph[n]

	for _, node := range dependencies {
		// fmt.Println("checking dependencies status:", node.Name, node.Status)
		if node.Status != statusDone {
			return false
		}
	}

	n.Status = statusReady
	fmt.Printf("%s is ready...\n", n.ID)

	return true
}

// start the task if the node is ready
func (n *WorkNode) TaskStart() {
	if n.Status == statusReady {
		n.Status = statusRunning
		fmt.Printf("%s is running...\n", n.ID)

		return
	}

	fmt.Printf("%s is not ready\n", n.ID)
}

func (n *WorkNode) DoWork() {
	took := time.Duration(n.Duration) * time.Second
	time.Sleep(took) // work took x time
}

// do work if it's already running
func (n *WorkNode) TaskDone() {
	if n.Status == statusRunning {
		n.DoWork() // work consumes time
		n.Status = statusDone

		close(n.done)

		return
	}

	fmt.Printf("%s is not running\n", n.ID)
}

func (n *WorkNode) IsDone() bool {
	return n.Status == statusDone
}

func (n *WorkNode) InDegree(g DiGraph) int {
	if deps, ok := g[n]; ok {
		return len(deps)
	}

	return 0
}

func (n *WorkNode) OutDegree(g DiGraph) int {
	return n.InDegree(g.Transpose())
}

// unblock means node dependencies are all ready
func (n *WorkNode) WaitForDependencies(g DiGraph) {
	dependencies := g[n]

	for _, node := range dependencies {
		<-node.done
	}
}

func (n *WorkNode) Run(g DiGraph) {
	defer timer(n.ID)()

	if n.TaskReady(g) {
		n.TaskStart()
		n.TaskDone()
	}
}

func main() {
	defer timer("all tasks")()

	t1 := NewWorkNode("task1")
	t2 := NewWorkNode("task2", Duration(2))
	t3 := NewWorkNode("task3")
	t4 := NewWorkNode("task4")
	t5 := NewWorkNode("task5", Duration(2))

	var g = DiGraph{
		t2: {t1, t5},
		t3: {t1},
		t4: {t3},
	}

	// TODO make trigger task running on `task_ready` event
	for node := range g.Nodes() {
		node := node

		go func() {
			node.WaitForDependencies(g)
			node.Run(g)
		}()
	}

	for _, node := range g.FinalNodes() {
		<-node.done
	}
}

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s is done...(took %v)\n", name, time.Since(start))
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
