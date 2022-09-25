def check(n)
  upper  = 2 ** n
  result = upper * (upper - 1) / 2

  result
end

p check(32) # 9223372034707292160
