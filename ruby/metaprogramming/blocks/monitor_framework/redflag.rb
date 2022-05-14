# run this script every 1 seconds to check all events
# watch -d -n1 "ruby ruby/metaprogramming/blocks/monitor_framework/redflag.rb"

# page 96

def event(name, &block)
  @events[name] = block
end

def setup(&block)
  @setups << block
end

script_dir = File.dirname(__FILE__)
Dir.glob("#{script_dir}/events/*events.rb").each do |file|
  @setups = []
  @events = {}
  load file

  @events.each_pair do |name, event|
    env = Object.new
    @setups.each do |setup|
      env.instance_eval &setup
    end

    puts "#{Time.now} [WARN] #{name}" if env.instance_eval &event
  end
end

