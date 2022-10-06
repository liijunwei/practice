require 'rspec'
require 'pry'

class Person
  def initialize(female = nil, code = nil)
    @female = female
    @code = code
  end

  def self.create_female
    Female.new
  end

  def self.create_male
    Male.new
  end

  def female?
    @female
  end

  def code
    @code
  end
end

class Female < Person
  def initialize(female = nil, code = nil)
    super(true, 'F')
  end

  def female?
    @female
  end

  def code
    @code
  end
end

class Male < Person
  def initialize(female = nil, code = nil)
    super(false, 'M')
  end

  def female?
    @female
  end

  def code
    @code
  end
end

RSpec.describe Female do
  let(:pe) { Person.create_female }

  it 'is female' do
    expect(pe.female?).to be_truthy
    expect(pe.code).to eq("F")
  end
end

RSpec.describe Male do
  let(:pe) { Person.create_male }

  it 'is male' do
    expect(pe.female?).to be_falsey
    expect(pe.code).to eq("M")
  end
end
