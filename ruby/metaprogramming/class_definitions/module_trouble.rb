# page 129

module MyModule
  def self.my_method
    "hello"
  end
end

class MyClass
  include MyModule
end

module MyModuleRefactored
  def my_method
    "hello"
  end
end

class MyClassRefactored
  extend MyModuleRefactored
end

require 'rspec'

RSpec.describe "module trouble" do
  describe "my_method" do
    it 'raises error' do
      expect { MyClass.my_method }.to raise_error(NoMethodError)
    end

    it 'works' do
      expect(MyClassRefactored.my_method).to eq("hello")
    end
  end
end
