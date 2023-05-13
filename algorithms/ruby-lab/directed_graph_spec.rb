require_relative './directed_graph'

# gem install rspec
# rspec ./directed_graph_spec.rb

RSpec.describe DirectedGraph do
  it 'works for numbers' do
    graph = described_class.new(4)
    graph.add_edge(0, 1)
    graph.add_edge(1, 2)
    graph.add_edge(1, 3)

    expect(graph.edges_count).to eq(3)
    expect(graph.adj(0)).to eq([1])
    expect(graph.adj(1)).to eq([2, 3])
    expect(graph.adj(2)).to eq([])
    expect(graph.adj(3)).to eq([])
    expect(graph.vertices).to eq([0, 1, 2, 3])

    expect(graph.in_degree(0)).to eq(0)
    expect(graph.in_degree(1)).to eq(1)
    expect(graph.in_degree(2)).to eq(1)
    expect(graph.in_degree(3)).to eq(1)
    expect(graph.out_degree(0)).to eq(1)
    expect(graph.out_degree(1)).to eq(2)
    expect(graph.out_degree(2)).to eq(0)
    expect(graph.out_degree(3)).to eq(0)

    expect(graph.cycles.empty?).to eq(true)
    graph.add_edge(3, 0)
    expect(graph.cycles.empty?).to eq(false)
  end
end
