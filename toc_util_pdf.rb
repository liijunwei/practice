require 'pdf-reader'

# https://github.com/yob/pdf-reader
# not quite handy :(

reader = PDF::Reader.new("#{ENV['HOME']}/Desktop/99-bottles.pdf")
reader.pages[1..5].map {|page| page.text.to_s.split("\n").reject(&:blank?)}.flatten


toc = reader.pages[1..5]
            .map {|page| page.text.to_s.split("\n")
            .reject(&:blank?)}
            .flatten
            .reject {|line| line.match(/\A\s+\d\z/)}

puts toc
