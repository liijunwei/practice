require 'rspec'
require 'pry'

class User
  attr_accessor :name
  attr_accessor :email
end

class MyFactoryBot
  def self.define(&block)
    instance_exec(&block)
  end

  def self.create(model_sym)
    User.new
  end
end

MyFactoryBot.define do
  factory :user do
    require "pry"; binding.pry
    name { 'Bogs' }
    email { 'demo@example.com' }
  end
end

# rspec my_factory_bot.rb
RSpec.describe MyFactoryBot do
  describe '.create' do
    it 'creates a new instance of User' do
      user = described_class.create(:user)
      expect(user.name).to eq('Bogs')
      expect(user.email).to eq('demo@example.com')
    end
  end
end
