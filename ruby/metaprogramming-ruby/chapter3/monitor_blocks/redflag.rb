def event(name)
  puts "ALERT: #{name}" if yield
end

# require 'pry'; binding.pry

Dir.glob("#{File.dirname(__FILE__)}/*events.rb").each { |file| load file }

