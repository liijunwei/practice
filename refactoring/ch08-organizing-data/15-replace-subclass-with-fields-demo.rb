require 'rspec'
require 'pry'

class Person
  def self.create_female
    Female.new
  end

  def self.create_male
    Male.new
  end
end

class Female < Person
  def female?
    true
  end

  def code
    'F'
  end
end

class Male < Person
  def female?
    false
  end

  def code
    'M'
  end
end

RSpec.describe Female do
  subject { Person.create_female }

  it 'is female' do
    expect(subject.female?).to be_truthy
    expect(subject.code).to eq("F")
  end
end

RSpec.describe Male do
  subject { Person.create_male }

  it 'is male' do
    expect(subject.female?).to be_falsey
    expect(subject.code).to eq("M")
  end
end
