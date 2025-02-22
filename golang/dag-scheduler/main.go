package main

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"sync"
	"time"
)

// TODO detect critical path and print it

// Task 接口定义了任务的基本行为和数据
type Task interface {
	ID() string
	Execute() error
	Duration() time.Duration
	SetInput(key string, value interface{})
	GetOutput(key string) interface{}
}

// TaskContext 包含任务执行的上下文信息
type TaskContext struct {
	Inputs  map[string]interface{} // 从其他任务获取的输入
	Outputs map[string]interface{} // 当前任务的输出
}

// BaseTask 提供了Task接口的基本实现
type BaseTask struct {
	id       string
	duration time.Duration
	inputs   map[string]interface{}
	outputs  map[string]interface{}
	handler  func(inputs map[string]interface{}) (map[string]interface{}, error)
}

var _ Task = &BaseTask{} // 确保BaseTask实现了Task接口

func NewBaseTask(id string, duration time.Duration, handler func(map[string]interface{}) (map[string]interface{}, error)) *BaseTask {
	return &BaseTask{
		id:       id,
		duration: duration,
		inputs:   make(map[string]interface{}),
		outputs:  make(map[string]interface{}),
		handler:  handler,
	}
}

func (t *BaseTask) ID() string {
	return t.id
}

func (t *BaseTask) Execute() error {
	assert(t.inputs != nil, "task inputs map is nil", t.id)
	assert(t.outputs != nil, "task outputs map is nil", t.id)

	time.Sleep(t.duration)
	if t.handler != nil {
		outputs, err := t.handler(t.inputs)
		if err != nil {
			return err
		}
		assert(outputs != nil, "handler returned nil outputs map", t.id)
		t.outputs = outputs
	}
	return nil
}

func (t *BaseTask) Duration() time.Duration {
	return t.duration
}

func (t *BaseTask) SetInput(key string, value interface{}) {
	t.inputs[key] = value
}

func (t *BaseTask) GetOutput(key string) interface{} {
	return t.outputs[key]
}

// 添加监控相关的类型定义
type TaskStatus struct {
	ID        string
	Status    string
	StartTime time.Time
	Duration  time.Duration
}

type SchedulerStats struct {
	TotalTasks     int
	CompletedTasks int
	RunningTasks   []TaskStatus
	StartTime      time.Time
}

// 在scheduler结构中添加监控相关字段
type scheduler struct {
	sync.Mutex
	cond           *sync.Cond
	workQueue      []Task
	nodeStatuses   map[string]string
	graph          *fsmGraph
	indegrees      map[string]int
	completedCount int
	taskResults    map[string]Task
	// 添加监控相关字段
	startTime      time.Time
	taskStartTimes map[string]time.Time
	stopMonitor    chan struct{} // 用于停止监控的信号
}

func main() {
	_, f, _, _ := runtime.Caller(0)
	dir := filepath.Dir(f)
	filepath := filepath.Join(dir, fmt.Sprintf("./graph%s.txt", os.Args[1]))
	fmt.Println("Reading file:", filepath)

	g := newGraph(filepath)
	s := newScheduler(g)

	startTime := time.Now()

	var wg sync.WaitGroup
	numWorkers := 3
	for i := 0; i < numWorkers; i++ {
		w := newWorker(i, s)
		wg.Add(1)
		go w.run(&wg)
	}

	wg.Wait()
	close(s.stopMonitor) // 停止监控goroutine

	fmt.Println("spent:", time.Since(startTime))
}

// 修改newScheduler，初始化新增字段
func newScheduler(g *fsmGraph) *scheduler {
	assert(g != nil, "graph cannot be nil")
	assert(len(g.adj) > 0, "graph cannot be empty")

	// 确保图是无环的
	hasCycle, cyclePath := g.detectCycle()
	assert(!hasCycle, "graph must be acyclic", "cycle detected:", strings.Join(cyclePath, " -> "))

	s := &scheduler{
		graph:          g,
		workQueue:      make([]Task, 0),
		nodeStatuses:   make(map[string]string),
		indegrees:      make(map[string]int),
		taskResults:    make(map[string]Task),
		startTime:      time.Now(),
		taskStartTimes: make(map[string]time.Time),
		stopMonitor:    make(chan struct{}),
	}
	s.cond = sync.NewCond(&s.Mutex)

	vertices := g.vertices()
	assert(len(vertices) > 0, "graph must not be empty")

	// 验证图的完整性
	for _, node := range vertices {
		deps := g.getDependencies(node)
		for _, dep := range deps {
			assert(g.adj[dep] != nil, "invalid graph: missing dependency", dep)
		}
	}

	// 为每个顶点创建对应的任务
	for _, node := range vertices {
		task := NewBaseTask(node, time.Second, func(inputs map[string]interface{}) (map[string]interface{}, error) {
			// 示例：将所有输入连接起来作为输出
			var inputVals []string
			for _, input := range inputs {
				inputVals = append(inputVals, fmt.Sprint(input))
			}
			return map[string]interface{}{
				"result": fmt.Sprintf("task-%s-processed-%v", node, strings.Join(inputVals, "+")),
			}, nil
		})

		s.indegrees[node] = g.inDegree(node)
		if g.inDegree(node) == 0 {
			s.nodeStatuses[node] = "ready"
			s.workQueue = append(s.workQueue, task)
		} else {
			s.nodeStatuses[node] = "pending"
		}
	}

	assert(len(s.workQueue) > 0, "graph must have at least one entry point")

	// 启动监控goroutine
	go s.monitor()

	return s
}

