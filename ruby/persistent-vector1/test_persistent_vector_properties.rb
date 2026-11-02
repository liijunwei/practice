# frozen_string_literal: true

require "minitest/autorun"
require_relative "persistent_vector"

# 极简 property-based 测试：不引入 Rantly/PropCheck 这类 gem，
# 用固定种子的随机生成 + 对照 Ruby Array 验证「性质」，而不是写死的个例。
# shortcut: 没有 shrinking（失败时靠 seed/round 重跑定位），也只在有限输入范围里抽样。
class PersistentVectorProperties < Minitest::Test
  BITS = PersistentVector::BITS
  # trie 的深度变化都发生在 32 的幂附近，bug 大多藏在这些点上。
  BOUNDARY_SIZES = [0, 1, 31, 32, 33, 64, 65, 1024, 1025, 1056, 1057, 32_768, 32_800, 32_801]

  # 性质一：参考模型一致 —— 随机操作序列下，向量始终等价于 Ruby Array；
  # 且所有历史版本之后都不会被污染（持久化）。
  def test_conforms_to_array_reference_model
    rng = Random.new(42)
    30.times do |round|
      vec = PersistentVector.empty
      arr = []
      versions = [vec]
      expected = [arr.dup]

      150.times do
        case rng.rand(4)
        when 0 # push
          x = rng.rand(1000)
          vec = vec.push(x)
          arr.push(x)
        when 1 # assoc
          next if arr.empty?
          i = rng.rand(arr.length)
          x = rng.rand(1000)
          vec = vec.assoc(i, x)
          arr[i] = x
        when 2 # pop
          next if arr.empty?
          vec = vec.pop
          arr.pop
        when 3 # 随机读
          next if arr.empty?
          i = rng.rand(arr.length)
          assert_equal arr[i], vec[i], "round #{round} 读下标 #{i}"
        end
        assert_equal arr, vec.to_a, "round #{round} 模型不一致"
        versions << vec
        expected << arr.dup
      end

      versions.each_with_index do |ver, i|
        assert_equal expected[i], ver.to_a, "round #{round} 第 #{i} 个历史版本被污染"
      end
    end
  end

  # 性质二：push / pop 互逆。
  def test_push_then_pop_roundtrip
    rng = Random.new(7)
    50.times do
      v = random_vector(rng)
      x = rng.rand(1000)
      assert_equal v, v.push(x).pop
    end
  end

  def test_pop_then_push_roundtrip
    rng = Random.new(8)
    50.times do
      v = random_nonempty(rng)
      assert_equal v, v.pop.push(v.last)
    end
  end

  # 性质三：assoc 只改一个位置；写回原值等于没改。
  def test_assoc_changes_only_target
    rng = Random.new(9)
    50.times do
      v = random_nonempty(rng)
      i = rng.rand(v.count)
      x = rng.rand(1000)
      w = v.assoc(i, x)
      assert_equal x, w[i]
      v.count.times { |j| assert_equal v[j], w[j] if j != i }
    end
  end

  def test_assoc_with_same_value_is_equal
    rng = Random.new(10)
    50.times do
      v = random_nonempty(rng)
      i = rng.rand(v.count)
      assert_equal v, v.assoc(i, v[i])
    end
  end

  # 性质四：to_a / from_a 往返不变。
  def test_from_a_to_a_roundtrip
    rng = Random.new(11)
    50.times do
      v = random_vector(rng)
      assert_equal v, PersistentVector.from_a(v.to_a)
    end
  end

  # 性质五：size / first / last 与 to_a 一致。
  def test_size_first_last_match
    rng = Random.new(12)
    50.times do
      v = random_vector(rng)
      assert_equal v.to_a.size, v.size
      assert_equal v.to_a.first, v.first
      assert_equal v.to_a.last, v.last
    end
  end

  # 性质六：hash / eql 契约 —— 相等的向量必须有相等的哈希。
  def test_hash_contract
    rng = Random.new(13)
    50.times do
      v = random_vector(rng)
      w = PersistentVector.from_a(v.to_a)
      assert_equal v, w
      assert_equal v.hash, w.hash
    end
  end

  # 性质七：越界一定抛 IndexError。
  def test_out_of_bounds_raises
    rng = Random.new(14)
    50.times do
      v = random_vector(rng)
      assert_raises(IndexError) { v[v.count] }
      assert_raises(IndexError) { v[-1 - rng.rand(1000)] }
      assert_raises(IndexError) { v.assoc(v.count, :x) }
      assert_raises(IndexError) { v.assoc(-1, :x) }
    end
    assert_raises(IndexError) { PersistentVector.empty.pop }
  end

  # 性质八：结构不变量 —— shift 是 BITS 的倍数、容量守恒（不会过深也不会溢出）。
  def test_structural_invariants
    rng = Random.new(15)
    50.times do
      v = random_vector(rng, max_len: 2000)
      next if v.empty?
      assert_equal 0, v.shift % BITS
      assert_equal 0, v.tail_offset % PersistentVector::WIDTH
      assert_equal v.count, v.tail_offset + v.tail.length
      assert_includes (1..PersistentVector::WIDTH), v.tail.length
      assert_operator v.tail_offset, :<=, (1 << (v.shift + BITS))
      if v.shift > 0
        assert_operator v.tail_offset, :>, (1 << v.shift)
      else
        assert_operator v.tail_offset, :<=, (1 << BITS)
      end
    end
  end

  # 性质九：整棵树由冻结的 Array 节点构成，叶子层不再有子节点。
  def test_tree_is_frozen_and_well_formed
    [1, 32, 33, 64, 65, 100, 1024, 1025, 1057].each do |n|
      v = PersistentVector.from_a((0...n).to_a)
      assert_predicate v.tail, :frozen? unless v.tail.nil?
      check_node(v.root, v.shift) unless v.root.nil?
    end
  end

  # 性质十：边界尺寸下，随机操作序列仍与 Ruby Array 一致。
  def test_boundary_sizes_match_array
    rng = Random.new(20)
    BOUNDARY_SIZES.each do |n|
      v = PersistentVector.from_a((0...n).to_a)
      arr = (0...n).to_a
      20.times do
        case rng.rand(4)
        when 0 # push
          x = rng.rand(1000)
          v = v.push(x)
          arr.push(x)
        when 1 # assoc
          next if arr.empty?
          i = rng.rand(arr.length)
          x = rng.rand(1000)
          v = v.assoc(i, x)
          arr[i] = x
        when 2 # pop
          next if arr.empty?
          v = v.pop
          arr.pop
        when 3 # 读
          next if arr.empty?
          i = rng.rand(arr.length)
          assert_equal arr[i], v[i], "size #{n} 读下标 #{i}"
        end
        assert_equal arr, v.to_a, "size #{n} 模型不一致"
      end
    end
  end

  # 性质十一：pop 跨过节点边界（tail 合并 / 根剥层）结果仍正确。
  def test_pop_across_node_boundaries
    [33, 65, 1057, 32_801].each do |n|
      v = PersistentVector.from_a((0...n).to_a)
      assert_equal (0...(n - 2)).to_a, v.pop.pop.to_a, "pop 跨节点边界 #{n}"
    end
  end

  # 性质十二：push 到满时会长出新的树层，内容仍正确。
  def test_push_grows_a_new_tree_level
    [64, 1056, 32_800].each do |n|
      v = PersistentVector.from_a((0...n).to_a)
      w = v.push(n)
      assert_equal (0..n).to_a, w.to_a
      assert_operator w.shift, :>, v.shift, "push 后应长出新的树层 #{n}"
    end
  end

  # 性质十三：不同构造路径（from_a / 逐次 push）得到同一个值。
  def test_independent_construction_paths_agree
    rng = Random.new(21)
    50.times do
      arr = Array.new(rng.rand(100)) { rng.rand(1000) }
      by_push = arr.reduce(PersistentVector.empty) { |v, x| v.push(x) }
      assert_equal PersistentVector.from_a(arr), by_push
      assert_equal arr, by_push.to_a
    end
  end

  # 性质十四：把 b 逐个 push 到 a 上，等价于数组拼接（Clojure 的 through-transient 类比）。
  def test_concat_via_append_matches_array_concat
    rng = Random.new(22)
    50.times do
      a = Array.new(rng.rand(50)) { rng.rand(1000) }
      b = Array.new(rng.rand(50)) { rng.rand(1000) }
      via_append = b.reduce(PersistentVector.from_a(a)) { |v, x| v.push(x) }
      assert_equal a + b, via_append.to_a
    end
  end

  # 性质十五：相等但构造路径不同的向量，哈希必须一致。
  def test_hash_consistent_across_construction_paths
    rng = Random.new(23)
    50.times do
      arr = Array.new(rng.rand(100)) { rng.rand(1000) }
      a = PersistentVector.from_a(arr).hash
      b = arr.reduce(PersistentVector.empty) { |v, x| v.push(x) }.hash
      assert_equal a, b
    end
  end

  # 性质十六：peek 语义 —— 刚 push 进去的值就是 last。
  def test_peek_last_consistent
    rng = Random.new(24)
    50.times do
      v = random_vector(rng)
      x = rng.rand(1000)
      assert_equal x, v.push(x).last
      assert_equal v.to_a.first, v.first unless v.empty?
    end
  end

  # 性质十七：Enumerable 的 reduce/map 与普通数组一致。
  def test_enumerable_matches_array
    rng = Random.new(25)
    50.times do
      v = random_vector(rng)
      assert_equal v.to_a.reduce(0) { |s, x| s + x }, v.reduce(0) { |s, x| s + x }
      assert_equal v.to_a.map { |x| x * 2 }, v.map { |x| x * 2 }
    end
  end

  private

  def random_vector(rng, max_len: 100)
    v = PersistentVector.empty
    rng.rand(max_len).times { v = v.push(rng.rand(1000)) }
    v
  end

  def random_nonempty(rng, max_len: 100)
    v = random_vector(rng, max_len: max_len)
    v.empty? ? v.push(0) : v
  end

  def check_node(node, level)
    assert_instance_of Array, node
    assert_predicate node, :frozen?
    return if level.zero?
    node.each { |child| check_node(child, level - BITS) unless child.nil? }
  end
end
