package main

import (
	"fmt"
	"strings"
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

type Graph map[*WorkNode]([]*WorkNode)

func (g Graph) Transpose() Graph {
	transposeGraph := Graph{}

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

func (g Graph) String() string {
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

func (g Graph) Nodes() map[*WorkNode]struct{} {
	nodes := make(map[*WorkNode]struct{})

	for current, dependencies := range g {
		nodes[current] = struct{}{}

		for _, dep := range dependencies {
			nodes[dep] = struct{}{}
		}
	}

	return nodes
}

// TODO Q: need to add mutex to WorkNode?
type WorkNode struct {
	Name     string
	Status   string
	Duration int64
	done     chan struct{}
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
		close(n.done)
		fmt.Printf("%s is done...\n", n.Name)

		return
	}

	fmt.Printf("%s is not running\n", n.Name)
}

func (n *WorkNode) IsDone() bool {
	return n.Status == statusDone
}

func (n *WorkNode) InDegree(g Graph) int {
	if deps, ok := g[n]; ok {
		return len(deps)
	}

	return 0
}

func (n *WorkNode) OutDegree(g Graph) int {
	return n.InDegree(g.Transpose())
}

func main() {
	defer timer("all tasks")()

	t1 := &WorkNode{
		Name:     "task1",
		Status:   statusHold,
		Duration: 1,
		done:     make(chan struct{}),
	}

	t2 := &WorkNode{
		Name:     "task2",
		Status:   statusHold,
		Duration: 2,
		done:     make(chan struct{}),
	}

	t3 := &WorkNode{
		Name:     "task3",
		Status:   statusHold,
		Duration: 1,
		done:     make(chan struct{}),
	}

	t4 := &WorkNode{
		Name:     "task4",
		Status:   statusHold,
		Duration: 1,
		done:     make(chan struct{}),
	}

	var taskGraph = Graph{
		t2: {t1},
		t3: {t1},
		t4: {t3},
	}

	g2 := taskGraph.Transpose()
	fmt.Println(taskGraph)
	fmt.Println(g2)

	fmt.Println("in degree of", t1.Name, t1.InDegree(taskGraph))
	fmt.Println("in degree of", t2.Name, t2.InDegree(taskGraph))
	fmt.Println("in degree of", t3.Name, t3.InDegree(taskGraph))
	fmt.Println("in degree of", t4.Name, t4.InDegree(taskGraph))

	fmt.Println("out degree of", t1.Name, t1.OutDegree(taskGraph))
	fmt.Println("out degree of", t2.Name, t2.OutDegree(taskGraph))
	fmt.Println("out degree of", t3.Name, t3.OutDegree(taskGraph))
	fmt.Println("out degree of", t4.Name, t4.OutDegree(taskGraph))
	fmt.Println()
	for n := range taskGraph.Nodes() {
		fmt.Println("all notes", n.Name)
	}

	for {
		// TODO make this non-blocking
		// TODO make trigger task running on `task_ready` event

		go func() {
			waitForDependencies(t1, &taskGraph)
			run(t1, &taskGraph)
		}()

		go func() {
			waitForDependencies(t2, &taskGraph)
			run(t2, &taskGraph)
		}()

		go func() {
			waitForDependencies(t3, &taskGraph)
			run(t3, &taskGraph)
		}()

		go func() {
			waitForDependencies(t4, &taskGraph)
			run(t4, &taskGraph)
		}()

		// Q: am I doing this right?
		finals := findFinalNodes(&taskGraph)
		for _, node := range finals {
			<-node.done
		}

		// TODO dfs iterate through all notes
		if t1.IsDone() || t2.IsDone() || t3.IsDone() || t4.IsDone() {
			fmt.Println("all work done")
			break
		}
	}
}

func findFinalNodes(g *Graph) []*WorkNode {
	var finals []*WorkNode

	for node := range g.Nodes() {
		if node.OutDegree(*g) == 0 {
			finals = append(finals, node)
			fmt.Println("we will wait for", node.Name)
		}
	}

	return finals
}

func waitForDependencies(n *WorkNode, g *Graph) {
	dependencies := (*g)[n]

	for _, node := range dependencies {
		<-node.done
	}

	fmt.Println("dependencies for", n.Name, "are met")
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
