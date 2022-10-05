class Order
  def initialize(customer_name)
    @customer = Customer.new(customer_name)
  end

  def customer=(customer_name)
    @customer = Customer.new(customer_name)
  end

  def customer_name
    @customer.name
  end
end

class Customer
  attr_reader :name

  def initialize(name)
    @name = name
  end
end

require 'pry'
require 'rspec'

RSpec.describe Order do
  def number_of_orders_for(orders, customer)
    orders.select { |order| order.customer_name == customer.name }.size
  end

  let(:customer_a) { Customer.new('a') }
  let(:customer_ab) { Customer.new('ab') }

  let(:orders) do
    os = []
    os << Order.new("a")
    os << Order.new("b")
    os << Order.new("c")
    os << Order.new("d")
    os
  end

  # client code
  describe 'number_of_orders_for' do
    specify do
      expect(number_of_orders_for(orders, customer_a)).to eq(1)
      expect(number_of_orders_for(orders, customer_ab)).to eq(0)
    end
  end
end

# At the moment Customer is a value object.
# Each order has its own customer object even if they are for the same conceptual customer.
# I want to change this so that if we have several orders for the same conceptual customer, they share a single customer object.
# For this case this means that there should be only one customer object for each customer name.
