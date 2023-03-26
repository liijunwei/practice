class DepthFirstSearch
  def initialize(graph, source)
    @graph = graph
    @source = source
    @marked = Array.new(graph.V, false)
    @count = 0

    dfs(graph, source)
  end

  def marked(w)
    @marked[w]
  end

  def count
    @count
  end

  private

  def dfs(graph, v)
    @marked[v] = true

    @count += 1

    # 不能改为 graph.adj(v).select {|w| !@marked[w]}.each {|w| dfs(graph, w)}
    # 还没看懂原因，大致是还没理解这里递归的含义和执行过程
    graph.adj(v).each do |w|
      dfs(graph, w) if !@marked[w]
    end
  end
end

require_relative './bit_matrix'

# The order of a graph is its number of vertices |V|
# The size of a graph is its number of edges |E|

v, e, *edges = File.readlines("data/tinyG.txt")

V = v.to_i
E = e.to_i
graph = BitMatrix.new(V)
edges.map {|e| e.strip.split(" ")}.each {|from, to| graph.set(from.to_i, to.to_i) }

source = ARGV[0].to_i
search = DepthFirstSearch.new(graph, source)

graph.vertices.select {|v| search.marked(v)}.each {|v| print("#{v} ")}
puts
print("NOT ") if search.count != graph.V
puts "connected"
