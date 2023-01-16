require 'rspec'
require 'pry'
require 'active_support/all'

class User
  attr_accessor :name
  attr_accessor :email
  attr_accessor :gender
end

class Website
  attr_accessor :url
end

class MyFactoryBot
  def self.define(&block)
    instance_exec(&block)
  end

  def self.create(model_sym)
    @factory.record
  end

  def self.factory(model_sym, &block)
    @factory = MyFactory.new(model_sym)
    @factory.instance_exec(&block)
  end
end

class MyFactory
  attr_reader :record

  def initialize(model_sym)
    @record = Kernel.const_get(model_sym.to_s.classify).new
  end

  def method_missing(attr, *args, &block)
    @record.send("#{attr}=", block.call)
  end
end

MyFactoryBot.define do
  factory :user do
    name { 'Bogs' }
    email { 'demo@example.com' }
    gender { 'male' }
  end
end

MyFactoryBot.define do
  factory :website do
    url { 'https://www.codewithjason.com/articles/' }
  end
end

# rspec my_factory_bot.rb
RSpec.describe MyFactoryBot do
  describe '.create' do
    it 'creates a new instance of User' do
      user = described_class.create(:user)
      expect(user.name).to eq('Bogs')
      expect(user.email).to eq('demo@example.com')
      expect(user.gender).to eq('male')
    end

    it 'creates a new instance of User' do
      website = described_class.create(:website)
      expect(website.url).to eq('https://www.codewithjason.com/articles/')
    end
  end
end
