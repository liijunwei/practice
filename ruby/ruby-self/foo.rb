class Foo
  def fizz
    # self.buzz
    buzz
  end

  private

  def buzz
    puts __LINE__
  end
end

# 私有方法服从一个简单的规则: 不能明确指定一个接受者来调用一个私有方法
# 即 每次调用一个私有方法时, 只能调用于隐含的接收者self上

Foo.new.fizz
