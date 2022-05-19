# page 114

class MyClass
  def my_atrribute=(value)
    @my_atrribute = value
  end

  def my_atrribute
    @my_atrribute
  end
end

class MyClassRefactored
  attr_accessor :my_atrribute
end

require 'rspec'

[MyClass, MyClassRefactored].each do |classname|
  RSpec.describe classname do
    describe "#my_atrribute=" do
      it "sets value for instance variable" do
        subject.my_atrribute = 99
        expect(subject.instance_variable_get("@my_atrribute")).to eq(99)
      end
    end

    describe "#my_atrribute" do
      it "is nil when not set" do
        expect(subject.instance_variable_get("@my_atrribute")).to be_nil
      end

      it "gets value for instance variable" do
        subject.my_atrribute = 99
        expect(subject.my_atrribute).to eq(99)
      end
    end
  end
end
