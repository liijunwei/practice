class SearchCriteria
  fields = [:author_id, :publisher_id, :isbn, :foo]
  attr_reader :author_id, :publisher_id, :isbn, :foo

  def self.hash_initializer(*attribute_names)
    define_method(:initialize) do |*args|
      data = args.first || {}
      attribute_names.each do |attribute_name|
        instance_variable_set "@#{attribute_name}", data[attribute_name]
      end
    end
  end

  hash_initializer :author_id, :publisher_id, :isbn, :foo
end

require 'rspec'
require 'pry'

RSpec.describe SearchCriteria do
  let(:hash) { {author_id: 1, publisher_id: 2, isbn: 'foobar', foo: 1} }

  it 'has instance variables initialized' do
    criteria = described_class.new(hash)

    expect(criteria.author_id).to eq(1)
    expect(criteria.publisher_id).to eq(2)
    expect(criteria.isbn).to eq('foobar')
    expect(criteria.isbn).to eq('foobar')
    expect(criteria.foo).to eq(1)
  end
end
