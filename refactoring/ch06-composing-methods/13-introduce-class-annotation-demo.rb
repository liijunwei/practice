class SearchCriteria
  attr_reader :author_id, :publisher_id, :isbn

  def initialize(hash)
    @author_id = hash[:author_id]
    @publisher_id = hash[:publisher_id]
    @isbn = hash[:isbn]
  end
end

require 'rspec'
require 'pry'

RSpec.describe SearchCriteria do
  let(:hash) { {author_id: 1, publisher_id: 2, isbn: 'foobar'} }

  it 'has instance variables initialized' do
    criteria = described_class.new(hash)

    expect(criteria.author_id).to eq(1)
    expect(criteria.publisher_id).to eq(2)
    expect(criteria.isbn).to eq('foobar')
  end
end
