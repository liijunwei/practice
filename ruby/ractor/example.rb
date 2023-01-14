tripple_number_ractor = Ractor.new do
  puts "I will receive a message soon"
  msg = Ractor.receive
  puts "I will return a tripple of what I receive"
  msg * 3
end
# I will receive a message soon
puts tripple_number_ractor.send(15) # mailman takes message to the door
# I will return a tripple of what I receive
puts tripple_number_ractor.take # mailman takes the response
# => 45
