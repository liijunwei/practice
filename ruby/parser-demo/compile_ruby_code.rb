# https://www.youtube.com/watch?v=ySuMOEVLaMw

# <  ruby 1.8
# source code -> lexer(to tokens) -> parser(to ast) -> interpreter

# >= ruby 1.9
# source code -> lexer(to tokens) -> parser(to ast) -> compile to bytecode -> RubyVM(execute bytecode)


foo = <<-RUBY
  def foo(a,b)
    a+b
  end

  foo(9,8)
RUBY

require 'ripper'
demo1 = Ripper.lex(foo)
demo2 = Ripper.sexp(foo)

# https://ruby-doc.org/core-2.6/RubyVM/InstructionSequence.html
demo3 = RubyVM::InstructionSequence.compile(foo,
  specialized_instruction: false,
  instructions_unification: false,
  operands_unification: false,
)

p demo3.to_a

demo4 = RubyVM::InstructionSequence.compile(foo).disasm
puts demo4

demo5 = RubyVM::InstructionSequence.compile("puts 9+8").disasm
puts demo5
require "pry"; binding.pry

# other apis to try: https://ruby-doc.org/core-2.6/RubyVM/
