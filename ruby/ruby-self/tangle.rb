module Printable
  def print
    puts "#{__LINE__} #{__method__}"
  end

  def prepare_cover
    puts "#{__LINE__} #{__method__}"
  end
end

module Document
  def print_to_screen
    prepare_cover
    format_for_screen
    print
  end

  def format_for_screen
    puts "#{__LINE__} #{__method__}"
  end

  def print
    puts "#{__LINE__} #{__method__}"
  end
end

class Book
  include Printable
  include Document
end

book = Book.new
book.print_to_screen
# puts Book.ancestors
