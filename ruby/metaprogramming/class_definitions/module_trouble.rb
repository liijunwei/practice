# page 129

module MyModule
  def self.my_method
    "hello"
  end
end

class MyClass
  include MyModule
end

require 'rspec'

RSpec.describe MyClass do
  describe "my_method" do
    it 'raises error' do
      expect { described_class.my_method }.to raise_error(NoMethodError)
    end
  end
end
