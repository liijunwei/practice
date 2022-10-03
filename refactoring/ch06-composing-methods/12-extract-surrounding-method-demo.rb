class Person
  attr_reader :mother, :children, :name

  def initialize(name, date_of_birth, date_of_death=nil, mother=nil)
    @name = name
    @mother = mother
    @date_of_birth = date_of_birth
    @date_of_death = date_of_death
    @children = []
    @mother.add_child(self) if @mother
  end

  def add_child(child)
    @children << child
  end

  def number_of_living_descendants
    count_descendants_matching { |descendant| descendant.alive? }
  end

  def number_of_descendants_named(name)
    count_descendants_matching { |descendant| descendant.name == name }
  end

  def alive?
    @date_of_death.nil?
  end

  protected

  def count_descendants_matching(&block)
    children.reduce(0) do |count, child|
      count += 1 if yield child
      count + child.count_descendants_matching(&block)
    end
  end
end

require 'rspec'
require 'pry'

# rspec 12-extract-surrounding-method-demo.rb
RSpec.describe Person do
  let(:mother) { described_class.new("mia", "1963-01-01", nil, nil) }

  before do
    described_class.new("tiny-mia1", "1993-01-01", nil, mother)
    described_class.new("tiny-mia2", "1994-01-01", nil, mother)
    described_class.new("tiny-mia2", "1998-01-01", nil, mother)
  end

  describe '#number_of_living_descendants' do
    specify { expect(mother.number_of_living_descendants).to eq(3) }
  end

  describe '#number_of_descendants_named' do
    specify { expect(mother.number_of_descendants_named('tiny-mia2')).to eq(2) }
    specify { expect(mother.number_of_descendants_named('tiny-mia1')).to eq(1) }
    specify { expect(mother.number_of_descendants_named('foobar')).to be_zero }
  end
end
