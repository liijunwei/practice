# frozen_string_literal: true

# 无 tail 的基线版本，只用于 benchmark 里对照 tail 优化（第 3 篇）省了多少。
# 行为和 PersistentVector 完全一致，差别仅是「没有把最后一段单独提出去」。
class PersistentVectorNoTail
  BITS  = 5
  WIDTH = 1 << BITS
  MASK  = WIDTH - 1

  attr_reader :count, :shift, :root

  def self.empty
    new(0, 0, nil)
  end

  def self.from_a(ary)
    ary.reduce(empty) { |v, x| v.push(x) }
  end

  def initialize(count, shift, root)
    @count = count
    @shift = shift
    @root  = root
    freeze
  end

  def [](index)
    unless (0...@count).cover?(index)
      raise IndexError, "index #{index} out of bounds for size #{@count}"
    end

    node  = @root
    level = @shift
    while level > 0
      node = node[(index >> level) & MASK]
      level -= BITS
    end
    node[index & MASK]
  end

  def push(value)
    return self.class.new(1, 0, new_path(0, 0, value)) if @count.zero?

    if @count == (1 << (@shift + BITS))
      grown = Array.new(WIDTH)
      grown[0] = @root
      self.class.new(@count + 1, @shift + BITS,
                     do_assoc(grown, @shift + BITS, @count, value))
    else
      self.class.new(@count + 1, @shift, do_assoc(@root, @shift, @count, value))
    end
  end

  def pop
    raise IndexError, "pop from empty vector" if @count.zero?
    return self.class.empty if @count == 1

    new_root  = pop_tail(@root, @shift, @count - 1)
    new_shift = @shift
    if new_shift > 0 && new_root[1].nil?
      new_root  = new_root[0]
      new_shift -= BITS
    end

    self.class.new(@count - 1, new_shift, new_root)
  end

  def to_a
    Array.new(@count) { |i| self[i] }
  end

  private

  def do_assoc(node, level, index, value)
    ret = node.dup
    if level.zero?
      ret[index & MASK] = value
    else
      subidx = (index >> level) & MASK
      ret[subidx] =
        if node[subidx]
          do_assoc(node[subidx], level - BITS, index, value)
        else
          new_path(level - BITS, index, value)
        end
    end
    ret.freeze
  end

  def new_path(level, index, value)
    ret = Array.new(WIDTH)
    if level.zero?
      ret[index & MASK] = value
    else
      ret[(index >> level) & MASK] = new_path(level - BITS, index, value)
    end
    ret.freeze
  end

  def pop_tail(node, level, index)
    subidx = (index >> level) & MASK
    if level.zero?
      return nil if subidx.zero?
      ret = node.dup
      ret[subidx] = nil
      return ret.freeze
    end

    child = pop_tail(node[subidx], level - BITS, index)
    if child.nil?
      return nil if subidx.zero?
      ret = node.dup
      ret[subidx] = nil
      ret.freeze
    else
      ret = node.dup
      ret[subidx] = child
      ret.freeze
    end
  end
end
