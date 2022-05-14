# run this script every 5 seconds to detect events

def monthly_
  
end

def event(name)
  puts "#{Time.now} ALERT: #{name}" if yield
end

# require 'pry'; binding.pry

Dir.glob("#{File.dirname(__FILE__)}/*events.rb").each { |file| load file }

