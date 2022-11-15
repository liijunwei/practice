def rank(key, sorted_list)
  low = 0
  high = sorted_list.size - 1

  while low <= high
    mid = (low + high) / 2

    if key < sorted_list[mid]
      high = mid - 1
    elsif key > sorted_list[mid]
      low = mid + 1
    else
      return mid
    end
  end

  return -1
end

def simple_rank(key, sorted_list)
  sorted_list.each_with_index do |e, i|
    return i if e == key
  end

  return -1
end

require 'test/unit'

class TestRank < Test::Unit::TestCase
  def setup
    @list = [1, 2, 3, 4].sort
  end

  def test_key_found
    assert_equal rank(1, @list), 0
    assert_equal rank(3, @list), 2
    assert_equal simple_rank(1, @list), 0
    assert_equal simple_rank(3, @list), 2
  end

  def test_key_not_found
    assert_equal rank(999, @list), -1
    assert_equal simple_rank(999, @list), -1
  end
end
