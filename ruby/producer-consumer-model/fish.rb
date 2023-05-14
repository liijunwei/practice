$stdout.sync = true

def print_char(char)
  loop { print(char) }
end

FSM = [
  {from: :A, to: :B, event: '<'},
  {from: :B, to: :C, event: '>'},
  {from: :C, to: :D, event: '<'},
  {from: :D, to: :A, event: '_'},
  {from: :A, to: :E, event: '>'},
  {from: :E, to: :F, event: '<'},
  {from: :F, to: :D, event: '<'}
]

MUTEX = Mutex.new
COND_VAR = Thread::ConditionVariable.new

@current_state = :A

def next_state(event)
  transition = FSM.find do |transition|
    transition[:from] == @current_state && transition[:event] == event
  end

  # assert
  raise "invalid transition, current_state: #{@current_state}, event: #{event}" if transition.nil?

  transition[:to]
end

def main
  thread_pattern = '<<<>>>___'

  threads = thread_pattern.chars.map do |char|
    Thread.new { print_char(char) }
  end

  threads.each { |t| t.join }
end

main
