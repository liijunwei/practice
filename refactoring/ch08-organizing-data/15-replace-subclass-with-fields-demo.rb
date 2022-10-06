require 'rspec'
require 'pry'

class Person
  def initialize(female = nil, code = nil)
    @female = female
    @code = code
  end

  def self.create_female
    new(true, 'F')
  end

  def self.create_male
    new(false, 'M')
  end

  def female?
    @female
  end

  def code
    @code
  end
end

RSpec.describe Person do
  let(:f) { Person.create_female }
  let(:m) { Person.create_male }

  it 'is female' do
    expect(f.female?).to be_truthy
    expect(f.code).to eq("F")
  end

  it 'is male' do
    expect(m.female?).to be_falsey
    expect(m.code).to eq("M")
  end
end
