# run this script every 1 seconds to check all events
# watch -d -n1 "ruby ruby/metaprogramming/blocks/monitor_framework/redflag_v2.rb"

# page 98

# This is too much T.T

require 'redis'

def redis
  @instance ||= Redis.new
end

lambda {
  setups = []
  events = {}

  Kernel.send :define_method, :event do |name, &block|
    events[name] = block
  end

  Kernel.send :define_method, :setup do |&block|
    setups << block
  end

  Kernel.send :define_method, :each_event do |&block|
    events.each_pair do |name, event_block|
      block.call name, event_block
    end
  end

  Kernel.send :define_method, :each_setup do |&block|
    setups.each do |setup|
      block.call setup
    end
  end
}.call

script_dir = File.dirname(__FILE__)
Dir.glob("#{script_dir}/events/*events.rb").each do |file|
  load file

  each_event do |name, event_block|
    env = Object.new
    each_setup do |setup|
      env.instance_eval &setup
    end

    puts "#{Time.now} [WARN] #{name}" if env.instance_eval &event_block
  end
end

