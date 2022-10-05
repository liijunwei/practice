require 'pry'
require 'rspec'

class Node
  attr_reader :network

  def initialize(network)
    @network = network
  end
end

class Network
  attr_reader :name

  def initialize(name)
    @name = name
  end
end

class NetworkResult
  attr_reader :old_networks
  attr_reader :nodes
  attr_accessor :name

  def initialize
    @old_networks, @nodes = [], []
  end

  def name
    @old_networks.collect { | network | network.name }.join(" - ")
  end
end

# rspec 06-replace-hash-with-object-demo.rb
RSpec.describe NetworkResult do
  specify do
    network = Network.new("foobar")
    node = Node.new(network)

    new_network = NetworkResult.new
    new_network.old_networks << node.network
    new_network.nodes << node

    expect(new_network.name).to eq('foobar')
  end
end
