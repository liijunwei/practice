# page 129

module MyModule
  def my_method
    "hello"
  end
end

class MyClass
  extend MyModule
end

require 'rspec'

RSpec.describe MyClass do
  describe "my_method" do
    let(:obj) { Object.new }

    it 'works' do
      expect(described_class.my_method).to eq("hello")
    end
  end


end
