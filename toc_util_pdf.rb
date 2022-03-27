require 'pdf-reader'

# https://github.com/yob/pdf-reader
# not quite handy :(

reader = PDF::Reader.new("#{ENV['HOME']}/Desktop/99-bottles.pdf")
reader.pages[2..5].each {|page| puts page.text}
