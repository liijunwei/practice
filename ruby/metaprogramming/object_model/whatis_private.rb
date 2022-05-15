class C
  def method_public
    self.method_private
  end

  private

  def method_private
    "demo"
  end
end

# 为啥 ruby 2.6.5 会报错
# 为啥 ruby 3.1.1 不会报错
p C.new.method_public

