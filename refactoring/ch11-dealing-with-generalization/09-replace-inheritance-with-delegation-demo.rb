require 'forwardable'
require 'pry'
require 'rspec'

class Policy
  extend Forwardable
  def_delegators :@rules, :size, :empty?, :[] # unclear

  attr_reader :name

  def initialize(name)
    @name = name
    @rules = {}
  end

  def <<(rule)
    key = rule.attribute
    @rules[key] ||= []
    @rules[key] << rule
  end

  def apply(account)
    @rules.each do |attribute, rules|
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

RSpec.describe Policy do
  let(:account) { 'foobar@policy' }

  specify do
    policy = Policy.new('kyc_terms')
    policy << Rule.new('age_must_be_greater_than', 18)
    policy << Rule.new('age_must_be_greater_than', 28)
    policy << Rule.new('country_must_be', 'china')
    policy << Rule.new('account_balance_must_be_greater_than', 10_000)
    require "pry"; binding.pry

    policy.apply(account)
  end
end

RSpec.describe Rule do
  let(:account) { 'foobar@rule' }

  specify do
    rule = described_class.new('age_must_be_greater_than', 18)
    expect(rule.attribute).to eq('age_must_be_greater_than')
    expect(rule.default_value).to eq(18)
    expect(rule.apply(account)).to eq("applying rule age_must_be_greater_than for foobar@rule")
  end
end
