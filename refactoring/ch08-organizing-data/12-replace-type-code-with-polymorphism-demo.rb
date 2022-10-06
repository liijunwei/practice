require 'pry'
require 'rspec'

module MountainBike
  TIRE_WIDTH_FACTOR = 1
  FRONT_SUSPENSION_FACTOR = 1
  REAR_SUSPENSION_FACTOR = 1

  def initialize(params)
    @commission = 1
    @base_price = 1
    @front_suspension_price = 1
    @rear_suspension_price = 1
    @rear_fork_travel = 1
    params.each { |key, value| instance_variable_set "@#{key}", value }
  end

  def off_road_ability
    result = @tire_width * TIRE_WIDTH_FACTOR

    if @type_code == :front_suspension || @type_code == :full_suspension
      result += @front_fork_travel * FRONT_SUSPENSION_FACTOR
    end

    if @type_code == :full_suspension
      result += @rear_fork_travel * REAR_SUSPENSION_FACTOR
    end

    result
  end

  def price
    case @type_code
    when :rigid
      (1 + @commission) * @base_price
    when :front_suspension
      (1 + @commission) * @base_price + @front_suspension_price
    when :full_suspension
      (1 + @commission) * @base_price + @front_suspension_price + @rear_suspension_price
    end
  end
end

class RigidMountainBike
  include MountainBike
end

class FrontSuspensionMountainBike
  include MountainBike
end

class FullSuspensionMountainBike
  include MountainBike
end

RSpec.describe MountainBike do
  specify do
    bike1 = RigidMountainBike.new(:type_code => :rigid, :tire_width => 2.5)
    bike2 = FrontSuspensionMountainBike.new(:type_code => :front_suspension, :tire_width => 2, :front_fork_travel => 3)
    bike3 = FullSuspensionMountainBike.new(:type_code => :full_suspension, :tire_width => 2, :front_fork_travel => 3)

    bike1.off_road_ability
    bike1.price
    bike2.off_road_ability
    bike2.price
    bike3.off_road_ability
    bike3.price
  end
end
