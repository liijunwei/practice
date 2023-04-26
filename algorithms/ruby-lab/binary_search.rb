def rank(key, a)
  low = 0
  high = a.size - 1

  while low <= high
    mid = (low + high) / 2

    if key < a[mid]
      high = mid - 1
    elsif key > a[mid]
      low = mid + 1
    else
      return mid
    end
  end

  return -1
end

def simple_rank(key, a)
  a.each_with_index do |e, i|
    return i if e == key
  end

  return -1
end

def recursive_rank(key, a)
  r_rank(key, a, 0, a.size - 1)
end

def r_rank(key, a, low, high)
  return -1 if low > high

  mid = (low + high) / 2

  if key < a[mid]
    r_rank(key, a, low, mid-1)
  elsif key > a[mid]
    r_rank(key, a, mid+1, high)
  else
    mid
  end
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
    assert_equal recursive_rank(1, @list), 0
    assert_equal recursive_rank(3, @list), 2
  end

  def test_key_not_found
    assert_equal rank(999, @list), -1
    assert_equal simple_rank(999, @list), -1
    assert_equal recursive_rank(999, @list), -1
  end
end
