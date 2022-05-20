# https://medium.com/rubycademy/understanding-the-eigenclass-in-less-than-5-minutes-dcb8ca223eb4
# https://gist.github.com/mehdi-farsi/4f0923d833b93fe38fac87462b1cd8dd#file-eigenclass_03-rb

# TODO
# 越看越淦!!!

class Greeting
  def self.hello
    'hello world!'
  end

  def self.eigenclass
    class << self
      self
    end
  end
end

p Greeting.eigenclass      # => #<Class:Greeting>
p Greeting.eigenclass.name # => nil

p Greeting.singleton_methods                  # => [:hello, :eigenclass]
p Greeting.eigenclass.instance_methods(false) # => [:hello, :eigenclass]

