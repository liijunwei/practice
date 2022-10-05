row = []
row[0] = "Liverpool"
row[1] = "15"
puts "before: #{row.inspect}"

class Performance
  attr_accessor :name
  attr_accessor :wins

  def initialize
    @data = []
  end

  def []=(index, value)
    @data.insert(index, value)
  end

  def [](index)
    @data[index]
  end
end

row = Performance.new
row.name = "Liverpool"
row.wins = "15"
puts "after:  #{row.inspect}"
