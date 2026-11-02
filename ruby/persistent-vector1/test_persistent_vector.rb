# frozen_string_literal: true

require "minitest/autorun"
require_relative "persistent_vector"

class PersistentVectorTest < Minitest::Test
  def test_lookup
    v = PersistentVector.from_a((0...1000).to_a)
    assert_equal 500, v[500]
    assert_equal 999, v.nth(999)
  end

  def test_lookup_out_of_bounds
    v = PersistentVector.from_a([1, 2, 3])
    assert_raises(IndexError) { v[3] }
    assert_raises(IndexError) { v[-1] }
  end

  def test_assoc_returns_new_vector_and_keeps_original
    v  = PersistentVector.from_a((0...1000).to_a)
    v2 = v.assoc(500, :beef)
    assert_equal :beef, v2[500]
    assert_equal 500, v[500]
  end

  def test_pop_equality_but_not_identity
    v = PersistentVector.from_a((0...50).to_a)
    a = v.pop
    b = v.pop
    assert_equal a, b
    refute_same a.root, b.root
  end

  def test_pop_empty_raises
    assert_raises(IndexError) { PersistentVector.empty.pop }
  end

  def test_path_copying_shares_untouched_subtrees
    big  = PersistentVector.from_a((1..100).to_a)
    big2 = big.assoc(5, :x)
    refute_same big.root, big2.root
    refute_same big.root[0], big2.root[0]
    assert_same big.root[1], big2.root[1]
  end

  def test_push_grows_new_root_on_overflow
    v32 = PersistentVector.from_a((0...32).to_a)
    v33 = v32.push(32)
    assert_equal 5, v33.shift
    assert_same v32.root, v33.root[0]
  end

  def test_vector_and_nodes_are_frozen
    big = PersistentVector.from_a((1..100).to_a)
    assert_predicate big, :frozen?
    assert_predicate big.root, :frozen?
  end

  def test_enumerable
    v = PersistentVector.from_a((1..10).to_a)
    assert_equal (1..10).map { |x| x * 2 }, v.map { |x| x * 2 }
    assert_equal [2, 4, 6, 8, 10], v.select(&:even?)
  end

  def test_empty_first_last_and_equality
    e = PersistentVector.empty
    assert_predicate e, :empty?
    assert_nil e.first
    assert_nil e.last

    v = PersistentVector.from_a([1, 2, 3])
    assert_equal 1, v.first
    assert_equal 3, v.last
    assert_equal PersistentVector[1, 2, 3], v
  end

end
