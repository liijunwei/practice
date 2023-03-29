LOCK = Mutex.new

$stdout.sync = true
$stdin.sync = true

def can_produce?
  @count < NUM_OF_NEST
end

def can_consume?
  @count > 0
end

def t_produce
  loop do
    LOCK.lock
    if can_produce?
      @count += 1
      print('(')
    end
    LOCK.unlock
  end
end

def t_consume
  loop do
    LOCK.lock
    if can_consume?
      @count -= 1
      print(')')
    end
    LOCK.unlock
  end
end

@count = 0
NUM_OF_NEST = ARGV[0].to_i
num_of_thread = ARGV[1].to_i

threads = []

num_of_thread.times do
  threads << Thread.new { t_produce }
  threads << Thread.new { t_consume }
end

threads.each {|t| t.join}
