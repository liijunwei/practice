class Order
  attr_accessor :customer
end

class Customer
 def initialize
    @orders = Set.new
 end
end
