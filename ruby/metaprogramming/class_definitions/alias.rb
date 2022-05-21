# page 131

class MyClass
  def my_method
    "hello"
  end

  alias :m :my_method          # alias 是一个关键字, 不是方法
  alias_method :m2, :my_method # Module#alias_method 是一个方法, 功能和alias关键字一样
end

require 'rspec'

RSpec.describe MyClass do
  describe "my_method" do
    it 'has an alias of m' do
      expect(subject.my_method).to eq("hello")
      expect(subject.m).to eq("hello")
    end

    it 'has an alias of m2' do
      expect(subject.my_method).to eq("hello")
      expect(subject.m2).to eq("hello")
    end
  end
end