// 添加监控相关方法
func (s *scheduler) monitor() {
	ticker := time.NewTicker(10 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-s.stopMonitor:
			return
		case <-ticker.C:
			stats := s.getStats()
			s.printStats(stats)
		}
	}
}

func (s *scheduler) getStats() SchedulerStats {
	s.Lock()
	defer s.Unlock()

	stats := SchedulerStats{
		TotalTasks:     len(s.nodeStatuses),
		CompletedTasks: s.completedCount,
		StartTime:      s.startTime,
		RunningTasks:   make([]TaskStatus, 0),
	}

	// 收集正在运行的任务信息
	for id, status := range s.nodeStatuses {
		if status == "running" {
			if startTime, ok := s.taskStartTimes[id]; ok {
				stats.RunningTasks = append(stats.RunningTasks, TaskStatus{
					ID:        id,
					Status:    status,
					StartTime: startTime,
					Duration:  time.Since(startTime),
				})
			}
		}
	}

	return stats
}

func (s *scheduler) printStats(stats SchedulerStats) {
	// 清除从光标到屏幕底部的内容
	fmt.Print("\033[2J")
	// 将光标移动到屏幕顶部
	fmt.Print("\033[H")

	// 打印统计信息
	fmt.Printf("=== Scheduler Stats ===\n")
	fmt.Printf("Total Tasks: %d\n", stats.TotalTasks)
	fmt.Printf("Completed: %d (%.1f%%)\n",
		stats.CompletedTasks,
		float64(stats.CompletedTasks)/float64(stats.TotalTasks)*100)
	fmt.Printf("Running Time: %v\n", time.Since(stats.StartTime))

	if len(stats.RunningTasks) > 0 {
		fmt.Printf("\nRunning Tasks:\n")
		for _, task := range stats.RunningTasks {
			fmt.Printf("- %s (running for %v)\n", task.ID, task.Duration)
		}
	}

	// 保存光标位置
	fmt.Print("\033[s")
	fmt.Println()
}

// 修改getTask方法，记录任务开始时间
func (s *scheduler) getTask() (Task, bool) {
	s.Lock()
	defer s.Unlock()

	for len(s.workQueue) == 0 && s.completedCount < len(s.nodeStatuses) {
		s.cond.Wait()
	}

	if s.completedCount >= len(s.nodeStatuses) {
		return nil, false
	}

	assert(len(s.workQueue) > 0, "work queue is empty but completedCount < total")
	task := s.workQueue[0]
	assert(task != nil, "nil task in work queue")

	s.workQueue = s.workQueue[1:]
	taskID := task.ID()

	// 验证任务状态转换
	assert(s.nodeStatuses[taskID] == "ready", "task must be ready before running", taskID)
	assert(s.indegrees[taskID] == 0, "running task must have zero indegree", taskID)
	s.nodeStatuses[taskID] = "running"

	// 验证依赖任务的完成状态
	for _, dep := range s.graph.getDependencies(taskID) {
		assert(s.nodeStatuses[dep] == "done", "dependency not completed", dep, "->", taskID)
		depTask, ok := s.taskResults[dep]
		assert(ok, "dependency task result not found", dep, "->", taskID)
		assert(depTask != nil, "dependency task is nil", dep, "->", taskID)

		result := depTask.GetOutput("result")
		assert(result != nil, "dependency task has no output", dep, "->", taskID)
		task.SetInput(fmt.Sprintf("%s.result", dep), result)
	}

	s.taskResults[taskID] = task
	s.taskStartTimes[taskID] = time.Now()
	return task, true
}

// 修改completeTask方法，清理任务时间记录
func (s *scheduler) completeTask(task Task) {
	s.Lock()
	defer s.Unlock()

	assert(task != nil, "cannot complete nil task")
	taskID := task.ID()

	// 验证任务状态
	assert(s.nodeStatuses[taskID] == "running", "can only complete running tasks", taskID)
	assert(task.GetOutput("result") != nil, "completed task must have output", taskID)

	s.nodeStatuses[taskID] = "done"
	s.completedCount++
	assert(s.completedCount <= len(s.nodeStatuses), "completed count exceeds total tasks")

	// 检查后继节点
	for _, next := range s.graph.adj[taskID] {
		s.indegrees[next]--
		assert(s.indegrees[next] >= 0, "indegree must not be negative", next)

		if s.indegrees[next] == 0 {
			assert(s.nodeStatuses[next] == "pending",
				"task with zero indegree must be pending", next)
			s.nodeStatuses[next] = "ready"

			// 创建新任务
			nextTask := NewBaseTask(next, time.Second, func(inputs map[string]interface{}) (map[string]interface{}, error) {
				assert(len(inputs) > 0, "task must have inputs from dependencies", next)
				// 示例处理函数
				var inputVals []string
				for _, input := range inputs {
					inputVals = append(inputVals, fmt.Sprint(input))
				}
				return map[string]interface{}{
					"result": fmt.Sprintf("task-%s-processed-%v", next, strings.Join(inputVals, "+")),
				}, nil
			})
			s.workQueue = append(s.workQueue, nextTask)
		}
	}

	delete(s.taskStartTimes, taskID)
	s.cond.Broadcast()
}

