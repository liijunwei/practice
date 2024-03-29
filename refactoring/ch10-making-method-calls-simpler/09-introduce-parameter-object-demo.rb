require 'pry'
require 'rspec'

class Account
  def initialize
    @charges = []
  end

  def add_charge(charge)
    @charges << charge
  end

  def total_charge
    @charges.inject(0) { |total, charge| total + charge.total }
  end
end

class Charge
  attr_reader :base_price, :tax_rate, :imported

  def initialize(base_price, tax_rate, imported)
    @base_price = base_price
    @tax_rate = tax_rate
    @imported = imported
  end

  def total
    total = base_price + base_price * tax_rate
    total += base_price * 0.1 if imported
    total
  end
end

RSpec.describe Account do
  let(:account) { subject }
  specify do
    account.add_charge(Charge.new(5, 0.1, true))
    account.add_charge(Charge.new(12, 0.125, false))

    expect(account.total_charge).to eq(19.5)
  end
end
