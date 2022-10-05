puts 'before'
row = []
row[0] = "Liverpool"
row[1] = "15"
p row

puts '=' * 20
puts 'after'
class Performance
  attr_accessor :name
  attr_accessor :wins
end

row = Performance.new
row.name = "Liverpool"
row.wins = "15"
p row
