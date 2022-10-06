require 'rspec'
require 'pry'

class Person
  def initialize(female, code)
    @female = female
    @code = code
  end

  def female?
    @female
  end

  def code
    @code
  end
end

RSpec.describe Person do
  let(:f) { Person.new(true, 'F') }
  let(:m) { Person.new(false, 'M') }

  it 'is female' do
    expect(f.female?).to be_truthy
    expect(f.code).to eq("F")
  end

  it 'is male' do
    expect(m.female?).to be_falsey
    expect(m.code).to eq("M")
  end
end
