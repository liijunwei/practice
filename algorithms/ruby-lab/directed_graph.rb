class DirectedGraph
  attr_reader :edges_count

  def initialize(vertices_count)
    @vertices_count = vertices_count
    @edges_count = 0
    @adj = Array.new(vertices_count) { Set.new }
  end

  def add_edge(v, w)
    @adj[v].add(w)
    @edges_count += 1
  end

  def adj(v)
    @adj[v].to_a
  end

  def in_degree(v)
    @adj.select { |w| w.include?(v) }.count
  end

  def out_degree(v)
    adj(v).count
  end

  def cycles
    @marked = Array.new(@vertices_count) { false }
    @edge_to = Array.new(@vertices_count)
    @on_stack = Array.new(@vertices_count) { false }
    @cycle = []
    @vertices_count.times { |s| dfs(s) if !@marked[s] }

    @cycle.reverse
  end

  def vertices
    @vertices_count.times.to_a
  end

  private

  def have_cycle?
    !@cycle.empty?
  end

  def dfs(v)
    @on_stack[v] = true
    @marked[v] = true

    adj(v).each do |w|
      if have_cycle?
        return
      elsif !@marked[w]
        @edge_to[w] = v
        dfs(w)
      elsif @on_stack[w]
        @cycle = []

        x = v
        while x != w
          @cycle << x
          x = @edge_to[x]
        end

        @cycle << w
        @cycle << v
      else
        # :nocov:
        # do nothing
        # :nocov:
      end
    end

    @on_stack[v] = false
  end
end
