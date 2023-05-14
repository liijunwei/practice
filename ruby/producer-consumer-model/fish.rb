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
  {from: :F, to: :D, event: '>'}
]

def main
  thread_pattern = '<<<>>>___'

  threads = thread_pattern.chars.map do |char|
    Thread.new { print_char(char) }
  end

  threads.each { |t| t.join }
end

main
