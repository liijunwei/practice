foo = <<-RUBY
  def foo(a,b)
    a+b
  end

  foo(9,8)
RUBY

require 'ripper'
demo1 = Ripper.lex(foo)

require 'parser/current'
demo2 = Parser::CurrentRuby.parse(foo)

require 'ruby_parser'
demo3 = RubyParser.new.parse(foo)

require "pry"; binding.pry
