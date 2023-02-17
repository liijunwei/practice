class SelectionSorter
  def sort(a)
    length = a.length

    a.each_with_index do |e, i|
      min_index = i

      j = i + 1

      while (j < length)
        min_index = j if a[j] < a[min_index]
        j += 1
      end

      swap(a, i, min_index)
    end

    a
  end

  private

  def swap(a, i, j)
    tmp = a[i]
    a[i] = a[j]
    a[j] = tmp
  end
end

require 'test/unit'

# ruby SelectionSorter.rb
class TestSelectionSorter < Test::Unit::TestCase
  def setup
    @list1 = ["S", "O", "R", "T", "E", "X", "A", "M", "P", "L", "E"]
    @list2 = ["A", "E", "E", "L", "M", "O", "P", "R", "S", "T", "X"]
  end

  def test_sort_work_as_expected
    assert_equal @list2, SelectionSorter.new.sort(@list1)
  end
end
