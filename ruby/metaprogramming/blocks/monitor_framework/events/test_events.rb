event "the sky is falling" do
  @sky_height < 300
end

event "it's getting closer" do
  @sky_height < @mountains_height
end

event "starting raining..." do
  true
end

event "you have beed slained" do
  true
end

setup do
  puts "#{Time.now} [INFO] Setting up sky"
  @sky_height = redis.get("sky_height").to_i
end

setup do
  puts "#{Time.now} [INFO] Setting up mountains"
  @mountains_height = 200
end

