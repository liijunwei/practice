class C
  def method_public
    self.method_private
  end

  private

  def method_private
    "demo"
  end
end

p C.new.method_public

