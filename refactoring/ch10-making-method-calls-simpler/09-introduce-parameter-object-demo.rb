require 'pry'
require 'rspec'

class Account
  def initialize
    @charges = []
  end

  def add_charge(base_price, tax_rate, imported, charge = nil)
    total = base_price + base_price * tax_rate
    total += base_price * 0.1 if imported
    @charges << total
  end

  def total_charge
    @charges.inject(0) { |total, charge| total + charge }
  end
end

class Charge
  attr_accessor :base_price, :tax_rate, :imported

  def initialize(base_price, tax_rate, imported)
    @base_price = base_price
    @tax_rate = tax_rate
    @imported = imported
  end
end

RSpec.describe Account do
  let(:account) { subject }
  specify do
    account.add_charge(5, 0.1, true, Charge.new(5, 0.1, true))
    account.add_charge(12, 0.125, false, Charge.new(12, 0.125, false))
    total = account.total_charge
    expect(total).to eq(19.5)
  end
end
