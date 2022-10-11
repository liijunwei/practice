class Policy < Hash
  attr_reader :name

  def initialize(name)
    @name = name
  end

  def <<(rule)
    key = rule.attribute
    self[key] ||= []
    self[key] << rule
  end

  def apply(account)
    self.each do |attribute, rules|
      rules.each { |rule| rule.apply(account) }
    end
  end
end

class Rule
  attr_reader :attribute, :default_value

  def initialize(attribute, default_value)
    @attribute, @default_value = attribute, default_value
  end

  def apply(account)
    "applying rule #{attribute} for #{account}"
  end
end

require 'pry'
require 'rspec'

RSpec.describe Policy do
end

RSpec.describe Rule do
  let(:account) { 'foobar' }

  specify do
    rule = described_class.new('age_must_be_greater_than', 18)
    expect(rule.attribute).to eq('age_must_be_greater_than')
    expect(rule.default_value).to eq(18)
    expect(rule.apply(account)).to eq("applying rule age_must_be_greater_than for foobar")
  end
end
