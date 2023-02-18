N = ARGV[0].to_i

h = 1

result = []

while h < N/3 do
  h = 3*h + 1
  result << h
end

p result
