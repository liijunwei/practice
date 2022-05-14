# run this script every 1 seconds to check all events
# watch -d -n1 "ruby ruby/metaprogramming/blocks/monitor_framework/redflag.rb"

# page 96

def event(name, &block)
  puts "#{Time.now} ALERT: #{name}" if yield
  
end

def setup
end

# require 'pry'; binding.pry

script_dir = File.dirname(__FILE__)
Dir.glob("#{script_dir}/events/*events.rb").each { |file| load file }

