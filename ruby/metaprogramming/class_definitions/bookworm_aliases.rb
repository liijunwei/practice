# page 135

require 'timeout'

class Amazon
  def reviews_of(book)
    sleep 2 # mock slow response

    ["good", "bad"] * 11
  end
end

class Amazon
  alias :old_reviews_of :reviews_of

  def reviews_of(book)
    Timeout.timeout(1) { old_reviews_of(book) }
  rescue Timeout::Error => e
    warn "#{e.class} #{e.message}"
    []
  end
end

def deserves_a_look?(book)
  amazon = Amazon.new
  amazon.reviews_of(book).size > 20
end

book = "war and peace"
p deserves_a_look?(book)

