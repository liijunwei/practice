package main

import (
	"fmt"
	"strings"
	"time"
)

// TODO add tests

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
			depNames = append(depNames, dep.Name)
		}

		sb.WriteString(fmt.Sprintf("%s depends on %s\n", current.Name, depNames))
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

type WorkNodeOption func(*WorkNode)

// TODO Q: need to add mutex to WorkNode?
type WorkNode struct {
	Name     string
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
		Name:     name,
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
func (n *WorkNode) TaskReady(graph *DiGraph) bool {
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
		close(n.done)
		fmt.Printf("%s is done...\n", n.Name)

		return
	}

	fmt.Printf("%s is not running\n", n.Name)
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
			waitForDependencies(node, &g)
			run(node, &g)
		}()
	}

	// Q: am I doing this right?
	finals := findFinalNodes(&g)
	for _, node := range finals {
		<-node.done
	}
}

func findFinalNodes(g *DiGraph) []*WorkNode {
	var finals []*WorkNode

	for node := range g.Nodes() {
		if node.OutDegree(*g) == 0 {
			finals = append(finals, node)
			fmt.Println("we will wait for", node.Name)
		}
	}

	return finals
}

func waitForDependencies(n *WorkNode, g *DiGraph) {
	dependencies := (*g)[n]

	for _, node := range dependencies {
		<-node.done
	}

	fmt.Println("dependencies for", n.Name, "are met")
}

func run(n *WorkNode, g *DiGraph) {
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
