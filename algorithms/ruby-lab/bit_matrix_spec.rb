require 'rspec'
require_relative './bit_matrix'

# rspec bit_matrix_spec.rb
RSpec.describe BitMatrix do
  it 'has correct graph implementation' do
    graph = BitMatrix.new(13)
    graph.set(0, 5)
    graph.set(4, 3)
    graph.set(0, 1)
    graph.set(9, 12)
    graph.set(6, 4)
    graph.set(5, 4)
    graph.set(0, 2)
    graph.set(11, 12)
    graph.set(9, 10)
    graph.set(0, 6)
    graph.set(7, 8)
    graph.set(9, 11)
    graph.set(5, 3)

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

  it 'works for string vertices' do
    vertices = %w[pending detected confirmed failed]

    graph = BitMatrix.new(vertices.size)
    graph.set(vertices.index('pending'), vertices.index('detected'))
    graph.set(vertices.index('detected'), vertices.index('confirmed'))
    graph.set(vertices.index('detected'), vertices.index('failed'))

    expect(graph.V).to eq(4)
    expect(graph.E).to eq(3)

    expect(graph.adj(vertices.index('pending')).sort).to eq([vertices.index('detected')])
    expect(graph.adj(vertices.index('detected')).sort).to eq([vertices.index('pending'), vertices.index('confirmed'), vertices.index('failed')])
    expect(graph.adj(vertices.index('confirmed')).sort).to eq([vertices.index('detected')])
    expect(graph.adj(vertices.index('failed')).sort).to eq([vertices.index('detected')])
  end
end
