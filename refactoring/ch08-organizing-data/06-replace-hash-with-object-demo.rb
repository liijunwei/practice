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
  def initialize
    @old_networks, @nodes = [], []
  end

  def [](attribute)
    instance_variable_get "@#{attribute}"
  end

  def []=(attribute, value)
    instance_variable_set "@#{attribute}", value
  end
end

# rspec 06-replace-hash-with-object-demo.rb
RSpec.describe NetworkResult do
  specify do
    network = Network.new("foobar")
    node = Node.new(network)

    new_network = NetworkResult.new
    new_network[:old_networks] << node.network
    new_network[:nodes] << node
    new_network[:name] = new_network[:old_networks].collect do |network|
      network.name
    end.join(" - ")
  end
end
