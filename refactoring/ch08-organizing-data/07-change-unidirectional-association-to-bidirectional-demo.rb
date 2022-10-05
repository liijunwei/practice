require 'set'
require 'pry'
require 'rspec'

# customer->orders
# one to many

class Order
  attr_reader :customer

  def customer=(value)
    customer.friend_orders.subtract(self) unless customer.nil?
    @customer = value
    customer.friend_orders.add(self) unless customer.nil?
  end
end

# Because the order will take charge
# I need to add a helper method to the customer
# that allows direct access to the orders collection.
class Customer
  attr_reader :orders

  def initialize
    @orders = Set.new
  end

  def add_order(order)
    @orders << order
  end

  def friend_orders
    #should only be used by Order when modifying the association
    @orders
  end
end

# rspec 07-change-unidirectional-association-to-bidirectional-demo.rb
RSpec.describe Order do
  it 'responds to customer' do
    o = Order.new

    expect(o.respond_to?(:customer)).to be_truthy
  end

  it 'gets customer info' do
    c = Customer.new
    o1 = Order.new
    o2 = Order.new

    # 双向关联的责任交由 order 里的 `customer=` 维护
    o1.customer = c
    o2.customer = c

    expect(o1.customer).to eq(c)
    expect(o2.customer).to eq(c)
    expect(c.orders).to eq(Set.new([o1, o2]))
  end

  it 'does not have customer' do
    o1 = Order.new

    o1.customer = nil
    expect(o1.customer).to be_nil
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
