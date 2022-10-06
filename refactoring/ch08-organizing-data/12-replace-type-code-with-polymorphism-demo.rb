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
end

class RigidMountainBike
  include MountainBike

  def price
    (1 + @commission) * @base_price
  end

  def off_road_ability
    @tire_width * TIRE_WIDTH_FACTOR
  end
end

class FrontSuspensionMountainBike
  include MountainBike

  def price
    (1 + @commission) * @base_price + @front_suspension_price
  end

  def off_road_ability
    result = @tire_width * TIRE_WIDTH_FACTOR
    result += @front_fork_travel * FRONT_SUSPENSION_FACTOR

    result
  end
end

class FullSuspensionMountainBike
  include MountainBike

  def price
    (1 + @commission) * @base_price + @front_suspension_price + @rear_suspension_price
  end

  def off_road_ability
    result = @tire_width * TIRE_WIDTH_FACTOR
    result += @front_fork_travel * FRONT_SUSPENSION_FACTOR
    result += @rear_fork_travel * REAR_SUSPENSION_FACTOR

    result
  end
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
