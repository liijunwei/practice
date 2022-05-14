require 'test/unit'
require_relative 'using'

# ruby ruby/metaprogramming/blocks/using_test.rb

class TestUsing < Test::Unit::TestCase
  class Resource
    def dispose
      @dispose = true
    end

    def disposed?
      @dispose
    end
  end

  def test_disposes_of_resources
    r = Resource.new
    using(r) {}

    assert r.disposed?
  end

  def test_disposes_of_resources_jn_case_of_exception
    r = Resource.new

    assert_raise(Exception) {
      using(r) {
        raise Exception
      }
    }

    assert r.disposed?
  end
end
