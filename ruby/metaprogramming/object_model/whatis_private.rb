class C
  def method_public
    self.method_private
  end

  private

  def method_private
    "demo"
  end
end

# https://www.ruby-lang.org/en/news/2019/12/25/ruby-2-7-0-released/
# when ruby version <  2.7 =>  private method `method_private' called(NoMethodError)
# when ruby version >= 2.7 =>  OK
p C.new.method_public

