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

# client code
orders =
customer =
orders.select { |order| order.customer_name == customer.name }.size

# At the moment Customer is a value object.
# Each order has its own customer object even if they are for the same conceptual customer.
# I want to change this so that if we have several orders for the same conceptual customer, they share a single customer object.
# For this case this means that there should be only one customer object for each customer name.
