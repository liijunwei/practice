def event(name)
  puts "ALERT: #{name}" if yield
end

Dir.glob("*event.rb").each { |file| load file }

