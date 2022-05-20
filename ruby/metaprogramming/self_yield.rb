class Person
  attr_accessor :name, :age
  def initialize
    yield self
  end
end

require 'rspec'

RSpec.describe Person do
  describe "#initialize" do
    let(:person) do
      described_class.new do |p|
        p.name = "Bob"
        p.age  = 99
      end
    end

    it "is initialized without error" do
      expect(person.name).to eq("Bob")
      expect(person.age).to eq(99)
    end

    it "raises an error" do
      expect do
        described_class.new do |p|
          p.name   = "Bob"
          p.age    = 99
          p.gender = "male"
        end
      end.to raise_error(NoMethodError)
    end
  end
end