type worker struct {
	id int
	s  *scheduler
}

func newWorker(id int, s *scheduler) *worker {
	return &worker{
		id: id,
		s:  s,
	}
}

func (w *worker) run(wg *sync.WaitGroup) {
	defer wg.Done()

	for {
		task, ok := w.s.getTask()
		if !ok {
			return
		}

		assert(task != nil, "worker received nil task")
		err := task.Execute()
		if err != nil {
			fmt.Printf("Worker %d failed task %s: %v\n", w.id, task.ID(), err)
			continue
		}

		assert(task.GetOutput("result") != nil,
			"task execution produced no output", task.ID())
		w.s.completeTask(task)
	}
}

// fsmGraph is a directed graph of a finite state machine
type fsmGraph struct {
	name string
	adj  map[string][]string // map[from] -> tos
}

func newGraph(filepath string) *fsmGraph {
	input, err := os.Open(filepath)
	boom(err, "failed to open graph file", filepath)

	defer input.Close()

	data, err := io.ReadAll(input)
	boom(err, "failed to read graph file", filepath)

	graph := &fsmGraph{
		name: filepath,
		adj:  make(map[string][]string),
	}

	for _, line := range strings.Split(string(data), "\n") {
		if line == "" {
			continue
		}

		tokens := strings.Split(line, " ")
		assert(len(tokens) == 2, "invalid graph edge format", line)
		from, to := tokens[0], tokens[1]
		graph.addEdge(from, to)
	}

	// 检测环
	if hasCycle, cyclePath := graph.detectCycle(); hasCycle {
		// spew.Dump(cyclePath)
		fmt.Println("Error: Cycle detected in the graph!")
		fmt.Printf("Cycle path: %s\n", strings.Join(cyclePath, " -> "))
		fmt.Println("The graph must be acyclic (DAG) for the scheduler to work.")
		os.Exit(1)
	}

	return graph
}

func (g *fsmGraph) addEdge(from, to string) {
	g.adj[from] = append(g.adj[from], to)
}

func (g *fsmGraph) vertices() []string {
	states := make([]string, 0, len(g.adj))
	seen := make(map[string]bool)

	var visitAll func(items []string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(g.adj[item])
				states = append(states, item)
			}
		}
	}

	keys := make([]string, 0, len(g.adj))
	for key := range g.adj {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	visitAll(keys)

	return states
}

func (g *fsmGraph) inDegree(state string) int {
	count := 0
	for _, tos := range g.adj {
		for _, to := range tos {
			if to == state {
				count++
			}
		}
	}
	return count
}

func (g *fsmGraph) getDependencies(node string) []string {
	deps := []string{}
	for from, tos := range g.adj {
		for _, to := range tos {
			if to == node {
				deps = append(deps, from)
			}
		}
	}
	return deps
}

func (g *fsmGraph) detectCycle() (bool, []string) {
	visited := make(map[string]bool) // 所有访问过的节点
	inPath := make(map[string]bool)  // 当前DFS路径上的节点
	path := make([]string, 0)        // 用于记录环路径

	var dfs func(node string) bool
	dfs = func(node string) bool {
		visited[node] = true
		inPath[node] = true
		path = append(path, node)

		for _, next := range g.adj[node] {
			if !visited[next] {
				if dfs(next) {
					return true // 在子树中发现环
				}
			} else if inPath[next] {
				// 找到环，将环的起点加入路径
				path = append(path, next)
				return true
			}
		}

		// 回溯时从路径中移除当前节点
		path = path[:len(path)-1]
		inPath[node] = false
		return false
	}

	// 对每个未访问的节点进行DFS
	for node := range g.adj {
		if !visited[node] {
			if dfs(node) {
				// 找到环，处理路径使其从环的起点开始
				cycleStart := path[len(path)-1]
				for i, node := range path {
					if node == cycleStart && i < len(path)-1 {
						path = append(path[i:], path[i])
						return true, path
					}
				}
				return true, path
			}
		}
	}

	return false, nil
}

func boom(e error, msg ...string) {
	if e != nil {
		fmt.Println(strings.Join(msg, " "), e)
		panic(e)
	}
}

func assert(condition bool, msg ...string) {
	if !condition {
		fmt.Println(strings.Join(msg, " "), "assert failed")
		panic("assert failed")
	}
}

// 条件变量，万能并行计算框架
//
// for {
// 	var work
//
// 	lock()
// 	if !(has_new_work() || all_done()) {
// 		unlock()
// 		continue

// 	}

// 	if work == nil {
// 		break
// 	}

// 	work.run()
// 	release_work()

// 	unlock()
// }
