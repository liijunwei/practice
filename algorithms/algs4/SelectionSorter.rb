class SelectionSorter
  def sort(a)
    a.each_with_index do |e, i|
      min_index = find_min_index(i, a)

      swap(a, i, min_index)
    end

    a
  end

  private

  def find_min_index(curr_min_index, a)
    min_index = curr_min_index

    cursor = curr_min_index + 1
    length = a.length

    while (cursor < length)
      min_index = cursor if a[cursor] < a[min_index]
      cursor += 1
    end

    min_index
  end

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
