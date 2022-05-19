# page 115

class Book
  def initialize
    @title = "Different Seasons"
    @subtitle = "Rita Heyworth and Shawshank Redemption"
  end

  def GetTitle
    @title
  end

  def title2
    @subtitle
  end

  def LEND_TO_USER(user)
    "Lending #{@title} to #{user}"
  end
end

class BookRefactored
  def initialize
    @title = "Different Seasons"
    @subtitle = "Rita Heyworth and Shawshank Redemption"
  end

  def title
    @title
  end

  def subtitle
    @subtitle
  end

  def lend_to(user)
    "Lending #{@title} to #{user}"
  end

  def self.deprecate(old_method, new_method)
    define_method(old_method) do |*args, &block|
      warn "Warning: #{old_method}() is deprecated. Please use #{new_method}()."
      send(new_method, *args, &block)
    end
  end

  deprecate :GetTitle,     :title
  deprecate :title2,       :subtitle
  deprecate :LEND_TO_USER, :lend_to
end
