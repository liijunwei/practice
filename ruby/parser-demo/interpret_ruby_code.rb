# https://www.youtube.com/watch?v=ySuMOEVLaMw

# <  ruby 1.8
# source code -> lexer(to tokens) -> parser(to ast) -> interpreter

# >= ruby 1.9
# source code -> lexer(to tokens) -> parser(to ast) -> compile to bytecode


foo = <<-RUBY
  def foo(a,b)
    a+b
  end

  foo(9,8)
RUBY

require 'ripper'
demo1 = Ripper.lex(foo)

demo2 = Ripper.sexp(foo)

require "pry"; binding.pry
