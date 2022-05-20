# page 129

module MyModule
  def my_method
    "hello"
  end
end

require 'rspec'

RSpec.describe Object do
  before do
    class << subject
      include MyModule
    end
  end

  describe "my_method" do
    it 'works' do
      expect(subject.my_method).to eq("hello")
      expect(subject.singleton_methods).to eq([:my_method])
    end
  end
end



