event "the sky is falling" do
  @sky_height < 300
end

event "it's getting closer" do
  @sky_height < @mountains_height
end

event "starting raining..." do
  true
end

setup do
  puts "#{Time.now} [INFO] Setting up sky"
  @sky_height = 100
end

setup do
  puts "#{Time.now} [INFO] Setting up mountains"
  @mountains_height = 200
end

