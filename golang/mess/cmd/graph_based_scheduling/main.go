package main

import (
	"fmt"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)

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

// TODO add tests
// TODO what's the relationship between this and toposort?
// Q: how to see the result of single thread execution?
// Q: how to make use of producer / consumer pattern, is it necessary/helpful?

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

func main() {
	defer timer("all tasks")()

	ta := NewWorkNode("task#A")
	tb := NewWorkNode("task#B")
	tc := NewWorkNode("task#C")
	td := NewWorkNode("task#D")
	te := NewWorkNode("task#E")
	tf := NewWorkNode("task#F")
	tg := NewWorkNode("task#G")
	th := NewWorkNode("task#H")
	ti := NewWorkNode("task#I")
	tj := NewWorkNode("task#J")
	tk := NewWorkNode("task#K")
	tl := NewWorkNode("task#L")
	tm := NewWorkNode("task#M")

	graph := DiGraph{
		ta: {tc, td},
		tb: {tc, td},
		tc: {tg, th},
		td: {te, tf},
		te: {th},
		tf: {ti},
		tg: {tj},
		th: {tj},
		ti: {tj},
		tj: {tk},
		tj: {tl},
		tj: {tm},
		tk: {},
		tl: {},
		tm: {},
	}

	reverseGrapth := graph.Transpose()

	for node := range graph.Nodes() {
		node := node

		go func() {
			node.Run(reverseGrapth)
		}()
	}

	for _, node := range graph.FinalNodes() {
		<-node.done
	}
}

// TODO cycle detection
type DiGraph map[*WorkNode]([]*WorkNode)

func (g DiGraph) Transpose() DiGraph {
	transposeGraph := DiGraph{}

	for from, downstreams := range g {
		for _, ds := range downstreams {
			if _, ok := transposeGraph[ds]; ok {
				transposeGraph[ds] = append(transposeGraph[ds], from)
				continue
			}

			transposeGraph[ds] = []*WorkNode{from}
		}
	}

	return transposeGraph
}

func (g DiGraph) String() string {
	var sb strings.Builder

	for from, downstreams := range g {
		dsIDs := make([]string, 0, len(downstreams))
		for _, dep := range downstreams {
			dsIDs = append(dsIDs, dep.ID)
		}

		sb.WriteString(fmt.Sprintf("from %s to %s\n", from.ID, dsIDs))
	}

	return sb.String()
}

func (g DiGraph) Nodes() map[*WorkNode]struct{} {
	nodes := make(map[*WorkNode]struct{})

	for from, downstreams := range g {
		nodes[from] = struct{}{}

		for _, dep := range downstreams {
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

	cond sync.Cond
}

func Duration(d int64) WorkNodeOption {
	return func(node *WorkNode) {
		node.Duration = d
	}
}

func NewWorkNode(name string, options ...WorkNodeOption) *WorkNode {
	mu := &sync.Mutex{}

	node := &WorkNode{
		ID:       name,
		Status:   statusHold,
		Duration: 1,
		done:     make(chan struct{}),
		cond:     *sync.NewCond(mu),
	}

	for _, o := range options {
		o(node)
	}

	return node
}

// TODO need to synchronize
// check whether dependencies are all done
// mark to ready if yes
func (n *WorkNode) TaskReady(dependencies []*WorkNode) bool {
	n.cond.L.Lock()
	defer n.cond.L.Unlock()

	for n.PendingDependenciesCountNotZero(dependencies) {
		n.cond.Wait()
	}

	n.Status = statusReady
	fmt.Printf("%s is ready...\n", n.ID)

	return true
}

func (n *WorkNode) PendingDependenciesCountNotZero(dependencies []*WorkNode) bool {
	var count uint32

	for _, node := range dependencies {
		if node.Status != statusDone {
			atomic.AddUint32(&count, 1)
		}
	}

	return count != 0
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
	return n.OutDegree(g.Transpose())
}

func (n *WorkNode) OutDegree(g DiGraph) int {
	if downstreams, ok := g[n]; ok {
		return len(downstreams)
	}

	return 0
}

func (n *WorkNode) Run(reverseGrapth DiGraph) {
	defer timer(n.ID)()

	dependencies := reverseGrapth[n]

	// unblock means node dependencies are all ready
	for _, node := range dependencies {
		<-node.done
	}

	if n.TaskReady(dependencies) {
		n.TaskStart()
		n.TaskDone()
	}
}

func timer(name string) func() {
	start := time.Now()

	return func() {
		fmt.Printf("%s is done...(took %v)\n", name, time.Since(start))
	}
}
