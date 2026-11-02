# frozen_string_literal: true

# 一个「持久化向量」(persistent vector)：不可变、按下标随机访问的序列。
#
# 核心思想（来自 Hypirion 的系列文章）：
#   - 不用一整个可变数组，而是用一棵「有序 trie」存元素。
#   - 每个节点是 32 叉（WIDTH = 32）：内部节点存 32 个子节点，叶子存 32 个值。
#   - 下标被切成 5 bit 一段（BITS = 5），从高位到低位决定每一步走哪个分支。
#   - 所有修改都「只复制路径」(path copying)：从根到被改位置的那条链复制一份，
#     其余节点原样共享，所以旧向量永远不变，新旧向量共享大部分结构。
#
# 复杂度：查找 / 更新 / 追加 / 弹出都是 O(log32 n)。深度极小（10 亿元素才 6 层），
# 所以常被称作「实际上的 O(1)」。
class PersistentVector
  BITS  = 5
  WIDTH = 1 << BITS   # 32
  MASK  = WIDTH - 1   # 0b11111，取低 5 位

  attr_reader :count, :shift, :root

  # 空向量。shift 的约定：shift = BITS * (深度 - 1)，0 表示根节点本身就是叶子。
  def self.empty
    new(0, 0, nil)
  end

  def self.[](*items)
    items.reduce(empty) { |v, x| v.push(x) }
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

  def size
    @count
  end
  alias length size

  def empty?
    @count.zero?
  end

  # 按下标查找。沿着 shift 逐层向下，每次取 5 个 bit 决定走哪个分支。
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
  alias nth []

  # 更新下标 index 处的值，返回新向量；原向量不变。
  def assoc(index, value)
    unless (0...@count).cover?(index)
      raise IndexError, "index #{index} out of bounds for size #{@count}"
    end

    PersistentVector.new(@count, @shift, do_assoc(@root, @shift, index, value))
  end

  # 在末尾追加一个值，返回新向量；原向量不变。
  def push(value)
    return PersistentVector.new(1, 0, new_path(0, 0, value)) if @count.zero?

    # 根节点已满（元素数 == WIDTH^深度）时，长出一层根，把旧根挂到 0 号子节点。
    if @count == (1 << (@shift + BITS))
      grown = Array.new(WIDTH)
      grown[0] = @root
      PersistentVector.new(@count + 1, @shift + BITS,
                           do_assoc(grown, @shift + BITS, @count, value))
    else
      PersistentVector.new(@count + 1, @shift,
                           do_assoc(@root, @shift, @count, value))
    end
  end
  alias conj push

  # 弹出最后一个元素，返回新向量；原向量不变。
  def pop
    raise IndexError, "pop from empty vector" if @count.zero?
    return PersistentVector.empty if @count == 1

    new_root  = pop_tail(@root, @shift, @count - 1)
    new_shift = @shift

    # 根只剩一个孩子时，去掉这层根，避免无用的深度。
    if new_shift > 0 && new_root[1].nil?
      new_root  = new_root[0]
      new_shift -= BITS
    end

    PersistentVector.new(@count - 1, new_shift, new_root)
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

  # 从根往下，把 index 路径上的节点都复制一份，在叶子处写入 value。
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
          new_path(level - BITS, index, value)  # 追加时还没建出来的分支
        end
    end
    ret.freeze
  end

  # 生成一条从当前层到叶子的全新路径（沿路只有 index 这一个元素）。
  def new_path(level, index, value)
    ret = Array.new(WIDTH)
    if level.zero?
      ret[index & MASK] = value
    else
      ret[(index >> level) & MASK] = new_path(level - BITS, index, value)
    end
    ret.freeze
  end

  # 删除最后一个元素：路径复制到底，叶子清空后向上传播「空节点」。
  def pop_tail(node, level, index)
    subidx = (index >> level) & MASK
    if level.zero?
      return nil if subidx.zero?          # 这个叶子被删空了
      ret = node.dup
      ret[subidx] = nil
      return ret.freeze
    end

    child = pop_tail(node[subidx], level - BITS, index)
    if child.nil?
      return nil if subidx.zero?          # 整个子树都被删空了
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
