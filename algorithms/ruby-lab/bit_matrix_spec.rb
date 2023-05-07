require 'rspec'
require_relative './bit_matrix'

# rspec bit_matrix_spec.rb
RSpec.describe BitMatrix do
  v, _e, *edges = File.readlines("../algs4/data/tinyG.txt")

  let(:graph) { BitMatrix.new(v.to_i) }

  before do
    edges.map {|e| e.strip.split(" ")}.each {|from, to| graph.set(from.to_i, to.to_i) }
  end

  it 'has correct graph implementation' do
    expect(graph.V).to eq(13)
    expect(graph.E).to eq(13)

    expect(graph.adj(0).sort).to eq([1, 2, 5, 6])
    expect(graph.adj(1).sort).to eq([0])
    expect(graph.adj(2).sort).to eq([0])
    expect(graph.adj(3).sort).to eq([4, 5])
    expect(graph.adj(4).sort).to eq([3, 5, 6])
    expect(graph.adj(5).sort).to eq([0, 3, 4])
    expect(graph.adj(6).sort).to eq([0, 4])
    expect(graph.adj(7).sort).to eq([8])
    expect(graph.adj(8).sort).to eq([7])
    expect(graph.adj(9).sort).to eq([10, 11, 12])
    expect(graph.adj(10).sort).to eq([9])
    expect(graph.adj(11).sort).to eq([9, 12])
    expect(graph.adj(12).sort).to eq([9, 11])

    expect(graph.vertices.sort).to eq((0..12).to_a)
  end
end
