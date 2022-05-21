class Fibber
  def initialize
    @cache = {}
  end

  def fib(n)
    if n < 2
      n
    else
      @cache[n] ||= fib(n - 1) + fib(n - 2)
    end
  end
end

# time ruby algorithms/fibonacci_sequence/cached_version.rb 40
# 0.05s
p Fibber.new.fib ARGV[0].to_i

