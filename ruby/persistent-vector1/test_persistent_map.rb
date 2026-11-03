# frozen_string_literal: true

require "minitest/autorun"
require_relative "persistent_map"

class PersistentMapTest < Minitest::Test
  def test_basic_ops
    m = PersistentMap[:a, 1, :b, 2, :c, 3]
    assert_equal 1, m[:a]
    assert_equal 3, m[:c]
    assert_nil m[:missing]
    assert_equal 3, m.size
    assert m.key?(:a)
    refute m.key?(:missing)
    assert_equal %i[a b c], m.keys.sort
    assert_equal [1, 2, 3], m.values.sort
  end

  def test_assoc_dissoc_are_immutable
    m  = PersistentMap.from_h(a: 1, b: 2)
    m2 = m.assoc(:c, 3).dissoc(:a)
    assert_equal({ b: 2, c: 3 }, m2.to_h)
    assert_equal({ a: 1, b: 2 }, m.to_h)
  end

  def test_fetch
    m = PersistentMap[:a, 1]
    assert_equal 1, m.fetch(:a)
    assert_equal :dflt, m.fetch(:x, :dflt)
    assert_equal :blk, m.fetch(:x) { |k| :blk }
    assert_raises(KeyError) { m.fetch(:x) }
  end

  def test_collision_keys_share_leaf_bucket
    klass = Class.new do
      def initialize(id)
        @id = id
      end

      def hash
        42
      end

      attr_reader :id
    end
    a = klass.new(1)
    b = klass.new(2)
    m = PersistentMap.from_h(a => "a", b => "b")
    assert_equal "a", m[a]
    assert_equal "b", m[b]
    assert_equal 2, m.size
  end

  def test_order_independent_equality_and_hash
    m1 = PersistentMap.from_h(a: 1, b: 2, c: 3)
    m2 = PersistentMap.from_h(c: 3, a: 1, b: 2)
    assert_equal m1, m2
    assert_equal m1.hash, m2.hash
    assert_equal m1, { a: 1, b: 2, c: 3 }
    assert_equal m1.hash, { a: 1, b: 2, c: 3 }.hash
  end

  def test_merge_select_reject
    m = PersistentMap[:a, 1, :b, 2]
    assert_equal({ a: 1, b: 2, c: 3 }, m.merge(c: 3).to_h)
    assert_equal({ a: 1 }, m.select { |_, v| v.odd? }.to_h)
    assert_equal({ b: 2 }, m.reject { |_, v| v.odd? }.to_h)
  end

  def test_differential_vs_ruby_hash
    rng = Random.new(7)
    30.times do |round|
      pm = PersistentMap.empty
      h = {}
      versions = [pm]
      snapshots = [h.dup]

      200.times do
        key = rng.rand(20)
        case rng.rand(3)
        when 0, 1 # assoc
          val = rng.rand(1000)
          pm = pm.assoc(key, val)
          h[key] = val
        when 2 # dissoc
          pm = pm.dissoc(key)
          h.delete(key)
        end
        assert_equal h, pm.to_h, "round #{round} 模型不一致"
        versions << pm
        snapshots << h.dup
      end

      versions.each_with_index do |v, i|
        assert_equal snapshots[i], v.to_h, "round #{round} 第 #{i} 个历史版本被污染"
      end
    end
  end
end
