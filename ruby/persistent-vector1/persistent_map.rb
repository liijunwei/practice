# frozen_string_literal: true

# 不可变的持久化哈希表（persistent hash map）。
#
# 和刚做的 persistent vector 同一个套路：32 叉 bit-partitioned trie + 路径复制。
# 区别是这里按下标的角色由「键的 hash」承担：hash 的低 30 位切成 6 段、每段 5 bit，
# 从高到低决定走哪个槽；到了最底层，哈希相同的键放进一个小的 [k, v] 桶里线性查找。
# 所有修改只复制路径，其余节点共享，所以旧 map 永远不变。
#
# 用法上对应 Ruby Hash 的读接口；写接口是 assoc / dissoc（返回新 map，不改原值）。
# shortcut: 叶子用线性桶处理冲突，且不做节点压缩、不保留插入顺序；
#           键特别多或要严格替代 Hash 时，再换 bitmap 节点 + 有序结构。
class PersistentMap
  BITS  = 5
  WIDTH = 1 << BITS          # 32
  MASK  = WIDTH - 1          # 低 5 位
  LEVELS = 6                 # 覆盖 hash 低 30 位
  SHIFT0 = BITS * (LEVELS - 1)
  FULL_MASK = (1 << (BITS * LEVELS)) - 1

  attr_reader :count, :root

  def self.empty
    new(0, nil)
  end

  def self.[](*pairs)
    raise ArgumentError, "odd number of arguments for PersistentMap" if pairs.length.odd?
    m = empty
    pairs.each_slice(2) { |k, v| m = m.assoc(k, v) }
    m
  end

  def self.from_h(h)
    h.reduce(empty) { |m, (k, v)| m.assoc(k, v) }
  end

  def initialize(count, root)
    @count = count
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

  def [](key)
    return nil if @root.nil?
    node  = @root
    shift = SHIFT0
    hash  = key.hash & FULL_MASK
    while shift > 0
      node = node[(hash >> shift) & MASK]
      return nil if node.nil?
      shift -= BITS
    end
    node.each { |(k, v)| return v if k.eql?(key) }
    nil
  end

  def key?(key)
    return false if @root.nil?
    node  = @root
    shift = SHIFT0
    hash  = key.hash & FULL_MASK
    while shift > 0
      node = node[(hash >> shift) & MASK]
      return false if node.nil?
      shift -= BITS
    end
    node.any? { |(k, _)| k.eql?(key) }
  end
  alias include? key?
  alias has_key? key?
  alias member? key?

  def fetch(key, *default)
    return self[key] if key?(key)
    return default[0] if default.length == 1
    return yield(key) if block_given?
    raise KeyError, "key not found: #{key.inspect}"
  end

  def assoc(key, value)
    hash = key.hash & FULL_MASK
    if @root.nil?
      new_root = new_path(SHIFT0, hash, key, value)
      added = true
    else
      new_root, added = do_assoc(@root, SHIFT0, hash, key, value)
    end
    PersistentMap.new(@count + (added ? 1 : 0), new_root)
  end

  def dissoc(key)
    return self unless key?(key)
    return PersistentMap.empty if @count == 1
    hash = key.hash & FULL_MASK
    PersistentMap.new(@count - 1, do_dissoc(@root, SHIFT0, hash, key))
  end

  def each
    return enum_for(:each) unless block_given?
    each_node(@root, SHIFT0) { |k, v| yield k, v } if @root
    self
  end
  include Enumerable

  def keys
    map { |k, _| k }
  end

  def values
    map { |_, v| v }
  end

  def to_h
    h = {}
    each { |k, v| h[k] = v }
    h
  end

  def merge(other)
    h = other.respond_to?(:to_h) ? other.to_h : other
    h.reduce(self) { |m, (k, v)| m.assoc(k, v) }
  end

  def select
    return enum_for(:select) unless block_given?
    m = self.class.empty
    each { |k, v| m = m.assoc(k, v) if yield(k, v) }
    m
  end

  def reject
    return enum_for(:reject) unless block_given?
    m = self.class.empty
    each { |k, v| m = m.assoc(k, v) unless yield(k, v) }
    m
  end

  def inspect
    "{#{map { |k, v| "#{k.inspect}=>#{v.inspect}" }.join(", ")}}"
  end
  alias to_s inspect

  def ==(other)
    case other
    when PersistentMap
      other.count == @count && all? { |k, v| other.key?(k) && other[k] == v }
    when Hash
      other.size == @count && all? { |k, v| other.key?(k) && other[k] == v }
    else
      false
    end
  end
  alias eql? ==

  def hash
    # shortcut: O(n) 哈希；Hash 的 hash 与插入顺序无关，所以这里与无序的 == 一致。
    to_h.hash
  end

  private

  def do_assoc(node, shift, hash, key, value)
    if shift.zero?
      i = node.index { |(k, _)| k.eql?(key) }
      if i
        leaf = node.dup
        leaf[i] = [key, value]
        [leaf.freeze, false]
      else
        [(node + [[key, value]]).freeze, true]
      end
    else
      idx = (hash >> shift) & MASK
      child = node[idx]
      new_child, added =
        if child
          do_assoc(child, shift - BITS, hash, key, value)
        else
          [new_path(shift - BITS, hash, key, value), true]
        end
      ret = node.dup
      ret[idx] = new_child
      [ret.freeze, added]
    end
  end

  def do_dissoc(node, shift, hash, key)
    if shift.zero?
      leaf = node.reject { |(k, _)| k.eql?(key) }
      leaf.empty? ? nil : leaf.freeze
    else
      idx = (hash >> shift) & MASK
      child = do_dissoc(node[idx], shift - BITS, hash, key)
      if child.nil?
        return nil if node.count { |c| !c.nil? } == 1
        ret = node.dup
        ret[idx] = nil
        ret.freeze
      else
        ret = node.dup
        ret[idx] = child
        ret.freeze
      end
    end
  end

  def new_path(shift, hash, key, value)
    if shift.zero?
      [[key, value]].freeze
    else
      arr = Array.new(WIDTH)
      arr[(hash >> shift) & MASK] = new_path(shift - BITS, hash, key, value)
      arr.freeze
    end
  end

  def each_node(node, shift, &blk)
    if shift.zero?
      node.each { |(k, v)| blk.call(k, v) }
    else
      node.each { |child| each_node(child, shift - BITS, &blk) unless child.nil? }
    end
  end
end
