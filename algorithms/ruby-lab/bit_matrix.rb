# graph data structure
# author: tenderlove
# source code: https://tenderlovemaking.com/2023/03/19/bitmap-matrix-and-undirected-graphs-in-ruby.html
class BitMatrix
  def initialize size
    @size = size
    size = (size + 7) & -8 # round up to the nearest multiple of 8
    @row_bytes = size / 8
    @buffer = "\0".b * (@row_bytes * size)
  end

  def initialize_copy other
    @buffer = @buffer.dup
  end

  def set x, y
    raise IndexError if y >= @size || x >= @size

    x, y = [y, x].sort

    row = x * @row_bytes
    column_byte = y / 8
    column_bit = 1 << (y % 8)

    @buffer.setbyte(row + column_byte, @buffer.getbyte(row + column_byte) | column_bit)
  end

  def set? x, y
    raise IndexError if y >= @size || x >= @size

    x, y = [y, x].sort

    row = x * @row_bytes
    column_byte = y / 8
    column_bit = 1 << (y % 8)

    (@buffer.getbyte(row + column_byte) & column_bit) != 0
  end

  def each_pair
    return enum_for(:each_pair) unless block_given?

    @buffer.bytes.each_with_index do |byte, i|
      row = i / @row_bytes
      column = i % @row_bytes
      8.times do |j|
        if (1 << j) & byte != 0
          yield [row, (column * 8) + j]
        end
      end
    end
  end

  def to_dot
    "graph g {\n" + each_pair.map { |x, y| "#{x} -- #{y};" }.join("\n") + "\n}"
  end

  def edges
    @edges ||= each_pair.to_a
  end

  # 和v相邻的所有顶点
  def adj(v)
    v1 = edges.select {|from, _to| from == v}.map {|_from, to| to}
    v2 = edges.select {|_from, to| to   == v}.map {|from, _to| from}

    v1 + v2
  end

  def vertices
    edges.flatten.uniq
  end

  def V
    vertices.count
  end

  def E
    edges.count
  end
end
