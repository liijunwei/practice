# run this script every 5 seconds to check all events

def event(name)
  puts "#{Time.now} ALERT: #{name}" if yield
end

# require 'pry'; binding.pry

Dir.glob("#{File.dirname(__FILE__)}/*events.rb").each { |file| load file }

