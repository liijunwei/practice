class Roulette
  def method_missing(name, *args)
    person = name.to_s.capitalize
    number = 0

    3.times do
      number = rand(10) + 1
      puts "#{number}..."
    end

    "#{person} got a #{number}"
  end
end

number_of = Roulette.new
puts number_of.bob
puts number_of.frank

# why "stack level too deep (SystemStackError)" ?
# because number variable(method) is out of scope, and is missing all the way down

# how to fix?
# initialize `number` before the loop
