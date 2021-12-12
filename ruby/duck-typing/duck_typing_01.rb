# https://www.youtube.com/watch?v=eJZFA72pvEw
class DuckTyping01
  def add_one(s, one)
    s << one
  end
end

result = DuckTyping01.new.add_one([1, 2, 3], 4)
puts result.to_s
result = DuckTyping01.new.add_one("This is the deep ", "end")
puts result.to_s
