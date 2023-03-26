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

  def connected?
    @count == @graph.V
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
require 'rspec'

# rspec ruby-lab/depth_first_search.rb
RSpec.describe DepthFirstSearch do
  v, _e, *edges = File.readlines("data/tinyG.txt")

  let(:graph) { BitMatrix.new(v.to_i) }
  let(:search) { DepthFirstSearch.new(graph, source) }
  let(:source) { 0 }

  before do
    edges.map {|e| e.strip.split(" ")}.each {|from, to| graph.set(from.to_i, to.to_i) }
  end

  context 'when source is 0' do
    specify do
      expect(graph.vertices.select {|v| search.marked(v)}.sort).to eq([0,1,2,3,4,5,6])
      expect(search.connected?).to eq(false)
    end
  end

  context 'when source is 7' do
    let(:source) { 7 }

    specify do
      expect(graph.vertices.select {|v| search.marked(v)}.sort).to eq([7,8])
      expect(search.connected?).to eq(false)
    end
  end

  context 'when source is 9' do
    let(:source) { 9 }

    specify do
      expect(graph.vertices.select {|v| search.marked(v)}.sort).to eq([9,10,11,12])
      expect(search.connected?).to eq(false)
    end
  end

  context 'when new edges added' do
    before do
      graph.set(6, 7)
      graph.set(8, 9)
    end

    specify do
      expect(graph.vertices.select {|v| search.marked(v)}.sort).to eq([0,1,2,3,4,5,6,7,8,9,10,11,12])
      expect(search.connected?).to eq(true)
    end
  end
end
