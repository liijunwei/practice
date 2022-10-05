class Order
  def initialize(customer_name)
    @customer = Customer.create(customer_name)
  end

  def customer=(customer_name)
    @customer = Customer.create(customer_name)
  end

  def customer_name
    @customer.name
  end
end

class Customer
  attr_reader :name

  Instances = {}

  def initialize(name)
    @name = name
  end

  def self.create(name)
    Customer.new(name)
  end
end

require 'pry'
require 'rspec'

# rspec 03-change-value-to-reference.rb
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
