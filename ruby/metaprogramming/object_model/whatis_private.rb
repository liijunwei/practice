class C
  def method_public
    self.method_private
  end

  private

  def method_private
    "demo"
  end
end

# when ruby version <  2.7 =>  private method `method_private' called(NoMethodError)
# when ruby version >= 2.7 =>  private method `method_private' called(NoMethodError)
p C.new.method_public

