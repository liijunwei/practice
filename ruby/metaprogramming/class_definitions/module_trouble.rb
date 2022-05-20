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

RSpec.describe "module trouble" do
  describe "my_method" do
    it 'raises error' do
      expect { MyClass.my_method }.to raise_error(NoMethodError)
    end
  end
end
