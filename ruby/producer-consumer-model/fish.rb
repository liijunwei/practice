$stdout.sync = true

def print_char(char)
  loop { print(char) }
end

def main
  thread_pattern = '<<<>>>___'

  threads = thread_pattern.chars.map do |char|
    Thread.new { print_char(char) }
  end

  threads.each { |t| t.join }
end

main
