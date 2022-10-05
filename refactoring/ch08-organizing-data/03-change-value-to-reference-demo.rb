class Order
  def initialize(customer_name)
    @customer = Customer.with_name(customer_name)
  end

  def customer=(customer_name)
    @customer = Customer.with_name(customer_name)
  end

  def customer_name
    @customer.name
  end
end

class Customer
  attr_reader :name

  # Then I decide whether to create customers on the fly when asked or to create them in advance.
  Instances = {}

  def initialize(name)
    @name = name
  end

  def self.with_name(name)
    Instances[name]
  end

  def self.load_customers
    require 'yaml'
    YAML.load_file("./03-change-value-to-reference-customers.yml").each do |c|
      new(c).store
    end
  end

  def store
    Instances[name] = self
  end
end

require 'pry'
require 'rspec'

# rspec 03-change-value-to-reference-demo.rb
RSpec.describe Order do
  def number_of_orders_for(orders, customer)
    orders.select { |order| order.customer_name == customer.name }.size
  end

  before { Customer.load_customers }

  let(:customer_a) { Customer.new('a') }
  let(:customer_ab) { Customer.new('ab') }

  let(:orders) do
    os = []
    os << Order.new("a")
    os << Order.new("c")
    os << Order.new("d")
    os << Order.new("e")
    os
  end

  # client code
  describe 'number_of_orders_for' do
    specify do
      expect(number_of_orders_for(orders, customer_a)).to eq(1)
      expect(number_of_orders_for(orders, customer_ab)).to eq(0)
    end
  end

  describe 'Customer Instances' do
    it 'has pre-loaded customers' do
      Customer.load_customers
      expect(Customer::Instances.keys.count).to eq(5)
    end
  end
end
