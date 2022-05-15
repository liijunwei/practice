module Printable
  def print
  end

  def prepare_cover
  end
end

module Document
  def print_to_screen
    prepare_cover
    format_for_screen
    print
  end

  def format_for_screen
  end

  def print
  end
end

class Book
  include Document
  include Printable

  p ancestors
end

b = Book.new
p b.method(:print)

# p Book.ancestors # [Book, Printable, Document, Object, Kernel, BasicObject]


