class Amazon
  def reviews_of(book)
    sleep 3 # mock slow response

    ["good", "bad"] * 11
  end
end

class Amazon
  alias :old_reviews_of :reviews_of

  def reviews_of(book)
    start      = Time.now
    result     = old_reviews_of(book)
    time_taken = Time.now - start
    if time_taken > 2
      raise "reviews_of() took more than #{time_taken} seconds"
    end
    result
  rescue => e
    puts "#{e.class} #{e.message}"
    []
  end
end

def deserves_a_look?(book)
  amazon = Amazon.new
  amazon.reviews_of(book).size > 20
end

book = "war and peace"
p deserves_a_look?(book)

