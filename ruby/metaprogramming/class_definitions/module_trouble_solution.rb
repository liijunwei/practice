# page 129

module MyModule
  def my_method
    "hello"
  end
end

class MyClass
  class << self
    include MyModule
  end
end

require 'rspec'

RSpec.describe "module trouble" do
  describe "my_method" do
    it 'works' do
      expect(MyClass.my_method).to eq("hello")
    end
  end
end
