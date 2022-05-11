# https://refactoring.guru/design-patterns/adapter/ruby/example#lang-features

def client(target)
  "#{target.request}"
end

class Target
  def request
    "Target: This is the original request..."
  end
end

class UnfitTarget
  def special_request
    "UnfitTarget: I have some special behavior that you might want, but my interface is not aligned, you'd better use an adapter..."
  end

  def method_missing(m, *args, &block)
    "UnfitTarget: I don't understand..."
  end
end

# The adapter knows how `UnfitTarget` works
class UnfitTargetAdapter < Target
  def initialize(adaptee)
    @adaptee = adaptee
  end

  def request
    @adaptee.special_request
  end
end

puts "when client don't need any special behavior"
puts client(Target.new)
puts

puts "when client need some special behavior, but their api is not consitent with existing ones"
puts client(UnfitTarget.new)
puts

puts "So I introduce an adapter..."
wanted_target = UnfitTarget.new
puts client(UnfitTargetAdapter.new(wanted_target))
puts


