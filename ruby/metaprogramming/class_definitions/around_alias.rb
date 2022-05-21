# page 132

class String
  alias :real_length :length

  def length
    real_length > 5 ? "long" : "short"
  end
end

require 'rspec'

RSpec.describe String do
  describe "enhanced string" do
    let(:str) { "War and Peace" }

    it 'uses original length()' do
      expect(str.real_length).to eq(13)
    end

    it 'uses enhanced length()' do
      expect(str.length).to eq("long")
    end
  end
end
