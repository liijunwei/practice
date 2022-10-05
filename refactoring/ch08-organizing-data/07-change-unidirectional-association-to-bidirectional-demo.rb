require 'set'

class Order
  attr_reader :customer
end

class Customer
  attr_reader :orders

  def initialize
    @orders = Set.new
  end

  def add_order(order)
    @orders << order
  end
end

# rspec 07-change-unidirectional-association-to-bidirectional-demo.rb
RSpec.describe Order do
  it 'responds to customer' do
    o = Order.new

    expect(o.respond_to?(:customer)).to be_truthy
  end
end

RSpec.describe Customer do
  it 'responds to orders' do
    c = Customer.new
    o1 = Order.new
    o2 = Order.new

    c.add_order(o1)
    c.add_order(o2)

    expect(c.orders).to eq(Set.new([o1, o2]))
  end
end
