class Fibber
  def fib(n)
    if n < 2
      n
    else
      fib(n - 1) + fib(n - 2)
    end
  end
end

# time ruby algorithms/fibonacci_sequence/no_cache_version.rb 40
# 9.8s
p Fibber.new.fib ARGV[0].to_i

