LOCK = Mutex.new

$stdout.sync = true
$stdin.sync = true

def can_produce?
  @count < @num_of_nest
end

def can_consume?
  @count > 0
end

def t_producer
  loop do
    LOCK.synchronize do
      if can_produce?
        @count += 1
        print('(')
      end
    end
  end
end

def t_consumer
  loop do
    LOCK.synchronize do
      if can_consume?
        @count -= 1
        print(')')
      end
    end
  end
end

@count = 0
@num_of_nest = ARGV[0].to_i
@num_of_thread = ARGV[1].to_i

# require "pry"; binding.pry

threads = []

@num_of_thread.times do
  threads << Thread.new { t_producer }
  threads << Thread.new { t_consumer }
end

threads.each {|t| t.join}
