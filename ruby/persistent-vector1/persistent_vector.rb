# frozen_string_literal: true

# 持久化向量（带 tail 优化版，对应 Hypirion 系列第 3 篇）。
#
# 结构：
#   - 树：32 叉 bit-partitioned trie，存「前面」的 tail_offset 个元素；所有叶子都是满 32 槽。
#   - tail：直接挂在向量头部的最后一段（1..32 个元素），不放进树里。
#   - 查找 / 更新：下标 < tail_offset 走树，否则直接读 tail。
#   - 追加：tail 没满就只复制 tail；满了才把整片满叶子并进树，新 tail 只放一个新元素。
#   - 弹出：tail 多于 1 个元素就只复制 tail；只剩 1 个时把树的最右叶子提出来当新 tail。
#
# 为什么快：每 32 次追加/弹出只有 1 次真正动树，其余 31 次是 O(1) 的小数组复制。
class PersistentVector
  BITS  = 5
  WIDTH = 1 << BITS   # 32
  MASK  = WIDTH - 1   # 0b11111，取低 5 位

  attr_reader :count, :shift, :root, :tail, :tail_offset

  def self.empty
    new(0, 0, nil, nil, 0)
  end

  def self.[](*items)
    items.reduce(empty) { |v, x| v.push(x) }
  end

  def self.from_a(ary)
    ary.reduce(empty) { |v, x| v.push(x) }
  end

  def initialize(count, shift, root, tail, tail_offset)
    @count = count
    @shift = shift
    @root = root
    @tail = tail
    @tail_offset = tail_offset
    freeze
  end

  def size
    @count
  end
  alias length size

  def empty?
    @count.zero?
  end

  # 下标在 tail 里就直接读 tail，否则沿树逐层取 5 个 bit 定位。
  def [](index)
    unless (0...@count).cover?(index)
      raise IndexError, "index #{index} out of bounds for size #{@count}"
    end

    if index >= @tail_offset
      @tail[index - @tail_offset]
    else
      node  = @root
      level = @shift
      while level > 0
        node = node[(index >> level) & MASK]
        level -= BITS
      end
      node[index & MASK]
    end
  end
  alias nth []

  # 更新一个下标；在树里就路径复制，在 tail 里就只复制 tail。
  def assoc(index, value)
    unless (0...@count).cover?(index)
      raise IndexError, "index #{index} out of bounds for size #{@count}"
    end

    if index >= @tail_offset
      new_tail = @tail.dup
      new_tail[index - @tail_offset] = value
      PersistentVector.new(@count, @shift, @root, new_tail.freeze, @tail_offset)
    else
      PersistentVector.new(@count, @shift, do_assoc(@root, @shift, index, value),
                           @tail, @tail_offset)
    end
  end

  def push(value)
    return PersistentVector.new(1, 0, nil, [value].freeze, 0) if @count.zero?

    if @tail.length < WIDTH
      new_tail = @tail.dup
      new_tail[@tail.length] = value
      return PersistentVector.new(@count + 1, @shift, @root, new_tail.freeze, @tail_offset)
    end

    # tail 满了：把这片满叶子并进树，新 tail 只放这一个值。
    new_root, new_shift = push_tail(@root, @shift, @tail_offset, @tail)
    PersistentVector.new(@count + 1, new_shift, new_root, [value].freeze,
                         @tail_offset + WIDTH)
  end
  alias conj push

  def pop
    raise IndexError, "pop from empty vector" if @count.zero?
    return PersistentVector.empty if @count == 1

    if @tail.length > 1
      new_tail = @tail[0...-1].freeze
      return PersistentVector.new(@count - 1, @shift, @root, new_tail, @tail_offset)
    end

    # tail 只剩 1 个：把树的最右叶子提出来当新 tail。
    if @shift.zero?
      promoted  = @root
      new_root  = nil
      new_shift = 0
    else
      new_root, promoted = pop_leaf(@root, @shift, @tail_offset - 1)
      if new_root.nil?
        new_shift = 0
      elsif new_root[1].nil?   # 根只剩一个孩子，剥掉这层根
        new_root  = new_root[0]
        new_shift = @shift - BITS
      else
        new_shift = @shift
      end
    end

    PersistentVector.new(@count - 1, new_shift, new_root, promoted, @tail_offset - WIDTH)
  end

  def first
    empty? ? nil : self[0]
  end

  def last
    empty? ? nil : self[@count - 1]
  end

  def each
    return enum_for(:each) unless block_given?
    @count.times { |i| yield self[i] }
    self
  end
  include Enumerable

  def to_a
    Array.new(@count) { |i| self[i] }
  end

  def inspect
    to_a.inspect
  end
  alias to_s inspect

  def ==(other)
    other.is_a?(PersistentVector) && other.count == @count &&
      (0...@count).all? { |i| self[i] == other[i] }
  end
  alias eql? ==

  def hash
    # shortcut: O(n) 哈希，学习/小规模够用；把大向量当哈希键时再换成缓存哈希。
    to_a.hash
  end

  private

  # 树内更新：路径复制到叶子，写入 value。树的叶子都是满的，所以沿路一定非空。
  def do_assoc(node, level, index, value)
    ret = node.dup
    if level.zero?
      ret[index & MASK] = value
    else
      subidx = (index >> level) & MASK
      ret[subidx] = do_assoc(node[subidx], level - BITS, index, value)
    end
    ret.freeze
  end

  # 把一片满叶子并进树，返回 [新根, 新 shift]。
  def push_tail(node, shift, index, leaf)
    return [leaf, 0] if node.nil?

    if index == (1 << (shift + BITS))   # 树满了，先长一层根
      grown = Array.new(WIDTH)
      grown[0] = node
      [push_leaf(grown, shift + BITS, index, leaf), shift + BITS]
    else
      [push_leaf(node, shift, index, leaf), shift]
    end
  end

  # 在树的叶子层插入一片满叶子。index 是 32 的倍数，所以直接落到叶槽。
  def push_leaf(node, level, index, leaf)
    ret = node.dup
    subidx = (index >> level) & MASK
    if level == BITS
      ret[subidx] = leaf
    else
      child = node[subidx] ? push_leaf(node[subidx], level - BITS, index, leaf)
                           : new_path_node(level - BITS, index, leaf)
      ret[subidx] = child
    end
    ret.freeze
  end

  # 生成一条只含这片叶子的全新路径。
  def new_path_node(level, index, leaf)
    ret = Array.new(WIDTH)
    if level == BITS
      ret[(index >> level) & MASK] = leaf
    else
      ret[(index >> level) & MASK] = new_path_node(level - BITS, index, leaf)
    end
    ret.freeze
  end

  # 从树里摘掉最右叶子，返回 [新根(可能 nil), 被摘出的叶子]。
  def pop_leaf(node, level, index)
    subidx = (index >> level) & MASK
    if level == BITS
      leaf = node[subidx]
      ret = node.dup
      ret[subidx] = nil
      return [nil, leaf] if subidx.zero?
      return [ret.freeze, leaf]
    end

    child, leaf = pop_leaf(node[subidx], level - BITS, index)
    if child.nil?
      return [nil, leaf] if subidx.zero?
      ret = node.dup
      ret[subidx] = nil
      [ret.freeze, leaf]
    else
      ret = node.dup
      ret[subidx] = child
      [ret.freeze, leaf]
    end
  end
end
