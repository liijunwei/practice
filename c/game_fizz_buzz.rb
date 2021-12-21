# http://wiki.c2.com/?FizzBuzzTest

result =
1.upto(100).map do |n|
  line = ""

  line += "Fizz" if n % 3 == 0
  line += "Buzz" if n % 5 == 0

  line.empty? ? n : line
end

puts result
