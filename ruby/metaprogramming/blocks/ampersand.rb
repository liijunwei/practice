def math(a, b)
  yield(a, b)
end

def teach_match(a, b, &operation)
  puts "Let's do the math:"
  puts math(a, b, &operation)
end

teach_match(2, 3) { |x, y| x * y }
