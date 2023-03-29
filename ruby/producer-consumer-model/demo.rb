LOCK = Mutex.new

$stdout.sync = true
$stdin.sync = true

class Counter
  attr_reader :mutex

  def initialize(n)
    @mutex = Mutex.new
    @count = 0
    @n = n
  end

  def increment
    @count += 1
  end

  def decrement
    @count -= 1
  end

  def can_produce?
    @count < @n
  end

  def can_consume?
    @count > 0
  end
end

def t_produce(counter)
  loop do
    counter.mutex.lock

    if counter.can_produce?
      counter.increment
      print('(')
    end

    counter.mutex.unlock
  end
end

def t_consume(counter)
  loop do
    counter.mutex.lock

    if counter.can_consume?
      counter.decrement
      print(')')
    end

    counter.mutex.unlock
  end
end

def main
  n = ARGV[0].to_i
  c = Counter.new(n)
  t = ARGV[1].to_i

  threads = []

  t.times do
    threads << Thread.new { t_produce(c) }
    threads << Thread.new { t_consume(c) }
  end

  threads.each {|t| t.join}
end

main
