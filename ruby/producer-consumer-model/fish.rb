$stdout.sync = true

# ref: https://jyywiki.cn/OS/2023/build/lect9.ipynb

# Q: what's the difference between `MUTEX.lock {}` and `MUTEX.lock + MUTEX.unlock` ?
# Q: what's the meaning of @quota?

FSM = [
  {from: :A, to: :B, event: '<'},
  {from: :B, to: :C, event: '>'},
  {from: :C, to: :D, event: '<'},
  {from: :D, to: :A, event: '_'},
  {from: :A, to: :E, event: '>'},
  {from: :E, to: :F, event: '<'},
  {from: :F, to: :D, event: '>'}
]

MUTEX = Mutex.new
COND_VAR = Thread::ConditionVariable.new

@current_state = :A
@quota = 1

def next_state(event)
  transition = FSM.find do |transition|
    transition[:from] == @current_state && transition[:event] == event
  end

  transition && transition[:to]
end

def can_print?(event)
  next_state(event) && @quota > 0
end

def fish_before(event)
  MUTEX.lock

  while !can_print?(event) do
    COND_VAR.wait(MUTEX)
  end
  @quota -= 1

  MUTEX.unlock
end

def fish_after(event)
  MUTEX.lock

  @quota += 1
  @current_state = next_state(event)
  # assert @current_state
  raise "invalid transition, current_state: #{@current_state}, event: #{event}" if @current_state.nil?
  COND_VAR.broadcast

  MUTEX.unlock
end

def fish_thread(char)
  loop do
    fish_before(char)
    print(char)
    fish_after(char)
  end
end

def main
  thread_pattern = '<<<>>>___'

  threads = thread_pattern.chars.map do |char|
    Thread.new { fish_thread(char) }
  end

  threads.each { |t| t.join }
end

main
