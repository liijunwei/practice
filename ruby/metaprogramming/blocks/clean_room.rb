# page 85

class CleanRoom
  def complex_calculation
    1
    # 11
  end

  def do_something
    puts "aaaaaaaa"
  end
end

clean_room = CleanRoom.new
clean_room.instance_eval do
  if complex_calculation > 10
    do_something
  end
end

