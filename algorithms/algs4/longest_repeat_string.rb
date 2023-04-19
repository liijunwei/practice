
# TODO make this right

class SuffixArray
  attr_reader :suffixes

  def initialize(text)
    @suffixes = get_suffixes_of(text)
  end

  def index(i)
    suffixes[i][:k]
  end

  def lcp(i)
    lcp_suffix(suffixes[i][:v], suffixes[i - 1][:v])
  end

  def lcp_suffix(suffix_s, suffix_t)
    max = [suffix_s.size, suffix_t.size].max

    max.times do |i|
      if suffix_s[i] != suffix_t[i]
        return i
      end
    end

    max
  end

  def select(i)
    suffixes[i][:v]
  end

  private

  def get_suffixes_of(text)
    upper_limit = text.size - 1
    text.chars.map.with_index {|_char, index| {k: index, v: text[index..upper_limit]}}.sort_by {|e| e[:v]}
  end
end

class LRS
  def execute
    text = ARGF.read

    n = text.size
    sa = SuffixArray.new(text)
    lrs = ""

    n.times do |i|
      length = sa.lcp(i)

      if(length > lrs.size)
        lrs = sa.select(i)[0..length]
      end
    end

    puts "text:\n#{text}"
    puts "lrs:\n#{lrs}"
    puts
  end
end

LRS.new.execute
