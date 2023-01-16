class MyFactoryBot
  def self.define
  end
end

MyFactoryBot.define do
  factory :user do
    name { FFaker::Name.name }
    email { FFaker::Internet.unique.email }
  end
end

require 'rspec'
require 'pry'

# rspec my_factory_bot.rb
RSpec.describe MyFactoryBot do
end
